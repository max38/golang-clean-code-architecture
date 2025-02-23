package crudusecase

import (
	"fmt"

	entities "github.com/max38/golang-clean-code-architecture/src/domain/entities/crud"
	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
	crudrepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/crud"
)

type ICRUDUsecase interface {
	// Get
	Retrive(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, id string) (interface{}, error)

	// List
	RetriveList(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, pageNumber int, pageSize int, filter map[string]interface{}) (*entities.DataListResponse, error)

	// Create
	Create(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, entityModel interface{}) (interface{}, error)

	// Update
	Update(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, entityModel interface{}) (interface{}, error)

	// Delete
	Delete(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, entityModel interface{}) error

	// Describe
	DescribeDataSource(entityModelSlug string, schema string) (map[string]interface{}, error)
}

type crudUsecase struct {
}

func CrudUsecase() ICRUDUsecase {
	return &crudUsecase{}
}

func (c *crudUsecase) Retrive(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, id string) (interface{}, error) {
	var crudRepository, errorInit = crudrepository.GetCRUDRepository(entityModelSlug)
	if errorInit != nil {
		return nil, errorInit
	}
	return crudRepository.GetOneById(id)
}

func (c *crudUsecase) RetriveList(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, pageNumber int, pageSize int, filter map[string]interface{}) (*entities.DataListResponse, error) {
	var crudRepository, errorInit = crudrepository.GetCRUDRepository(entityModelSlug)

	if errorInit != nil {
		return nil, errorInit
	}
	return crudRepository.GetList(pageNumber, pageSize, filter)
}
func (c *crudUsecase) Create(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, entityModel interface{}) (interface{}, error) {
	crudRepository, errorInit := crudrepository.GetCRUDRepository(entityModelSlug)
	if errorInit != nil {
		return nil, errorInit
	}

	if crudDataModel, ok := entityModel.(*entities.ICRUDDataModel); ok {
		return crudRepository.Create(crudDataModel)
	} else {

		return nil, fmt.Errorf("invalid type: expected *entities.ICRUDDataModel, got %T", entityModel)
	}
}

func (c *crudUsecase) Update(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, entityModel interface{}) (interface{}, error) {
	var crudRepository, errorInit = crudrepository.GetCRUDRepository(entityModelSlug)
	if errorInit != nil {
		return nil, errorInit
	}

	switch v := entityModel.(type) {
	case *entities.ICRUDDataModel:
		return crudRepository.Update(v)
	default:
		return nil, fmt.Errorf("invalid entity model type")
	}
}

func (c *crudUsecase) Delete(userPermissionModel *entitymodels.UserPermissionModel, entityModelSlug string, entityModel interface{}) error {
	var crudRepository, errorInit = crudrepository.GetCRUDRepository(entityModelSlug)
	if errorInit != nil {
		return errorInit
	}
	switch v := entityModel.(type) {
	case *entities.ICRUDDataModel:
		return crudRepository.Delete(v)
	default:
		return fmt.Errorf("invalid entity model type for delete")
	}
}

func (c *crudUsecase) DescribeDataSource(entityModelSlug string, schema string) (map[string]interface{}, error) {
	var crudRepository, errorInit = crudrepository.GetCRUDRepository(entityModelSlug)

	if errorInit != nil {
		return nil, errorInit
	}
	return crudRepository.DescribeDataSource(schema), nil
}
