# go-array

## Element
支持 int、int32、int64、uint、uint32、uint64、float32、float64、string 之间的相互转换

示例:
```go
package main

import (
    "fmt"
    "github.com/goperate/convert/core/element"
)

func main()  {
    e := elemen.NewElement("12")
    fmt.Println(e.ToInt())
    fmt.Println(e.Error) //可以通过Error查看是否转换失败了
}
```

## Array
支持 []int、[]int32、[]int64、[]uint、[]uint32、[]uint64、[]float32、[]float64、[]string 之间的相互转换及去重

示例:
```go
package main

import (
	"fmt"
	"github.com/goperate/convert/core/array"
)

func main() {
	arr := array.NewArray([]string{"1", "2", "3", "5", "3"})
	fmt.Println(arr.ToIntArray())
	fmt.Println(arr.ToFloat32Map()) //返回每个值出现的数量
	arr.Deduplication(true) //支持对返回结果进行去重
	fmt.Println(arr.ToIntArray())
}

```

## Dict
支持生成复杂map及从map中取值

示例:
```go
package main

import (
	"fmt"
	"github.com/goperate/convert/core/dict"
)

func main() {
	d := dict.NewDict(nil)
	d.Set("1", false, 2, "3", 4, 5)
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set(2, true, 3, "4", 5)
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set(3, false, 3, 5)
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set(2.3, true, 3, "x")
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set("2", true, "3", "xx")
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set("3", true, 3, 4)
	fmt.Println(d.Data, "----------", d.Value.Type())
	fmt.Println(d.Get(3, 4, 0), d.Error)
	fmt.Println(d.GetElement(3, "x", 0).ToString(), d.Error)
}
```
```
map[2:map[3:map[4:map[5:1]]]] ---------- map[int]map[string]map[int]map[int]string
map[2:map[3:map[4:map[5:1]]] 3:map[5:3 4:map[5:[2]]]] ---------- map[int]interface {}
map[2:map[3:map[4:map[5:1]]] 3:map[4:map[5:[2]]]] ---------- map[int]interface {}
map[2:map[3:map[4:map[5:1]]] 3:map[5:3 4:map[5:[2]] x:[2.3]]] ---------- map[int]interface {}
map[2:map[3:map[4:map[5:1]]] 3:map[5:3 4:map[5:[2]] x:[2.3]] 3:map[xx:[2]]] ---------- map[interface {}]interface {}
map[2:map[3:map[4:map[5:1]]] 3:map[4:[3] 5:3 4:map[5:[2]] x:[2.3]] 3:map[xx:[2]]] ---------- map[interface {}]interface {}
3 <nil>
2.3 <nil>
```
