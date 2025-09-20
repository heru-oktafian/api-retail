package routes

import (
	"os"

	"github.com/heru-oktafian/api-retail/controllers"
	"github.com/heru-oktafian/scafold/framework"
	"github.com/heru-oktafian/scafold/middlewares"
)

// SysAuthRoutes mengatur rute untuk autentikasi (register dan login)
func SysAuthRoutes(app *framework.Fiber) {
	JWTSecret := os.Getenv("JWT_SECRET_KEY")
	// Load Secret Key from environment
	auth := app.Group("/api")

	// auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)
	auth.Post("/logout", controllers.Logout)
	auth.Post("/set_branch", controllers.SetBranch)
	auth.Get("/profile", middlewares.Protected(JWTSecret), middlewares.AuthorizeRole("administrator", "superadmin"), controllers.GetProfile)
	// auth.Get("/list_branches", middleware.Protected(), controllers.GetBranchByUserId)
}

func CmbBranchRoutes(app *framework.Fiber) {
	app.Get("/api/branches_combo", controllers.CmbBranch)
}

func SysMenuRoutes(app *framework.Fiber) {
	JWTSecret := os.Getenv("JWT_SECRET_KEY")

	app.Get("/menus", func(c *framework.Ctx) error {
		if err := middlewares.Protected(JWTSecret)(c); err != nil {
			return err
		}
		if err := middlewares.AuthorizeRole("operator", "cashier", "finance", "superadmin", "administrator")(c); err != nil {
			return err
		}
		return controllers.GetMenus(c)
	})
}
