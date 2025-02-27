package internal

import (
	"os"
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestSaveQRCode(t *testing.T) {
	testData := "test-qr-code"

	// Cleanup before the test
	os.RemoveAll("data")

	// Run the function
	err := SaveQRCode(testData)
	if err != nil {
		t.Fatalf("SaveQRCode failed: %v", err)
	}

	// Check if directory was created
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		t.Fatalf("Expected 'data' directory to be created, but it does not exist")
	}

	// Check if QR code file was created
	if _, err := os.Stat("data/whatsapp_qr.png"); os.IsNotExist(err) {
		t.Fatalf("Expected QR code file to be created, but it does not exist")
	}

	// Cleanup after the test
	os.RemoveAll("data")
}

func TestSaveQRCode_DirectoryCreationFails(t *testing.T) {
	// Create a file named "data" to force mkdir failure
	os.RemoveAll("data")
	_, err := os.Create("data")
	if err != nil {
		t.Fatalf("Failed to create a file named 'data': %v", err)
	}
	defer os.Remove("data")

	err = SaveQRCode("test-data")
	if err == nil {
		t.Fatalf("Expected error due to directory creation failure, but got nil")
	}
}

func TestSaveQRCode_QRCodeWriteFails(t *testing.T) {
	// Override the QR code size to an invalid value to trigger an error
	invalidSize := -1
	err := qrcode.WriteFile("test", qrcode.Medium, invalidSize, "data/whatsapp_qr.png")

	if err == nil {
		t.Fatalf("Expected QR code writing to fail, but it succeeded")
	}
}
