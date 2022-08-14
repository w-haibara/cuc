package listview

import (
	"github.com/olekukonko/tablewriter"
	"github.com/w-haibara/cuc/pkg/iostreams"
)

type ListView struct {
	*tablewriter.Table
	Columns Columns
}

type Columns struct {
	Keys   *[]Key
	Fields map[string][]string
}

type Key struct {
	Text   string
	Colors tablewriter.Colors
}

func New(io iostreams.IOStreams) ListView {
	table := tablewriter.NewWriter(io.Out)
	view := ListView{
		Table: table,
		Columns: Columns{
			Keys:   &[]Key{},
			Fields: map[string][]string{},
		},
	}
	return view
}

func (view ListView) SetKeys(keys []Key) {
	for _, key := range keys {
		*view.Columns.Keys = append(*view.Columns.Keys, key)
	}
}

func (view ListView) AddFields(fields map[string][]string) {
	for key, field := range fields {
		view.Columns.Fields[key] = field
	}
}

func (view ListView) Render() {
	data := [][]string{}
	for i := range view.Columns.Fields["ID"] {
		row := make([]string, 0, len(*view.Columns.Keys))
		for _, key := range *view.Columns.Keys {
			row = append(row, view.Columns.Fields[key.Text][i])
		}
		data = append(data, row)
	}
	view.AppendBulk(data)

	keys := []string{}
	for _, key := range *view.Columns.Keys {
		keys = append(keys, key.Text)
	}
	view.Table.SetHeader(keys)

	colors := []tablewriter.Colors{}
	for _, key := range *view.Columns.Keys {
		colors = append(colors, key.Colors)
	}
	view.SetColumnColor(colors...)

	view.Table.Render()
}
