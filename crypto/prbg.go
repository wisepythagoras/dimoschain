package crypto

import (
	"crypto/hmac"

	"github.com/wisepythagoras/dimoschain/utils"
)

// PRBG is the struct representing our pseud-random byte generator object.
type PRBG struct {
	index  int64
	Seed   []byte
	buffer []byte
}

// Next gets the next set of random bytes.
func (p *PRBG) Next(n int) []byte {
	// The payload contains the seed, the current buffer and the index.
	payload := append(p.Seed, p.buffer...)
	payload = append(payload, utils.Int64ToBytes(p.index)...)

	// Now we create the HMAC and write the payload.
	h := hmac.New(HashStrategy, p.Seed)
	h.Write(payload)

	p.index++

	// Return the checksum.
	return h.Sum(nil)[:n]
}

// NextInt64 gets the next set of random integer.
func (p *PRBG) NextInt64(n int) int64 {
	// Get the next set of bytes and return them as an integer.
	return utils.BytesToInt64(p.Next(n))
}
