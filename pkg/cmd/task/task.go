package task

import (
	"context"
	"fmt"

	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/iostreams"
	"github.com/w-haibara/cuc/pkg/view/color"
	"github.com/w-haibara/cuc/pkg/view/listview"
)

type TaskOptions struct {
	ListID   string
	Archived bool
}

func NewCmdTask(opts TaskOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "List tasks in a list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.ListID = args[0]
			return taskRun(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Archived, "archived", "a", false, "Filter by archived status")

	return cmd
}

func taskRun(opts TaskOptions) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	taskOpts := clickup.GetTasksOptions{
		Archived: opts.Archived,
	}
	tasks, _, err := client.Tasks.GetTasks(ctx, opts.ListID, &taskOpts)
	if err != nil {
		return err
	}

	view := listview.New(iostreams.IO)
	view.SetKeys([]listview.Key{
		{
			Text: "ID",
			ColorScheme: listview.ColorScheme{
				Style:   color.Bold,
				FgColor: color.FgHiGreen,
			},
		},
		{
			Text: "CustomID",
			ColorScheme: listview.ColorScheme{
				Style:   color.Bold,
				FgColor: color.FgHiGreen,
			},
		},
		{
			Text: "Name",
			ColorScheme: listview.ColorScheme{
				Style: color.Bold,
			},
		},
		{
			Text: "Points",
			ColorScheme: listview.ColorScheme{
				Style: color.Bold,
			},
		},
	})
	fields := map[string][]string{}
	for _, task := range tasks {
		fields["ID"] = append(fields["ID"], task.ID)
		fields["CustomID"] = append(fields["CustomID"], task.CustomID)
		fields["Name"] = append(fields["Name"], task.Name)
		fields["Points"] = append(fields["Points"], fmt.Sprintf("%d", task.Points))
	}
	view.AddFields(fields)
	view.Render()

	return nil
}
