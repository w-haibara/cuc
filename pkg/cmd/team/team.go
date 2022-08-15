package team

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
)

type TeamOptions struct {
}

func NewCmdTeam(opts TeamOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team",
		Args:  cobra.ExactArgs(0),
		Short: "List teams in a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return teamRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr())
		},
	}

	return cmd
}

func teamRun(opts TeamOptions, out, errOut io.Writer) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	for _, team := range client.Teams {
		fmt.Fprintf(out, "%s, %s\n", team.Name, team.ID)
	}

	return nil
}
