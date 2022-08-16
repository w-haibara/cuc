package list

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/internal/config"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type ListOptions struct {
}

func NewCmdConfigList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Print a list of configuration keys and values",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), jsonview.JsonFlag(cmd))
		},
	}

	return cmd
}

func listRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
	for _, cfg := range config.Configs() {
		val, err := config.Get(cfg.Key)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "%s=%s\n", cfg.Key, val)
	}

	return nil
}
