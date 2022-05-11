package app

import (
	"github.com/pkg/errors"
	"github.com/tapokshot/shelld1core/httpServer"
	"github.com/tapokshot/shelld1core/log"
)

type App struct {
	server *httpServer.Server
	Config *Config
	Log    *log.Logger
}

type initHandlersFunc func(server *httpServer.Server) error

const (
	defaultConfigPath = "config/default.yml"
)

func NewWithConfig(config *Config) (*App, error) {
	l, err := configureLogger(config)
	if err != nil {
		return nil, errors.Wrap(err, "error create logger")
	}

	server, err := httpServer.New(l)
	if err != nil {
		return nil, errors.Wrap(err, "error create http server")
	}

	return &App{
		Config: config,
		Log:    l,
		server: server,
	}, nil
}

// New create app with default configuration from
func New() (*App, error) {
	defaultConfig, err := CreateConfig(defaultConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "error create defaultConfig")
	}
	return NewWithConfig(defaultConfig)
}

func (a *App) Run() error {
	return a.server.Run(a.Config.BindAddr)
}

func (a *App) InitHttpHandlers(f initHandlersFunc) error {
	return f(a.server)
}

func configureLogger(config *Config) (*log.Logger, error) {
	rootLog, err := log.NewLogger(config.LoggerCfg)

	if err != nil {
		return nil, err
	}
	log.SetRootLog(rootLog)
	return rootLog, nil
}
