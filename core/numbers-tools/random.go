package numberstools

import (
	"crypto/rand"
	"math/big"
)

func GetRandomIntInRange(min, max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(int64(min) + n.Int64())
}

func GetRandomInt(maxN int) int {
	return GetRandomIntInRange(0, maxN)
}

func GetRandomIntSlice(len int) []int {
	slice := make([]int, 0, len)
	for i := 0; i < len; i++ {
		slice = append(slice, GetRandomInt(100))
	}
	return slice
}
