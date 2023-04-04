package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Đăng nhập")
	myWindow.Resize(fyne.NewSize(400, 400))

	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	loginButton := widget.NewButton("Đăng nhập", func() {
	})

	registerButton := widget.NewButton("Đăng kí", func() {
		registerWindow := myApp.NewWindow("Đăng kí")
		registerWindow.Resize(fyne.NewSize(400, 400))

		registerUsernameEntry := widget.NewEntry()
		registerPasswordEntry := widget.NewPasswordEntry()
		registerConfirmPasswordEntry := widget.NewPasswordEntry()

		registerButton := widget.NewButton("Đăng kí", func() {
		})

		registerContent := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			widget.NewLabel("Đăng kí"),
			registerUsernameEntry,
			registerPasswordEntry,
			registerConfirmPasswordEntry,
			registerButton,
		)

		registerWindow.SetContent(registerContent)
		registerWindow.Show()
	})

	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		widget.NewLabel("Đăng nhập"),
		usernameEntry,
		passwordEntry,
		loginButton,
		widget.NewLabel("Hoặc"),
		registerButton,
	)

	myWindow.SetContent(content)
	myWindow.Show()
	startButton := widget.NewButton("Bắt đầu", func() {
		btn00 := widget.NewButton("", func() {})
		btn01 := widget.NewButton("", func() {})
		btn02 := widget.NewButton("", func() {})
		btn10 := widget.NewButton("", func() {})
		btn11 := widget.NewButton("", func() {})
		btn12 := widget.NewButton("", func() {})
		btn20 := widget.NewButton("", func() {})
		btn21 := widget.NewButton("", func() {})
		btn22 := widget.NewButton("", func() {})

		board := container.NewGridWithColumns(3,
			btn00, btn01, btn02,
			btn10, btn11, btn12,
			btn20, btn21, btn22,
		)

		var turn int
		turn = 0

		btn00.OnTapped = func() {
			fmt.Println("")
			if btn00.Text == "" {
				if turn%2 == 0 {
					btn00.SetText("X")
				} else {
					btn00.SetText("O")
				}
				turn++
			}
		}
		btn01.OnTapped = func() {
			fmt.Println("")
			if btn01.Text == "" {
				if turn%2 == 0 {
					btn01.SetText("X")
				} else {
					btn01.SetText("O")
				}
				turn++
			}
		}
		btn02.OnTapped = func() {
			fmt.Println("")
			if btn02.Text == "" {
				if turn%2 == 0 {
					btn02.SetText("X")
				} else {
					btn02.SetText("O")
				}
				turn++
			}
		}
		btn10.OnTapped = func() {
			fmt.Println("")
			if btn10.Text == "" {
				if turn%2 == 0 {
					btn10.SetText("X")
				} else {
					btn10.SetText("O")
				}
				turn++
			}
		}
		btn11.OnTapped = func() {
			fmt.Println("")
			if btn11.Text == "" {
				if turn%2 == 0 {
					btn11.SetText("X")
				} else {
					btn11.SetText("O")
				}
				turn++
			}
		}
		btn12.OnTapped = func() {
			fmt.Println("")
			if btn12.Text == "" {
				if turn%2 == 0 {
					btn12.SetText("X")
				} else {
					btn12.SetText("O")
				}
				turn++
			}
		}
		btn20.OnTapped = func() {
			fmt.Println("")
			if btn20.Text == "" {
				if turn%2 == 0 {
					btn20.SetText("X")
				} else {
					btn20.SetText("O")
				}
				turn++
			}
		}
		btn21.OnTapped = func() {
			fmt.Println("")
			if btn21.Text == "" {
				if turn%2 == 0 {
					btn21.SetText("X")
				} else {
					btn21.SetText("O")
				}
				turn++
			}
		}
		btn22.OnTapped = func() {
			fmt.Println("")
			if btn22.Text == "" {
				if turn%2 == 0 {
					btn22.SetText("X")
				} else {
					btn22.SetText("O")
				}
				turn++
			}
		}
		myWindow.SetContent(board)
		myWindow.Resize(fyne.NewSize(300, 300))
		myWindow.ShowAndRun()
	})
	menuWindow := myApp.NewWindow("Menu")

	historyButton := widget.NewButton("Lịch sử", func() {

	})

	menuContent := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		startButton,
		historyButton,
	)

	menuWindow.SetContent(menuContent)

	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myApp.Run()
}
