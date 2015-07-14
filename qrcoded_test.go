package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")

	if buffer.Len() == 0 {
		t.Errorf("Generated QRCode has no data")
	}

	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected Error")
}

func TestGenerateQRCodePropagatsErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2386")

	if err == nil || err.Error() != "Expected Error" {
		t.Errorf("Error not propagated correctly, got %v", err)
	}
}
