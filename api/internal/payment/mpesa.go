package payment

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
)

type STKPushRequest struct {
    BusinessShortCode string `json:"BusinessShortCode"`
    Password          string `json:"Password"`
    Timestamp         string `json:"Timestamp"`
    TransactionType   string `json:"TransactionType"`
    Amount            string `json:"Amount"`
    PartyA            string `json:"PartyA"`
    PartyB            string `json:"PartyB"`
    PhoneNumber       string `json:"PhoneNumber"`
    CallBackURL       string `json:"CallBackURL"`
    AccountReference  string `json:"AccountReference"`
    TransactionDesc   string `json:"TransactionDesc"`
}

type STKPushResponse struct {
    MerchantRequestID string `json:"MerchantRequestID"`
    CheckoutRequestID string `json:"CheckoutRequestID"`
    ResponseCode      string `json:"ResponseCode"`
    ResponseDescription string `json:"ResponseDescription"`
    CustomerMessage   string `json:"CustomerMessage"`
}

// Initiate STK Push payment request
func InitiateSTKPush(phoneNumber string, amount string, accountRef string, callbackURL string) (*STKPushResponse, error) {
    // Prepare request data (BusinessShortCode, Password, Timestamp etc. should be generated or stored securely)
    reqBody := STKPushRequest{
        BusinessShortCode: os.Getenv("MPESA_SHORTCODE"),
        Password:          os.Getenv("MPESA_PASSWORD"),
        Timestamp:         time.Now().Format("20060102150405"),
        TransactionType:   "CustomerPayBillOnline",
        Amount:            amount,
        PartyA:            phoneNumber,
        PartyB:            os.Getenv("MPESA_SHORTCODE"),
        PhoneNumber:       phoneNumber,
        CallBackURL:       callbackURL,
        AccountReference:  accountRef,
        TransactionDesc:   "Ticket Payment",
    }

    jsonData, err := json.Marshal(reqBody)
    if err != nil {
        return nil, err
    }

    // Example endpoint (replace with actual Daraja endpoint)
    url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+os.Getenv("MPESA_ACCESS_TOKEN"))

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var stkResp STKPushResponse
    if err := json.NewDecoder(resp.Body).Decode(&stkResp); err != nil {
        return nil, err
    }

    if stkResp.ResponseCode != "0" {
        return &stkResp, fmt.Errorf("mpesa error: %s", stkResp.ResponseDescription)
    }

    return &stkResp, nil
}
