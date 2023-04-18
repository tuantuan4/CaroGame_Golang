package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Game struct {
	Board    [9]string
	Winner   string
	NextTurn string
}

func main() {
	myApp := app.New()

	historyLabel := widget.NewLabel("")

	historyButton := widget.NewButton("Lịch sử", func() {
		resp, err := http.Get("http://localhost:8080/v1/games/GetHistory/5")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		var games []Game
		err = json.NewDecoder(resp.Body).Decode(&games)
		if err != nil {
			fmt.Println(err)
			return
		}

		historyText := "Lịch sử trò chơi Tic Tac Toe:\n\n"
		for _, game := range games {
			historyText += fmt.Sprintf("Winner: %s, NextTurn: %s\n", game.Winner, game.NextTurn)
			historyText += fmt.Sprintf("Board: %v\n\n", game.Board)
		}
		historyLabel.SetText(historyText)
	})

	content := container.New(layout.NewVBoxLayout(), historyButton, historyLabel)

	myWindow := myApp.NewWindow("Tic Tac Toe History")
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
