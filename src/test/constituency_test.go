package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	entityuser "github.com/max38/golang-clean-code-architecture/src/domain/entities/user"
	"github.com/stretchr/testify/assert"
)

func TestGetConstituenciesHandler(t *testing.T) {
	// Call Get Constituencies API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/constituencies", baseURL), "GET", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var constituencies []entityconstituency.ConstituencyEntity
	err = json.Unmarshal(*body, &constituencies)
	assert.NoError(t, err)
}

func CreateConstituency(t *testing.T, userLoginResponse *entityuser.UserLoginResponse, constituencyCreateRequest entityconstituency.ConstituencyCreateRequest) (*entityconstituency.ConstituencyEntity, error) {
	// Create Request Body
	requestBody, err := json.Marshal(constituencyCreateRequest)
	assert.NoError(t, err)

	// Call Create Constituency API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/constituencies", baseURL), "POST", requestBody, &userLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, statusCode)

	// Convert JSON to Struct
	var constituencyResponse entityconstituency.ConstituencyEntity
	err = json.Unmarshal(*body, &constituencyResponse)
	assert.NoError(t, err)

	return &constituencyResponse, nil
}

func TestCreateConstituencyHandler(t *testing.T) {

	// Register User
	var userRegisterRequest = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: 1,
	}

	// Call Register and login user API
	userRegisterAndLoginResponse, err := RegisterAndLoginUser(t, userRegisterRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userRegisterAndLoginResponse)

	constituencyCreateRequest := entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	// Call Create Constituency API
	constituencyResponse, err := CreateConstituency(t, userRegisterAndLoginResponse, constituencyCreateRequest)
	assert.NoError(t, err)
	assert.Equal(t, constituencyCreateRequest.Name, constituencyResponse.Name)
	assert.Equal(t, constituencyCreateRequest.Province, constituencyResponse.Province)
}

func TestGetConstituencyByConstituencyIdHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: 1,
	}

	// Call Register and login user API
	userRegisterAndLoginResponse, err := RegisterAndLoginUser(t, userRegisterRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userRegisterAndLoginResponse)

	constituencyCreateRequest := entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	// Call Create Constituency API
	constituencyResponse, err := CreateConstituency(t, userRegisterAndLoginResponse, constituencyCreateRequest)
	assert.NoError(t, err)
	assert.Equal(t, constituencyCreateRequest.Name, constituencyResponse.Name)
	assert.Equal(t, constituencyCreateRequest.Province, constituencyResponse.Province)

	// Call Get Constituency By Constituency Id API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/constituencies/%d", baseURL, constituencyResponse.Id), "GET", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Read Response Body
	var constituency entityconstituency.ConstituencyEntity
	err = json.Unmarshal(*body, &constituency)
	assert.NoError(t, err)
	assert.NotNil(t, constituency)
	assert.Equal(t, constituencyResponse.Id, constituency.Id)
	assert.Equal(t, constituencyResponse.Name, constituency.Name)
	assert.Equal(t, constituencyResponse.Province, constituency.Province)
}

func TestUpdateConstituencyHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: 1,
	}

	// Call Register and login user API
	userRegisterAndLoginResponse, err := RegisterAndLoginUser(t, userRegisterRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userRegisterAndLoginResponse)

	constituencyCreateRequest := entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	// Call Create Constituency API
	constituencyResponse, err := CreateConstituency(t, userRegisterAndLoginResponse, constituencyCreateRequest)
	assert.NoError(t, err)
	assert.Equal(t, constituencyCreateRequest.Name, constituencyResponse.Name)
	assert.Equal(t, constituencyCreateRequest.Province, constituencyResponse.Province)

	var constituencyForUpdate = entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	var constituencyUpdateRequest = entityconstituency.ConstituencyUpdateRequest{
		Id:                        constituencyResponse.Id,
		ConstituencyCreateRequest: constituencyForUpdate,
	}

	// Create Request Body
	requestBody, err := json.Marshal(constituencyUpdateRequest)
	assert.NoError(t, err)

	// Call Update Constituency API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/constituencies", baseURL), "PUT", requestBody, &userRegisterAndLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var constituencyUpdatedResponse entityconstituency.ConstituencyEntity
	err = json.Unmarshal(*body, &constituencyUpdatedResponse)
	assert.NoError(t, err)

	assert.Equal(t, constituencyForUpdate.Name, constituencyUpdatedResponse.Name)
	assert.Equal(t, constituencyForUpdate.Province, constituencyUpdatedResponse.Province)
}

func TestDeleteConstituencyHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: 1,
	}

	// Call Register and login user API
	userRegisterAndLoginResponse, err := RegisterAndLoginUser(t, userRegisterRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userRegisterAndLoginResponse)

	constituencyCreateRequest := entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	// Call Create Constituency API
	constituencyResponse, err := CreateConstituency(t, userRegisterAndLoginResponse, constituencyCreateRequest)
	assert.NoError(t, err)
	assert.Equal(t, constituencyCreateRequest.Name, constituencyResponse.Name)
	assert.Equal(t, constituencyCreateRequest.Province, constituencyResponse.Province)

	// Call Delete Constituency API
	_, statusCode, err := CallAPI(t, fmt.Sprintf("%s/constituencies/%d", baseURL, constituencyResponse.Id), "DELETE", nil, &userRegisterAndLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
}
