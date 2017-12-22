package container

import (
	"sync"
)

type Set struct {
	items   []interface{}
	rwMutex *sync.RWMutex
}

func NewSet() *Set {
	retSet := &Set{}
	retSet.rwMutex = new(sync.RWMutex)
	return retSet
}

func (m *Set) Add(value interface{}) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	isHave := false
	for _, v := range m.items {
		if v == value {
			isHave = true
		}
	}
	if !isHave {
		m.items = append(m.items, value)
	}
}

func (m *Set) Delete(value interface{}) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	for k, v := range m.items {
		if v == value {
			m.items = append(m.items[:k], m.items[k+1:]...)
			break
		}
	}
}

func (m *Set) Has(value interface{}) bool {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	isHave := false
	for _, v := range m.items {
		if v == value {
			isHave = true
		}
	}
	return isHave
}

func (m *Set) GetAll() []interface{} {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	var ret []interface{}
	for _, v := range m.items {
		ret = append(ret, v)
	}
	return ret
}

func (m *Set) Len() int {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	ret := len(m.items)
	return ret
}
