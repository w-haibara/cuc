package listview

import (
	"bytes"
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/k0kubun/pp"
	"github.com/olekukonko/tablewriter"
	"github.com/w-haibara/cuc/pkg/iostreams"
	"github.com/w-haibara/cuc/pkg/view/color"
)

type ListView struct {
	*tablewriter.Table
	IO      iostreams.IOStreams
	Columns Columns
	buf     *bytes.Buffer
}

type Columns struct {
	Keys   *[]Key
	Fields map[string][]string
}

type Key struct {
	Text        string
	ColorScheme ColorScheme
}

func New(io iostreams.IOStreams) ListView {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	view := ListView{
		Table: table,
		IO:    io,
		Columns: Columns{
			Keys:   &[]Key{},
			Fields: map[string][]string{},
		},
		buf: &buf,
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
		colors = append(colors, key.ColorScheme.C().(tablewriter.Colors))
	}
	view.SetColumnColor(colors...)

	view.setStyle()
	view.Table.Render()

	pp.Println(view.buf.String())

	out, err := glamour.Render(view.buf.String(), "dark")
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprint(view.IO.Out, out)
}

func (view ListView) setStyle() {
	// Markdown style
	view.Table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	view.Table.SetCenterSeparator("|")
}

type ColorScheme struct {
	Style, FgColor, BgColor color.Color
}

func (cs ColorScheme) C() any {
	colors := make(tablewriter.Colors, 0, 3)

	switch cs.Style {
	case color.Normal:
		colors = append(colors, tablewriter.Normal)
	case color.Bold:
		colors = append(colors, tablewriter.Bold)
	case color.Underline:
		colors = append(colors, tablewriter.UnderlineSingle)
	default:
	}

	switch cs.FgColor {
	case color.FgBlack:
		colors = append(colors, tablewriter.FgBlackColor)
	case color.FgRed:
		colors = append(colors, tablewriter.FgRedColor)
	case color.FgGreen:
		colors = append(colors, tablewriter.FgGreenColor)
	case color.FgYellow:
		colors = append(colors, tablewriter.FgYellowColor)
	case color.FgBlue:
		colors = append(colors, tablewriter.FgBlueColor)
	case color.FgMagenta:
		colors = append(colors, tablewriter.FgMagentaColor)
	case color.FgCyan:
		colors = append(colors, tablewriter.FgCyanColor)
	case color.FgWhite:
		colors = append(colors, tablewriter.FgWhiteColor)
	case color.FgHiBlack:
		colors = append(colors, tablewriter.FgHiBlackColor)
	case color.FgHiRed:
		colors = append(colors, tablewriter.FgHiRedColor)
	case color.FgHiGreen:
		colors = append(colors, tablewriter.FgHiGreenColor)
	case color.FgHiYellow:
		colors = append(colors, tablewriter.FgHiYellowColor)
	case color.FgHiBlue:
		colors = append(colors, tablewriter.FgHiBlueColor)
	case color.FgHiMagenta:
		colors = append(colors, tablewriter.FgHiMagentaColor)
	case color.FgHiCyan:
		colors = append(colors, tablewriter.FgHiCyanColor)
	case color.FgHiWhite:
		colors = append(colors, tablewriter.FgHiWhiteColor)
	default:
	}

	switch cs.BgColor {
	case color.BgBlack:
		colors = append(colors, tablewriter.BgBlackColor)
	case color.BgRed:
		colors = append(colors, tablewriter.BgRedColor)
	case color.BgGreen:
		colors = append(colors, tablewriter.BgGreenColor)
	case color.BgYellow:
		colors = append(colors, tablewriter.BgYellowColor)
	case color.BgBlue:
		colors = append(colors, tablewriter.BgBlueColor)
	case color.BgMagenta:
		colors = append(colors, tablewriter.BgMagentaColor)
	case color.BgCyan:
		colors = append(colors, tablewriter.BgCyanColor)
	case color.BgWhite:
		colors = append(colors, tablewriter.BgWhiteColor)
	case color.BgHiBlack:
		colors = append(colors, tablewriter.BgHiBlackColor)
	case color.BgHiRed:
		colors = append(colors, tablewriter.BgHiRedColor)
	case color.BgHiGreen:
		colors = append(colors, tablewriter.BgHiGreenColor)
	case color.BgHiYellow:
		colors = append(colors, tablewriter.BgHiYellowColor)
	case color.BgHiBlue:
		colors = append(colors, tablewriter.BgHiBlueColor)
	case color.BgHiMagenta:
		colors = append(colors, tablewriter.BgHiMagentaColor)
	case color.BgHiCyan:
		colors = append(colors, tablewriter.BgHiCyanColor)
	case color.BgHiWhite:
		colors = append(colors, tablewriter.BgHiWhiteColor)
	default:
	}

	return colors
}
