package main

import (
	"fmt"
	"os"

  "github.com/kmelow/shpn/task"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list list.Model
}

func New() *model {
  return &model{}
}

func (m *model) initList(w, h int) {
	Tasks := []list.Item{
		task.Task{Tit: "Raspberry Pi’s", Desc: "I have ’em all over my house"},
		task.Task{Tit: "Nutella", Desc: "It's good on toast"},
		task.Task{Tit: "Bitter melon", Desc: "It cools you down"},
		task.Task{Tit: "Nice socks", Desc: "And by that I mean socks without holes"},
		task.Task{Tit: "Terrycloth", Desc: "In other words, towel fabric"},
	}

	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), w, h)
  m.list.Title = "Backlog"
  m.list.SetItems(Tasks)
}


func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func main() {
	m := New()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
