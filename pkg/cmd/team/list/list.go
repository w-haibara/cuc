package team

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/ui/jsonui"
	"github.com/w-haibara/cuc/pkg/util"
)

type ListOptions struct {
}

func NewCmdTeamList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.ExactArgs(0),
		Short: "List teams in a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return teamRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	return cmd
}

func teamRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	if jsonFlag {
		if err := jsonui.NewModel(client.Teams).Render(); err != nil {
			return err
		}

		return nil
	}

	for _, team := range client.Teams {
		fmt.Fprintf(out, "%s, %s\n", team.Name, team.ID)
	}

	return nil
}
