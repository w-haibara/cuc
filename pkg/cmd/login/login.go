package login

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/iostreams"
)

type LoginOptions struct {
}

func NewCmdLogin(opts LoginOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Args:  cobra.ExactArgs(0),
		Short: "Authenticate with ClickUp",
		RunE: func(cmd *cobra.Command, args []string) error {
			return loginRun(opts)
		},
	}

	return cmd
}

func loginRun(opts LoginOptions) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintf(iostreams.IO.Out, "Authentication Success\nUser: %s\n", client.User.Username)
	for _, team := range client.Teams {
		fmt.Fprintf(iostreams.IO.Out, "Team: %s", team.Name)
	}

	return nil
}
