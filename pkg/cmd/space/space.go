package space

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
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
			return spaceRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), jsonview.JsonFlag(cmd))
		},
	}

	return cmd
}

func spaceRun(opts SpaceOptions, out, errOut io.Writer, jsonFlag bool) error {
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
		obj := make([]map[string]string, len(spaces))
		for i, space := range spaces {
			obj[i] = map[string]string{
				"ID":                space.ID,
				"Name":              space.Name,
				"Private":           strconv.FormatBool(space.Private),
				"MultipleAssignees": strconv.FormatBool(space.MultipleAssignees),
				"Archived":          strconv.FormatBool(space.Archived),
			}
		}
		jsonview.Render(out, obj)
		return nil
	}

	for _, space := range spaces {
		fmt.Fprintf(out, "%s, %s\n", space.Name, space.ID)
	}

	return nil
}
