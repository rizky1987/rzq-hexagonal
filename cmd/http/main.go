package main

import (
	"errors"
	"net/http"
	"os"
	"rzq-hexagonal/infrastructure/factory"
	"rzq-hexagonal/infrastructure/router"
)

func main() {

	services := factory.NewServiceFactory()
	startFramework(os.Getenv("APP_FRAMEWORK"), os.Getenv("APP_PORT"), services)

}

func startFramework(frameworkName, port string, servicesFactory *factory.ServicesFactory) error {

	app := router.NewEchoRouter()
	app.RegisterMiddleware()
	err := app.Start(port)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
