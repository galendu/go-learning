package transform

type Transformer interface {
	Hash(string, int) uint64
}
