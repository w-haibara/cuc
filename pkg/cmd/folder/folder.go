package folder

import (
	"github.com/spf13/cobra"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/folder/list"
)

type FolderOptions struct {
	cmdList.ListOptions
}

func NewCmdFolder(opts FolderOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "folder <command>",
		Short: "Work with ClickUp folders",
	}

	cmd.AddCommand(cmdList.NewCmdFolderList(opts.ListOptions))

	return cmd
}
