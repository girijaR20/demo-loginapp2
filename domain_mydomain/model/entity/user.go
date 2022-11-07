package entity

import (
	"demo-loginapp2/domain_mydomain/model/errorenum"
	"fmt"
)

type User struct {
	ID       string `` //
	Username string
	Password string
	Status   string
}

type UserRequest struct {
	Username string
	Password string
}

func NewUser(req UserRequest) (*User, error) {
	if req.Username == "" {
		return nil, errorenum.UsernameMustNotEmpty
	}
	fmt.Println("UsernameMustNotEmpty", req.Username)
	if req.Password == "" {
		return nil, errorenum.PasswordMustNotEmpty
	}
	fmt.Println("PasswordMustNotEmpty", req.Password)

	var obj User

	obj.Username = req.Username
	obj.Password = req.Password

	if req.Username == "paxel" && req.Password == "paxel" {
		//	fmt.Println("Username:paxel && Password:paxel", req.Username, req.Password)
		obj.Status = "loggedinsuccessfully"
	} else {
		obj.Status = "loggedinfailed"
	}

	// assign value here

	err := obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil

}

func (r *User) Validate() error {
	return nil
}
