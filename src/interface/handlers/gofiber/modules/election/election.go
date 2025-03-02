package election

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entityelection "github.com/max38/golang-clean-code-architecture/src/domain/entities/election"
	gofiberentities "github.com/max38/golang-clean-code-architecture/src/interface/handlers/gofiber/entities"
	electionusecase "github.com/max38/golang-clean-code-architecture/src/usecases/election"
)

type electionHandlersErrorCode string

const (
	addResultError electionHandlersErrorCode = "election-001"
	getResultError electionHandlersErrorCode = "election-002"
)

type IElectionHandler interface {
	NewResult(c *fiber.Ctx) error
	GetResult(c *fiber.Ctx) error
	Scoreboard(c *fiber.Ctx) error
}

type electionHandler struct {
	electionUsecase electionusecase.IElectionUsecase
}

func ElectionHandler(electionUsecase electionusecase.IElectionUsecase) IElectionHandler {
	return &electionHandler{
		electionUsecase: electionUsecase,
	}
}

// @Summary NewResult Election
// @Description NewResult Election
// @Tags Election
// @Accept application/json
// @Produce json
// @Success 201 {object} entityelection.ElectionNewresultRequest
// @Failure 400 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/election/result [post]
// @Param request body entityelection.ElectionNewresultRequest true "query params"
func (h *electionHandler) NewResult(c *fiber.Ctx) error {
	// Parse request
	electionResultRequest := new(entityelection.ElectionNewresultRequest)
	if err := c.BodyParser(electionResultRequest); err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(addResultError),
			err.Error(),
		).Response()
	}

	err := h.electionUsecase.AddResult(*electionResultRequest)
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(addResultError),
			err.Error(),
		).Response()
	}

	return gofiberentities.NewResponse(c).Success(fiber.StatusCreated, electionResultRequest).Response()
}

// @Summary GetResult Election
// @Description NewResult Election
// @Tags Election
// @Accept application/json
// @Produce json
// @Success 200 {object} entityelection.ElectionNewresultRequest
// @Failure 400 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/election/result/{id} [get]
func (h *electionHandler) GetResult(c *fiber.Ctx) error {
	entityId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(getResultError),
			err.Error(),
		).Response()
	}

	responseData, err := h.electionUsecase.GetResult(int(entityId))
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(getResultError),
			err.Error(),
		).Response()
	}

	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, responseData).Response()
}

// @Summary GetResult Election
// @Description Scoreboard Election
// @Tags Election
// @Accept application/json
// @Produce json
// @Success 200 {object} entityelection.Scoreboard
// @Failure 400 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/election/scoreboard [get]
func (h *electionHandler) Scoreboard(c *fiber.Ctx) error {
	responseData, err := h.electionUsecase.GetScoreboard()
	if err != nil {
		gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(getResultError),
			err.Error(),
		).Response()
	}

	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, responseData).Response()
}
