package main

import (
	"time"

	"github.com/a-pavithraa/product-service-golang/api"
	"github.com/a-pavithraa/product-service-golang/util"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	log "github.com/sirupsen/logrus"
)

func configureLogging(logLevel string) error {
	l, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(l)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
	})
	return nil
}

func main() {
	settings := util.LoadAppConfig()
	configureLogging(settings.LogLevel)
	server := api.NewApiServer(settings)

	err := server.Start(settings.Port)
	if err != nil {
		log.Fatal(err)
	}

}
