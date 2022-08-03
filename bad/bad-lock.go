package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex
var criticalSection int

// f1 will call f2 with read lock
func f1() int {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return f2()
}

// f2 will read critical section with read lock again
func f2() int {
	// sleep for 1s to give write lock a chance to be called before 2nd read lock is called
	time.Sleep(1 * time.Second)
	rwLock.RLock()
	defer rwLock.RUnlock()
	return criticalSection
}

func main() {
	go func() {
		// when write lock is called, it will wait the current read locks to be released.
		// no new read locks can be acquired from now onwards unless write lock is acquired and released.
		rwLock.Lock()
		criticalSection += 1
		time.Sleep(10 * time.Second)
		rwLock.Unlock()
	}()
	fmt.Println(f1())
}

