package main

import (
	"fmt"
	"github.com/goperate/convert/core/array"
)

func main() {
	arr := array.NewArray([]string{"1", "2", "3", "5", "3"})
	fmt.Println(arr.ToIntArray())
	fmt.Println(arr.ToFloat32Map()) //返回每个值出现的数量
	arr.Deduplication(true)         //支持对返回结果进行去重
	fmt.Println(arr.ToIntArray())
	fmt.Println(arr.ToInterfaceArray())
}
