package tickets

import (
	"fmt"

	utils "github.com/brk-a/madaraka_express/api/internal/utils"
)

func GenerateTicket(userName, tripDetails, bookingID string) ([]byte, error) {
	qrData := fmt.Sprintf("BookingID:%s;User:%s;Trip:%s", bookingID, userName, tripDetails)
	qrCode, err := utils.GenerateQRCode(qrData)
	if err != nil {
		return nil, err
	}

	pdfBytes, err := utils.GenerateTicketPDF(userName, tripDetails, bookingID, qrCode)
	if err != nil {
		return nil, err
	}
	return pdfBytes, nil
}
