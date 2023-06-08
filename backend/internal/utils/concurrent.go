package utils

import (
	"sync"
)

type mutex interface {
	*sync.RWMutex | *sync.Mutex

	Lock()
	Unlock()
}

// exec any func with read only lock
func ExecRWMutex(m *sync.RWMutex, f func()) {
	m.RLock()
	f()
	m.RUnlock()
}

// exec any func with full lock
func ExecMutex[T mutex](m T, f func()) {
	m.Lock()
	f()
	m.Unlock()
}

