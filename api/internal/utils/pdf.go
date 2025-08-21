package utils

import (
    "bytes"
    "fmt"
    gopdf "github.com/signintech/gopdf"
)

func GenerateTicketPDF(userName, tripDetails, bookingID string, qrCode []byte) ([]byte, error) {
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
    pdf.AddPage()
    pdf.SetFont("Arial", "", 14)

    pdf.Cell(nil, fmt.Sprintf("Ticket for %s", userName))
    pdf.Br(20)
    pdf.Cell(nil, tripDetails)
    pdf.Br(20)
    pdf.Cell(nil, fmt.Sprintf("Booking ID: %s", bookingID))
    pdf.Br(20)

    // Add QR code image
    imgHolder, err := gopdf.ImageHolderByBytes(qrCode)
    if err != nil {
        return nil, err
    }
    pdf.ImageByHolder(imgHolder, 50, 100, nil)

    var buf bytes.Buffer
    err = pdf.Write(&buf)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
