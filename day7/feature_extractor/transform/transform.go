package transform

import (
	"strconv"

	farm "github.com/dgryski/go-farm"
	"github.com/spaolacci/murmur3"
)

// 把每一个特征的每一个取值映射到[0,2^64]上,几乎保证不同的特征取值映射后也不同
type Transformer interface {
	Hash(string, int) uint64
}

type Murmur struct{}

func (Murmur) Hash(str string, featureId int) uint64 {
	ss := strconv.FormatInt(int64(featureId), 10) + ":" + str
	hash := murmur3.New64()
	hash.Write([]byte(ss))
	z := hash.Sum64()
	hash.Reset()
	return z
}

type FarmHash struct{}

func (FarmHash) Hash(str string, featureId int) uint64 {
	return farm.Hash64WithSeed([]byte(str), uint64(featureId))
}
