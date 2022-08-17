package jsonview

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/MakeNowJust/heredoc/v2"
)

func Render(out io.Writer, obj any) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Fprintf(out, heredoc.Docf(`
		{
			"error": %s
		}
		`), err.Error())
	}

	fmt.Fprintln(out, string(b))
}
