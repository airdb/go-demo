package cow

import (
	"sync"
	"sync/atomic"
)

// AtomicMap 原子Map实现
//
// 利用atomic.Value原子(无锁)的加载数据
// 利用Copy-on-Write实现数据更新
type AtomicMap struct {
	mu    sync.Mutex
	clean atomic.Value
}

func (m *AtomicMap) Load(key interface{}) (interface{}, bool) {
	data, _ := m.clean.Load().(map[interface{}]interface{})
	v, ok := data[key]
	return v, ok
}

func (m *AtomicMap) Store(key, value interface{}) {
	m.mu.Lock()
	dirty := m.dirty()
	dirty[key] = value
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *AtomicMap) dirty() map[interface{}]interface{} {
	data, _ := m.clean.Load().(map[interface{}]interface{})
	dirty := make(map[interface{}]interface{}, len(data)+1)

	for k, v := range data {
		dirty[k] = v
	}
	return dirty
}

func (m *AtomicMap) LoadOrStore(key, value interface{}) (interface{}, bool) {
	data, _ := m.clean.Load().(map[interface{}]interface{})
	v, ok := data[key]
	if ok {
		return v, ok
	}

	m.mu.Lock()
	// Lock阻塞获取锁期间,可能数据已经存在，再次Load检查数据
	data, _ = m.clean.Load().(map[interface{}]interface{})
	v, ok = data[key]
	if !ok {
		dirty := m.dirty()
		dirty[key] = value
		v = value
		m.clean.Store(dirty)
	}
	m.mu.Unlock()
	return v, ok
}

func (m *AtomicMap) Delete(key interface{}) {
	m.mu.Lock()
	dirty := m.dirty()
	delete(dirty, key)
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *AtomicMap) Range(f func(key, value interface{}) (shouldContinue bool)) {
	data, _ := m.clean.Load().(map[interface{}]interface{})
	for k, v := range data {
		if !f(k, v) {
			break
		}
	}
}
