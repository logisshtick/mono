package main

import (
    "fmt"
    "image/png"
    "log"
    "net/http"

    "github.com/skip2/go-qrcode"
)

func generateQRCode(w http.ResponseWriter, r *http.Request) {
    // Generate the QR code with the user's login information
    loginInfo ... не ебу
    qrCode, err := qrcode.New(loginInfo, qrcode.Medium)
    if err != nil {
        log.Fatal(err)
    }

    // Create a PNG image from the QR code
    pngImage := qrCode.Image(256)

    // Set the appropriate headers for the response
    w.Header().Set("Content-Type", "image/png")

    // Write the image data to the response writer
    err = png.Encode(w, pngImage)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/generate-qr", generateQRCode)
    fmt.Println("Server is running on http://localhost:не помню/generate-qr")
    log.Fatal(http.ListenAndServe(":8080", nil))
} 
