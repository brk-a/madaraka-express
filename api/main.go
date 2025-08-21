package main

import (
	"log"

	db "github.com/brk-a/madaraka_express/api/internal/db"
	router "github.com/brk-a/madaraka_express/api/router" // Import the router package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Recommended for UI interaction
	"github.com/joho/godotenv"                    // For loading .env file
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, assuming environment variables are set.")
	}

	// Connect to database
	err = db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure DB connection is closed when main exits

	app := fiber.New()

	// Add CORS middleware (adjust origins for production)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Or specify your Next.js UI origin, e.g., "http://localhost:3000"
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup API routes
	router.SetupRoutes(app) // Call the setup function from the router package

	log.Fatal(app.Listen(":8080"))
}
