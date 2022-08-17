package team

import (
	"github.com/spf13/cobra"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/team/list"
)

type TeamOptions struct {
	cmdList.ListOptions
}

func NewCmdTeam(opts TeamOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team <command>",
		Short: "Work with ClickUp teams",
	}

	cmd.AddCommand(cmdList.NewCmdTeamList(opts.ListOptions))

	return cmd
}
