package main

import (
	"fmt"
	"github.com/goperate/convert/core/sorted"
)

func main() {
	a := []int{1, 2, 3, 7, 2, 5}
	sorted.Sort(a)
	fmt.Println(a)
	b := []float32{1, 2, 3, 7, 2, 5}
	sorted.Sort(b, false)
	fmt.Println(b)
	c := []string{"xx", "as", "cd"}
	sorted.Sort(c)
	fmt.Println(c)
}
