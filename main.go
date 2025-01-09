package main

import (
	// "bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ProtonMail/gopenpgp/v3/crypto"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"io/ioutil"
	"net/http"
	"os"
)

type login struct {
	token              string
	tokenStoreLocation string
	gpgEncryption      bool
	warningMessage     string
}

func main() {
	login := login{token: "", tokenStoreLocation: "", gpgEncryption: true, warningMessage: ""}

	pgp := crypto.PGP()

	form := huh.NewForm(
		huh.NewGroup(huh.NewInput().Description("Token").Value(&login.token).Placeholder("Token").Suggestions([]string{"lip_"}).EchoMode(huh.EchoModePassword).Validate(func(value string) error {
			if len(value) == 0 {
				return errors.New("The token can't be empty.")
			}

			headers := map[string]string{
				"Authorization": "Bearer " + value,
				"Content-Type":  "application/json",
			}

			// body := map[string]string{
			// 	"text": "hello",
			// }

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

			for key, value := range headers {
				req.Header.Set(key, value)
			}

			client := &http.Client{}
			resp, _ := client.Do(req)
			defer resp.Body.Close()

			respBody, _ := ioutil.ReadAll(resp.Body)

			var respMap map[string]string
			json.Unmarshal(respBody, &respMap)

			value, containsKey := respMap["error"]

			if containsKey {
				return errors.New("The given token is invalid.")
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
				Value(&login.tokenStoreLocation),
			huh.NewConfirm().
				Title("Encrypt with PGP").
				Affirmative("Yes").
				Negative("No").
				Value(&login.gpgEncryption).Validate(func(value bool) error {
				file, _ := os.Create(login.tokenStoreLocation + "/token")

				if value {
					encHandle, _ := pgp.Encryption().Password([]byte("hunter2")).New()
					pgpMessage, _ := encHandle.Encrypt([]byte("my message"))
					armored, _ := pgpMessage.ArmorBytes()

					file.WriteString(string(armored))
				} else {
					file.WriteString(login.token)
				}

				return nil
			}),
			// Validate(),
			huh.NewNote().Description(login.warningMessage),
		),
	).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
