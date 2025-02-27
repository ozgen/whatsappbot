package internal

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
)

// SaveQRCode generates and saves a QR code image
func SaveQRCode(data string) error {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			return fmt.Errorf("failed to create data directory: %v", err)
		}
	}

	return qrcode.WriteFile(data, qrcode.Medium, 256, "data/whatsapp_qr.png")
}
