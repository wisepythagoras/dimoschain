package main

import (
	"os"
	"github.com/mdp/qrterminal"
)

// GetQRCode generates a QR code and displays in on the terminal.
func GetQRCode(data string) {
	config := qrterminal.Config{
		Level: qrterminal.M,
		Writer: os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig(data, config)
}

