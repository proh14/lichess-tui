package config

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/proh14/lichess-tui/internal/lichess"
	"github.com/proh14/lichess-tui/internal/security"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"errors"
	"os/user"
)

type Config struct {
	token                  string
	shouldTokenBeEncrypted bool   `yaml:"shouldTokenBeEncrypted"`
	tokenPath              string `yaml:"tokenPath"`
}

var cfg *Config

func setupConfig(path string) {
	cfg = &Config{}

	var token string
	var tokenPath string
	var shouldEncrypt bool = true

	currentUser, _ := user.Current()
	homeDir := currentUser.HomeDir

	tokenField := huh.NewInput().Description("Token").Placeholder("lip_").EchoMode(huh.EchoModePassword).Validate(lichess.ValidateToken).Value(&token)
	tokenFileField := huh.NewInput().Description("Location of the token persistence file").Placeholder("/home/username/lichess-tui-token").Value(&tokenPath).Validate(func(value string) error {
				if len(value) == 0 {
					return errors.New("The filename can't be empty.")
				}

				if string(value[0]) == "~" {
					return errors.New("You may need to replace ~ with " + homeDir + ".")
				}

				_, err := os.Stat(filepath.Dir(value))
				if os.IsNotExist(err) {
					return errors.New("The directory doesn't exist.")
				}

				return nil
			})

	shouldEncryptConfirm := huh.NewConfirm().
		Title("Encrypt with PGP").
		Value(&shouldEncrypt)

	group := huh.NewGroup(tokenField, tokenFileField, shouldEncryptConfirm)

	form := huh.NewForm(group).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		log.Fatalf("Error running form: %v", err)
	}

	originalToken := token

	if shouldEncrypt {
		var pgpPassword string
		pgpPasswordInput := huh.NewInput().Description("PGP Password").EchoMode(huh.EchoModePassword).Value(&pgpPassword)
		form = huh.NewForm(huh.NewGroup(pgpPasswordInput)).WithProgramOptions(tea.WithAltScreen())
		err = form.Run()

		if err != nil {
			log.Fatalf("Error running form: %v", err)
		}

		token = security.EncryptToken(token, pgpPassword)
	}

	file, err := os.Create(tokenPath)
	file.WriteString(token)

	defer file.Close()

	cfg.token = originalToken
	cfg.shouldTokenBeEncrypted = shouldEncrypt
	cfg.tokenPath = tokenPath

	err = SaveConfig(path)
}

func SaveConfig(path string) error {
	data, err := yaml.Marshal(cfg)

	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	os.MkdirAll(dir, 0755)

	err = os.WriteFile(path, data, 0644)
	return err
}

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			setupConfig(path)
			return nil
		}
	}

	err = yaml.Unmarshal(data, cfg)

	return err
}
