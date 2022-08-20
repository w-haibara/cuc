package listui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
	"github.com/w-haibara/cuc/pkg/ui/errui"
)

type Model struct {
	List list.Model
	Cmd  func() tea.Msg
}

func NewModel(title string, cmd func() tea.Msg) Model {
	m := Model{
		List: list.New(nil, list.NewDefaultDelegate(), 1, 1),
		Cmd:  cmd,
	}
	m.List.Title = title
	m.List.StartSpinner()
	return m
}

func (m Model) Render() error {
	return ui.Render(m)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Cmd,
		m.List.StartSpinner(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetSize(msg.Width, msg.Height)

	case Msg:
		m.List.StopSpinner()
		m.List.SetItems(msg.Items)
		m.List.Title = msg.Title
		m.List.Filter = m.filter

	case tea.KeyMsg:
		tea.Println("-->", msg.String())

	case error:
		return errui.NewModel(msg), cmd
	}

	return m, cmd
}

func (m Model) View() string {
	return m.List.View()
}

func (m Model) filter(term string, targets []string) []list.Rank {
	return list.DefaultFilter(term, targets)
}

type Msg struct {
	Title string
	Items []list.Item
}

func NewMsg(title string, items []list.Item) Msg {
	return Msg{
		Title: title,
		Items: items,
	}
}

type Item struct {
	title string
	desc  string
}

func MakeItems(size int) *[]list.Item {
	items := make([]list.Item, 0, size)
	return &items
}

func AppendItem(items *[]list.Item, title, desc string) {
	*items = append(*items, Item{
		title: title,
		desc:  desc,
	})
}

func (item Item) Title() string       { return item.title }
func (item Item) Description() string { return item.desc }
func (item Item) FilterValue() string { return item.title }
