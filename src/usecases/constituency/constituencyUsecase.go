package constituencyusecase

import (
	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	repositories "github.com/max38/golang-clean-code-architecture/src/domain/repositories"
)

type IConstituencyUsecase interface {
	GetConstituencies() (*[]entityconstituency.ConstituencyEntity, error)
	GetConstituencyByConstituencyId(constituencyId uint) (*entityconstituency.ConstituencyEntity, error)
	CreateConstituency(constituencyCreateRequest *entityconstituency.ConstituencyCreateRequest) (*entityconstituency.ConstituencyEntity, error)
	UpdateConstituency(constituencyUpdateRequest *entityconstituency.ConstituencyUpdateRequest) (*entityconstituency.ConstituencyEntity, error)
	DeleteConstituency(constituencyId uint) error
}

type constituencyUsecase struct {
	constituencyRepository repositories.IConstituencyRepository
}

func ConstituencyUsecase(constituencyRepository repositories.IConstituencyRepository) IConstituencyUsecase {
	return &constituencyUsecase{
		constituencyRepository: constituencyRepository,
	}
}

func (u *constituencyUsecase) GetConstituencies() (*[]entityconstituency.ConstituencyEntity, error) {
	var constituenciesEntity, err = u.constituencyRepository.GetConstituencies()
	if err != nil {
		return nil, err
	}

	var constituencies []entityconstituency.ConstituencyEntity
	for _, constituency := range *constituenciesEntity {
		constituencies = append(constituencies, *constituency.ToEntity())
	}

	return &constituencies, nil
}

func (u *constituencyUsecase) CreateConstituency(constituencyCreateRequest *entityconstituency.ConstituencyCreateRequest) (*entityconstituency.ConstituencyEntity, error) {

	var constituencyEntity, err = u.constituencyRepository.InsertConstituency(constituencyCreateRequest)
	if err != nil {
		return nil, err
	}

	return constituencyEntity, nil
}

func (u *constituencyUsecase) GetConstituencyByConstituencyId(constituencyId uint) (*entityconstituency.ConstituencyEntity, error) {
	var constituencyEntity, errorGetConstituency = u.constituencyRepository.GetConstituencyByConstituencyId(constituencyId)
	if errorGetConstituency != nil {
		return nil, errorGetConstituency
	}

	return constituencyEntity.ToEntity(), nil
}

func (u *constituencyUsecase) UpdateConstituency(constituencyUpdateRequest *entityconstituency.ConstituencyUpdateRequest) (*entityconstituency.ConstituencyEntity, error) {
	var constituencyEntity, errorUpdateConstituency = u.constituencyRepository.UpdateConstituency(constituencyUpdateRequest)
	if errorUpdateConstituency != nil {
		return nil, errorUpdateConstituency
	}

	return constituencyEntity, nil
}

func (u *constituencyUsecase) DeleteConstituency(constituencyId uint) error {
	return u.constituencyRepository.DeleteConstituency(constituencyId)
}
