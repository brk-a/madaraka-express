package handlers

import (
	"fmt"
	"log"
	"os"

	payment "github.com/brk-a/madaraka_express/api/internal/payment"
	fiber "github.com/gofiber/fiber/v2"
)

type MpesaSTKPushRequest struct {
	PhoneNumber string `json:"phone_number"`
	Amount      string `json:"amount"`
	BookingID   int    `json:"booking_id"` // To link payment to a booking
}

// ProcessMpesaSTKPush initiates an M-Pesa STK Push
func ProcessMpesaSTKPush(c *fiber.Ctx) error {
	var req MpesaSTKPushRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// You'd typically fetch booking details to get the actual amount and reference
	// For now, using direct amount from request and booking ID as reference
	accountRef := fmt.Sprintf("Booking-%d", req.BookingID)
	callbackURL := os.Getenv("MPESA_CALLBACK_URL") // Set this in your .env

	resp, err := payment.InitiateSTKPush(req.PhoneNumber, req.Amount, accountRef, callbackURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":              "STK Push initiated successfully",
		"merchant_request_id":  resp.MerchantRequestID,
		"checkout_request_id":  resp.CheckoutRequestID,
		"response_description": resp.ResponseDescription,
	})
}

// MpesaCallback handles the M-Pesa payment confirmation callback (requires separate route definition)
func MpesaCallback(c *fiber.Ctx) error {
	// This handler will receive the M-Pesa payment confirmation
	// TODO: Parse callback, update payment status in DB, trigger ticket generation and email
	var callbackData map[string]interface{}
	if err := c.BodyParser(&callbackData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid callback data"})
	}

	log.Printf("M-Pesa Callback Received: %+v\n", callbackData)
	// Here, you'd extract payment details, find the corresponding booking,
	// update payment status in the 'payments' table, and if successful,
	// call tickets.GenerateTicket and tickets.SendTicketEmail
	// Example:
	// if paymentStatus == "success" {
	//     pdfBytes, err := tickets.GenerateTicket(user.Name, trip.Details, booking.ID)
	//     if err == nil {
	//         tickets.SendTicketEmail(user.Email, pdfBytes)
	//     }
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Callback received"})
}
