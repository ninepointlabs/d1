package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"d1/internal/api"
	"d1/internal/config"
	"d1/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Parse CLI arguments
	// If arguments are provided after flags, treat as quick post.
	flag.Parse()
	args := flag.Args()

	// Handle 'config' command
	if len(args) > 0 && args[0] == "config" {
		if err := config.ConfigurePrompt(); err != nil {
			log.Fatalf("Configuration failed: %v", err)
		}
		return
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize API client (now using Email)
	// You need to set D1_SENDER_EMAIL, D1_SENDER_PASSWORD, and D1_TARGET_EMAIL
	// in your environment for this to work.
	client := api.New(cfg.SenderEmail, cfg.SenderPassword, cfg.TargetEmail)

	if len(args) > 0 {
		// --- Quick Post Mode ---
		text := strings.Join(args, " ")
		fmt.Println("Posting quick entry...")
		if err := client.CreateEntry(text, []string{"quick-post"}); err != nil {
			log.Fatalf("Error creating entry: %v", err)
		}
	} else {
		// --- Compose Mode (TUI) ---
		p := tea.NewProgram(tui.NewModel(client), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error processing TUI: %v", err)
			os.Exit(1)
		}
	}
}
