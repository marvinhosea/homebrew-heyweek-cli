package models

import (
	"context"
	"errors"
	"fmt"
	"heyweek/cli/pkg/api"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	title     string
	longInput string
	input     textinput.Model
	err       error
}

type errMsg error

var _ tea.Model = (*State)(nil)

func (m State) Init() tea.Cmd {
	return textinput.Blink
}

func (m State) View() string {
	if m.title != "" {
		return fmt.Sprintf("The title is %s", m.title)
	}

	return "Add"
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyCtrlBackslash:
			ctx, cancel := context.WithTimeoutCause(context.Background(), 40*time.Second, errors.New("timedout"))
			defer cancel()
			cl := api.NewClient()
			cl.Request(ctx, "POST")
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, tea.Quit
	}

	return m, nil
}

func NewState() (*State, error) {
	ti := textinput.New()
	ti.CharLimit = 30
	ti.Focus()
	ti.Placeholder = "Give me a title"

	return &State{input: ti, err: nil}, nil
}
