package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func main() {
	var name string
	var store_location string = "~/.local/share"
	var gpg_encryption bool = false
	var file string

	form := huh.NewForm(
		huh.NewGroup(huh.NewInput().Description("Token").Value(&name).Placeholder("Token"), // ,
			huh.NewInput().Description("Where to store?").Value(&store_location).Placeholder("~/.local/share"),
			huh.NewFilePicker().
				Description("Where to store the token?").
				DirAllowed(true).
				FileAllowed(false).
				ShowPermissions(false).
				ShowHidden(true).
				CurrentDirectory("/home").
				Value(&file),
			huh.NewConfirm().
				Title("Encrypt with GPG").
				Affirmative("Yes").
				Negative("No").
				Value(&gpg_encryption),
		)).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
