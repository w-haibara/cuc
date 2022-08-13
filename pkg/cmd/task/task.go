package task

import (
	"context"
	"fmt"

	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/iostreams"
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

	for _, task := range tasks {
		fmt.Fprintf(iostreams.IO.Out, "%s, %s, %v\n", task.Name, task.ID, task.Archived)
	}

	return nil
}
