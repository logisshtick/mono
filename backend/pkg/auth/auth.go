// some kind of abbreviation for shoosh)
//
//	m   - map
//	Ms  - maps
//	mu  - mutex
//	rmu - read-write mutex
//	g   - goroutine
//	t   - token
//	Ts  - tokens
//	T   - type
package auth

import (
	"sync"
)

const (
	// size for token caching hashmaps
	mapSize = 2048

	// tokens max life time in seconds
	// refresh token should live more
	// than access token
	accessTLife  int64 = 1800
	refreshTLife int64 = 2592000
)

type token struct {
	date     int64
	id       int
	deviceId uint64
}

// another data type for local mutexes
type tokenMs struct {
	// uint64 coz xxHash uses it
	accessTs    map[uint64]token
	accessTsRmu sync.RWMutex

	// string hashing is really slow in go map
	// but we need access it only if
	// accessTLife time expired
	refreshTs    map[string]token
	refreshTsRmu sync.RWMutex
}

func newTokenMs() tokenMs {
	return tokenMs{
		accessTs:  make(map[uint64]token, mapSize),
		refreshTs: make(map[string]token, mapSize),
	}
}

var (
	// is package inited
	inited = false
	maps   = newTokenMs()
)

// private init method mostly used
// to set vars based on constant
// and not to do it in frequently
// called functions
func init() {
	hashVarSet()
}

func Init() error {
	if inited {
		return nil
	}

	err := initChecks()
	if err != nil {
		return err
	}

	inited = true
	return nil
}
