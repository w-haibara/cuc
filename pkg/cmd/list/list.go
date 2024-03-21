package list

import (
	"github.com/spf13/cobra"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/list/list"
)

type ListOptions struct {
	cmdList.ListOptions
}

func NewCmdList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "Work with ClickUp lists",
	}

	cmd.AddCommand(cmdList.NewCmdListList(opts.ListOptions))

	return cmd
}
