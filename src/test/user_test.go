package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	entityuser "github.com/max38/golang-clean-code-architecture/src/domain/entities/user"
	"github.com/stretchr/testify/assert"
)

func RegisterUser(t *testing.T, userRegisterRequest entityuser.UserRegisterRequest) (*entityuser.UserEntity, error) {
	// Convert to JSON
	requestBody, err := json.Marshal(userRegisterRequest)
	assert.NoError(t, err)

	// Call Register User API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/users/register", baseURL), "POST", requestBody, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, statusCode)

	// Convert JSON to Map
	var user *entityuser.UserEntity
	err = json.Unmarshal(*body, &user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userRegisterRequest.Email, user.Email)
	assert.Equal(t, userRegisterRequest.FirstName, user.FirstName)
	assert.Equal(t, userRegisterRequest.LastName, user.LastName)
	assert.Equal(t, userRegisterRequest.ConstituencyId, user.ConstituencyId)

	return user, err
}

func LoginUser(t *testing.T, userLoginRequest entityuser.UserLoginRequest) (*entityuser.UserLoginResponse, error) {
	// Convert to JSON
	requestBody, err := json.Marshal(userLoginRequest)
	assert.NoError(t, err)

	// Call Login User API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/users/login", baseURL), "POST", requestBody, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Map
	var user *entityuser.UserLoginResponse
	err = json.Unmarshal(*body, &user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.User)
	assert.NotNil(t, user.Token)
	assert.NotEmpty(t, user.Token.Access)
	assert.NotEmpty(t, user.Token.Refresh)
	assert.Equal(t, userLoginRequest.Email, user.User.Email)

	return user, err
}

func RegisterAndLoginUser(t *testing.T, userRegisterRequest entityuser.UserRegisterRequest) (*entityuser.UserLoginResponse, error) {
	// Call Register User API
	_, err := RegisterUser(t, userRegisterRequest)
	assert.NoError(t, err)

	// Login User
	var userLoginRequest = entityuser.UserLoginRequest{
		Email:    userRegisterRequest.Email,
		Password: userRegisterRequest.Password,
	}

	// Call Login User API
	userLoginResponse, err := LoginUser(t, userLoginRequest)
	assert.NoError(t, err)

	return userLoginResponse, nil
}

func TestRegisterUser(t *testing.T) {
	// Create Request Body
	var userRegisterRequest = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: 1,
	}

	// Call API
	userRegisterResponse, err := RegisterUser(t, userRegisterRequest)

	// Check Response
	assert.NoError(t, err)
	assert.Equal(t, userRegisterRequest.Email, userRegisterResponse.Email)
	assert.Equal(t, userRegisterRequest.FirstName, userRegisterResponse.FirstName)
	assert.Equal(t, userRegisterRequest.LastName, userRegisterResponse.LastName)
	assert.Equal(t, userRegisterRequest.ConstituencyId, userRegisterResponse.ConstituencyId)
}

func TestLoginUser(t *testing.T) {
	// Create Request Body
	var userRegisterRequest = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: 1,
	}

	// Register User
	_, err := RegisterUser(t, userRegisterRequest)
	assert.NoError(t, err)

	// Create Request Body
	var userLoginRequest = entityuser.UserLoginRequest{
		Email:    userRegisterRequest.Email,
		Password: userRegisterRequest.Password,
	}

	// Call API
	userLoginResponse, err := LoginUser(t, userLoginRequest)

	// Check Response
	assert.NoError(t, err)
	assert.NotEmpty(t, userLoginResponse)
	assert.NotEmpty(t, userLoginResponse.User)
	assert.NotEmpty(t, userLoginResponse.Token)
	assert.Equal(t, userRegisterRequest.Email, userLoginResponse.User.Email)
	assert.Equal(t, userRegisterRequest.FirstName, userLoginResponse.User.FirstName)
	assert.Equal(t, userRegisterRequest.LastName, userLoginResponse.User.LastName)
	assert.Equal(t, userRegisterRequest.ConstituencyId, userLoginResponse.User.ConstituencyId)
}
