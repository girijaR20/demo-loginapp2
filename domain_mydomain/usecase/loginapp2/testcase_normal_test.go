package loginapp2

import (
	"context"
	"testing"

	"demo-loginapp2/domain_mydomain/model/entity"
)

type mockOutportNormal struct {
	t *testing.T
}

// TestCaseNormal is for the case where ...
// explain the purpose of this test here with human readable naration...
func TestCaseNormal(t *testing.T) {

	ctx := context.Background()

	mockOutport := mockOutportNormal{
		t: t,
		//status: User,
	}

	res, err := NewUsecase(&mockOutport).Execute(ctx, InportRequest{
		Username: "paxel",
		Password: "paxel",
	})

	if err != nil {
		t.Errorf("%v", err.Error())
		t.FailNow()
	}

	t.Logf("%v", res)

}

func (r *mockOutportNormal) SaveUser(ctx context.Context, obj *entity.User) error {

	return nil
}
func (r *mockOutportNormal) LoginUser(ctx context.Context, obj *entity.User) (bool, error) {
	return false, nil
}
