FROM golang:1.15-alpine AS build-env

WORKDIR /app

RUN apk update && apk add --no-cache git alpine-sdk bash
RUN apk add --no-cache ca-certificates && update-ca-certificates
RUN adduser -D -g '' appuser

ADD . /app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ir-remote-backend

FROM alpine:3.9
WORKDIR /app
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /etc/passwd /etc/passwd
COPY --from=build-env /app/ir-remote-backend /app/
ENV GIN_MODE release

USER appuser

EXPOSE 9000

ENTRYPOINT [ "/app/ir-remote-backend" ]
CMD [ "server" ]
