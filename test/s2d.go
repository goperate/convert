package main

import (
	"fmt"
	"github.com/goperate/convert/core/array"
)

type A struct {
	Id   int
	Id2  uint
	Name string
}

func test1() {
	data := []*map[string]int{{
		"id": 122,
	}}
	arr := array.NewObjArray(data, "id")
	res := arr.ToIdMap().(map[int]*map[string]int)
	fmt.Println(res)
	res2 := arr.ToIdMapArray().(map[int][]*map[string]int)
	fmt.Println(res2)
	fmt.Println()
}

func test2() {
	data := []map[int]float32{{
		22: 122,
	}}
	arr := array.NewObjArray(data, 22)
	res := arr.ToIdMap().(map[float32]map[int]float32)
	fmt.Println(res)
	res2 := arr.ToIdMapArray().(map[float32][]map[int]float32)
	fmt.Println(res2)
	fmt.Println()
}

func test3() {
	data := []*A{{1, 2, "a"}, {2, 3, "b"}}
	arr := array.NewObjArray(data, "Id")
	res := arr.ToIdMap().(map[int]*A)
	fmt.Println(res)
	res2 := arr.ToIdMapArray().(map[int][]*A)
	fmt.Println(res2)
	fmt.Println()
}

func test4() {
	data := []A{{1, 2, "a"}, {2, 3, "b"}, {3, 2, "aa"}}
	arr := array.NewObjArray(data, "Id2")
	res := arr.ToIdMap().(map[uint]A)
	fmt.Println(res)
	res2 := arr.ToIdMapArray().(map[uint][]A)
	fmt.Println(res2)
	fmt.Println()
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
