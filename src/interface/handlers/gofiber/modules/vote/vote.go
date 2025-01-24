package gofibervotehandler

import (
	"github.com/gofiber/fiber/v2"

	entityuser "github.com/max38/golang-clean-code-architecture/src/domain/entities/user"
	entityvote "github.com/max38/golang-clean-code-architecture/src/domain/entities/vote"
	gofiberentities "github.com/max38/golang-clean-code-architecture/src/interface/handlers/gofiber/entities"
	candidateusecase "github.com/max38/golang-clean-code-architecture/src/usecases/candidate"
	voteusecase "github.com/max38/golang-clean-code-architecture/src/usecases/vote"
)

type voteHandlersErrorCode string

const (
	createVoteError voteHandlersErrorCode = "votes-001"
)

type IVoteHandler interface {
	CreateVote(c *fiber.Ctx) error
}

type voteHandler struct {
	voteUsecase      voteusecase.IVoteUsecase
	candidateUsecase candidateusecase.ICandidateUsecase
}

func VoteHandler(voteUsecase voteusecase.IVoteUsecase, candidateUsecase candidateusecase.ICandidateUsecase) IVoteHandler {
	return &voteHandler{
		voteUsecase:      voteUsecase,
		candidateUsecase: candidateUsecase,
	}
}

// @Summary Create Vote
// @Description Create Vote
// @Tags Vote
// @Accept application/json
// @Produce json
// @Success 201 {object} entityvote.VoteEntity
// @Failure 400 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/votes [post]
// @Param request body entityvote.VoteCreateRequest true "query params"
func (h *voteHandler) CreateVote(c *fiber.Ctx) error {
	// Parse request
	var voteCreateRequest = new(entityvote.VoteCreateRequest) // Pointer
	if err := c.BodyParser(voteCreateRequest); err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createVoteError),
			err.Error(),
		).Response()
	}

	// Get user id from claims
	var userEntity = c.Locals("user").(*entityuser.UserEntity)

	var voteCreateWithUserRequest entityvote.VoteCreateWithUserRequest
	voteCreateWithUserRequest.CandidateId = voteCreateRequest.CandidateId
	voteCreateWithUserRequest.UserId = userEntity.Id

	// Check if user already voted
	exists, checkErr := h.voteUsecase.CheckExistsVoteByUserId(voteCreateWithUserRequest.UserId)
	if checkErr != nil {
		// check exists vote error
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(createVoteError),
			checkErr.Error(),
		).Response()
	}

	if exists {
		// user already voted
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createVoteError),
			"user already voted",
		).Response()
	}

	// Get Candidate by id
	var candidate, error = h.candidateUsecase.GetCandidateByCandidateId(voteCreateWithUserRequest.CandidateId)
	if error != nil {
		// get candidate error
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createVoteError),
			error.Error(),
		).Response()
	}

	// Check candidate exists
	if candidate == nil {
		// candidate not found
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createVoteError),
			"candidate not found",
		).Response()
	}

	// Check candidate is in user constituency
	if candidate.ConstituencyId != userEntity.ConstituencyId {
		// candidate not in user constituency
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createVoteError),
			"candidate not in user constituency",
		).Response()
	}

	// Insert vote
	var voteEntity, err = h.voteUsecase.InsertVote(&voteCreateWithUserRequest)
	if err != nil {
		// insert vote error
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createVoteError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusCreated, voteEntity).Response()
}
