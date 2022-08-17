package root

import (
	"github.com/spf13/cobra"
	cmdConfig "github.com/w-haibara/cuc/pkg/cmd/config"
	cmdExtension "github.com/w-haibara/cuc/pkg/cmd/extension"
	cmdExtensionExec "github.com/w-haibara/cuc/pkg/cmd/extension/exec"
	cmdFolder "github.com/w-haibara/cuc/pkg/cmd/folder"
	cmdList "github.com/w-haibara/cuc/pkg/cmd/list"
	cmdLogin "github.com/w-haibara/cuc/pkg/cmd/login"
	cmdSpace "github.com/w-haibara/cuc/pkg/cmd/space"
	cmdTask "github.com/w-haibara/cuc/pkg/cmd/task"
	cmdTeam "github.com/w-haibara/cuc/pkg/cmd/team"
	"github.com/w-haibara/cuc/pkg/config"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type Command struct {
	*cobra.Command
}

func NewCmdRoot() *Command {
	return &Command{newCmdRoot()}
}

func (cmd *Command) ExecuteC() (*Command, error) {
	c, err := cmd.Command.ExecuteC()
	if err != nil {
		if jsonview.JsonFlag(cmd.Command) {
			jsonview.Render(cmd.OutOrStderr(), map[string]string{
				"error": err.Error(),
			})
			return cmd, nil
		}

		return cmd, err
	}

	return &Command{c}, nil
}

func newCmdRoot() *cobra.Command {
	config.Init()

	cmd := &cobra.Command{
		Use:   "cuc <command> <subcommand> [flags]",
		Short: "ClickUp CLI",

		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().Bool("version", false, "Show cuc version")
	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.PersistentFlags().Bool("json", false, "Output JSON")

	// Child commands
	cmd.AddCommand(cmdConfig.NewCmdConfig(
		cmdConfig.ConfigOptions{},
	))

	cmd.AddCommand(cmdLogin.NewCmdLogin(
		cmdLogin.LoginOptions{},
	))

	cmd.AddCommand(cmdTeam.NewCmdTeam(
		cmdTeam.TeamOptions{},
	))

	cmd.AddCommand(cmdSpace.NewCmdSpace(
		cmdSpace.SpaceOptions{},
	))

	cmd.AddCommand(cmdFolder.NewCmdFolder(
		cmdFolder.FolderOptions{},
	))

	cmd.AddCommand(cmdList.NewCmdList(
		cmdList.ListOptions{},
	))

	cmd.AddCommand(cmdTask.NewCmdTask(
		cmdTask.TaskOptions{},
	))

	cmd.AddCommand(cmdExtension.NewCmdExtension(
		cmdExtension.ExtensionOptions{
			ExecOptions: cmdExtensionExec.ExecOptions{
				Cmd: *cmd,
			},
		},
	))

	return cmd
}
