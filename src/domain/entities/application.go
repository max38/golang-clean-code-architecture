package entities

import (
	candidateusecase "github.com/max38/golang-clean-code-architecture/src/usecases/candidate"
	constituencyusecase "github.com/max38/golang-clean-code-architecture/src/usecases/constituency"
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
	voteusecase "github.com/max38/golang-clean-code-architecture/src/usecases/vote"
)

type ApplicationEntity struct {
	UserUsecase         userusecase.IUserUsecase
	ConstituencyUsecase constituencyusecase.IConstituencyUsecase
	CandidateUsecase    candidateusecase.ICandidateUsecase
	VoteUsecase         voteusecase.IVoteUsecase
}
