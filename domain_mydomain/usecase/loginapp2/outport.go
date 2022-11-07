package loginapp2

import "demo-loginapp2/domain_mydomain/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveUserRepo
}
