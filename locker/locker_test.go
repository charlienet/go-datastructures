package locker

import "testing"

func TestLocker(t *testing.T) {
	l := &Locker{}

	l.Lock()
	println(l)
	l.Unlock()
	l.Synchronize()
	l.RLock()
	println(l)
	l.RUnlock()
}
