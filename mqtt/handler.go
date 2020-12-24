package mqtt

import (
	"encoding/json"
	"strings"

	"github.com/cbrand/ir-remote-backend/protocol"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// NewHandlerFromOptions returns a mqtt client for a specific handler.
func NewHandlerFromOptions(options *mqtt.ClientOptions) *Handler {
	client := mqtt.NewClient(options)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		panic(token.Error())
	}
	return NewHandler(client)
}

// NewHandler returns a new mqtt handler backed by the passed mqtt client.
func NewHandler(client mqtt.Client) *Handler {
	return &Handler{
		mqttClient:   client,
		remoteStatus: map[string]*RemoteStatus{},
	}
}

// Handler supports to send and retrieve data via MQTT to and from the connected remotes
type Handler struct {
	mqttClient   mqtt.Client
	remoteStatus map[string]*RemoteStatus
}

// Monitor starts checking for the specified remote the online status and returns the current
// configuration.
func (handler *Handler) Monitor(remote *protocol.Remote) (*RemoteStatus, error) {
	remoteStatus, ok := handler.remoteStatus[remote.GetMqttTopicPrefix()]
	var err error = nil
	if !ok {
		handler.connectIfNecessary()
		remoteStatus := &RemoteStatus{}
		handler.remoteStatus[remote.GetMqttTopicPrefix()] = remoteStatus
		remoteStatus.Online = false
		token := handler.mqttClient.Subscribe(handler.topicFor(remote, "livesign"), 1, func(client mqtt.Client, message mqtt.Message) {
			liveSignMessage := &LiveSignMessage{}
			err := json.Unmarshal(message.Payload(), liveSignMessage)
			if err != nil {
				return
			}
			message.Ack()

			remoteStatus, ok := handler.remoteStatus[remote.GetMqttTopicPrefix()]
			if !ok {
				return
			}
			remoteStatus.Lifesign = liveSignMessage.Datetime.ToIso8601()
			remoteStatus.Update()
		})
		token.Wait()
		err = token.Error()
	}
	return remoteStatus, err
}

// SendScene puts the provided scene to the MQTT server for the remote to be sent
func (handler *Handler) SendScene(remote *protocol.Remote, scene *protocol.Scene) error {
	payload := SerializeScene(scene)
	err := handler.publishPayload(remote, payload)
	return err
}

// SendCommand sends the passed command to the passed remote
func (handler *Handler) SendCommand(remote *protocol.Remote, command *protocol.Command) error {
	payload := SerializeCommand(command)
	err := handler.publishPayload(remote, payload)
	return err
}

// publishPayload publishes the given payload to the MQTT Broker for sending to the IR.
func (handler *Handler) publishPayload(remote *protocol.Remote, payload map[string]interface{}) error {
	handler.connectIfNecessary()
	topic := handler.topicFor(remote, "ir/command")
	serializedPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	token := handler.mqttClient.Publish(topic, 1, false, serializedPayload)
	token.Wait()
	return token.Error()
}

// connectIfNecessary checks if the connection to the mqtt server has been established. If not, it tries to connect
func (handler *Handler) connectIfNecessary() error {
	mqttClient := handler.mqttClient
	if mqttClient.IsConnected() {
		return nil
	}

	token := mqttClient.Connect()
	token.Wait()
	return token.Error()
}

// topicFor gives the fully configured topic for the given remote which should be used for publishing or subscribing to data.
func (handler *Handler) topicFor(remote *protocol.Remote, topicPostfix string) string {
	mqttTopicPrefix := remote.GetMqttTopicPrefix()
	fixedTopicPrefix := strings.TrimSuffix(mqttTopicPrefix, "/")
	fixedTopicPostfix := strings.TrimPrefix(topicPostfix, "/")
	return strings.Join([]string{fixedTopicPrefix, fixedTopicPostfix}, "/")
}
