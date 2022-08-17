package view

import "github.com/w-haibara/cuc/pkg/view/listview"

type View interface {
	Render()
}

func NewListView(title string, size int) listview.ListView {
	return listview.NewListView(title, size)
}
