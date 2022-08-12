package root

import (
	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/internal/config"
	cmdConfig "github.com/w-haibara/cuc/pkg/cmd/config"
	cmdLogin "github.com/w-haibara/cuc/pkg/cmd/login"
	cmdTeam "github.com/w-haibara/cuc/pkg/cmd/team"
)

func NewCmdRoot() *cobra.Command {
	config.Init()

	cmd := &cobra.Command{
		Use:   "cuc <command> <subcommand> [flags]",
		Short: "ClickUp CLI",

		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().Bool("version", false, "Show cuc version")
	cmd.PersistentFlags().Bool("help", false, "Show help for command")

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

	return cmd
}
