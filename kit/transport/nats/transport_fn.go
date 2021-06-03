package nats

import (
	"github.com/keshareenv/go-base/utils/log"
	natn "github.com/nats-io/nats.go"
)

// DisconnectErrorCallback is called when the connection to
// nats server is lost
func DisconnectErrorCallback(logger log.Logger) natn.ConnErrHandler {
	return func(nc *natn.Conn, err error) {
		logger.Error("disconnected from nats", log.Error(err))
	}
}

// ReconnectCallback is called when the connection to nats
// server is re-established
func ReconnectCallback(logger log.Logger) natn.ConnHandler {
	return func(nc *natn.Conn) {
		logger.Info("Got reconnected", log.String("url", nc.ConnectedUrl()))
	}
}
