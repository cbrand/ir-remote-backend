package cli

import (
	cli "github.com/urfave/cli/v2"

	"github.com/getsentry/sentry-go"
)

// GetApp returns the app for returning the cli app
func GetApp() *cli.App {
	app := &cli.App{
		Name:  "ir-remote-backend",
		Usage: "The ir remote backend service.",
		Before: func(context *cli.Context) error {
			err := sentry.Init(sentry.ClientOptions{})
			return err
		},
		Commands: []*cli.Command{
			ServerCommand(),
			AddDeviceCommand(),
		},
	}
	return app
}
