package handlers

import (
	db "github.com/brk-a/madaraka_express/api/internal/db"
	utils "github.com/brk-a/madaraka_express/api/internal/utils"
	fiber "github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	// Insert user into DB
	_, err = db.DB.Exec(c.Context(),
		`INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4)`,
		req.FirstName, req.LastName, req.Email, hashedPassword,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
