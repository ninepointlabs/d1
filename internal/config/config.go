package config

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/zalando/go-keyring"
	"golang.org/x/term"
)

const (
	KeyringService = "d1-journal"
	KeySenderEmail = "sender-email"
	KeySenderPass  = "sender-pass"
	KeyTargetEmail = "target-email"
)

type Config struct {
	SenderEmail    string
	SenderPassword string
	TargetEmail    string
}

func Load() (*Config, error) {
	// 1. Try Environment Variables first (override)
	cfg := &Config{
		SenderEmail:    os.Getenv("D1_SENDER_EMAIL"),
		SenderPassword: os.Getenv("D1_SENDER_PASSWORD"),
		TargetEmail:    os.Getenv("D1_TARGET_EMAIL"),
	}

	// 2. Fallback to Keyring
	if cfg.SenderEmail == "" {
		if val, err := keyring.Get(KeyringService, KeySenderEmail); err == nil {
			cfg.SenderEmail = val
		}
	}
	if cfg.SenderPassword == "" {
		if val, err := keyring.Get(KeyringService, KeySenderPass); err == nil {
			cfg.SenderPassword = val
		}
	}
	if cfg.TargetEmail == "" {
		if val, err := keyring.Get(KeyringService, KeyTargetEmail); err == nil {
			cfg.TargetEmail = val
		}
	}

	// 3. Validation
	if cfg.SenderEmail == "" || cfg.SenderPassword == "" || cfg.TargetEmail == "" {
		return nil, errors.New("credentials missing: please run 'd1 config' to set them, or set D1_* env vars")
	}

	return cfg, nil
}

// ConfigurePrompt interactively asks for credentials and saves them to the keyring
func ConfigurePrompt() error {
	fmt.Println("Interactive Setup (credentials will be stored in system keyring)")
	fmt.Println("---------------------------------------------------------------")

	var sender, pass, target string

	fmt.Print("Your Gmail Address (Sender): ")
	fmt.Scanln(&sender)

	fmt.Print("Your Gmail App Password: ")
	bytePass, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	pass = string(bytePass)
	fmt.Println() // Newline after password input

	fmt.Print("Day One Journal Email (Target): ")
	fmt.Scanln(&target)

	if sender == "" || pass == "" || target == "" {
		return errors.New("all fields are required")
	}

	// Save to keyring
	if err := keyring.Set(KeyringService, KeySenderEmail, sender); err != nil {
		return fmt.Errorf("failed to save email: %w", err)
	}
	if err := keyring.Set(KeyringService, KeySenderPass, pass); err != nil {
		return fmt.Errorf("failed to save password: %w", err)
	}
	if err := keyring.Set(KeyringService, KeyTargetEmail, target); err != nil {
		return fmt.Errorf("failed to save target email: %w", err)
	}

	fmt.Println("Credentials saved successfully!")
	return nil
}
