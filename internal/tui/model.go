package tui

import (
	"fmt"
	"strings"

	"d1/internal/api"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	textarea textarea.Model
	apiClient *api.Client
	err      error
	quitting bool
}

func NewModel(client *api.Client) Model {
	ta := textarea.New()
	ta.Placeholder = "Write your journal entry here..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 0 // Unlimited
	ta.ShowLineNumbers = false

	return Model{
		textarea:  ta,
		apiClient: client,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
			return m, tea.Quit
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlS:
			// Save logic
			text := m.textarea.Value()
			if text != "" {
				// Perform save (this should strictly be a command if it takes time)
				// For simplicity in this foundational model, we'll verify it's not empty 
				// and then quit or clear. Ideally, return a Cmd that calls the API.
				return m, SaveEntry(m.apiClient, text)
			}
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case tea.WindowSizeMsg:
		m.textarea.SetWidth(msg.Width)
		m.textarea.SetHeight(msg.Height - 2) // Reserve space for status bar

	case saveSuccessMsg:
		// Handle successful save
		return m, tea.Quit

	case errorMsg:
		m.err = msg.err
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	
	// Status bar styling
	statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(1)
	
	wordCount := len(strings.Fields(m.textarea.Value()))
	status := fmt.Sprintf("Words: %d | Ctrl+S to Send | Esc to Quit", wordCount)

	if m.err != nil {
		status = fmt.Sprintf("Error: %v | Ctrl+S to Retry | Esc to Quit", m.err)
		statusStyle = statusStyle.Foreground(lipgloss.Color("196")) // Red
	}

	return fmt.Sprintf(
		"%s\n%s",
		m.textarea.View(),
		statusStyle.Render(status),
	)
}

// Msg types for handling async save
type saveSuccessMsg struct{}
type errorMsg struct{ err error }

func SaveEntry(client *api.Client, text string) tea.Cmd {
	return func() tea.Msg {
		err := client.CreateEntry(text, []string{"tui-post"})
		if err != nil {
			return errorMsg{err}
		}
		return saveSuccessMsg{}
	}
}
