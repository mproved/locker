package locker

import (
	"sync"
	"time"
)

type Locker struct {
	locks     map[any]struct{}
	mutex     sync.Mutex
	sleepTime time.Duration
	maxTries  int
}

func (m *Locker) Lock(key any) bool {
	tries := 0

	for {
		if tries >= m.maxTries {
			return false
		}

		m.mutex.Lock()
		if _, ok := m.locks[key]; ok {
			m.mutex.Unlock()
			time.Sleep(m.sleepTime)
			tries += 1
		} else {
			m.locks[key] = struct{}{}
			m.mutex.Unlock()
			return true
		}
	}
}

func (m *Locker) Unlock(key any) {
	m.mutex.Lock()
	delete(m.locks, key)
	m.mutex.Unlock()
}

func NewLocker(sleepTime time.Duration, maxTries int) *Locker {
	return &Locker{
		locks:     make(map[any]struct{}),
		mutex:     sync.Mutex{},
		sleepTime: sleepTime,
		maxTries:  maxTries,
	}
}
