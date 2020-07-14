package ring

import "sync"

type Ring struct {
	value interface{}
	lock  sync.RWMutex
	next  *Ring
	prev  *Ring
}

func NewRing(value interface{}) *Ring {
	r := &Ring{
		value: value,
	}
	r.next = r
	r.prev = r
	return r
}

func (r *Ring) Insert(value interface{}) {
	newRing := NewRing(value)
	r.InsertRing(newRing)
}

func (r *Ring) InsertRing(newRing *Ring) {
	r.lock.Lock()
	defer r.lock.Unlock()
	newRing.next = r.next
	r.next = newRing
	newRing.prev = r
	if newRing.next == r {
		r.prev = newRing
	}
}

func (r *Ring) Remove() {
	if r.next == r.prev {
		return
	}
	r.prev.lock.Lock()
	defer r.prev.lock.Unlock()
	r.next.lock.Lock()
	defer r.next.lock.Unlock()
	r.prev.next = r.next
	r.next.prev = r.prev

	r.lock.Lock()
	defer r.lock.Unlock()
	r.prev = nil
	r.next = nil
}

func (r *Ring) Next() *Ring {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.next
}

func (r *Ring) Prev() *Ring {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.prev
}

func (r *Ring) Value() interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.value
}
