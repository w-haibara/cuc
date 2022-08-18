package list

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
	TeamID string
}

func NewCmdSpaceList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List spaces in a team",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.TeamID = args[0]
			return spaceRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	return cmd
}

func spaceRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	spaces, _, err := client.Spaces.GetSpaces(ctx, opts.TeamID)
	if err != nil {
		return err
	}

	if jsonFlag {
		if err := jsonui.NewJsonModel(spaces).Render(); err != nil {
			return err
		}

		return nil
	}

	for _, space := range spaces {
		fmt.Fprintf(out, "%s, %s\n", space.Name, space.ID)
	}

	return nil
}
