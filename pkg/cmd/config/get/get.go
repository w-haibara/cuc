package get

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/config"
	"github.com/w-haibara/cuc/pkg/ui/jsonui"
	"github.com/w-haibara/cuc/pkg/util"
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
			return getRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	return cmd
}

func getRun(opts GetOptions, out, errOut io.Writer, jsonFlag bool) error {
	val, err := config.Get(opts.Key)
	if err != nil {
		return err
	}

	if jsonFlag {
		obj := map[string]string{
			opts.Key: val,
		}
		if err := jsonui.NewJsonModel(obj).Render(); err != nil {
			return err
		}

		return nil
	}

	if val != "" {
		fmt.Fprintf(out, "%s\n", val)
	}
	return nil
}
