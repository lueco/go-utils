package container

type MapSet struct {
	inerMap *Map
}

func NewMapSet() (retMapSet *MapSet) {
	retMapSet = &MapSet{}
	retMapSet.inerMap = NewMap()
	return
}

func (m *MapSet) GetSet(key interface{}) (retSet *Set, retOk bool) {
	retSet = nil
	retOk = false
	if ret, ok := m.inerMap.Get(key); ok {
		retSet = ret.(*Set)
		retOk = ok
	}
	return
}

func (m *MapSet) Has(key interface{}, value interface{}) (retSet *Set, retOk bool) {
	retSet = nil
	retOk = false
	if set, ok := m.GetSet(key); ok {
		retSet = set
		retOk = retSet.Has(value)
	}
	return
}

func (m *MapSet) Add(key interface{}, value interface{}) {
	var retSet *Set
	if set, ok := m.GetSet(key); !ok {
		retSet = NewSet()
		m.inerMap.Add(key, retSet)
	} else {
		retSet = set
	}
	retSet.Add(value)
}

func (m *MapSet) Delete(key interface{}, value interface{}) (retSet *Set, retOk bool) {
	retOk = false
	retSet = nil
	if set, ok := m.GetSet(key); ok {
		set.Delete(value)
		retOk = ok
		if set.Len() == 0 {
			m.inerMap.Delete(key)
		} else {
			retSet = set
		}
	}
	return
}

func (m *MapSet) DeleteSet(key interface{}) {
	m.inerMap.Delete(key)
}

func (m *MapSet) GetAllSet() []*Set {
	var ret []*Set
	for _, v := range m.inerMap.GetAll() {
		if set, ok := v.(*Set); ok {
			ret = append(ret, set)
		}
	}
	return ret
}

func (m *MapSet) GetAllKey() []interface{} {
	var ret []interface{}
	for key, _ := range m.inerMap.GetAll() {
		ret = append(ret, key)
	}
	return ret
}

func (m *MapSet) GetAll() []interface{} {
	var ret []interface{}
	for _, set := range m.GetAllSet() {
		ret = append(ret, set.GetAll()...)
	}
	return ret
}
