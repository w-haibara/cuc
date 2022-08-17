package list

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type ListOptions struct {
	SpaceID  string
	Archived bool
}

func NewCmdFolderList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List folders in a space",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.SpaceID = args[0]
			return listRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), jsonview.JsonFlag(cmd))
		},
	}

	cmd.Flags().BoolVarP(&opts.Archived, "archived", "a", false, "limit to archived folders")

	return cmd
}

func listRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	folders, _, err := client.Folders.GetFolders(ctx, opts.SpaceID, opts.Archived)
	if err != nil {
		return err
	}

	if jsonFlag {
		jsonview.Render(out, folders)
		return nil
	}

	for _, folder := range folders {
		fmt.Fprintf(out, "%s, %s, %v\n", folder.Name, folder.ID, folder.Archived)
	}

	return nil
}
