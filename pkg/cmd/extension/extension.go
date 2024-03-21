package extension

import (
	"github.com/spf13/cobra"
	cmdExec "github.com/w-haibara/cuc/pkg/cmd/extension/exec"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/extension/list"
)

type ExtensionOptions struct {
	cmdList.ListOptions
	cmdExec.ExecOptions
}

func NewCmdExtension(opts ExtensionOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extension <command>",
		Short: "Manage cuc extension",
	}

	cmd.AddCommand(cmdList.NewCmdExtensionList(opts.ListOptions))
	cmd.AddCommand(cmdExec.NewCmdExtensionExec(opts.ExecOptions))

	return cmd
}
