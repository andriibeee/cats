package rest

import (
	"cats/internal/domain/usecase"
	"cats/internal/rest/handlers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	cuc *usecase.CatsUseCase
	muc *usecase.MissionsUseCase
	app *fiber.App
}

func New(
	cuc *usecase.CatsUseCase,
	muc *usecase.MissionsUseCase,
) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Status(code)

			return ctx.JSON(map[string]interface{}{
				"error": err.Error(),
			})
		},
	})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(logger.New())

	validate := validator.New(validator.WithRequiredStructEnabled())

	cr := handlers.NewCatsHandlers(cuc, validate)
	mc := handlers.NewMissionsHandlers(muc, validate)

	app.Route("/cats", cr.Routes)
	app.Route("/missions", mc.Routes)

	return &Server{
		app: app,
	}

}

func (s *Server) Run(port string) error {
	return s.app.Listen(port)
}
