package exec

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/extension"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type ExecOptions struct {
	Cmd        cobra.Command
	ExtCmdName string
	Args       []string
}

func NewCmdExtensionExec(opts ExecOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec <extension command>",
		Short: "Execute an installed extension",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.ExtCmdName = args[0]
			opts.Args = args[1:]
			return execRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), jsonview.JsonFlag(cmd))
		},
	}

	return cmd
}

func execRun(opts ExecOptions, out, errOut io.Writer, jsonFlag bool) error {
	ext := extension.NewExtension(&opts.Cmd)
	if err := ext.Setup(); err != nil {
		return err
	}

	if err := ext.Run(opts.ExtCmdName, opts.Args); err != nil {
		return err
	}

	return nil
}
