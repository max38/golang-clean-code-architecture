package entities

import (
	constituencyusecase "github.com/max38/golang-clean-code-architecture/src/usecases/constituency"
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
)

type ApplicationEntity struct {
	UserUsecase         userusecase.IUserUsecase
	ConstituencyUsecase constituencyusecase.IConstituencyUsecase
}
