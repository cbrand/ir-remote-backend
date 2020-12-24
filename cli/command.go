package cli

import (
	"fmt"
	"net"

	cli "github.com/urfave/cli/v2"

	"github.com/cbrand/ir-remote-backend/handler"
	mqttBackend "github.com/cbrand/ir-remote-backend/mqtt"
	"github.com/cbrand/ir-remote-backend/protocol"
	"google.golang.org/grpc"
)

func remoteHandlerFor(context *cli.Context) (*handler.Server, error) {
	clientOptions, err := MQTTConfigFromContext(context)
	if err != nil {
		return nil, err
	}

	redisClient := redisBackendFor(context)
	mqttClient := mqttBackend.NewHandlerFromOptions(clientOptions)
	remoteHandler := handler.NewServer(redisClient, mqttClient)
	return remoteHandler, nil
}

// ServerCommand returns the command for running the server.
func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:     "server",
		Aliases:  []string{"s"},
		Category: "server",
		Usage:    "Runs the GRPC server to connect the RPC devices.",
		Flags:    ConcatSlices(ListeningServerFlags(), MQTTFlags(), RedisFlags()),
		Action: func(context *cli.Context) error {
			listening, err := net.Listen("tcp", fmt.Sprintf("%s:%d", context.String("host"), context.Int("port")))
			if err != nil {
				return err
			}

			grpcServer := grpc.NewServer()
			remoteHandler, err := remoteHandlerFor(context)
			if err != nil {
				return err
			}
			protocol.RegisterRemoteServiceServer(grpcServer, remoteHandler)

			err = grpcServer.Serve(listening)
			if err != nil {
				return err
			}
			fmt.Println("Server started successfully")

			return nil
		},
	}
}

// AddDeviceCommand exposes the command to add a device to be managed by this backend
func AddDeviceCommand() *cli.Command {
	return &cli.Command{
		Name:    "add-remote",
		Aliases: []string{"ar"},
		Flags: ConcatSlices(RedisFlags(), MQTTFlags(), []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "The name of the remote which should be added",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "mqtt-prefix",
				Usage:    "The prefix of the remote for publishing mqtt commands",
				Required: true,
			},
		}),
		Category: "setup",
		Usage:    "Create a remote and store its value to the backend",
		Action: func(context *cli.Context) error {
			remoteHandler, err := remoteHandlerFor(context)
			if err != nil {
				return err
			}
			remote, err := remoteHandler.AddRemote(context.String("name"), context.String("mqtt-prefix"))
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprintf("New remote added (UUID %s, Name %s, Prefix %s)", remote.GetId(), remote.GetName(), remote.GetMqttTopicPrefix()))
			return nil
		},
	}
}
