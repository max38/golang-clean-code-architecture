package gofiberserver

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	gofiberentities "github.com/max38/golang-clean-code-architecture/src/interface/handlers/gofiber/entities"
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
)

type middlewareErrorCode string

const (
	routerCheckError middlewareErrorCode = "middlware-001"
	jwtAuthError     middlewareErrorCode = "middlware-002"
	paramsCheckError middlewareErrorCode = "middlware-003"
	authorizeError   middlewareErrorCode = "middlware-004"
	apiKeyError      middlewareErrorCode = "middlware-005"
)

type IMiddlewares interface {
	Cors() fiber.Handler
	Logger() fiber.Handler
	RouterCheck() fiber.Handler
	JwtAuth(userUsecase userusecase.IUserUsecase) fiber.Handler
}

type middlewares struct {
}

func InitMiddlewares() IMiddlewares {
	return &middlewares{}
}

func (h *middlewares) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "http://127.0.0.1:3000",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "X-Custom-Header",
		MaxAge:           86400,
	})
}

func (h *middlewares) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} - ${method} ${path}\n",
		TimeFormat: "2006/01/02/ 15:04:05",
		TimeZone:   "UTC",
	})
}

func (h *middlewares) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return gofiberentities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(routerCheckError),
			"router not found",
		).Response()
	}
}

func (h *middlewares) JwtAuth(userUsecase userusecase.IUserUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token = strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

		var userEntity, userPermission, errorAuthentication = userUsecase.UserAuthentication(token)

		if errorAuthentication != nil {
			return gofiberentities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(jwtAuthError),
				"invalid token or token expired",
			).Response()
		}
		// Set User
		c.Locals("user", userEntity)
		c.Locals("user_permission", userPermission)
		return c.Next()
	}
}
