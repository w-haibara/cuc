package space

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/iostreams"
)

type SpaceOptions struct {
	TeamID string
}

func NewCmdSpace(opts SpaceOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "space",
		Short: "List spaces in a team",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.TeamID = args[0]
			return spaceRun(opts)
		},
	}

	return cmd
}

func spaceRun(opts SpaceOptions) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	spaces, _, err := client.Spaces.GetSpaces(ctx, opts.TeamID)
	if err != nil {
		return err
	}

	for _, space := range spaces {
		fmt.Fprintf(iostreams.IO.Out, "%s, %s\n", space.Name, space.ID)
	}

	return nil
}
