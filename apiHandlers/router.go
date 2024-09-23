package apiHandlers

import (
	"autodock-be/api"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/autodock-be/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")

	RouteMappings(group)
	DefaultMappings(defaultGroup)

}
func RouteMappings(cg fiber.Router) {
	cg.Post("/CreateWebHook",api.CreateRepoWebhook)
	cg.Post("/webhook",api.WebHookHandlerApi)
}
func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "autodock-be service is up and running", "version": "1.0"})
	})
}
