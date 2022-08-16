package jsonview

import (
	"encoding/json"
	"fmt"
	"io"
)

func Render(out io.Writer, obj any) error {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	fmt.Fprintln(out, string(b))

	return nil
}
