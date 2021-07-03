package main

import (
	"fmt"
	"github.com/goperate/convert/core/set"
	"reflect"
)

type C struct {
	A int
}

func main() {
	fmt.Println(reflect.Zero(reflect.TypeOf("1")))

	var a string
	set.Default(&a, "10")
	fmt.Println(a)

	var b map[string]interface{}
	set.New(&b)
	b["a"] = 1
	fmt.Println(b)

	var c *C
	c = &C{10}
	fmt.Println(c)
	set.New(&c)
	fmt.Println(*c)
}
