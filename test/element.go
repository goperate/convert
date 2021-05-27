package main

import (
	"fmt"
	"github.com/goperate/convert/core/element"
)

func main() {
	e := element.NewElement("12")
	fmt.Println(e.ToInt())
	fmt.Println(e.Error) //可以通过Error查看是否转换失败了
}
