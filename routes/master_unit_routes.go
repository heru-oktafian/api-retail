package routes

import (
	"os"

	"github.com/heru-oktafian/api-retail/controllers"
	"github.com/heru-oktafian/scafold/framework"
	"github.com/heru-oktafian/scafold/middlewares"
	// "github.com/heru-oktafian/scafold/responses"
)

// MasterUnitRoutes mengatur rute-rute untuk resource unit
// func MasterUnitRoutes(app *framework.Fiber) {
// 	// Load Secret Key from environment
// 	JWTSecret := os.Getenv("JWT_SECRET_KEY")

// 	// Group routes that require authentication
// 	unitAPI := app.Group("/api/units")

// 	unitAPI.Get("/", controllers.GetAllUnit, middlewares.Protected(JWTSecret),
// 		middlewares.AuthorizeRole("operator", "cashier", "finance", "superadmin", "administrator"))

// 	unitAPI.Get("/:id", controllers.GetUnitByID, middlewares.Protected(JWTSecret),
// 		middlewares.AuthorizeRole("operator", "cashier", "finance", "superadmin", "administrator"))

// 	unitAPI.Post("/", controllers.CreateUnit,
// 		middlewares.AuthorizeRole("operator", "superadmin", "administrator"))

// 	unitAPI.Put("/:id", controllers.UpdateUnit,
// 		middlewares.AuthorizeRole("operator", "superadmin", "administrator"))

// 	unitAPI.Delete("/:id", controllers.DeleteUnit,
// 		middlewares.AuthorizeRole("superadmin", "administrator"))

// 	unitAPI.Get("/coba", func(c *framework.Ctx) error {
// 		return responses.JSONResponse(c, 200, "Coba endpoint hit", nil)
// 	})
// }

func MasterUnitRoutes(app *framework.Fiber) {
	// Load Secret Key from environment
	JWTSecret := os.Getenv("JWT_SECRET_KEY")

	// Grup rute yang dilindungi JWT
	unitAPI := app.Group("/api/units", middlewares.Protected(JWTSecret))

	// GET /api/units -> ambil semua unit
	unitAPI.Get("/", controllers.GetAllUnit)

	// GET /api/units/:id -> ambil unit berdasarkan ID
	unitAPI.Get("/:id", controllers.GetUnitByID)

	// POST /api/units -> buat unit baru
	unitAPI.Post("/", controllers.CreateUnit)

}

func CobaRoutes(app *framework.Fiber) {
	app.Group("/coba", func(c *framework.Ctx) error {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "Coba endpoint hit",
			"data":    nil,
		})
	})
}
