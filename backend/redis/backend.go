package redis

import (
	"os"
	"strings"
	"time"

	"github.com/cbrand/ir-remote-backend/backend"
	"github.com/cbrand/ir-remote-backend/backend/generic"
	"github.com/go-redis/redis"
)

// New returns a backend which uses for its storage redis
func New(options *redis.Options, prefix string) backend.Backend {
	rdb := redis.NewClient(options)
	redisBytesBackend := &Backend{
		client: rdb,
		prefix: prefix,
	}
	return generic.New(redisBytesBackend)
}

// Backend stores ir remote information in redis
type Backend struct {
	client *redis.Client
	prefix string
}

// prefixedKey used for calculating a key form the given payload.
func (backend *Backend) prefixedKey(key string) string {
	return strings.Join([]string{backend.prefix, key}, "-")
}

// Get returns the byte payload for the specific key
func (backend *Backend) Get(key string) ([]byte, error) {
	prefixedKey := backend.prefixedKey(key)
	command := backend.client.Get(prefixedKey)
	if command.Err() == redis.Nil {
		return nil, os.ErrNotExist
	}
	return command.Bytes()
}

// Set sets the byte payload for the specific key
func (backend *Backend) Set(key string, value []byte) error {
	prefixedKey := backend.prefixedKey(key)
	command := backend.client.Set(prefixedKey, value, time.Duration(0))
	return command.Err()
}

// Delete removes the specified key from the backend
func (backend *Backend) Delete(key string) error {
	prefixedKey := backend.prefixedKey(key)
	command := backend.client.Del(prefixedKey)
	if command.Err() == redis.Nil {
		return os.ErrNotExist
	}
	return command.Err()
}
