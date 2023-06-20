/*
simple generic thread safe limiter
also can be used as action limiter
*/
package limiter

import (
	"time"
	"sync"
	"github.com/logisshtick/mono/internal/utils"

	"golang.org/x/exp/constraints"
)

const (
	// default hashmap size for first allocation
	defaultMapSize = 2048

	// max hashmap len returned by len()
	// before clean up
	defaultMaxMapLen = 16384

	// default time of actions
	defaultMaxTime = 3600

	// default count of actions
	defaultMaxCount = 30

	// default value)
	Default = -1
)

type action struct {
	deltaTime int64
	count int
}

type Limiter[T constraints.Ordered] struct {
	m map[T]action
	mu sync.RWMutex
	maxTime int64
	maxCount int64
	maxMapLen int
}

// make new limiter for type T with maxCount for all actions
//
// if mapSize < 0 it sets to default map size
// also u can use limiter.Default const
//
// if maxMapLen is 0 means that the maximum map size is unlimited 
// and clean up will never happen
// also u can use limiter.Default const
func New[T constraints.Ordered](maxCount int, maxTime int64, mapSize, maxMapLen int) *Limiter[T] {
	if maxCount <= 0 {
		maxCount = defaultMaxCount
	}
	if mapSize <= 0 {
		mapSize = defaultMapSize
	}
	if maxMapLen < 0 {
		maxMapLen = defaultMaxMapLen
	}

	return &Limiter[T]{
		m: make(map[T]action, mapSize),
		maxTime: maxTime,
		maxCount: int64(maxCount),
		maxMapLen: maxMapLen,
	}
}

