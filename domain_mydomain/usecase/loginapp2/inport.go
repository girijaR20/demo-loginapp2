package loginapp2

import (
	"context"

	"demo-loginapp2/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Username string
	Password string
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Status string
}
