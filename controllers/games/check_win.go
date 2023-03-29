package games

import (
	"Caro_Game/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitMatrix(matrix [][]int) {
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}
}

func IsWinner(matrix [][]int) int {
	// kiem tra hang`
	for i := 0; i < 3; i++ {
		if matrix[i][0] == matrix[i][1] && matrix[i][1] == matrix[i][2] &&
			matrix[i][0] != 0 && matrix[i][1] != 0 && matrix[i][2] != 0 {
			return matrix[i][0]
		}
	}
	// kiem tra cot
	for i := 0; i < 3; i++ {
		if matrix[0][i] == matrix[1][i] && matrix[1][i] == matrix[2][i] &&
			matrix[0][i] != 0 && matrix[1][i] != 0 && matrix[2][i] != 0 {
			return matrix[0][i]
		}
	}
	//kiem tra duong cheo chinh
	if matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] &&
		matrix[0][0] != 0 && matrix[1][1] != 0 && matrix[2][2] != 0 {
		return matrix[0][0]
	}
	// kiem tra duong cheo phu
	if matrix[0][2] == matrix[1][1] && matrix[1][1] == matrix[2][0] &&
		matrix[0][2] != 0 && matrix[1][1] != 0 && matrix[2][0] != 0 {
		return matrix[0][2]
	}
	return -1
}
func IsEmpty(matrix [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
func TwoPlayerDraw(matrix [][]int, idWinner int) (bool, int, int) {
	if !IsEmpty(matrix) && idWinner < 0 {
		counts := make(map[int]int)
		for _, row := range matrix {
			for _, val := range row {
				counts[val]++
			}
		}
		var unique []int
		for key := range counts {
			unique = append(unique, key)
		}
		return true, unique[0], unique[1]
	}
	return false, -1, -1
}

func PlayerLose(matrix [][]int, idWinner int) int {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] != idWinner && matrix[i][j] != 0 {
				return matrix[i][j]
				break
			}
		}
	}
	return -1
}

func CheckWin(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		matrix := make([][]int, 3)
		InitMatrix(matrix)

		var result []models.Moves
		idGame, err := strconv.Atoi(ctx.Param("id_game"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id not type",
			})
			return
		}
		if err := db.Where("game_id = ?", idGame).Find(&result).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "game not found",
			})
			return
		}
		//init x,y to matrix
		for i := range result {
			matrix[result[i].XCoordinate][result[i].YCoordinate] = result[i].PlayerId
		}
		// check result
		idWinner := IsWinner(matrix)
		ok, id1, id2 := TwoPlayerDraw(matrix, idWinner)
		if ok {
			stringRes := "2 player draw"
			db.Model(&models.User{}).Where("id = ?", id1).UpdateColumn("win", gorm.Expr("draw + ?", 1))
			db.Model(&models.User{}).Where("id = ?", id2).UpdateColumn("lose", gorm.Expr("draw + ?", 1))
			ctx.JSON(200, gin.H{
				"message": stringRes,
			})
		} else if idWinner > 0 {
			idLoser := PlayerLose(matrix, idWinner)
			//update bang games
			db.Model(&models.Games{}).Where("id = ?", idGame).Updates(models.Games{WinnerId: idWinner, LoserId: idLoser})
			//update bang users
			db.Model(&models.User{}).Where("id = ?", idWinner).UpdateColumn("win", gorm.Expr("win + ?", 1))
			db.Model(&models.User{}).Where("id = ?", idLoser).UpdateColumn("lose", gorm.Expr("lose + ?", 1))

			ctx.JSON(200, gin.H{
				"data":    idWinner,
				"IDLoser": idLoser,
			})
		}

	}
}
