package dimos

type Blockchain struct {
	Height int64 `json: "h"`
	CurrentHash []byte `json: "ch"`
}
