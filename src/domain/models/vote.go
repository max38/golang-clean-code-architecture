package entitymodels

import (
	entityvote "github.com/max38/golang-clean-code-architecture/src/domain/entities/vote"
	sharedcrud "github.com/max38/golang-clean-code-architecture/src/shared/crud"
)

type VoteModel struct {
	entityvote.VoteEntity
	Candidate CandidateModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"candidate"`
	User      UserModel      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}

func (u *VoteModel) TableName() string {
	return "votes"
}

func (u *VoteModel) ToEntity() *entityvote.VoteEntity {
	return &u.VoteEntity
}

func (u *VoteModel) Datasource() sharedcrud.DatasourceType {
	return sharedcrud.DatasourcePostgresql
}
