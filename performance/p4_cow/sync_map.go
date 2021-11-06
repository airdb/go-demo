package cow

import "sync"

// SyncMap 标准库sync.Map重命名
type SyncMap struct {
	sync.Map
}
