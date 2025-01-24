package entitymodels

import (
	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	sharedcrud "github.com/max38/golang-clean-code-architecture/src/shared/crud"
)

type CandidateModel struct {
	entitycandidate.CandidateEntity
	Constituency ConstituencyModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"constituency"`
}

func (u *CandidateModel) TableName() string {
	return "candidates"
}

func (u *CandidateModel) ToEntity() *entitycandidate.CandidateEntity {
	return &u.CandidateEntity
}

func (u *CandidateModel) Datasource() sharedcrud.DatasourceType {
	return sharedcrud.DatasourcePostgresql
}
