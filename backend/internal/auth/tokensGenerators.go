// here i generate xxhash and uuid
// based on system random generator
// /dev/random or /dev/urandom
// i can afford it since everything
// runs on linux
package auth

import (
	"bufio"
	crand "crypto/rand"
	"errors"
	"github.com/cespare/xxhash"
	"github.com/google/uuid"
	"math/rand"
	"os"
)

const (
	// should random be crypt safe?
	// if u dont know u can fuck urself
	cryptSafeRand = false

	// amount of bytes from what we
	// generate xxhash and uuid
	// more = better. but slower
	// >= 256
	bytesCount = 4096

	// amount of retryes if random is not work
	retryesAmount = 256
)

var (
	// or /dev/random
	randFile = "/dev/urandom"

	// or crand.Read
	randFunc = rand.Read

	// some kind of errors
	ErrRandBytesGen = errors.New("Error with random bytes generation. simply try to call func again.")
)

// called in init() func
func hashVarSet() {
	if cryptSafeRand {
		randFile = "/dev/random"
		randFunc = crand.Read
	}
}

// should be called in Init()
// coz we should check is randFile exist
// and can we read it of no returns error
func initChecks() error {
	f, err := os.Open(randFile)
	if err != nil {
		return err
	}
	br := bufio.NewReader(f)

	_, err = br.ReadByte()
	if err != nil {
		return err
	}

	return nil
}

func randByteSlice(bytes []byte) error {
	f, err := os.Open(randFile)
	if err != nil {
		return randByteSliceFallback(bytes)
	}
	br := bufio.NewReader(f)

	for i := 0; i < bytesCount; i++ {
		b, err := br.ReadByte()
		if err != nil {
			return randByteSliceFallback(bytes)
		}
		bytes[i] = b
	}

	return nil
}

// if we really get error in randByteSlice()
// i dont think so but i wanna have a fallback
func randByteSliceFallback(bytes []byte) error {
	// shoosh dont beat me plzzzzzzzzzz
	// in go i cant generate correct random number
	// without error and return constant value i cant too
	// because it can rewrite another user data
	for i := 0; i < retryesAmount; i++ {
		_, err := randFunc(bytes)
		if err == nil {
			break
		}
		return ErrRandBytesGen
	}
	return nil
}

func genHash() (uint64, error) {
	bytes := make([]byte, bytesCount)
	err := randByteSlice(bytes)
	if err != nil {
		return 0, err
	}
	return xxhash.Sum64(bytes), nil
}

func genUuid() (string, error) {
	bytes := make([]byte, bytesCount)
	err := randByteSlice(bytes)
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
