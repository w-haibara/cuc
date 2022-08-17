package listview

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListView struct {
	Title     string
	ListItems []ListItem
}

func NewListView(title string, size int) ListView {
	return ListView{
		Title:     title,
		ListItems: make([]ListItem, 0, size),
	}
}

func (view *ListView) AppendItem(title, desc string) {
	view.ListItems = append(view.ListItems, ListItem{
		Title: title,
		Desc:  desc,
	})
}

func (view *ListView) Render() error {
	items := make([]list.Item, len(view.ListItems))
	for i, v := range view.ListItems {
		items[i] = item{v.Title, v.Desc}
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	if view.Title != "" {
		m.list.Title = view.Title
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		return err
	}

	return nil
}

type ListItem struct {
	Title string
	Desc  string
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
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
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
