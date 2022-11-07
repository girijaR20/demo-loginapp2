package userapi

import (
	"github.com/gin-gonic/gin"

	"demo-loginapp2/domain_mydomain/usecase/loginapp2"
	"demo-loginapp2/shared/infrastructure/config"
	"demo-loginapp2/shared/infrastructure/logger"
)

type Controller struct {
	Router          gin.IRouter
	Config          *config.Config
	Log             logger.Logger
	Loginapp2Inport loginapp2.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/loginapp2", r.authorized(), r.loginapp2Handler())
}
