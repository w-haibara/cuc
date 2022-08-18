package list

import (
	"context"
	"fmt"
	"io"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/raksul/go-clickup/clickup"
	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/ui/jsonui"
	"github.com/w-haibara/cuc/pkg/ui/listui"
	"github.com/w-haibara/cuc/pkg/util"
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
	if jsonFlag {
		tasks, err := getTasks(opts)
		if err != nil {
			return err
		}

		if err := jsonui.NewJsonModel(tasks).Render(); err != nil {
			return err
		}

		return nil
	}

	fn := func() tea.Msg {
		tasks, err := getTasks(opts)
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			return fmt.Errorf("there are no tasks")
		}

		title := fmt.Sprintf("Tasks in [%s]", tasks[0].List.Name)
		msg := listui.NewListMsg(title)
		for _, task := range tasks {
			customID := ""
			if task.CustomID != "" {
				customID = fmt.Sprintf("[%s] ", task.CustomID)
			}

			status := task.Status.Status
			points := fmt.Sprintf("%d points", task.Points)

			msg.AppendItem(
				customID+task.Name,
				strings.Join([]string{status, points}, " "),
			)
		}

		return *msg
	}

	if err := listui.NewListModel("Tasks in ...", fn).Render(); err != nil {
		return err
	}

	return nil
}

func getTasks(opts ListOptions) ([]clickup.Task, error) {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	taskOpts := clickup.GetTasksOptions{
		Archived: opts.Archived,
	}
	tasks, _, err := client.Tasks.GetTasks(ctx, opts.ListID, &taskOpts)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
