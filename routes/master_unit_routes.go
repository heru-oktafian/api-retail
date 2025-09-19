package routes

import (
	"os"

	"github.com/heru-oktafian/api-retail/controllers"
	"github.com/heru-oktafian/scafold/framework"
	"github.com/heru-oktafian/scafold/middlewares"
)

// MasterUnitRoutes mengatur rute-rute untuk resource unit
func MasterUnitRoutes(app *framework.App) {
	// Load Secret Key from environment
	JWTSecret := os.Getenv("JWT_SECRET_KEY")

	// Group routes that require authentication
	unitAPI := app.Group("/api/units", middlewares.Protected(JWTSecret))

	unitAPI.Get("/", controllers.GetAllUnit,
		middlewares.AuthorizeRole("operator", "cashier", "finance", "superadmin", "administrator"))

	unitAPI.Get("/:id", controllers.GetUnitByID,
		middlewares.AuthorizeRole("operator", "cashier", "finance", "superadmin", "administrator"))

	unitAPI.Post("/", controllers.CreateUnit,
		middlewares.AuthorizeRole("operator", "superadmin", "administrator"))

	unitAPI.Put("/:id", controllers.UpdateUnit,
		middlewares.AuthorizeRole("operator", "superadmin", "administrator"))

	unitAPI.Delete("/:id", controllers.DeleteUnit,
		middlewares.AuthorizeRole("superadmin", "administrator"))

	unitAPI.Get("/coba", func(c *framework.Ctx) error {
		c.SendString(200, "Halo dari Framework heru di port "+os.Getenv("PORT"))
		return nil
	})
}
