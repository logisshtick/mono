// package for constant and vars sharing
// between endpoints and other packages
package vars

import (
	"errors"
)

const (
	MaxHttpBodyLen = 16384
)

var (
	ErrBodyLenIsTooBig      = errors.New("Request body too len for processing.")
	ErrInputJsonIsIncorrect = errors.New("Input json is incorrect.")
	ErrBodyReadingFailed    = errors.New("Error while req body reading.")
)
