package main

import (
	"fmt"
	"github.com/goperate/convert/core/elemen"
)

func main() {
	e := elemen.NewElement("12")
	fmt.Println(e.ToInt())
	fmt.Println(e.Error) //可以通过Error查看是否转换失败了
}
