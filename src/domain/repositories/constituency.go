package repositories

import (
	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
)

type IConstituencyRepository interface {
	GetConstituencies() (*[]entitymodels.ConstituencyModel, error)
	GetConstituencyByConstituencyId(constituencyId uint) (*entitymodels.ConstituencyModel, error)
	InsertConstituency(constituencyRegisterRequest *entityconstituency.ConstituencyCreateRequest) (*entityconstituency.ConstituencyEntity, error)
	UpdateConstituency(constituencyUpdateRequest *entityconstituency.ConstituencyUpdateRequest) (*entityconstituency.ConstituencyEntity, error)
	DeleteConstituency(constituencyId uint) error
}
