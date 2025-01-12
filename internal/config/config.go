package config

import (
	"errors"
	"github.com/charmbracelet/huh"
	"github.com/proh14/lichess-tui/internal/lichess"
	"github.com/proh14/lichess-tui/internal/security"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Token                  string `yaml:"-"`
	ShouldTokenBeEncrypted bool   `yaml:"shouldTokenBeEncrypted"`
	TokenPath              string `yaml:"tokenPath"`
}

var cfg *Config

func setupConfig(path string) {
	var token string
	var tokenPath string = AddDataDir(RELATIVE_TOKEN_PATH)
	var shouldEncrypt bool = true

	tokenField := huh.NewInput().
		Description("Token").
		Placeholder("lip_").
		EchoMode(huh.EchoModePassword).
		Validate(lichess.ValidateToken).
		Value(&token)

	tokenFileField := huh.NewInput().
		Description("Location of the token persistence file").
		Value(&tokenPath).
		Validate(validateFile)

	shouldEncryptConfirm := huh.NewConfirm().
		Title("Encrypt with PGP").
		Value(&shouldEncrypt)

	group := huh.NewGroup(tokenField, tokenFileField, shouldEncryptConfirm)

	form := huh.NewForm(group)

	err := form.Run()
	if err != nil {
		log.Fatalf("Error running form: %v", err)
	}

	originalToken := token

	if shouldEncrypt {
		var pgpPassword string

		pgpPasswordInput := huh.NewInput().
    Description("PGP Password").
    EchoMode(huh.EchoModePassword).
    Value(&pgpPassword)

		form = huh.NewForm(huh.NewGroup(pgpPasswordInput))

		err = form.Run()

		if err != nil {
			log.Fatalf("Error running form: %v", err)
		}

		token = security.EncryptToken(token, pgpPassword)
	}

	file, _ := os.Create(tokenPath)
	file.WriteString(token)

	defer file.Close()

	cfg.Token = originalToken
	cfg.ShouldTokenBeEncrypted = shouldEncrypt
	cfg.TokenPath = tokenPath

	err = SaveConfig(path)
}

func validateFile(value string) error {
	if len(value) == 0 {
		return errors.New("The filename can't be empty.")
	}

	if string(value[0]) == "~" {
		homeDir := os.Getenv("HOME")
		return errors.New("You may need to replace ~ with " + homeDir + ".")
	}

	_, err := os.Stat(filepath.Dir(value))
	if os.IsNotExist(err) {
		return errors.New("The directory doesn't exist.")
	}

	return nil
}

func SaveConfig(path string) error {
	data, err := yaml.Marshal(cfg)

	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	os.MkdirAll(dir, 0755)

	file, _ := os.Create(path)
	file.WriteString(string(data))

	return err
}

func LoadConfig(path string) error {
	cfg = &Config{}
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
