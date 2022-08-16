package team

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type TeamOptions struct {
}

func NewCmdTeam(opts TeamOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team",
		Args:  cobra.ExactArgs(0),
		Short: "List teams in a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return teamRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), jsonview.JsonFlag(cmd))
		},
	}

	return cmd
}

func teamRun(opts TeamOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	if jsonFlag {
		obj := make([]map[string]string, len(client.Teams))
		for i, team := range client.Teams {
			obj[i] = map[string]string{
				"ID":   team.ID,
				"Name": team.Name,
			}
		}

		if err := jsonview.Render(out, obj); err != nil {
			return err
		}

		return nil
	}

	for _, team := range client.Teams {
		fmt.Fprintf(out, "%s, %s\n", team.Name, team.ID)
	}

	return nil
}
