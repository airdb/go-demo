package main

import (
  "fmt"
  "sort"
)

func main() {
  strs := []string{"c","a","b"}
  sort.Strings(strs)
  fmt.Println("Strings:",strs)

  ints :=[]int{7, 2, 4}
  sort.Ints(ints)
  fmt.Println("ints:", ints)

  a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4}
  sort.Sort(sort.Float64Slice(a))
  fmt.Println(a)

  keys := []int{3, 2, 8, 1}
  sort.Sort(sort.Reverse(sort.IntSlice(keys)))
  fmt.Println("ints reverse:", keys)

  b := []float64{3.3, 2.2, 8.01, 1.0}
  sort.Sort(sort.Reverse(sort.Float64Slice(b)))
  fmt.Println("float reverse:",b)
}
