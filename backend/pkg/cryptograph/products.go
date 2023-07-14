package cryptograph

import (
	"crypto/sha256"
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/google/uuid"
)

// hash password with crypto safe hash func
func HashPass(pass string) string {
	sum := sha256.Sum256([]byte(pass))
	return fmt.Sprintf("%x", sum)
}

// generate random pow
func GenRandPow(length int) (string, error) {
	bytes := make([]byte, length/2)
	err := RandByteSlice(bytes)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%x", bytes), nil
}

// generate random xxhash
func GenRandHash() (uint64, error) {
	bytes := make([]byte, bytesCount)
	err := RandByteSlice(bytes)
	if err != nil {
		return 0, err
	}
	return xxhash.Sum64(bytes), nil
}

// generate random uuid with google package
func GenRandUuid() (string, error) {
	bytes := make([]byte, bytesCount)
	err := RandByteSlice(bytes)
	if err != nil {
		return "", err
	}
	uid, err := uuid.FromBytes(bytes)
	if err != nil {
		for i := 0; i < retryesAmount; i++ {
			uid, err = uuid.NewUUID()
			if err == nil {
				break
			}
			return "", err
		}
	}
	return uid.String(), nil
}
