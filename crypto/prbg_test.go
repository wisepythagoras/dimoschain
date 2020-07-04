package crypto

import (
	"testing"
)

// TestPrbgGeneration tests getting specific bytes.
func TestPrbgGeneration(t *testing.T) {
	// Define our pseudo-random byte generator.
	prbg := PRBG{
		Seed: []byte("test"),
	}

	bytes := prbg.Next(2)

	if bytes[0] != 127 && bytes[1] != 210 {
		t.Errorf("Expected [127, 210] got [%d, %d]\n", bytes[0], bytes[1])
	}

	bytes = prbg.Next(2)

	if bytes[0] != 232 && bytes[1] != 235 {
		t.Errorf("Expected [232, 235] got [%d, %d]\n", bytes[0], bytes[1])
	}
}

// TestPrbgGetUInt tests getting the next uint64 out of the PRBG.
func TestPrbgGetUInt(t *testing.T) {
	// Define our pseudo-random byte generator.
	prbg := PRBG{
		Seed: []byte("test"),
	}

	// Get the next uint64 which has the length of 5 digits.
	num := prbg.NextUInt64(8)

	if num != 4571189930856731263 {
		t.Errorf("Expected 4571189930856731263 got %d\n", num)
	}
}
