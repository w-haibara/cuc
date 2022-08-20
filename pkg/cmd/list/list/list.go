package list

import (
	"context"
	"fmt"
	"io"

	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/ui/jsonui"
	"github.com/w-haibara/cuc/pkg/util"
)

type ListOptions struct {
	FolderID   string
	Archived   bool
	Folderless bool
}

func NewCmdListList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List lists in a folder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.FolderID = args[0]
			return listRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	cmd.Flags().BoolVarP(&opts.Archived, "archived", "a", false, "Filter by archived status")
	cmd.Flags().BoolVarP(&opts.Folderless, "folderless", "f", false, "Filter by folderless list")

	return cmd
}

func listRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	var lists []clickup.List
	if opts.Folderless {
		lists, _, err = client.Lists.GetFolderlessLists(ctx, opts.FolderID, opts.Archived)
		if err != nil {
			return err
		}
	} else {
		lists, _, err = client.Lists.GetLists(ctx, opts.FolderID, opts.Archived)
		if err != nil {
			return err
		}
	}

	if jsonFlag {
		if err := jsonui.NewModel(lists).Render(); err != nil {
			return err
		}

		return nil
	}

	for _, list := range lists {
		fmt.Fprintf(out, "%s, %s, %v\n", list.Name, list.ID, list.Archived)
	}

	return nil
}
