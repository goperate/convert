package main

import (
	"fmt"
	"github.com/goperate/convert/core/combination"
)

func main() {
	obj := combination.NewCartesianProduct([]int{1, 2, 3}, []string{"a", "b"}, []float32{2, 4, 6})
	for i := 0; i < 30; i++ {
		fmt.Println(obj.Next())
	}
}
