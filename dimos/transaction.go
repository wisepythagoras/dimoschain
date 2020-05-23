package dimos

type Transaction struct {
	Hash []byte `json: "h"`
	Amount float64 `json: "a"`
}
