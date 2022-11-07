package application

import (
	"demo-loginapp2/domain_mydomain/controller/userapi"
	"demo-loginapp2/domain_mydomain/gateway/inmemory"
	"demo-loginapp2/domain_mydomain/usecase/loginapp2"
	"demo-loginapp2/shared/driver"
	"demo-loginapp2/shared/infrastructure/config"
	"demo-loginapp2/shared/infrastructure/logger"
	"demo-loginapp2/shared/infrastructure/server"
	"demo-loginapp2/shared/util"
)

type loginapp struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c loginapp) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewLoginapp() func() driver.RegistryContract {

	const appName = "loginapp"

	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData(appName, appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

		datasource := inmemory.NewGateway(log, appData, cfg)

		return &loginapp{
			httpHandler: &httpHandler,
			controller: &userapi.Controller{
				Log:             log,
				Config:          cfg,
				Router:          httpHandler.Router,
				Loginapp2Inport: loginapp2.NewUsecase(datasource),
			},
		}

	}
}
