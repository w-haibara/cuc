package config

import (
	"github.com/spf13/cobra"
	cmdGet "github.com/w-haibara/cuc/pkg/cmd/config/get"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/config/list"
	cmdSet "github.com/w-haibara/cuc/pkg/cmd/config/set"
)

type ConfigOptions struct {
	cmdGet.GetOptions
	cmdSet.SetOptions
	cmdList.ListOptions
}

func NewCmdConfig(opts ConfigOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage configuration for cuc",
	}

	cmd.AddCommand(cmdGet.NewCmdConfigGet(opts.GetOptions))
	cmd.AddCommand(cmdSet.NewCmdConfigSet(opts.SetOptions))
	cmd.AddCommand(cmdList.NewCmdConfigList(opts.ListOptions))

	return cmd
}
