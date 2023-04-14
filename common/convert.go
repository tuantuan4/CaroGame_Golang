package common

import (
	"fmt"
	"strconv"
)

func StringToInt(number string) int {
	marks, err := strconv.Atoi(number)

	if err != nil {
		fmt.Println("Error during conversion")
		return -1
	}
	return marks
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func ConvertSecondsToHMS(seconds float64) string {
	hours := int(seconds) / 3600

	// Lấy dư để tính số phút
	remainingSeconds := int(seconds) % 3600
	minutes := remainingSeconds / 60

	// Lấy dư để tính số giây
	remainingSeconds %= 60
	result := IntToString(int(hours)) + " giờ " +
		IntToString(int(minutes)) + " phút " +
		IntToString(int(remainingSeconds)) + " giây"
	return result
}
