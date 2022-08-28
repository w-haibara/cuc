package util

import (
	"strings"

	"github.com/spf13/cobra"
)

func JsonFlag(cmd *cobra.Command) bool {
	json, err := cmd.Flags().GetBool("json")
	if err != nil {
		panic(err.Error())
	}

	return json
}

func StringsJoin(elms []string, sep string, empty string) string {
	e := make([]string, 0, len(elms))
	for _, v := range elms {
		if v != empty {
			e = append(e, v)
		}
	}

	return strings.Join(e, sep)
}
