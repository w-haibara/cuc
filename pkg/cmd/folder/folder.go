package space

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
)

type FolderOptions struct {
	SpaceID  string
	Archived bool
}

func NewCmdFolder(opts FolderOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "folder",
		Short: "List folders in a space",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.SpaceID = args[0]
			return spaceRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr())
		},
	}

	cmd.Flags().BoolVarP(&opts.Archived, "archived", "a", false, "limit to archived folders")

	return cmd
}

func spaceRun(opts FolderOptions, out, errOut io.Writer) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	folders, _, err := client.Folders.GetFolders(ctx, opts.SpaceID, opts.Archived)
	if err != nil {
		return err
	}

	for _, folder := range folders {
		fmt.Fprintf(out, "%s, %s, %v\n", folder.Name, folder.ID, folder.Archived)
	}

	return nil
}
