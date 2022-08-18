package list

import (
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/extension"
	"github.com/w-haibara/cuc/pkg/ui/jsonui"
	"github.com/w-haibara/cuc/pkg/util"
)

type ListOptions struct {
}

func NewCmdExtensionList(opts ListOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List installed extension commands",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	return cmd
}

func listRun(opts ListOptions, out, errOut io.Writer, jsonFlag bool) error {
	scripts, err := extension.List()
	if err != nil {
		return err
	}

	if jsonFlag {
		if err := jsonui.NewJsonModel(scripts).Render(); err != nil {
			return err
		}

		return nil
	}

	for _, info := range scripts {
		fmt.Fprintln(out, info.Name, info.UpdatedAt.Format(time.RFC822), info.Size, info.Path)
	}

	return nil
}
