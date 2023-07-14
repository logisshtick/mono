package utils

import (
	"errors"
	"os"
)

const (
	// false == left
	// true  == right
	powLeftOrRight = false
)

// concat pow with string
func PowCat(str, pow string) string {
	if powLeftOrRight {
		return str + pow
	}
	return pow + str
}

// get global pow
func GetGlobalPow() (string, error) {
	powStr := os.Getenv("GLOBAL_POW")
	if powStr == "" {
		return "", errors.New("GLOBAL_POW env var is empty")
	}
	return powStr, nil
}
