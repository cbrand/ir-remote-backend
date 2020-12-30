package handler

import (
	"context"

	"github.com/cbrand/ir-remote-backend/backend"
	"github.com/cbrand/ir-remote-backend/mqtt"
	"github.com/cbrand/ir-remote-backend/protocol"
	uuid "github.com/satori/go.uuid"
)

// NewServer returns a new server object backed by the passed configuration
func NewServer(storage backend.Backend, mqttHandler *mqtt.Handler) *Server {
	return &Server{
		storage:     storage,
		mqttHandler: mqttHandler,
	}
}

// Server is the main server running the logic for handling
type Server struct {
	protocol.UnimplementedRemoteServiceServer
	storage     backend.Backend
	mqttHandler *mqtt.Handler
}

// AddRemote is not publicly available. It allows to add a new remote configuration to the backend store.
func (server *Server) AddRemote(remoteName string, remoteMQTTPrefix string) (*protocol.Remote, error) {
	newID := uuid.NewV4()
	newRemote := &protocol.Remote{
		Id:              newID.String(),
		Name:            remoteName,
		MqttTopicPrefix: remoteMQTTPrefix,
	}
	return newRemote, server.storage.Set(newRemote)
}

// GetRemote implements the service to return the specific remote
func (server *Server) GetRemote(ctx context.Context, remoteQueryParams *protocol.RemoteQueryParams) (*protocol.Remote, error) {
	return server.storage.Get(remoteQueryParams.GetId())
}

// GetStatus implements the status check of a specific remote.
func (server *Server) GetStatus(ctx context.Context, remoteQueryParams *protocol.RemoteQueryParams) (*protocol.RemoteStatus, error) {
	remote, err := server.GetRemote(ctx, remoteQueryParams)
	if err != nil {
		return nil, err
	}
	remoteStatus, err := server.mqttHandler.Monitor(remote)
	if err != nil {
		return nil, err
	}
	remoteStatus.Update()

	returnRemoteStatus := &protocol.RemoteStatus{
		RemoteId: remote.GetId(),
		Lifesign: remoteStatus.Lifesign,
		Online:   remoteStatus.Online,
	}
	return returnRemoteStatus, nil
}

// theaterResponseOf returns a theater resonse from the passed slice
func theaterResponseOf(theaters []*protocol.Theater) *protocol.TheatersResponse {
	response := &protocol.TheatersResponse{
		Theaters: theaters,
	}
	return response
}

// generateTheatersResponseOf returns the theater configuration.
func (server *Server) generateTheatersResponseOf(remoteID string) (*protocol.TheatersResponse, error) {
	theaters, err := server.storage.GetTheaters(remoteID)
	if err != nil {
		return nil, err
	}
	return theaterResponseOf(theaters), nil
}

// GetTheatersOf returns a list of the theaters for the passed query configuration.
func (server *Server) GetTheatersOf(ctx context.Context, remoteQueryParams *protocol.RemoteQueryParams) (*protocol.TheatersResponse, error) {
	return server.generateTheatersResponseOf(remoteQueryParams.GetId())
}

// SetTheater adds the theater to the specified configuration
// rpc setTheater(setTheaterQueryParams) returns (TheatersResponse);
func (server *Server) SetTheater(ctx context.Context, setTheaterQueryParams *protocol.SetTheaterQueryParams) (*protocol.TheatersResponse, error) {
	err := server.storage.SetTheater(setTheaterQueryParams.GetRemoteId(), setTheaterQueryParams.GetTheater())
	if err != nil {
		return nil, err
	}
	return server.generateTheatersResponseOf(setTheaterQueryParams.GetRemoteId())
}

// RemoveTheater removes a theater by their id from a specific remote.
// rpc removeTheater(removeTheaterQueryParams) returns (TheatersResponse);
func (server *Server) RemoveTheater(ctx context.Context, removeTheaterQueryParams *protocol.RemoveTheaterQueryParams) (*protocol.TheatersResponse, error) {
	err := server.storage.DeleteTheater(removeTheaterQueryParams.GetRemoteId(), removeTheaterQueryParams.GetTheaterId())
	if err != nil {
		return nil, err
	}
	return server.generateTheatersResponseOf(removeTheaterQueryParams.GetRemoteId())
}

// PlayScene takes the specified scene and plays them against the specified remote.
func (server *Server) PlayScene(ctx context.Context, playSceneParams *protocol.PlaySceneParams) (*protocol.PlayResponse, error) {
	remote, err := server.storage.Get(playSceneParams.GetRemoteId())
	response := &protocol.PlayResponse{
		Ok: false,
	}
	if err != nil {
		return response, err
	}
	err = server.mqttHandler.SendScene(remote, playSceneParams.GetScene())
	if err != nil {
		return response, err
	}
	response.Ok = true
	return response, err
}

// PlayCommand takes the specified ocmmand and plays them against the specified remote.
func (server *Server) PlayCommand(ctx context.Context, playCommandParams *protocol.PlayCommandParams) (*protocol.PlayResponse, error) {
	remote, err := server.storage.Get(playCommandParams.GetRemoteId())
	response := &protocol.PlayResponse{
		Ok: false,
	}
	if err != nil {
		return response, err
	}
	err = server.mqttHandler.SendCommand(remote, playCommandParams.GetCommand())
	if err != nil {
		return response, err
	}
	response.Ok = true
	return response, err
}

// GetIscpStatusOf returns the iscp status of the passed remote information.
func (server *Server) GetIscpStatusOf(ctx context.Context, remoteQueryParams *protocol.RemoteQueryParams) (*protocol.IscpStatusResponse, error) {
	remote, err := server.storage.Get(remoteQueryParams.GetId())
	response := &protocol.IscpStatusResponse{
		Devices: []*protocol.IscpStatus{},
	}
	if err != nil {
		return response, err
	}

	iscpStatus, err := server.mqttHandler.GetIscpStatus(remote)
	if err != nil {
		return response, err
	}

	protocolStatusSlice := []*protocol.IscpStatus{}
	for _, status := range iscpStatus {
		protocolStatus := &protocol.IscpStatus{
			Identifier:     status.Identifier,
			ModelName:      status.ModelName,
			AreaCode:       status.AreaCode,
			DeviceCategory: status.DeviceCategory,
			IscpPort:       status.IscpPort,
		}
		protocolStatusSlice = append(protocolStatusSlice, protocolStatus)
	}
	response.Devices = protocolStatusSlice
	return response, nil
}
