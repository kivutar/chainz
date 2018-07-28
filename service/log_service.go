package service

import (
	"os"

	"github.com/kivutar/chainz/context"
	"github.com/op/go-logging"
)

func NewLogger(config *context.Config) *logging.Logger {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(config.LogFormat)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.INFO, "")
	if config.DebugMode {
		backendLeveled.SetLevel(logging.DEBUG, "")
	}

	logging.SetBackend(backendLeveled)
	logger := logging.MustGetLogger(config.AppName)
	return logger
}
