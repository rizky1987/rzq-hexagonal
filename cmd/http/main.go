package main

import (
	"errors"
	"net/http"
	"os"
	echoHandler "rzq-hexagonal/adapter/http/handler/echo"
	fiberHandler "rzq-hexagonal/adapter/http/handler/fiber"
	"rzq-hexagonal/infrastructure/constanta"
	"rzq-hexagonal/infrastructure/factory"
	"rzq-hexagonal/infrastructure/router"
)

func main() {

	services := factory.NewServiceFactory()
	startFramework(os.Getenv("APP_FRAMEWORK"), os.Getenv("APP_PORT"), services)

}

func startFramework(frameworkName, port string, servicesFactory *factory.ServicesFactory) error {

	switch frameworkName {
	case constanta.FrameworkTypeEcho:

		app := router.NewEchoRouter()
		app.RegisterMiddleware()
		app.RegisterRoutes(&echoHandler.EchoHandler{servicesFactory})
		err := app.Start(port)

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	case constanta.FrameworkTypeFiber:
		app := router.NewFiberRouter()
		app.RegisterMiddleware()
		app.RegisterRoutes(&fiberHandler.FiberHandler{servicesFactory})
		err := app.Start(port)

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	default:
		return errors.New("please choice your framework")
	}

	return nil
}
