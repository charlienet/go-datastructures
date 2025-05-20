package locker

import "sync"

type Locker struct {
	l ilocker
}

func (l *Locker) Synchronize() *Locker {
	
	l.l = &rwLocker{}
	return l
}

func (l *Locker) Lock() {
	if l.l == nil {
		l.l = &emptyLocker{}
	}

	l.l.Lock()
}

func (l *Locker) RLock() {
	if l.l == nil {
		l.l = &emptyLocker{}
	}

	l.l.RLock()
}

func (l *Locker) Unlock() {
	l.l.Unlock()
}

func (l *Locker) RUnlock() {
	l.l.RUnlock()
}

type ilocker interface {
	Lock()
	RLock()
	Unlock()
	RUnlock()
}

type rwLocker struct {
	l sync.RWMutex
}

func (l *rwLocker) Lock() {
	l.l.Lock()
}

func (l *rwLocker) RLock() {
	l.l.RLock()
}

func (l *rwLocker) Unlock() {
	l.l.Unlock()
}

func (l *rwLocker) RUnlock() {
	l.l.RUnlock()
}

type emptyLocker struct {
}

func (l *emptyLocker) Lock() {
}

func (l *emptyLocker) RLock() {
}

func (l *emptyLocker) Unlock() {
}

func (l *emptyLocker) RUnlock() {
}
