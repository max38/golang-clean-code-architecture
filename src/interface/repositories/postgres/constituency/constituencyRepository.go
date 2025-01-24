package postgresconstituencyrepository

import (
	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
	repositories "github.com/max38/golang-clean-code-architecture/src/domain/repositories"
	postgreshandler "github.com/max38/golang-clean-code-architecture/src/infrastructure/database/postgres"
)

type constituencyRepository struct {
	dbHandler postgreshandler.IPostgresHandler
}

func ConstituencyRepository() repositories.IConstituencyRepository {
	var dbHandler = postgreshandler.PostgresHandler()
	return &constituencyRepository{
		dbHandler: dbHandler,
	}
}

func (u *constituencyRepository) GetConstituencies() (*[]entitymodels.ConstituencyModel, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var constituenciesModel = []entitymodels.ConstituencyModel{}
	var dbResult = dbConnector.Find(&constituenciesModel)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &constituenciesModel, nil
}

func (u *constituencyRepository) GetConstituencyByConstituencyId(constituencyId uint) (*entitymodels.ConstituencyModel, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var constituencyModel = entitymodels.ConstituencyModel{}
	var dbResult = dbConnector.Where("id = ?", constituencyId).First(&constituencyModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &constituencyModel, nil
}

func (u *constituencyRepository) InsertConstituency(constituencyRegisterRequest *entityconstituency.ConstituencyCreateRequest) (*entityconstituency.ConstituencyEntity, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var constituencyModel = entitymodels.ConstituencyModel{
		ConstituencyEntity: entityconstituency.ConstituencyEntity{
			Name:     constituencyRegisterRequest.Name,
			Province: constituencyRegisterRequest.Province,
		},
	}

	var dbResult = dbConnector.Create(&constituencyModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var constituencyEntity = *constituencyModel.ToEntity()

	return &constituencyEntity, nil
}

func (u *constituencyRepository) UpdateConstituency(constituencyUpdateRequest *entityconstituency.ConstituencyUpdateRequest) (*entityconstituency.ConstituencyEntity, error) {
	var dbConnector = u.dbHandler.GetConnector()

	var constituencyModel = entitymodels.ConstituencyModel{}
	var dbResult = dbConnector.Where("id = ?", constituencyUpdateRequest.Id).First(&constituencyModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	constituencyModel.Name = constituencyUpdateRequest.Name
	constituencyModel.Province = constituencyUpdateRequest.Province

	dbResult = dbConnector.Save(&constituencyModel)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	var constituencyEntity = *constituencyModel.ToEntity()

	return &constituencyEntity, nil
}

func (u *constituencyRepository) DeleteConstituency(constituencyId uint) error {
	var dbConnector = u.dbHandler.GetConnector()

	var dbResult = dbConnector.Where("id = ?", constituencyId).Delete(&entitymodels.ConstituencyModel{})
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}
