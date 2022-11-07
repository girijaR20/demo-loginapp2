package loginapp2

import (
	"context"
	"demo-loginapp2/domain_mydomain/model/entity"
	//"fmt"
)

//go:generate mockery --name Outport -output mocks/

type loginapp2Interactor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &loginapp2Interactor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *loginapp2Interactor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	userObj, err := entity.NewUser(entity.UserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveUser(ctx, userObj)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Status :", userObj)
	//res.Status = userObj
	//return res, nil
	//err = r.outport.LoginUser(ctx, userObj)
	//if err != nil{
	//return nil,err
	//}
	//!
	res.Status = userObj.Status
	return res, nil
}
