package collections

import (
	"fmt"

	"github.com/idsulik/go-collections/v2/ringbuffer"
)

func RunCollectionRingBuffer() {
	// Create a new ring buffer with capacity 3
	rb := ringbuffer.New[string](3)

	// Add some items
	rb.Write("first")
	rb.Write("second")
	rb.Write("third")

	// Buffer is now full
	fmt.Println(rb.IsFull()) // Output: true

	// Read an item
	value, ok := rb.Read()
	if ok {
		fmt.Println(value) // Output: "first"
	}

	// Now we can write another item
	rb.Write("fourth")

	// Read all remaining items
	for !rb.IsEmpty() {
		if value, ok := rb.Read(); ok {
			fmt.Println(value)
		}
	}
}
