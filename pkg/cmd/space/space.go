package space

import (
	"github.com/spf13/cobra"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/space/list"
)

type SpaceOptions struct {
	cmdList.ListOptions
}

func NewCmdSpace(opts SpaceOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "space <command>",
		Short: "Work with ClickUp spaces",
	}

	cmd.AddCommand(cmdList.NewCmdSpaceList(opts.ListOptions))

	return cmd
}
