package main

import (
	"encoding/json"
	"fmt"
	"github.com/goperate/convert/core/dict"
)

type A struct {
	Id   int
	Name string
}

func main() {
	data := []*map[string]int{{
		"Id": 122,
	}}
	res := dict.StructList2Map(data, "Id").(map[int]*map[string]int)
	b, _ := json.Marshal(res)
	fmt.Println(res, "\n", string(b))
}
