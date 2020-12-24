package backend

import (
	"github.com/cbrand/ir-remote-backend/protocol"
)

// RemoteID specifies an ID of a remote for interaction and requests.
type RemoteID string

// Backend represents the interface of the backend which stores the
// necessary IR remote information
type Backend interface {
	// Get returns the remote for the given ID. This is also used for authentication purposes
	// thus the ID shouldn't be to obvious.
	Get(remoteID string) (*protocol.Remote, error)
	// Set updates a specified remote configuraiton in the backend.
	Set(remote *protocol.Remote) error
	// Delete removes the specified remote of the given ID.
	Delete(remoteID string) error
	// GetTheaters returns a list of all theaters stored for the given remote.
	GetTheaters(remoteID string) ([]*protocol.Theater, error)
	// SetTheater sets or updates a theater for the provided remote.
	SetTheater(remoteID string, theater *protocol.Theater) error
	// DeleteTheater sets or updates a theater for the provided remote.
	DeleteTheater(remoteID string, theaterID string) error
}
