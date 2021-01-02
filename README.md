# IR Remote Backend #

[![IR Remote backend](https://img.shields.io/docker/pulls/cbrand/ir-remote-backend)](https://hub.docker.com/r/cbrand/ir-remote-backend) [![](https://img.shields.io/docker/image-size/cbrand/ir-remote-backend)](https://hub.docker.com/r/cbrand/ir-remote-backend) [![](https://img.shields.io/docker/v/cbrand/ir-remote-backend)](https://img.shields.io/docker/v/cbrand/ir-remote-backend)

Backend which allows communication between an ESP32 IR Remote communicating via a MQTT server and a Vue based frontend.

## Installation ##

A docker file is available on [Docker Hub](https://hub.docker.com/r/cbrand/ir-remote-backend).

The application requires a running MQTT server and a Redis server to function.

The easiest way is to use the embedded `docker-compose.yaml` file for configuration reference. The backend can
easily be started by running

```bash
docker-compose up
```

This starts the server and uses the latest image available on Docker hub.

The broker in this configuration doesn't use SSL in this configuration. For a production setup, SSL should be configured accordingly.

## Preliminary Software ##

This backend is build to function with
- [ir-remote-frontend](https://github.com/cbrand/ir-remote-frontend) - Vue.js backend frontend for interacting and configuring the IR remote
- [esp32-ir-remote](https://github.com/cbrand/esp32-ir-remote) - MicroPython project for using an ESP32 based Microcontroller for the IR app

## License ##

The software is published via the MIT license.
