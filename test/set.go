package main

import (
	"fmt"
	"github.com/goperate/convert/core/set"
)

func main() {
	var a string
	set.Default(&a, "10")
	fmt.Println(a)
}
