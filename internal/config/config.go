package config

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/proh14/lichess-tui/internal/lichess"
	"github.com/proh14/lichess-tui/internal/security"
	"gopkg.in/yaml.v2"
	"log"
	"os"
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
	var shouldEncrypt bool

	tokenField := huh.NewInput().Description("Token").Placeholder("lip_").EchoMode(huh.EchoModePassword).Validate(lichess.ValidateToken).Value(&token)
	tokenFileField := huh.NewFilePicker().
		Description("Where to store the token?").
		DirAllowed(true).
		FileAllowed(false).
		ShowPermissions(false).
		ShowHidden(true).
		CurrentDirectory("/home").
		Value(&tokenPath)
	shouldEncryptConfirm := huh.NewConfirm().
		Title("Encrypt with PGP").
		Affirmative("Yes").
		Negative("No").
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

		return
	}

	file, _ := os.Create(tokenPath + "/token")
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
