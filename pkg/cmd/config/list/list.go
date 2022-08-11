package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/internal/config"
	"github.com/w-haibara/cuc/pkg/iostreams"
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
			return listRun(opts)
		},
	}

	return cmd
}

func listRun(opts ListOptions) error {
	for _, cfg := range config.Configs() {
		val, err := config.Get(cfg.Key)
		if err != nil {
			return err
		}
		fmt.Fprintf(iostreams.IO.Out, "%s=%s\n", cfg.Key, val)
	}

	return nil
}
