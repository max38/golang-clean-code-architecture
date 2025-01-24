package postgrescandidaterepository

import (
	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
	repositories "github.com/max38/golang-clean-code-architecture/src/domain/repositories"
	postgreshandler "github.com/max38/golang-clean-code-architecture/src/infrastructure/database/postgres"
)

type candidateRepository struct {
	dbHandler postgreshandler.IPostgresHandler
}

func CandidateRepository() repositories.ICandidateRepository {
	var dbHandler = postgreshandler.PostgresHandler()
	return &candidateRepository{
		dbHandler: dbHandler,
	}
}

func (u *candidateRepository) GetGetCandidates() (*[]entitymodels.CandidateModel, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var constituenciesModel = []entitymodels.CandidateModel{}
	var dbResult = dbConnector.Find(&constituenciesModel)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &constituenciesModel, nil
}

func (u *candidateRepository) GetCandidateByCandidateId(candidateId uint) (*entitymodels.CandidateModel, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var candidateModel = entitymodels.CandidateModel{}
	var dbResult = dbConnector.Where("id = ?", candidateId).First(&candidateModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &candidateModel, nil
}

func (u *candidateRepository) GetCandidateByConstituencyId(constituencyId uint) (*[]entitymodels.CandidateModel, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var candidateModels = []entitymodels.CandidateModel{}
	var dbResult = dbConnector.Where("constituency_id = ?", constituencyId).Find(&candidateModels)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &candidateModels, nil
}

func (u *candidateRepository) InsertCandidate(candidateRegisterRequest *entitycandidate.CandidateCreateRequest) (*entitycandidate.CandidateEntity, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var candidateModel = entitymodels.CandidateModel{
		CandidateEntity: entitycandidate.CandidateEntity{
			Name:           candidateRegisterRequest.Name,
			ConstituencyId: candidateRegisterRequest.ConstituencyId,
		},
	}

	var dbResult = dbConnector.Create(&candidateModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var candidateEntity = *candidateModel.ToEntity()

	return &candidateEntity, nil
}

func (u *candidateRepository) UpdateCandidate(candidateUpdateRequest *entitycandidate.CandidateUpdateRequest) (*entitycandidate.CandidateEntity, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var candidateModel = entitymodels.CandidateModel{}
	var dbResult = dbConnector.Where("id = ?", candidateUpdateRequest.Id).First(&candidateModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	candidateModel.Name = candidateUpdateRequest.Name
	candidateModel.ConstituencyId = candidateUpdateRequest.ConstituencyId

	dbResult = dbConnector.Save(&candidateModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var candidateEntity = *candidateModel.ToEntity()

	return &candidateEntity, nil
}

func (u *candidateRepository) DeleteCandidate(candidateId uint) error {
	var dbConnector = u.dbHandler.GetConnector()

	var dbResult = dbConnector.Where("id = ?", candidateId).Delete(&entitymodels.CandidateModel{})
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}
