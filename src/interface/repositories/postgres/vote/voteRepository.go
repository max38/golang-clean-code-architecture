package postgresvoterepository

import (
	entityvote "github.com/max38/golang-clean-code-architecture/src/domain/entities/vote"
	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
	repositories "github.com/max38/golang-clean-code-architecture/src/domain/repositories"
	postgreshandler "github.com/max38/golang-clean-code-architecture/src/infrastructure/database/postgres"
)

type voteRepository struct {
	dbHandler postgreshandler.IPostgresHandler
}

func VoteRepository() repositories.IVoteRepository {
	var dbHandler = postgreshandler.PostgresHandler()
	return &voteRepository{
		dbHandler: dbHandler,
	}
}

func (u *voteRepository) CountVoteByCandidateId(candidateId uint) (*int64, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var count int64
	var dbResult = dbConnector.Model(&entitymodels.VoteModel{}).Where("candidate_id = ?", candidateId).Count(&count)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &count, nil
}

func (u *voteRepository) CheckExistsVoteByUserId(userId uint) (bool, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var count int64
	var dbResult = dbConnector.Model(&entitymodels.VoteModel{}).Where("user_id = ?", userId).Count(&count)
	if dbResult.Error != nil {
		return false, dbResult.Error
	}

	return count > 0, nil
}

func (u *voteRepository) InsertVote(voteRegisterRequest *entityvote.VoteCreateWithUserRequest) (*entityvote.VoteEntity, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var voteModel = entitymodels.VoteModel{
		VoteEntity: entityvote.VoteEntity{
			CandidateId: voteRegisterRequest.CandidateId,
			UserId:      voteRegisterRequest.UserId,
		},
	}

	var dbResult = dbConnector.Create(&voteModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var voteEntity = *voteModel.ToEntity()

	return &voteEntity, nil
}
