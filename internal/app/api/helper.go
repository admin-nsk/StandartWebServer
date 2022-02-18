package api

import (
	"github.com/admin-nsk/StandartWebServer/storage"

	"github.com/sirupsen/logrus"
	"net/http"
)

//Пытаемся от конфигурировать наш API instance
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

//пытаемся отконфигурировать маршрутизатор
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server started"))
	})
}

//Конфигурируем хранилище
func (a *API) configureStrageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.store = storage
	return nil
}
