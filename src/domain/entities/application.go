package entities

import (
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
	electionusecase "github.com/max38/golang-clean-code-architecture/src/usecases/election"
)

type ApplicationEntity struct {
	UserUsecase userusecase.IUserUsecase
	ElectionUsecase electionusecase.IElectionUsecase
}
