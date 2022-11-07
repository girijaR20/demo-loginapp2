package errorenum

import "demo-loginapp2/shared/model/apperror"

const (
	SomethingError apperror.ErrorType = "ER0000 something error"
	//ValidateUsername     apperror.ErrorType = "ER0001 validate username"
	//ValidatePassword     apperror.ErrorType = "ER0002 validate password"
	UsernameMustNotEmpty apperror.ErrorType = "ER0003 username must not empty"
	PasswordMustNotEmpty apperror.ErrorType = "ER0004 password must not empty"
	UsernameOrPassword   apperror.ErrorType = "ER0005 username or password"
)
