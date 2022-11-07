package inmemory

import (
	"context"
	"demo-loginapp2/domain_mydomain/model/entity"
	"demo-loginapp2/shared/driver"
	"demo-loginapp2/shared/infrastructure/config"
	"demo-loginapp2/shared/infrastructure/logger"
)

type gateway struct {
	Users []entity.User
	//Users1   []entity.User
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) SaveUser(ctx context.Context, obj *entity.User) error {
	r.log.Info(ctx, "called")
	r.Users = append(r.Users)

	return nil
}

//func (r *gateway) LoginUser(ctx context.Context, obj *entity.User) (bool, error) {
//r.log.Info(ctx, "called")
//for _, v := range r.Users {
//if v.Username == obj.Username {
//v.Password = obj.Password
//return true, nil
//}
//}

//r.Users = append(r.Users)

//return false, nil

//}
