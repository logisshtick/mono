// package for constant and vars sharing
// between endpoints and other packages
package vars

import (
	"errors"
)

var (
	ErrWithDb               = errors.New("An Error with db.")
	ErrBodyLenIsTooBig      = errors.New("Request body too len for processing.")
	ErrInputJsonIsIncorrect = errors.New("Input json is incorrect.")
	ErrBodyReadingFailed    = errors.New("Error while req body reading.")
	ErrFieldTooBig          = errors.New("Json field is too big.")
	ErrActionLimited        = errors.New("Action limited. Try another time.")
	ErrWithPowGen           = errors.New("Pow generation failed. Try another time.")
)
