package errors

import "github.com/pkg/errors"

const (
	errLoadConfigMsg                = "load config"
	errOpenDatabaseMsg              = "open database"
	errStartNatsSubscribeServiceMsg = "start nats subscribe service"
	errStartHTTPServerMsg           = "start http Server"
	errStartOrderServiceMsg         = "start order service"
)

func ErrLoadConfig(w error) error {
	return errors.Wrap(w, errLoadConfigMsg)
}

func ErrOpenDatabase(w error) error {
	return errors.Wrap(w, errOpenDatabaseMsg)
}

func ErrStartNatsSubscribeService(w error) error {
	return errors.Wrap(w, errStartNatsSubscribeServiceMsg)
}

func ErrStartHTTPServer(w error) error {
	return errors.Wrap(w, errStartHTTPServerMsg)
}

func ErrStartOrderService(w error) error {
	return errors.Wrap(w, errStartOrderServiceMsg)
}
