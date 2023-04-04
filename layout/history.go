package main

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()

    myWindow := myApp.NewWindow("Lịch sử trận đấu")


    historyButton := widget.NewButton("Lịch sử", func() {
		

        myWindow.SetContent(container.NewVBox())

        for i := 1; i <= 10; i++ {
            matchButton := widget.NewButton(fmt.Sprintf("Trận đấu %d", i), func() {
                fmt.Println("Bạn đã chọn trận đấu", i)
            })

            myWindow.Content().(*fyne.Container).Add(matchButton)
        }
    })

    content := container.NewVBox(
        historyButton,
    )
    myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400,400))

    myWindow.ShowAndRun()
}
