package jsonview

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func JsonFlag(cmd *cobra.Command) bool {
	json, err := cmd.Flags().GetBool("json")
	if err != nil {
		panic(err.Error())
	}

	return json
}

func Render(out io.Writer, obj any) error {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	fmt.Fprintln(out, string(b))

	return nil
}
