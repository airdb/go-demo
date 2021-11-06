// Package cow 并发安全Map实现
package cow

// SafeMap 并发安全Map接口
type SafeMap interface {
	// Load 加载Map中数据
	Load(interface{}) (interface{}, bool)

	// Store 存储key,value到Map中
	Store(key, value interface{})

	// LoadOrStore 加载key对应的值，如果没有则存储value并返回
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

	// Delete 删除key对应的值
	Delete(interface{})

	// Range 遍历Map键值进行对应操作
	Range(func(key, value interface{}) (shouldContinue bool))
}
