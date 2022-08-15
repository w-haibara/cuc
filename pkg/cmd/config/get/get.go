package get

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/internal/config"
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
			return getRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr())
		},
	}

	return cmd
}

func getRun(opts GetOptions, out, errOut io.Writer) error {
	val, err := config.Get(opts.Key)
	if err != nil {
		return err
	}

	if val != "" {
		fmt.Fprintf(out, "%s\n", val)
	}
	return nil
}
