package cli

import (
	"context"
	"strings"

	clientpkg "github.com/hashicorp/waypoint/internal/client"
	"github.com/hashicorp/waypoint/internal/pkg/flag"
	pb "github.com/hashicorp/waypoint/internal/server/gen"
	"github.com/hashicorp/waypoint/sdk/terminal"
	"github.com/posener/complete"
)

type HostnameRegisterCommand struct {
	*baseCommand
}

func (c *HostnameRegisterCommand) Run(args []string) int {
	// Initialize. If we fail, we just exit since Init handles the UI.
	if err := c.Init(
		WithArgs(args),
		WithFlags(c.Flags()),
		WithSingleApp(),
	); err != nil {
		return 1
	}

	hostname := ""
	if len(c.args) > 0 {
		hostname = c.args[0]
	}

	client := c.project.Client()
	err := c.DoApp(c.Ctx, func(ctx context.Context, app *clientpkg.App) error {
		resp, err := client.CreateHostname(ctx, &pb.CreateHostnameRequest{
			Hostname: hostname,
			Target: &pb.Hostname_Target{
				Target: &pb.Hostname_Target_Application{
					Application: &pb.Hostname_TargetApp{
						Application: app.Ref(),
						Workspace:   c.project.WorkspaceRef(),
					},
				},
			},
		})
		if err != nil {
			app.UI.Output(err.Error(), terminal.WithErrorStyle())
			return ErrSentinel
		}

		c.ui.Output(resp.Hostname.Fqdn, terminal.WithSuccessStyle())
		return nil
	})
	if err != nil {
		return 1
	}

	return 0
}

func (c *HostnameRegisterCommand) Flags() *flag.Sets {
	return c.flagSet(0, nil)
}

func (c *HostnameRegisterCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *HostnameRegisterCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *HostnameRegisterCommand) Synopsis() string {
	return "Register a hostname to route to your apps."
}

func (c *HostnameRegisterCommand) Help() string {
	helpText := `
Usage: waypoint hostname register [hostname]

  Register a hostname with the URL service to route to your apps.

  The URL service must be enabled and configured with the Waypoint server.
  This will output the fully qualified domain name that should begin
  routing immediately.

`

	return strings.TrimSpace(helpText)
}