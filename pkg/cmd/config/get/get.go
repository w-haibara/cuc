package get

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/internal/config"
	"github.com/w-haibara/cuc/pkg/iostreams"
)

type GetOptions struct {
	Key string
}

func NewCmdConfigGet(opts GetOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <key>",
		Short: "Print the value of a given configuration key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Key = args[0]
			return getRun(opts)
		},
	}

	return cmd
}

func getRun(opts GetOptions) error {
	val, err := config.Get(opts.Key)
	if err != nil {
		return err
	}

	if val != "" {
		fmt.Fprintf(iostreams.IO.Out, "%s\n", val)
	}
	return nil
}
