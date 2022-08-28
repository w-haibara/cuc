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
	"github.com/w-haibara/cuc/pkg/ui/message"
	"github.com/w-haibara/cuc/pkg/util"
	md "github.com/w-haibara/markdown-builder/markdown"
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

		if err := jsonui.NewModel(tasks).Render(); err != nil {
			return err
		}

		return nil
	}

	ui := listui.NewModel(
		"Tasks in ...",
		listuiCmd(opts),
		detailuiCmd(opts),
	)
	if err := ui.Render(); err != nil {
		return err
	}

	return nil
}

func listuiCmd(opts ListOptions) tea.Cmd {
	return func() tea.Msg {
		tasks, err := getTasks(opts)
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			return fmt.Errorf("there are no tasks")
		}

		rowTasks := make([]any, len(tasks))
		for i, task := range tasks {
			rowTasks[i] = task
		}

		items := listui.MakeItems(len(tasks))
		for _, task := range tasks {
			listui.AppendItem(items, taskTitle(task), taskDesc(task))
		}

		return message.InitListMsg{
			Title:    fmt.Sprintf("Tasks in [%s]", tasks[0].List.Name),
			Items:    *items,
			RowItems: rowTasks,
		}
	}
}

func detailuiCmd(opts ListOptions) func(data any) tea.Cmd {
	return func(data any) tea.Cmd {
		task, ok := data.(clickup.Task)
		if !ok {
			panic("invalid type")
		}

		return func() tea.Msg {
			assignees := func() string {
				names := make([]string, len(task.Assignees))
				for i, assignee := range task.Assignees {
					names[i] = assignee.Username
				}
				return strings.Join(names, ", ")
			}()

			m := md.NewMarkdown()
			m.Write(
				md.MustHeading(taskTitle(task), 1),
				md.Br(),
				md.RootList(
					strings.ToUpper(task.Status.Status),
					fmt.Sprintf("%d points", task.Points),
					"Addignees: "+assignees,
					"Priority: "+task.Priority.Priority,
					"Create At: "+task.DateCreated,
					"Due Date: "+task.DueDate,
					"ID: "+task.ID,
				),
				md.Br(),
				task.TextContent,
			)

			return message.ShowItemDetailMsg{
				MD: m.Render(),
			}
		}
	}
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

func taskTitle(task clickup.Task) string {
	customID := ""
	if task.CustomID != "" {
		customID = fmt.Sprintf("[%s] ", task.CustomID)
	}

	return customID + task.Name
}

func taskDesc(task clickup.Task) string {
	status := fmt.Sprintf("[%s]", strings.ToUpper(task.Status.Status))
	assignees := func() string {
		names := make([]string, len(task.Assignees))
		for i, assignee := range task.Assignees {
			names[i] = assignee.Username
		}
		return "[" + strings.Join(names, " ") + "]"
	}()
	priority := fmt.Sprintf("[%s]", task.Priority.Priority)
	points := fmt.Sprintf("[%d points]", task.Points)
	id := fmt.Sprintf("[ID:%s]", task.ID)

	return util.StringsJoin([]string{status, assignees, priority, points, id}, " ", "[]")
}
