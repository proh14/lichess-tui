package main

import (
	// "bytes"
	"encoding/json"
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"net/http"
	"io/ioutil"
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

			headers := map[string]string{
				"Authorization": "Bearer " + login.token,
				"Content-Type":  "application/json",
			}

			// Create the request body
			// body := map[string]string{
			// 	"text": "hello",
			// }

			// Serialize the body to JSON
			// bodyBytes, _ := json.Marshal(body)
			// if err != nil {
			// 	log.Fatalf("Error marshalling body: %v", err)
			// }

			// Make the POST request
			// url := "https://lichess.org/inbox/hoorad123"
			url := "https://lichess.org/api/account"
			// req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
			// req, _ := http.NewRequest("POST", url, nil)
			req, _ := http.NewRequest("GET", url, nil)

			// Set headers
			for key, value := range headers {
				req.Header.Set(key, value)
			}

			// Send the request
			client := &http.Client{}
			resp, _ := client.Do(req)
			defer resp.Body.Close()

			// Read the response body
			respBody, _ := ioutil.ReadAll(resp.Body)

			var respMap map[string]string
			json.Unmarshal(respBody, &respMap)

			value, containsKey := respMap["error"]

			if containsKey {
				return errors.New("No such token found.")
			} else {
				return nil
			}
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
	).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
