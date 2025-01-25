package entitymodels

import (
	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	sharedcrud "github.com/max38/golang-clean-code-architecture/src/shared/crud"
)

type ConstituencyModel struct {
	entityconstituency.ConstituencyEntity
}

func (u *ConstituencyModel) TableName() string {
	return "constituencies"
}

func (u *ConstituencyModel) ToEntity() *entityconstituency.ConstituencyEntity {
	return &u.ConstituencyEntity
}

func (u *ConstituencyModel) Datasource() sharedcrud.DatasourceType {
	return sharedcrud.DatasourcePostgresql
}
