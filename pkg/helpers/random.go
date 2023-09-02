package helpers

import (
	"math/rand"
	"time"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

const stringBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const numberBytes = "1234567890"

func RandomChain(chainType byte, n int) string {
	if (n <= 0) {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)

	if chainType == constants.RandomTypeNumber {
		for i := range b {
			b[i] = numberBytes[rand.Intn(len(numberBytes))]
		}
	} else {
		for i := range b {
			b[i] = stringBytes[rand.Intn(len(stringBytes))]
		}
	}
	return string(b)
}
