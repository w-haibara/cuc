package util

import "github.com/spf13/cobra"

func JsonFlag(cmd *cobra.Command) bool {
	json, err := cmd.Flags().GetBool("json")
	if err != nil {
		panic(err.Error())
	}

	return json
}
