package container

import (
	"sync"
)

type Map struct {
	items   map[interface{}]interface{}
	rwMutex *sync.RWMutex
}

func NewMap() *Map {
	retMap := &Map{}
	retMap.items = make(map[interface{}]interface{})
	retMap.rwMutex = new(sync.RWMutex)
	return retMap
}

func (m *Map) Get(key interface{}) (interface{}, bool) {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	v, ok := m.items[key]
	return v, ok
}

func (m *Map) Delete(key interface{}) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	delete(m.items, key)
}

func (m *Map) Add(key interface{}, value interface{}) {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	m.items[key] = value
}

func (m *Map) Len() int {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	return len(m.items)
}

func (m *Map) GetAll() map[interface{}]interface{} {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	ret := deepCopy(m.items)
	if retMap, ok := ret.(map[interface{}]interface{}); ok {
		return retMap
	}
	return nil
}

func (m *Map) GetCount() int {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	ret := len(m.items)
	return ret
}

func deepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[interface{}]interface{}); ok {
		newMap := make(map[interface{}]interface{})
		for k, v := range valueMap {
			newMap[k] = deepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = deepCopy(v)
		}

		return newSlice
	}

	return value
}
