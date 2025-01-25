package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	entityuser "github.com/max38/golang-clean-code-architecture/src/domain/entities/user"
	entityvote "github.com/max38/golang-clean-code-architecture/src/domain/entities/vote"
)

func TestGetVotesByConstituencyIdHandler(t *testing.T) {
	// Register User for create constituency
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

	// Create Constituency
	constituencyCreateRequest := entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	// Call Create Constituency API
	constituencyResponse, err := CreateConstituency(t, userRegisterAndLoginResponse, constituencyCreateRequest)
	assert.NoError(t, err)
	assert.Equal(t, constituencyCreateRequest.Name, constituencyResponse.Name)
	assert.Equal(t, constituencyCreateRequest.Province, constituencyResponse.Province)

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	// Register User using the same constituency
	var userRegisterSameConstituency = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Register and login user API
	userRegisterAndLoginSameConstituencyResponse, err := RegisterAndLoginUser(t, userRegisterSameConstituency)
	assert.NoError(t, err)
	assert.NotNil(t, userRegisterAndLoginSameConstituencyResponse)

	// Login User
	userLoginSameConstituencyRequest := entityuser.UserLoginRequest{
		Email:    userRegisterSameConstituency.Email,
		Password: userRegisterSameConstituency.Password,
	}

	// Call Login User API
	userLoginSameConstituencyResponse, err := LoginUser(t, userLoginSameConstituencyRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userLoginSameConstituencyResponse)

	// Create Vote Request
	voteCreateRequest := entityvote.VoteCreateRequest{
		CandidateId: candidateResponse.Id,
	}

	// Create Request Body
	requestBody, err := json.Marshal(voteCreateRequest)
	assert.NoError(t, err)

	// Call Create Vote API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/votes", baseURL), "POST", requestBody, &userLoginSameConstituencyResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, statusCode)

	// Convert JSON to Struct
	var voteResponse entityvote.VoteEntity
	err = json.Unmarshal(*body, &voteResponse)
	assert.NoError(t, err)

	// Call Get Votes By Constituency Id API
	body, statusCode, err = CallAPI(t, fmt.Sprintf("%s/votes/constituency/%d", baseURL, constituencyResponse.Id), "GET", nil, &userRegisterAndLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var votesByConstituencyIdResponse []entityvote.GetVotesByConstituencyIdResponse
	err = json.Unmarshal(*body, &votesByConstituencyIdResponse)
	assert.NoError(t, err)
}

func TestCreateVoteWithConstituencyCorrect(t *testing.T) {
	// Register User for create constituency
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

	// Create Constituency
	constituencyCreateRequest := entityconstituency.ConstituencyCreateRequest{
		Name:     RandomString(10),
		Province: RandomString(10),
	}

	// Call Create Constituency API
	constituencyResponse, err := CreateConstituency(t, userRegisterAndLoginResponse, constituencyCreateRequest)
	assert.NoError(t, err)
	assert.Equal(t, constituencyCreateRequest.Name, constituencyResponse.Name)
	assert.Equal(t, constituencyCreateRequest.Province, constituencyResponse.Province)

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	// Register User using the same constituency
	var userRegisterSameConstituency = entityuser.UserRegisterRequest{
		Email:          RandomEmail(),
		Password:       RandomString(10),
		FirstName:      RandomString(6),
		LastName:       RandomString(6),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Register and login user API
	userRegisterAndLoginSameConstituencyResponse, err := RegisterAndLoginUser(t, userRegisterSameConstituency)
	assert.NoError(t, err)
	assert.NotNil(t, userRegisterAndLoginSameConstituencyResponse)

	// Login User
	userLoginSameConstituencyRequest := entityuser.UserLoginRequest{
		Email:    userRegisterSameConstituency.Email,
		Password: userRegisterSameConstituency.Password,
	}

	// Call Login User API
	userLoginSameConstituencyResponse, err := LoginUser(t, userLoginSameConstituencyRequest)
	assert.NoError(t, err)
	assert.NotNil(t, userLoginSameConstituencyResponse)

	// Create Vote Request
	voteCreateRequest := entityvote.VoteCreateRequest{
		CandidateId: candidateResponse.Id,
	}

	// Create Request Body
	requestBody, err := json.Marshal(voteCreateRequest)
	assert.NoError(t, err)

	// Call Create Vote API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/votes", baseURL), "POST", requestBody, &userLoginSameConstituencyResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, statusCode)

	// Convert JSON to Struct
	var voteResponse entityvote.VoteEntity
	err = json.Unmarshal(*body, &voteResponse)
	assert.NoError(t, err)
}

func TestCreateVoteWithConstituencyIncorrect(t *testing.T) {

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

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	// Create Vote Request
	voteCreateRequest := entityvote.VoteCreateRequest{
		CandidateId: candidateResponse.Id,
	}

	// Create Request Body
	requestBody, err := json.Marshal(voteCreateRequest)
	assert.NoError(t, err)

	// Call Create Vote API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/votes", baseURL), "POST", requestBody, &userRegisterAndLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)

	// Convert JSON to Struct
	var voteResponse entityvote.VoteEntity
	err = json.Unmarshal(*body, &voteResponse)
	assert.NoError(t, err)
}
