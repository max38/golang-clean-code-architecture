package gofibercandidatehandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
	gofiberentities "github.com/max38/golang-clean-code-architecture/src/interface/handlers/gofiber/entities"
	candidateusecase "github.com/max38/golang-clean-code-architecture/src/usecases/candidate"
)

type candidateHandlersErrorCode string

const (
	createCandidateError candidateHandlersErrorCode = "candidates-001"
)

type ICandidateHandler interface {
	GetCandidates(c *fiber.Ctx) error
	GetCandidateByCandidateId(c *fiber.Ctx) error
	GetCandidatesByConstituencyId(c *fiber.Ctx) error
	CreateCandidate(c *fiber.Ctx) error
	UpdateCandidate(c *fiber.Ctx) error
	DeleteCandidate(c *fiber.Ctx) error
}

type candidateHandler struct {
	candidateUsecase candidateusecase.ICandidateUsecase
}

func CandidateHandler(candidateUsecase candidateusecase.ICandidateUsecase) ICandidateHandler {
	return &candidateHandler{
		candidateUsecase: candidateUsecase,
	}
}

// GetCandidates retrieves all candidates.
// @Summary Get Candidates
// @Description Get Candidates
// @Tags Candidate
// @Accept application/json
// @Produce json
// @Success 200 {array} entitycandidate.CandidateEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/candidates [get]
func (h *candidateHandler) GetCandidates(c *fiber.Ctx) error {
	// Pass through Use Case
	candidates, err := h.candidateUsecase.GetGetCandidates()
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, candidates).Response()
}

// GetCandidateByCandidateId retrieves the candidate's profile.
// @Summary Get Candidate By Candidate Id
// @Description Get Candidate By Candidate Id
// @Tags Candidate
// @Accept application/json
// @Produce json
// @Success 200 {object} entitycandidate.CandidateEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/candidates/{id} [get]
func (h *candidateHandler) GetCandidateByCandidateId(c *fiber.Ctx) error {
	// Parse request
	var candidateId = c.Params("id")

	// Pass thorugh Use Case
	candidateIdUint, parseErr := strconv.ParseUint(candidateId, 10, 32)
	if parseErr != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			"Invalid candidate ID",
		).Response()
	}
	var candidateEntity, err = h.candidateUsecase.GetCandidateByCandidateId(uint(candidateIdUint))
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, candidateEntity).Response()
}

// GetCandidatesByConstituencyId retrieves candidates by constituency ID.
// @Summary Get Candidates By Constituency Id
// @Description Get Candidates By Constituency Id
// @Tags Candidate
// @Accept application/json
// @Produce json
// @Success 200 {array} entitycandidate.CandidateEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/candidates/{id}/candidates [get]
func (h *candidateHandler) GetCandidatesByConstituencyId(c *fiber.Ctx) error {
	// Parse request
	var constituencyId = c.Params("id")

	// Pass through Use Case
	constituencyIdUint, parseErr := strconv.ParseUint(constituencyId, 10, 32)
	if parseErr != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			"Invalid constituency ID",
		).Response()
	}

	var candidates, err = h.candidateUsecase.GetCandidateByConstituencyId(uint(constituencyIdUint))
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, candidates).Response()
}

// @Summary Create Candidate
// @Description Create Candidate
// @Tags Candidate
// @Accept application/json
// @Produce json
// @Success 201 {object} entitycandidate.CandidateEntity
// @Failure 400 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/candidates [post]
// @Param request body entitycandidate.CandidateCreateRequest true "query params"
func (h *candidateHandler) CreateCandidate(c *fiber.Ctx) error {
	// Parse request
	var candidateCreateRequest = new(entitycandidate.CandidateCreateRequest) // Pointer
	if err := c.BodyParser(candidateCreateRequest); err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Pass thorugh Use Case
	var candidateEntity, err = h.candidateUsecase.CreateCandidate(candidateCreateRequest)
	if err != nil {
		// Email has been used
		// insert candidate error
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusCreated, candidateEntity).Response()
}

// UpdateCandidate updates.
// @Summary Update Candidate
// @Description Update Candidate
// @Tags Candidate
// @Accept application/json
// @Produce json
// @Success 200 {object} entitycandidate.CandidateEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/candidates [put]
func (h *candidateHandler) UpdateCandidate(c *fiber.Ctx) error {
	// Parse request
	var candidateUpdateRequest = new(entitycandidate.CandidateUpdateRequest) // Pointer
	if err := c.BodyParser(candidateUpdateRequest); err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	var candidateEntity, err = h.candidateUsecase.UpdateCandidate(candidateUpdateRequest)
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, candidateEntity).Response()
}

// DeleteCandidate deletes a candidate.
// @Summary Delete Candidate
// @Description Delete Candidate
// @Tags Candidate
// @Accept application/json
// @Produce json
// @Success 200 {object} gofiberentities.SuccessResponseType
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/candidates/{id} [delete]
func (h *candidateHandler) DeleteCandidate(c *fiber.Ctx) error {
	// Parse request
	var candidateId = c.Params("id")

	// Pass through Use Case
	candidateIdUint, err := strconv.ParseUint(candidateId, 10, 32)
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			"Invalid candidate ID",
		).Response()
	}
	err = h.candidateUsecase.DeleteCandidate(uint(candidateIdUint))
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createCandidateError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, "Candidate deleted successfully").Response()
}
