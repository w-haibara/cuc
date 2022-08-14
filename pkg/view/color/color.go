package color

type Color int

const (
	Normal Color = iota
	Bold
	Underline
)

const (
	FgBlack Color = iota + 10
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite

	FgHiBlack Color = iota + 20
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

const (
	BgBlack Color = iota + 30
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite

	BgHiBlack Color = iota + 40
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)
