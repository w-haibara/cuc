package list

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/iostreams"
)

type ListOptions struct {
	FolderID string
	Archived bool
}

func NewCmdList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List lists in a folder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.FolderID = args[0]
			return listRun(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Archived, "archived", "a", false, "Filter by archived status")

	return cmd
}

func listRun(opts ListOptions) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	lists, _, err := client.Lists.GetLists(ctx, opts.FolderID, opts.Archived)
	if err != nil {
		return err
	}

	for _, list := range lists {
		fmt.Fprintf(iostreams.IO.Out, "%s, %s, %v\n", list.Name, list.ID, list.Archived)
	}

	return nil
}
