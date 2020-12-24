package cli

import cli "github.com/urfave/cli/v2"

// ListeningServerFlags returns the flags for listening ports
func ListeningServerFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "host",
			Usage:    "The host which the GRPC server should listen to",
			EnvVars:  []string{"LISTENING_SERVER_HOST"},
			Value:    "0.0.0.0",
			Required: false,
		},
		&cli.IntFlag{
			Name:     "port",
			Usage:    "The port the GRPC server should listen to",
			EnvVars:  []string{"LISTENING_SERVER_PORT"},
			Value:    9111,
			Required: false,
		},
	}
}
