package handlers

import (
	"context"
	"fmt"

	db "github.com/brk-a/madaraka_express/api/internal/db"
	fiber "github.com/gofiber/fiber/v2"
)

type BookingRequest struct {
	UserID       int  `json:"user_id"`
	TripID       int  `json:"trip_id"`
	ReturnTripID *int `json:"return_trip_id,omitempty"` // optional for round-trip
	Seats        int  `json:"seats"`
}

func CreateBooking(c *fiber.Ctx) error {
	var req BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	ctx := context.Background()

	// Check seat availability for outbound trip
	available, err := checkSeatAvailability(ctx, req.TripID, req.Seats)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking seat availability"})
	}
	if !available {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Not enough seats available for outbound trip"})
	}

	// If round-trip, check return trip availability
	if req.ReturnTripID != nil {
		available, err = checkSeatAvailability(ctx, *req.ReturnTripID, req.Seats)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking return trip seat availability"})
		}
		if !available {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Not enough seats available for return trip"})
		}
	}

	// Insert booking for outbound trip
	bookingID, err := insertBooking(ctx, req.UserID, req.TripID, req.Seats)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating booking"})
	}

	// Insert booking for return trip if applicable
	if req.ReturnTripID != nil {
		_, err = insertBooking(ctx, req.UserID, *req.ReturnTripID, req.Seats)
		if err != nil {
			// Optionally rollback first booking or mark as partial booking
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating return trip booking"})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Booking created successfully", "booking_id": bookingID})
}

// Helper: Check seat availability for a trip
func checkSeatAvailability(ctx context.Context, tripID int, seatsRequested int) (bool, error) {
	// Assume max seats per trip is 100 (or query from trips table)
	const maxSeats = 100

	var bookedSeats int
	err := db.DB.QueryRow(ctx,
		`SELECT COALESCE(SUM(seats), 0) FROM bookings WHERE trip_id = $1 AND status = 'confirmed'`,
		tripID).Scan(&bookedSeats)
	if err != nil {
		return false, err
	}

	if bookedSeats+seatsRequested > maxSeats {
		return false, nil
	}
	return true, nil
}

// Helper: Insert booking into database
func insertBooking(ctx context.Context, userID, tripID, seats int) (int, error) {
	var bookingID int
	err := db.DB.QueryRow(ctx,
		`INSERT INTO bookings (user_id, trip_id, seats, status) VALUES ($1, $2, $3, 'confirmed') RETURNING id`,
		userID, tripID, seats).Scan(&bookingID)
	if err != nil {
		return 0, err
	}
	return bookingID, nil
}

// GetBooking handles fetching a single booking by ID
func GetBooking(c *fiber.Ctx) error {
	id := c.Params("id") // Get ID from route parameter [2][7]
	// TODO: Implement logic to fetch booking from DB
	return c.SendString(fmt.Sprintf("Get Booking with ID: %s", id))
}
