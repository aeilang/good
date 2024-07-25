package utils

import (
	"math/rand"
	"strconv"
)

func GetRandom() string {
	num := rand.Intn(9000) + 1000
	return strconv.Itoa(num)
}
