package gofiberconstituencyhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	entityconstituency "github.com/max38/golang-clean-code-architecture/src/domain/entities/constituency"
	gofiberentities "github.com/max38/golang-clean-code-architecture/src/interface/handlers/gofiber/entities"
	constituencyusecase "github.com/max38/golang-clean-code-architecture/src/usecases/constituency"
)

type constituencyHandlersErrorCode string

const (
	createConstituencyError constituencyHandlersErrorCode = "constituencys-001"
)

type IConstituencyHandler interface {
	GetConstituencies(c *fiber.Ctx) error
	GetConstituencyByConstituencyId(c *fiber.Ctx) error
	CreateConstituency(c *fiber.Ctx) error
	UpdateConstituency(c *fiber.Ctx) error
	DeleteConstituency(c *fiber.Ctx) error
}

type constituencyHandler struct {
	constituencyUsecase constituencyusecase.IConstituencyUsecase
}

func ConstituencyHandler(constituencyUsecase constituencyusecase.IConstituencyUsecase) IConstituencyHandler {
	return &constituencyHandler{
		constituencyUsecase: constituencyUsecase,
	}
}

// GetConstituencies retrieves all constituencies.
// @Summary Get Constituencies
// @Description Get Constituencies
// @Tags Constituency
// @Accept application/json
// @Produce json
// @Success 200 {array} entityconstituency.ConstituencyEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/constituencies [get]
func (h *constituencyHandler) GetConstituencies(c *fiber.Ctx) error {
	// Pass through Use Case
	constituencies, err := h.constituencyUsecase.GetConstituencies()
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, constituencies).Response()
}

// GetConstituencyByConstituencyId retrieves the constituency's profile.
// @Summary Get Constituency By Constituency Id
// @Description Get Constituency By Constituency Id
// @Tags Constituency
// @Accept application/json
// @Produce json
// @Success 200 {object} entityconstituency.ConstituencyEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/constituencies/{id} [get]
func (h *constituencyHandler) GetConstituencyByConstituencyId(c *fiber.Ctx) error {
	// Parse request
	var constituencyId = c.Params("id")

	// Pass thorugh Use Case
	constituencyIdUint, parseErr := strconv.ParseUint(constituencyId, 10, 32)
	if parseErr != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			"Invalid constituency ID",
		).Response()
	}
	var constituencyEntity, err = h.constituencyUsecase.GetConstituencyByConstituencyId(uint(constituencyIdUint))
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, constituencyEntity).Response()
}

// @Summary Create Constituency
// @Description Create Constituency
// @Tags Constituency
// @Accept application/json
// @Produce json
// @Success 201 {object} entityconstituency.ConstituencyEntity
// @Failure 400 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/constituencies [post]
// @Param request body entityconstituency.ConstituencyCreateRequest true "query params"
func (h *constituencyHandler) CreateConstituency(c *fiber.Ctx) error {
	// Parse request
	var constituencyCreateRequest = new(entityconstituency.ConstituencyCreateRequest) // Pointer
	if err := c.BodyParser(constituencyCreateRequest); err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	// Pass thorugh Use Case
	var constituencyEntity, err = h.constituencyUsecase.CreateConstituency(constituencyCreateRequest)
	if err != nil {
		// Email has been used
		// insert constituency error
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusCreated, constituencyEntity).Response()
}

// UpdateConstituency updates.
// @Summary Update Constituency
// @Description Update Constituency
// @Tags Constituency
// @Accept application/json
// @Produce json
// @Success 200 {object} entityconstituency.ConstituencyEntity
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/constituencies [put]
func (h *constituencyHandler) UpdateConstituency(c *fiber.Ctx) error {
	// Parse request
	var constituencyUpdateRequest = new(entityconstituency.ConstituencyUpdateRequest) // Pointer
	if err := c.BodyParser(constituencyUpdateRequest); err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	var constituencyEntity, err = h.constituencyUsecase.UpdateConstituency(constituencyUpdateRequest)
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, constituencyEntity).Response()
}

// DeleteConstituency deletes a constituency.
// @Summary Delete Constituency
// @Description Delete Constituency
// @Tags Constituency
// @Accept application/json
// @Produce json
// @Success 200 {object} gofiberentities.SuccessResponseType
// @Failure 401 {object} gofiberentities.ErrorResponseType
// @Router /api/v1/constituencies/{id} [delete]
func (h *constituencyHandler) DeleteConstituency(c *fiber.Ctx) error {
	// Parse request
	var constituencyId = c.Params("id")

	// Pass through Use Case
	constituencyIdUint, err := strconv.ParseUint(constituencyId, 10, 32)
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			"Invalid constituency ID",
		).Response()
	}
	err = h.constituencyUsecase.DeleteConstituency(uint(constituencyIdUint))
	if err != nil {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(createConstituencyError),
			err.Error(),
		).Response()
	}

	// Return response
	return gofiberentities.NewResponse(c).Success(fiber.StatusOK, "Constituency deleted successfully").Response()
}
