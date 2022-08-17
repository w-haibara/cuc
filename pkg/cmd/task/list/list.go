package list

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/util"
	"github.com/w-haibara/cuc/pkg/view"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type ListOptions struct {
	ListID   string
	Archived bool
}

func NewCmdTaskList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks in a list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.ListID = args[0]
			return taskRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	cmd.Flags().BoolVarP(&opts.Archived, "archived", "a", false, "Filter by archived status")

	return cmd
}

func taskRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
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

	if jsonFlag {
		jsonview.Render(out, tasks)
		return nil
	}

	if len(tasks) == 0 {
		return nil
	}

	view := view.NewListView(fmt.Sprintf("Tasks in [%s]", tasks[0].List.Name), len(tasks))
	for _, task := range tasks {
		customID := ""
		if task.CustomID != "" {
			customID = fmt.Sprintf("[%s] ", task.CustomID)
		}
		title := customID + task.Name

		status := task.Status.Status
		points := fmt.Sprintf("%d points", task.Points)
		desc := strings.Join([]string{status, points}, " ")

		view.AppendItem(title, desc)
	}
	view.Render()

	return nil
}
