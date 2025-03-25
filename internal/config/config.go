package config

import (
	"errors"
	"lichess-tui/internal/lichess"
	"lichess-tui/internal/security"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Token                  string `yaml:"-"`
	ShouldTokenBeEncrypted bool   `yaml:"shouldTokenBeEncrypted"`
	TokenPath              string `yaml:"tokenPath"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}

func setupConfig(path string) {
	var token string
	var tokenPath string = AddDataDir(RELATIVE_TOKEN_PATH)
	var shouldEncrypt bool = true

	tokenField := huh.NewInput().
		Description("Token").
		Placeholder("lip_").
		EchoMode(huh.EchoModePassword).
		Validate(func(value string) error {
			return lichess.ValidateToken(value, GetConfigPath())
		}).
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
	os.MkdirAll(dir, 0o755)

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

	loadToken()

	err = lichess.ValidateToken(cfg.Token, GetConfigPath())
	if err != nil {
		log.Fatalf("Invalid token: %v", err)
	}

	return err
}

func loadToken() {
	token, _ := os.ReadFile(cfg.TokenPath)
	if !cfg.ShouldTokenBeEncrypted {
		cfg.Token = string(token)
		return
	}

	pgpPassword := ""

	pgpPasswordInput := huh.NewInput().
		Description("PGP Password").
		EchoMode(huh.EchoModePassword).
		Value(&pgpPassword).
		Validate(func(password string) error {
			_, err := security.DecryptToken(string(token), password)
			if err != nil {
				return errors.New("Incorrect password.")
			}
			return nil
		})

	form := huh.NewForm(huh.NewGroup(pgpPasswordInput))

	err := form.Run()
	if err != nil {
		log.Fatalf("Error running form: %v", err)
	}

	cfg.Token, _ = security.DecryptToken(string(token), pgpPassword)
}
