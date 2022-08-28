package message

import "github.com/charmbracelet/bubbles/list"

type InitListMsg struct {
	Title    string
	Items    []list.Item
	RowItems []any
}

type InitDetailMsg struct {
	Data any
}

type ShowListMsg struct {
}

type ShowItemDetailMsg struct {
	MD string
}
