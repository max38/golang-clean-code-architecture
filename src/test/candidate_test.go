package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	userentity "github.com/max38/golang-clean-code-architecture/src/domain/entities/user"
)

func CreateCandidate(t *testing.T, userLoginResponse *userentity.UserLoginResponse, candidateCreateRequest entitycandidate.CandidateCreateRequest) (*entitycandidate.CandidateEntity, error) {
	// Create Request Body
	requestBody, err := json.Marshal(candidateCreateRequest)
	assert.NoError(t, err)

	// Call Create Candidate API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/candidates", baseURL), "POST", requestBody, &userLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, statusCode)

	// Convert JSON to Struct
	var candidateResponse entitycandidate.CandidateEntity
	err = json.Unmarshal(*body, &candidateResponse)
	assert.NoError(t, err)

	return &candidateResponse, nil
}

func TestGetCandidatesHandler(t *testing.T) {
	// Call Get Candidates API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/candidates", baseURL), "GET", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var candidates []entitycandidate.CandidateEntity
	err = json.Unmarshal(*body, &candidates)
	assert.NoError(t, err)
}

func TestCreateCandidateHandler(t *testing.T) {

	// Register User
	var userRegisterRequest = userentity.UserRegisterRequest{
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
	assert.NotNil(t, constituencyResponse)

	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)
	assert.Equal(t, candidateCreateRequest.Name, candidateResponse.Name)
	assert.Equal(t, candidateCreateRequest.ConstituencyId, candidateResponse.ConstituencyId)
}

func TestGetCandidatesByConstituencyIdHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = userentity.UserRegisterRequest{
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
	assert.NotNil(t, constituencyResponse)

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	// Call Get Candidates By Constituency Id API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/candidates/constituency/%d", baseURL, constituencyResponse.Id), "GET", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var candidates []entitycandidate.CandidateEntity
	err = json.Unmarshal(*body, &candidates)
	assert.NoError(t, err)
	assert.NotNil(t, candidates)
	assert.NotEmpty(t, candidates)
}

func TestGetCandidateByCandidateIdHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = userentity.UserRegisterRequest{
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
	assert.NotNil(t, constituencyResponse)

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	// Call Get Candidate By Candidate Id API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/candidates/%d", baseURL, candidateResponse.Id), "GET", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var candidate entitycandidate.CandidateEntity
	err = json.Unmarshal(*body, &candidate)
	assert.NoError(t, err)
	assert.NotNil(t, candidate)
	assert.Equal(t, candidateResponse.Id, candidate.Id)
}

func TestUpdateCandidateHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = userentity.UserRegisterRequest{
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
	assert.NotNil(t, constituencyResponse)

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	candidateForUpdateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Update Candidate
	candidateUpdateRequest := entitycandidate.CandidateUpdateRequest{
		Id:                     candidateResponse.Id,
		CandidateCreateRequest: candidateForUpdateRequest,
	}

	// Create Request Body
	requestBody, err := json.Marshal(candidateUpdateRequest)
	assert.NoError(t, err)

	// Call Update Candidate API
	body, statusCode, err := CallAPI(t, fmt.Sprintf("%s/candidates", baseURL), "PUT", requestBody, &userRegisterAndLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Convert JSON to Struct
	var candidate entitycandidate.CandidateEntity
	err = json.Unmarshal(*body, &candidate)
	assert.NoError(t, err)
	assert.NotNil(t, candidate)
	assert.Equal(t, candidateForUpdateRequest.Name, candidate.Name)
}

func TestDeleteCandidateHandler(t *testing.T) {
	// Register User
	var userRegisterRequest = userentity.UserRegisterRequest{
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
	assert.NotNil(t, constituencyResponse)

	// Create Candidate
	candidateCreateRequest := entitycandidate.CandidateCreateRequest{
		Name:           RandomString(10),
		ConstituencyId: constituencyResponse.Id,
	}

	// Call Create Candidate API
	candidateResponse, err := CreateCandidate(t, userRegisterAndLoginResponse, candidateCreateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, candidateResponse)

	// Call Delete Candidate API
	_, statusCode, err := CallAPI(t, fmt.Sprintf("%s/candidates/%d", baseURL, candidateResponse.Id), "DELETE", nil, &userRegisterAndLoginResponse.Token.Access)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
}
