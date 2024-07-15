package utils

import (
	"crypto/rand"
	"math/big"
)

const digits = "0123456789"

func GenerateNumber(length int) string {
	byteString := make([]byte, length)
	for i := range byteString {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return ""
		}
		byteString[i] = digits[randomIndex.Int64()]
	}
	return string(byteString)
}

func GenerateCardNumber() string {
	return GenerateNumber(16)
}

func GeneratePINNumber() string {
	return GenerateNumber(4)
}

func GenerateAccountNumber() string {
	return GenerateNumber(10)
}
