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

type tokenModel struct {
	token              string
	tokenStoreLocation string
}

type pgpEncryptionModel struct {
	pgpEncryption    bool
	pgpPassword      string
	pgpPasswordInput *huh.Input
}

// func (m pgpEncryptionModel) Init() tea.Cmd {
// 	return nil
// }
//
// func (m pgpEncryptionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	m.pgpPasswordInput.WithAccessible(!m.pgpEncryption)
//
// 	return m, nil
// }
//
// func (m pgpEncryptionModel) View() string {
// 	return ""
// }

func main() {
	tokenModel := tokenModel{token: "", tokenStoreLocation: ""}
	pgpEncryptionModel := pgpEncryptionModel{pgpEncryption: true, pgpPassword: "example"}
	pgpEncryptionModel.pgpPasswordInput = huh.NewInput().Description("Password").Value(&pgpEncryptionModel.pgpPassword)

	pgp := crypto.PGP()

	form := huh.NewForm(
		huh.NewGroup(huh.NewInput().Description("Token").Value(&tokenModel.token).Placeholder("lip_").EchoMode(huh.EchoModePassword).Validate(func(value string) error {
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
				Value(&tokenModel.tokenStoreLocation),
			huh.NewConfirm().
				Title("Encrypt with PGP").
				Affirmative("Yes").
				Negative("No").
				Value(&pgpEncryptionModel.pgpEncryption).Validate(func(value bool) error {
				file, _ := os.Create(tokenModel.tokenStoreLocation + "/token")
				//
				// pgpEncryptionModel.pgpPasswordInput.EchoMode(huh.EchoModeNone)

				if value {
					encHandle, _ := pgp.Encryption().Password([]byte(pgpEncryptionModel.pgpPassword)).New()
					pgpMessage, _ := encHandle.Encrypt([]byte(tokenModel.token))
					armored, _ := pgpMessage.ArmorBytes()

					file.WriteString(string(armored))
				} else {
					file.WriteString(tokenModel.token)
				}

				return nil
			}),
			pgpEncryptionModel.pgpPasswordInput,
		),
	).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
