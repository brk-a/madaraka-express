package router

import (
	handlers "github.com/brk-a/madaraka_express/api/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
	// You might want to add middleware like JWT auth here later
	// "github.com/brk-a/madaraka_express/api/internal/middleware"
)

// SetupRoutes wires up all the API handlers to their respective routes
func SetupRoutes(app *fiber.App) {
	// Create a base API group
	api := app.Group("/api") // All routes under this group will be prefixed with /api

	// User Routes
	userRoutes := api.Group("/users") // e.g., /api/users
	userRoutes.Post("/", handlers.RegisterUser)
	// userRoutes.Post("/login", handlers.LoginUser) // Assuming you'll add a login handler
	// userRoutes.Get("/:id", middleware.AuthRequired, handlers.GetUser) // Example with auth middleware

	// Trip Routes
	tripRoutes := api.Group("/trips") // e.g., /api/trips
	tripRoutes.Post("/", handlers.CreateTrip)
	tripRoutes.Get("/:id", handlers.GetTrip) // Assuming you'll implement GetTrip
	// tripRoutes.Get("/", handlers.GetAllTrips) // Assuming you'll implement GetAllTrips

	// Booking Routes
	bookingRoutes := api.Group("/bookings") // e.g., /api/bookings
	bookingRoutes.Post("/", handlers.CreateBooking)
	bookingRoutes.Get("/:id", handlers.GetBooking) // Assuming you'll implement GetBooking

	// Payment Routes
	paymentRoutes := api.Group("/payments")                            // e.g., /api/payments
	paymentRoutes.Post("/mpesa/stkpush", handlers.ProcessMpesaSTKPush) // Renamed from processPayment for clarity
	// paymentRoutes.Post("/mpesa/callback", handlers.MpesaCallback) // M-Pesa callback endpoint
}
