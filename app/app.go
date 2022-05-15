package app

import (
	"github.com/pkg/errors"
	"github.com/tapokshot/shelld1core/httpServer"
	"github.com/tapokshot/shelld1core/log"
)

type App struct {
	server *httpServer.Server
	Log    *log.Logger
}

type initHandlersFunc func(server *httpServer.Server) error

const (
	defaultConfigPath = "config/default.yml"
)

func New() (*App, error) {
	l, err := configureLogger()
	if err != nil {
		return nil, errors.Wrap(err, "error create logger")
	}

	server, err := httpServer.New(l)
	if err != nil {
		return nil, errors.Wrap(err, "error create http server")
	}

	return &App{
		Log:    l,
		server: server,
	}, nil
}

// New create app with default configuration from
//func New() (*App, error) {
//	defaultConfig, err := CreateConfig(defaultConfigPath)
//	if err != nil {
//		return nil, errors.Wrap(err, "error create defaultConfig")
//	}
//	return NewWithConfig(defaultConfig)
//}

func (a *App) Run(addr string) error {
	return a.server.Run(addr)
}

func (a *App) InitHttpHandlers(f initHandlersFunc) error {
	return f(a.server)
}

func configureLogger() (*log.Logger, error) {
	rootLog, err := log.NewLogger()

	if err != nil {
		return nil, err
	}
	log.SetRootLog(rootLog)
	return rootLog, nil
}
