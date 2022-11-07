package repository

import (
	"context"
	"demo-loginapp2/domain_mydomain/model/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
	//LoginUser(ctx context.Context, obj *entity.User) (bool, error)
}
