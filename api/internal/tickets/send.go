package tickets

import (
    "bytes"
    "fmt"
    "net/smtp"
)

func SendTicketEmail(toEmail string, pdfData []byte) error {
    from := "your-email@example.com"
    password := "your-email-password"

    smtpHost := "smtp.example.com"
    smtpPort := "587"

    subject := "Your Ticket Confirmation"
    body := "Thank you for booking with us. Please find your ticket attached."

    // Create email message with attachment (simplified)
    msg := bytes.Buffer{}
    msg.WriteString(fmt.Sprintf("From: %s\r\n", from))
    msg.WriteString(fmt.Sprintf("To: %s\r\n", toEmail))
    msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
    msg.WriteString("MIME-Version: 1.0\r\n")
    msg.WriteString(`Content-Type: multipart/mixed; boundary="BOUNDARY"` + "\r\n")
    msg.WriteString("\r\n--BOUNDARY\r\n")
    msg.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n")
    msg.WriteString(body + "\r\n")
    msg.WriteString("--BOUNDARY\r\n")
    msg.WriteString("Content-Type: application/pdf\r\n")
    msg.WriteString("Content-Disposition: attachment; filename=\"ticket.pdf\"\r\n\r\n")
    msg.Write(pdfData)
    msg.WriteString("\r\n--BOUNDARY--")

    auth := smtp.PlainAuth("", from, password, smtpHost)
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg.Bytes())
    return err
}
