package message

import "github.com/charmbracelet/bubbles/list"

type InitListMsg struct {
	Title       string
	Items       []list.Item
	ItemDetails ItemDetails
}

type ItemDetails struct {
	Keys []string
	Data *[]map[string]string
}

type ShowListMsg struct {
}

type ShowItemDetailMsg struct {
}
