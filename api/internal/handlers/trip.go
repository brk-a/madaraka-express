package handlers

import (
	"fmt"
	"time"

	db "github.com/brk-a/madaraka_express/api/internal/db"
	fiber "github.com/gofiber/fiber/v2"
)

type TripRequest struct {
	Origin        string     `json:"origin"`
	Destination   string     `json:"destination"`
	DepartureTime time.Time  `json:"departure_time"`
	ReturnTime    *time.Time `json:"return_time,omitempty"` // pointer for optional
	TripType      string     `json:"trip_type"`             // "one-way" or "round-trip"
}

func CreateTrip(c *fiber.Ctx) error {
	var req TripRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := db.DB.Exec(c.Context(),
		`INSERT INTO trips (origin, destination, departure_time, return_time, trip_type) VALUES ($1, $2, $3, $4, $5)`,
		req.Origin, req.Destination, req.DepartureTime, req.ReturnTime, req.TripType,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create trip"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Trip created successfully"})
}

func GetTrip(c *fiber.Ctx) error {
	id := c.Params("id") // Get ID from route parameter [2][7]
	// TODO: Implement logic to fetch trip from DB
	return c.SendString(fmt.Sprintf("Get Trip with ID: %s", id))
}
