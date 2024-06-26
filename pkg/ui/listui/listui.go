package listui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
	"github.com/w-haibara/cuc/pkg/ui/errui"
	"github.com/w-haibara/cuc/pkg/ui/mdui"
	"github.com/w-haibara/cuc/pkg/ui/message"
)

type Model struct {
	RowItems []any

	List list.Model
	Cmd  func() tea.Msg

	state state

	detail mdui.Model
}

func NewModel(title string, listCmd func() tea.Msg, detailCmd func(data any) tea.Cmd) Model {
	m := Model{
		List:   list.New(nil, list.NewDefaultDelegate(), 1, 1),
		Cmd:    listCmd,
		detail: mdui.NewModel(detailCmd),
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case message.ShowListMsg:
		m.state = defaultState

	case error:
		return errui.NewModel(msg), nil
	}

	switch m.state {
	case defaultState:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				detailEnabled := func() bool {
					if m.List.SelectedItem() == nil {
						return false
					}

					switch m.List.FilterState() {
					case list.Unfiltered:
						return true
					case list.Filtering:
						return false
					case list.FilterApplied:
						return true
					default:
						return true
					}
				}()
				if detailEnabled {
					m.state = showItemDetailState
					return m, func() tea.Msg {
						return message.InitDetailMsg{
							Data: m.RowItems[m.List.Index()],
						}
					}
				}
			}
		}

		m.List, cmd = m.List.Update(msg)
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.List.SetSize(msg.Width, msg.Height)

		case message.InitListMsg:
			m.List.StopSpinner()
			m.List.SetItems(msg.Items)
			m.List.Title = msg.Title
			m.List.Filter = m.filter
			m.RowItems = msg.RowItems
		}

		return m, cmd

	case showItemDetailState:
		m.detail, cmd = m.detail.Update(msg)

		return m, cmd

	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.state {
	case showItemDetailState:
		return m.detail.View()
	}

	return m.List.View()
}

func (m Model) filter(term string, targets []string) []list.Rank {
	return list.DefaultFilter(term, targets)
}

type state int

const (
	defaultState state = iota
	showItemDetailState
)

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
