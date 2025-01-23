package main

import (
	"boilerplate/api"
	"boilerplate/api/handler/user"
	"boilerplate/api/repository"
	"boilerplate/api/service/user/command"
	"boilerplate/api/service/user/query"
	"boilerplate/lib/database"
	env "boilerplate/lib/environment"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

func main() {
	// init log file
	// Set up lumberjack for log rotation
	logFile := &lumberjack.Logger{
		Filename:   "info.log", // Log file name
		MaxSize:    10,         // Maximum size in MB before rotation
		MaxBackups: 5,          // Maximum number of old log files to retain
		MaxAge:     30,         // Maximum age (days) before log files are deleted
		Compress:   true,       // Enable gzip compression for old log files
	}
	// output log to file & console
	// Set output to both file and console
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)

	// Load environment variables
	env.New(0) // Pass 0 if the env file is in the current directory

	// Get DSN from environment variables
	dsn := env.GetString(env.DsnKey)
	if dsn == "" {
		log.Fatal("DB_DSN not set in environment")
	}

	// Initialize Database
	database.InitDatabaseWithDSN(dsn)

	// Dependency Injection
	// Repos
	userRepo := repository.NewUserRepository(database.DB)

	// Services
	// user query
	getUserByIdService := query.NewGetUserByIdService(userRepo)

	// user command
	createUserService := command.NewCreateUserService(userRepo)

	// Handlers
	userHandler := user.NewUserHandler(getUserByIdService, createUserService)

	// Set up Fiber
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: env.GetString("cors.allow_origins"), // Chỉ cho phép các domain cụ thể
		AllowHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, " +
			"accept, origin, Cache-Control, X-Requested-With", // Header được phép
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS", // Phương thức HTTP được phép
		AllowCredentials: true,                              // Cho phép gửi cookies
	}))
	api.SetupRoutes(app, userHandler)

	// Get the port from the environment
	port := env.GetString(env.ServicePort)
	if port == "" {
		port = "8080" // Default port if not set
	}

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
