package main

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/go-zoox/fetch"
)

type login struct {
	token          string
	storeLocation  string
	gpgEncryption  bool
	warningMessage string
}

func main() {
	login := login{token: "", storeLocation: "", gpgEncryption: true, warningMessage: ""}

	form := huh.NewForm(
		huh.NewGroup(huh.NewInput().Description("Token").Value(&login.token).Placeholder("Token").EchoMode(huh.EchoModePassword).Validate(func(value string) error {
			if value == "" {
				return errors.New("The token can't be empty.")
			}

			// fetch.Headers.Set("Authorization", "Bearer " + "lip_EqXU9GUvR4a2bAzFdQM6")
			// fetch.Headers.Set("Content-Type", "application/json")
			
			response, _ := fetch.Get("https://lichess.org/inbox/hoorad123", &fetch.Config{
				Body: map[string]interface{}{
					"text": "hello",
				},
			})

			fmt.Println(response.JSON())

			return nil
		}),
			huh.NewFilePicker().
				Description("Where to store the token?").
				DirAllowed(true).
				FileAllowed(false).
				ShowPermissions(false).
				ShowHidden(true).
				CurrentDirectory("/home").
				Value(&login.storeLocation),
			huh.NewConfirm().
				Title("Encrypt with GPG").
				Affirmative("Yes").
				Negative("No").
				Value(&login.gpgEncryption),
			huh.NewNote().Description(login.warningMessage),
		),
		huh.NewGroup(
			huh.NewNote().Description("Token doesn't exist")),
	).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
