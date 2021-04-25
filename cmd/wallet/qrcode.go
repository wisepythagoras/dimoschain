package main

import (
	"github.com/mdp/qrterminal"
	"os"
)

// GetQRCode generates a QR code and displays in on the terminal.
func GetQRCode(data string) {
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig(data, config)
}
