package main

import (
	"fmt"
	"github.com/goperate/convert/core/set"
)

type Test struct {
	A int
}

func default_() {
	var a string
	set.Default(&a, "10")
	fmt.Println(a)

	var b int
	set.Default(&b, 10)
	fmt.Println(b)

	var c float32 = 1.28
	set.Default(&c, float32(2))
	fmt.Println(c)

	var d Test
	set.Default(&d, Test{A: 10})
	fmt.Println(d)

	//空指针必须传入指针的指针, 否则程序报错
	var e *Test
	set.Default(&e, &Test{A: 111})
	fmt.Println(e)

	var f []int
	set.Default(&f, []int{2222})
	fmt.Println(f)

	//长度为0的数组不会被覆盖
	g := make([]int, 0, 0)
	set.Default(&g, []int{2})
	fmt.Println(g)

	var h map[string]int
	set.Default(&h, map[string]int{"a": 666})
	fmt.Println(h)

	//初始化完成的空map不会被覆盖
	i := make(map[string]int)
	set.Default(&i, map[string]int{"a": 777})
	fmt.Println(i)
}

func new_() {
	var b map[string]interface{}
	set.New(&b)
	b["a"] = 1
	fmt.Println(b)

	var c *Test
	c = &Test{10}
	fmt.Println(c)
	set.New(&c)
	fmt.Println(*c)
}

func if_() {
	a := 10
	b := 3
	set.IF(&a, b > 0, 10000, -10)
	fmt.Println(a)
}

func main() {
	default_()
}
