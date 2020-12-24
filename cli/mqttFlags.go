package cli

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	cli "github.com/urfave/cli/v2"
)

func MQTTFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "mqtt-host",
			EnvVars:  []string{"MQTT_SERVER_HOST"},
			Required: true,
			Usage:    "The host of the MQTT server to connect to.",
		},
		&cli.IntFlag{
			Name:     "mqtt-port",
			EnvVars:  []string{"MQTT_SERVER_PORT"},
			Required: false,
			Value:    1883,
			Usage:    "The port of the MQTT server to connect to.",
		},
		&cli.StringFlag{
			Name:     "mqtt-client-id",
			EnvVars:  []string{"MQTT_CLIENT_ID"},
			Required: false,
			Value:    "ir-remote-backend",
			Usage:    "The client id for the MQTT config.",
		},
		&cli.StringFlag{
			Name:     "mqtt-username",
			EnvVars:  []string{"MQTT_AUTH_USERNAME"},
			Required: false,
			Usage:    "The username to authenticate against the MQTT broker.",
		},
		&cli.StringFlag{
			Name:     "mqtt-password",
			EnvVars:  []string{"MQTT_AUTH_PASSWORD"},
			Required: false,
			Usage:    "The password to authenticate against the MQTT broker.",
		},
		&cli.BoolFlag{
			Name:     "mqtt-use-tls",
			EnvVars:  []string{"MQTT_USE_TLS"},
			Usage:    "If the authentication should happen with mqtt.",
			Required: false,
			Value:    false,
		},
		&cli.StringFlag{
			Name:      "mqtt-tls-ca",
			EnvVars:   []string{"MQTT_TLS_CA"},
			Usage:     "The path to the CA file which should be used for authentication.",
			Required:  false,
			TakesFile: true,
		},
		&cli.StringFlag{
			Name:      "mqtt-tls-certificate",
			EnvVars:   []string{"MQTT_TLS_CERTIFICATE"},
			Usage:     "The path of the tls certificate for client side authentication.",
			Required:  false,
			TakesFile: true,
		},
		&cli.StringFlag{
			Name:      "mqtt-tls-private-key",
			EnvVars:   []string{"MQTT_TLS_PRIVATE_KEY"},
			Usage:     "The path of the tls private key for client side authentication.",
			Required:  false,
			TakesFile: true,
		},
	}
}

// MQTTTLSConfigFromContext extracts from the context the tls configuration to connect to MQTT
func MQTTTLSConfigFromContext(context *cli.Context) (*tls.Config, error) {
	caPath := context.String("mqtt-tls-ca")
	certificatePath := context.String("mqtt-tls-certificate")
	privateKeyPath := context.String("mqtt-tls-private-key")
	certpool := x509.NewCertPool()
	if len(caPath) > 0 {
		fmt.Println(fmt.Sprintf("SSL: loading ca from %s", caPath))
		pemCerts, err := ioutil.ReadFile(caPath)
		if err == nil {
			certpool.AppendCertsFromPEM(pemCerts)
		}
	}
	var cert tls.Certificate
	var err error

	certificates := []tls.Certificate{}
	if len(certificatePath) > 0 {
		fmt.Println(fmt.Sprintf("SSL: loading certificate from %s and private key from %s", certificatePath, privateKeyPath))
		cert, err = tls.LoadX509KeyPair(certificatePath, privateKeyPath)
		if err != nil {
			return nil, err
		}
		cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(cert.Leaf)

		certificates = append(certificates, cert)
	}
	tlsConfig := &tls.Config{
		RootCAs:      certpool,
		Certificates: certificates,
	}
	return tlsConfig, nil
}

// MQTTConfigFromContext extracts the mqtt configuration from the passed cli context's options.
func MQTTConfigFromContext(context *cli.Context) (*mqtt.ClientOptions, error) {
	clientOptions := mqtt.NewClientOptions()
	hostname := context.String("mqtt-host")
	port := context.Int("mqtt-port")
	clientID := context.String("mqtt-client-id")
	clientOptions.SetClientID(clientID)
	username := context.String("mqtt-username")
	clientOptions.SetUsername(username)
	password := context.String("mqtt-password")
	clientOptions.SetPassword(password)
	useTLS := context.Bool("mqtt-use-tls")
	if useTLS {
		tlsConfig, err := MQTTTLSConfigFromContext(context)
		if err != nil {
			return nil, err
		}
		clientOptions.SetTLSConfig(tlsConfig)
	}
	brokerURL := fmt.Sprintf("%s:%d", hostname, port)
	if useTLS {
		brokerURL = fmt.Sprintf("tcps://%s", brokerURL)
	} else {
		brokerURL = fmt.Sprintf("tcp://%s", brokerURL)
	}
	clientOptions.AddBroker(brokerURL)
	clientOptions.SetMaxReconnectInterval(1 * time.Second)
	fmt.Println(fmt.Sprintf("Connecting to MQTT server %s (tls: %t)", brokerURL, useTLS))
	clientOptions.SetAutoReconnect(true)
	clientOptions.SetResumeSubs(true)
	clientOptions.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connection to MQTT server established")
	}
	clientOptions.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Println(fmt.Sprintf("Lost connection to mqtt server: %s", err.Error()))
	}
	return clientOptions, nil
}
