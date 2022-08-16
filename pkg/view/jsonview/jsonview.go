package jsonview

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
)

func JsonFlag(cmd *cobra.Command) bool {
	json, err := cmd.Flags().GetBool("json")
	if err != nil {
		panic(err.Error())
	}

	return json
}

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
