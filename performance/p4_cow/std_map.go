package cow

type StdMap struct {
	inner map[interface{}]interface{}
}

func (m *StdMap) Load(key interface{}) (interface{}, bool) {
	v, ok := m.inner[key]
	return v, ok
}

func (m *StdMap) Store(key, value interface{}) {
}

func (m *StdMap) LoadOrStore(key, value interface{}) (interface{}, bool) {
	return nil, false
}

func (m *StdMap) Delete(key interface{}) {
}

func (m *StdMap) Range(f func(key, value interface{}) (shouldContinue bool)) {
}
