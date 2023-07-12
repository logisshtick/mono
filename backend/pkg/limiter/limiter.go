/*
simple generic thread safe limiter
also can be used as action limiter
*/
package limiter

import (
	"github.com/logisshtick/mono/pkg/mu"
	"sync"
	"runtime"
	"time"

	"golang.org/x/exp/constraints"
)

const (
	// default hashmap size for first allocation
	defaultMapLen = 2048

	// max hashmap len returned by len()
	// before clean up
	defaultMaxMapLen = 16384

	// how many objects Cleanup()
	// goroutine can clean at once
	defaultCleanAtOnce = 20

	// default time of actions
	defaultMaxTime = 3600

	// default count of actions
	defaultMaxCount = 30

	// default value)
	Default = -1
)

type action struct {
	deltaTime int64
	count     uint64
}

type Limiter[T constraints.Ordered] struct {
	m         map[T]action
	mu        sync.RWMutex
	maxTime   int64
	maxCount  uint64
	maxMapLen  int
	cleanAtOnce uint64
	cleaned bool
	cleanedMu sync.RWMutex
}

// make new limiter for type T with maxCount for all actions
//
// if mapSize < 0 it sets to default map size
// also u can use limiter.Default const
//
// if maxMapLen is 0 means that the maximum map size is unlimited
// and clean up will never happen
// also u can use limiter.Default const
func New[T constraints.Ordered](
	maxCount uint64, maxTime int64, 
	mapLen, maxMapLen int, cleanAtOnce uint64,
) *Limiter[T] {
	if maxCount <= 0 {
		maxCount = defaultMaxCount
	}
	if mapLen <= 0 {
		mapLen = defaultMapLen
	}
	if maxMapLen < 0 {
		maxMapLen = defaultMaxMapLen
	}
	if cleanAtOnce <= 0 {
		cleanAtOnce = defaultCleanAtOnce
	}

	return &Limiter[T]{
		m:         make(map[T]action, mapLen),
		maxTime:   maxTime,
		maxCount:  maxCount,
		maxMapLen: maxMapLen,
		cleanAtOnce: cleanAtOnce,
	}
}

func (l *Limiter[T]) Try(id T) bool {
	timeNow := time.Now().Unix()

	var (
		a  action
		ok bool
	)
	mu.ExecRWMutex(&l.mu, func() {
		a, ok = l.m[id]
	})
	if !ok {
		mu.ExecMutex(&l.mu, func() {
			l.m[id] = action{
				deltaTime: timeNow,
				count:     1,
			}
		})
		return true
	}
	if timeNow-a.deltaTime < l.maxTime && 
		a.count >= l.maxCount {
		return false
	}

	mu.ExecMutex(&l.mu, func() {
		l.m[id] = action{
			deltaTime: timeNow,
			count:     a.count + 1,
		}
	})

	if len(l.m) >= l.maxMapLen {
		go l.Clean()
	}

	return true
}


func (l *Limiter[T]) Clean() {
	var cleanedState bool
	mu.ExecRWMutex(&l.cleanedMu, func() {
		cleanedState = l.cleaned
	})
	if cleanedState {
		return
	}
	mu.ExecMutex(&l.cleanedMu, func() {
		l.cleaned = true
	})

	var i uint64
	mu.ExecMutex(&l.mu, func() {
		for key, val := range l.m {
			if i == l.cleanAtOnce {
				i = 0
				l.mu.Unlock()
				runtime.Gosched()
				l.mu.Lock()
			}

			timeNow := time.Now().Unix()
			if timeNow-val.deltaTime >= l.maxTime {
				delete(l.m, key)
			}
			i++
		}
	})
	mu.ExecMutex(&l.cleanedMu, func() {
		l.cleaned = false
	})
}
