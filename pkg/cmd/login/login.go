package login

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/client"
	"github.com/w-haibara/cuc/pkg/ui/jsonui"
	"github.com/w-haibara/cuc/pkg/util"
)

type LoginOptions struct {
}

func NewCmdLogin(opts LoginOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Args:  cobra.ExactArgs(0),
		Short: "Authenticate with ClickUp",
		RunE: func(cmd *cobra.Command, args []string) error {
			return loginRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	return cmd
}

func loginRun(opts LoginOptions, out, errOut io.Writer, jsonFlag bool) error {
	ctx := context.Background()
	c, err := newClient(ctx, out, errOut, jsonFlag)
	if err != nil {
		if jsonFlag {
			obj := map[string]string{
				"status": "ng",
				"error":  err.Error(),
			}
			if err := jsonui.NewModel(obj).Render(); err != nil {
				return err
			}
			return nil
		}

		return err
	}

	if jsonFlag {
		obj := map[string]string{"status": "ok"}
		if err := jsonui.NewModel(obj).Render(); err != nil {
			return err
		}

		return nil
	}

	fmt.Fprintf(out, "Authentication Success\nUser: %s\n", c.User.Username)
	for _, team := range c.Teams {
		fmt.Fprintf(out, "Team: %s", team.Name)
	}

	return nil
}

func newClient(ctx context.Context, out, errOut io.Writer, jsonFlag bool) (client.Client, error) {
	c, err := client.NewClient(ctx)
	if err != nil {
		if jsonFlag {
			return client.Client{}, err
		}

		fmt.Fprintln(errOut, "authentication failure:", err.Error())

		if err := client.SetupApiToken(); err != nil {
			return client.Client{}, err
		}

		return newClient(ctx, out, errOut, false)
	}

	return c, nil
}
