package ringo

import (
	"sync"
)

type Ringo struct {
	size    int
	current int
	items   []interface{}
	lock    sync.RWMutex
}

func New(size int) *Ringo {
	return &Ringo{
		size:  size,
		items: make([]interface{}, size),
	}
}

func (self *Ringo) Insert(value interface{}) interface{} {
	self.lock.Lock()

	result := self.items[self.current]
	self.items[self.current] = value
	if self.current == self.size-1 {
		self.current = 0
	} else {
		self.current++
	}

	self.lock.Unlock()

	return result
}

func (self *Ringo) DumpSize(size int) []interface{} {
	self.lock.RLock()

	tmp := make([]interface{}, size)
	ptr := self.current
	for i := 0; i < size; i++ {
		tmp[i] = self.items[ptr]
		if ptr == self.size-1 {
			ptr = 0
		} else {
			ptr++
		}
	}

	self.lock.RUnlock()

	return tmp
}

func (self *Ringo) Dump() []interface{} {
	return self.DumpSize(self.size)
}

func (self *Ringo) ShiftLeft() {
	self.lock.Lock()
	if self.current == self.size-1 {
		self.current = 0
	} else {
		self.current++
	}
	self.lock.Unlock()
}
