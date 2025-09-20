package main

import (
	log "log"
	os "os"
	"strconv"

	routes "github.com/heru-oktafian/api-retail/routes"
	config "github.com/heru-oktafian/scafold/config"
	env "github.com/heru-oktafian/scafold/env"
	framework "github.com/heru-oktafian/scafold/framework"
	middlewares "github.com/heru-oktafian/scafold/middlewares"
	utils "github.com/heru-oktafian/scafold/utils"
)

func main() {
	// Initialize timezone
	utils.InitTimezone()

	// Load .env file
	env.Load(".env")

	// Load Secret Key from environment
	JWTSecret := os.Getenv("JWT_SECRET_KEY")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET_KEY not set in .env file")
	}

	// Initialize database connection
	config.KoneksiPG(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Initialize Redis connection
	redisDB := 0
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		if dbInt, err := strconv.Atoi(dbStr); err == nil {
			redisDB = dbInt
		}
	}

	// Default to DB 0 if REDIS_DB is not set or invalid
	config.KoneksiRedis(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), redisDB)

	// Get port from environment
	serverPort := os.Getenv("PORT")

	// Start the application
	app := framework.New()

	app.Use(middlewares.CORS())
	app.Use(middlewares.Logger())

	// Routes
	routes.SysAuthRoutes(app)
	routes.CmbBranchRoutes(app)
	routes.SysMenuRoutes(app)
	routes.MasterUnitRoutes(app)

	// Start listening on the specified port
	app.Listen(":" + serverPort)
}
