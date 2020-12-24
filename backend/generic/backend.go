package generic

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/cbrand/ir-remote-backend/protocol"
)

// prefixedRemoteKey returns a remote key for a specific prefixed configuration.
func prefixedRemoteKey(key string) string {
	return strings.Join([]string{"remote", key}, "-")
}

// prefixedTheaterKey returns a prefixed key where the theathers of the remote are stored.
func prefixedTheaterKey(remoteID string) string {
	return strings.Join([]string{"theater", remoteID}, "-")
}

// BytesBackend represents the interface which underlying backends needs to implement for the
// generic backend to be used.
type BytesBackend interface {
	// Get returns the byte payload for the specific key
	Get(key string) ([]byte, error)
	// Set sets the byte payload for the specific key
	Set(key string, value []byte) error
	// Delete removes the specified key from the backend
	Delete(key string) error
}

// New returns a new backend for getting ir remote information using the passed
// bytes backend for persistence.
func New(byteBackend BytesBackend) *Backend {
	return &Backend{
		backend: byteBackend,
	}
}

// Backend implements a higher level representation of a backend and requires a lower level
// byte based backend for storing IR remote information.
type Backend struct {
	backend BytesBackend
}

// Get returns the remote protocol from the given configuration.
func (backend *Backend) Get(remoteID string) (*protocol.Remote, error) {
	backendRemoteID := prefixedRemoteKey(remoteID)
	bytesPayload, err := backend.backend.Get(backendRemoteID)
	if err != nil {
		return nil, err
	}

	remote := &protocol.Remote{}
	err = json.Unmarshal(bytesPayload, remote)
	if err != nil {
		return nil, err
	}

	return remote, nil
}

// Set updates a specified remote configuraiton in the backend.
func (backend *Backend) Set(remote *protocol.Remote) error {
	marshalledPayload, err := json.Marshal(remote)
	if err != nil {
		return err
	}

	backendRemoteID := prefixedRemoteKey(remote.GetId())
	return backend.backend.Set(backendRemoteID, marshalledPayload)
}

// Delete removes the specified remote of the given ID.
func (backend *Backend) Delete(remoteID string) error {
	backendRemoteID := prefixedRemoteKey(remoteID)
	theaterSortedKey := prefixedTheaterKey(remoteID)
	backend.backend.Delete(theaterSortedKey)
	return backend.backend.Delete(backendRemoteID)
}

// GetTheaters returns a list of all theaters stored for the given remote.
func (backend *Backend) GetTheaters(remoteID string) ([]*protocol.Theater, error) {
	_, err := backend.Get(remoteID)
	if err != nil {
		return nil, err
	}
	theaterRemoteKey := prefixedTheaterKey(remoteID)
	data, err := backend.backend.Get(theaterRemoteKey)
	if os.IsNotExist(err) {
		return []*protocol.Theater{}, nil
	}
	var theaters []*protocol.Theater
	err = json.Unmarshal(data, &theaters)
	if err != nil {
		return nil, err
	}
	return theaters, nil
}

// SetTheater sets or updates a theater for the provided remote.
func (backend *Backend) SetTheater(remoteID string, theater *protocol.Theater) error {
	theaters, err := backend.GetTheaters(remoteID)
	if err != nil {
		return err
	}
	theaterMapByID := TheaterSliceToMap(theaters)
	theaterMapByID[theater.GetId()] = theater
	return backend.setTheaterFromMap(remoteID, theaterMapByID)
}

// DeleteTheater sets or updates a theater for the provided remote.
func (backend *Backend) DeleteTheater(remoteID string, theaterID string) error {
	theaters, err := backend.GetTheaters(remoteID)
	if err != nil {
		return err
	}
	theaterMapByID := TheaterSliceToMap(theaters)
	delete(theaterMapByID, theaterID)
	return backend.setTheaterFromMap(remoteID, theaterMapByID)
}

// setTheaterFromMap updates the theaters in the backend
func (backend *Backend) setTheaterFromMap(remoteID string, theatersMap map[string]*protocol.Theater) error {
	theaters := TheaterMapToSortedSlice(theatersMap)
	return backend.setTheaterFromSlice(remoteID, theaters)
}

// setTheaterFromSlice updates the theaters in the backend
func (backend *Backend) setTheaterFromSlice(remoteID string, theaters []*protocol.Theater) error {
	theatersBytes, err := json.Marshal(theaters)
	if err != nil {
		return err
	}
	theaterStorageKey := prefixedTheaterKey(remoteID)
	return backend.backend.Set(theaterStorageKey, theatersBytes)
}
