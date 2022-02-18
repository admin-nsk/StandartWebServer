package api

import (
	"github.com/admin-nsk/StandartWebServer/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Base API server instance
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Поле для работы с хранилищем
	store *storage.Storage
}

//API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/conigure loggers, router, database
func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting api server at port:", api.config.BindAddr)

	//Конфигурирование маршрутизатора
	api.configureRouterField()

	//Конфигурируем хранилище
	if err := api.configureStrageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
