package main

import (
	"encoding/json"
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
		"Id": 122,
	}}
	res := array.ToIdMap(data, "Id").(map[int]*map[string]int)
	b, err := json.Marshal(res)
	fmt.Println(res, "\n", string(b), err)
}

func test2() {
	data := []map[int]float32{{
		22: 122,
	}}
	res := array.ToIdMap(data, 22).(map[float32]map[int]float32)
	b, err := json.Marshal(res)
	fmt.Println(res, "\n", string(b), err)
}

func test3() {
	data := []*A{{1, 2, "a"}, {2, 3, "b"}}
	res := array.ToIdMap(data, "Id").(map[int]*A)
	b, err := json.Marshal(res)
	fmt.Println(res, "\n", string(b), err)
}

func test4() {
	data := []A{{1, 2, "a"}, {2, 3, "b"}}
	res := array.ToIdMap(data, "Id2").(map[uint]A)
	b, err := json.Marshal(res)
	fmt.Println(res, "\n", string(b), err)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
