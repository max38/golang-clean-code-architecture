package repositories

import (
	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
)

type ICandidateRepository interface {
	GetGetCandidates() (*[]entitymodels.CandidateModel, error)
	GetCandidateByCandidateId(candidateId uint) (*entitymodels.CandidateModel, error)
	GetCandidateByConstituencyId(constituencyId uint) (*[]entitymodels.CandidateModel, error)
	InsertCandidate(candidateRegisterRequest *entitycandidate.CandidateCreateRequest) (*entitycandidate.CandidateEntity, error)
	UpdateCandidate(candidateUpdateRequest *entitycandidate.CandidateUpdateRequest) (*entitycandidate.CandidateEntity, error)
	DeleteCandidate(candidateId uint) error
}
