package cli

import (
	"fmt"

	"github.com/cbrand/ir-remote-backend/backend"
	redisBackend "github.com/cbrand/ir-remote-backend/backend/redis"
	"github.com/go-redis/redis"
	cli "github.com/urfave/cli/v2"
)

// RedisFlags returns the redis flag configurations.
func RedisFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "redis-host",
			EnvVars:  []string{"REDIS_SERVER_HOST"},
			Usage:    "The host to connect to the redis server",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "redis-port",
			EnvVars:  []string{"REDIS_SERVER_PORT"},
			Usage:    "The port to connect to the redis server",
			Required: false,
			Value:    6379,
		},
		&cli.StringFlag{
			Name:     "redis-password",
			EnvVars:  []string{"REDIS_PASSWORD"},
			Usage:    "The password used to connect against the redis server",
			Required: false,
		},
		&cli.IntFlag{
			Name:     "redis-database-number",
			EnvVars:  []string{"REDIS_DATABASE_NUMBER"},
			Usage:    "The number of the database on the redis server",
			Required: false,
			Value:    0,
		},
		&cli.StringFlag{
			Name:     "redis-key-prefix",
			EnvVars:  []string{"REDIS_STRING_PREFIX"},
			Usage:    "The key prefix for storing values into the redis backend",
			Required: false,
			Value:    "ir-remotes",
		},
	}
}

// redisBackendFor initialized the redis backend with the given config variables
func redisBackendFor(context *cli.Context) backend.Backend {
	redisFlags := RedisFlagsFromContext(context)
	return redisBackend.New(redisFlags, context.String("redis-key-prefix"))
}

// RedisFlagsFromContext can be used to create the appropiate redis options for the configuration.
func RedisFlagsFromContext(context *cli.Context) *redis.Options {
	options := &redis.Options{}
	options.Addr = fmt.Sprintf("%s:%d", context.String("redis-host"), context.Int("redis-port"))
	fmt.Println(fmt.Sprintf("Redis Backend: %s", options.Addr))
	options.Password = context.String("redis-password")
	options.DB = context.Int("redis-database-number")
	return options
}
