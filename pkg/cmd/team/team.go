package team

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
)

type TeamOptions struct {
	json bool
}

func NewCmdTeam(opts TeamOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team",
		Args:  cobra.ExactArgs(0),
		Short: "List teams in a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.json = true
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

	if opts.json {
		type _obj struct {
			ID   string
			Name string
		}
		obj := make([]_obj, len(client.Teams))
		for i, team := range client.Teams {
			obj[i].ID = team.ID
			obj[i].Name = team.Name
		}
		b, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(b))
		return nil
	}

	for _, team := range client.Teams {
		fmt.Fprintf(out, "%s, %s\n", team.Name, team.ID)
	}

	return nil
}
