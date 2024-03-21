package task

import (
	"github.com/spf13/cobra"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/task/list"
)

type TaskOptions struct {
	cmdList.ListOptions
}

func NewCmdTask(opts TaskOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task <command>",
		Short: "Work with ClickUp tasks",
	}

	cmd.AddCommand(cmdList.NewCmdTaskList(opts.ListOptions))

	return cmd
}
