package login

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/view/jsonview"
)

type LoginOptions struct {
}

func NewCmdLogin(opts LoginOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Args:  cobra.ExactArgs(0),
		Short: "Authenticate with ClickUp",
		RunE: func(cmd *cobra.Command, args []string) error {
			return loginRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), jsonview.JsonFlag(cmd))
		},
	}

	return cmd
}

func loginRun(opts LoginOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	c, err := newClient(ctx, out, errOut)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Authentication Success\nUser: %s\n", c.User.Username)
	for _, team := range c.Teams {
		fmt.Fprintf(out, "Team: %s", team.Name)
	}

	return nil
}

func newClient(ctx context.Context, out, errOut io.Writer) (client.Client, error) {
	c, err := client.NewClient(ctx)
	if err != nil {
		fmt.Fprintln(errOut, "authentication failure:", err.Error())

		if err := client.SetupApiToken(); err != nil {
			return client.Client{}, err
		}

		return newClient(ctx, out, errOut)
	}

	return c, nil
}
