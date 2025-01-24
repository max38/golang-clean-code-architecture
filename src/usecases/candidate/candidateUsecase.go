package candidateusecase

import (
	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	repositories "github.com/max38/golang-clean-code-architecture/src/domain/repositories"
)

type ICandidateUsecase interface {
	GetGetCandidates() (*[]entitycandidate.CandidateEntity, error)
	GetCandidateByCandidateId(candidateId uint) (*entitycandidate.CandidateEntity, error)
	GetCandidateByConstituencyId(constituencyId uint) (*[]entitycandidate.CandidateEntity, error)
	CreateCandidate(candidateCreateRequest *entitycandidate.CandidateCreateRequest) (*entitycandidate.CandidateEntity, error)
	UpdateCandidate(candidateUpdateRequest *entitycandidate.CandidateUpdateRequest) (*entitycandidate.CandidateEntity, error)
	DeleteCandidate(candidateId uint) error
}

type candidateUsecase struct {
	candidateRepository repositories.ICandidateRepository
}

func CandidateUsecase(candidateRepository repositories.ICandidateRepository) ICandidateUsecase {
	return &candidateUsecase{
		candidateRepository: candidateRepository,
	}
}

func (u *candidateUsecase) GetGetCandidates() (*[]entitycandidate.CandidateEntity, error) {
	var constituenciesEntity, err = u.candidateRepository.GetGetCandidates()
	if err != nil {
		return nil, err
	}

	var constituencies []entitycandidate.CandidateEntity
	for _, candidate := range *constituenciesEntity {
		constituencies = append(constituencies, *candidate.ToEntity())
	}

	return &constituencies, nil
}

func (u *candidateUsecase) CreateCandidate(candidateCreateRequest *entitycandidate.CandidateCreateRequest) (*entitycandidate.CandidateEntity, error) {

	var candidateEntity, err = u.candidateRepository.InsertCandidate(candidateCreateRequest)
	if err != nil {
		return nil, err
	}

	return candidateEntity, nil
}

func (u *candidateUsecase) GetCandidateByCandidateId(candidateId uint) (*entitycandidate.CandidateEntity, error) {
	var candidateEntity, errorGetCandidate = u.candidateRepository.GetCandidateByCandidateId(candidateId)
	if errorGetCandidate != nil {
		return nil, errorGetCandidate
	}

	return candidateEntity.ToEntity(), nil
}

func (u *candidateUsecase) GetCandidateByConstituencyId(constituencyId uint) (*[]entitycandidate.CandidateEntity, error) {
	var candidatesEntity, err = u.candidateRepository.GetCandidateByConstituencyId(constituencyId)
	if err != nil {
		return nil, err
	}

	var candidates []entitycandidate.CandidateEntity
	for _, candidate := range *candidatesEntity {
		candidates = append(candidates, *candidate.ToEntity())
	}

	return &candidates, nil
}

func (u *candidateUsecase) UpdateCandidate(candidateUpdateRequest *entitycandidate.CandidateUpdateRequest) (*entitycandidate.CandidateEntity, error) {
	var candidateEntity, errorUpdateCandidate = u.candidateRepository.UpdateCandidate(candidateUpdateRequest)
	if errorUpdateCandidate != nil {
		return nil, errorUpdateCandidate
	}

	return candidateEntity, nil
}

func (u *candidateUsecase) DeleteCandidate(candidateId uint) error {
	return u.candidateRepository.DeleteCandidate(candidateId)
}
