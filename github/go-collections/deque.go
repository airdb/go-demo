package collections

import (
	"fmt"

	"github.com/idsulik/go-collections/v2/deque"
)

func RunCollectionDeque() {
	d := deque.New[int](0)
	d.PushBack(1)
	d.PushFront(2)

	front, _ := d.PopFront()
	back, _ := d.PopBack()

	fmt.Println(front) // Output: 2
	fmt.Println(back)  // Output: 1
}
