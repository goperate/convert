[toc]

# 数字与字符串之间的任意转换
支持 int、int32、int64、uint、uint32、uint64、float32、float64、string 之间的相互转换

示例:
```go
package main

import (
    "fmt"
    "github.com/goperate/convert/core/element"
)

func main()  {
    e := element.NewElement("12")
    fmt.Println(e.ToInt())
    fmt.Println(e.Error) //可以通过Error查看是否转换失败了
}
```

# 数字数组与字符串数组之间的任意转换
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
# 结构体数组或map数组转map
取结构体或map中某个字段的值作为key生成map
```go
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
```
```
map[122:0xc000006030]
map[122:[0xc000006030]]

map[122:map[22:122]]
map[122:[map[22:122]]]

map[1:0xc000004a00 2:0xc000004a20]
map[1:[0xc000004a00] 2:[0xc000004a20]]

map[2:{3 2 aa} 3:{2 3 b}]
map[2:[{1 2 a} {3 2 aa}] 3:[{2 3 b}]]
```
# 自动生成map且自动推断map类型(使用reflect性能较低)
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
# 将多种格式的json字符串解析成数组
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/goperate/convert/core/types"
	"log"
)

type ArrayTest struct {
	ArrayNumber1  types.ArrayInt     `json:"array_number_1"`
	ArrayNumber2  types.ArrayInt32   `json:"array_number_2"`
	ArrayNumber3  types.ArrayInt64   `json:"array_number_3"`
	ArrayNumber4  types.ArrayUint    `json:"array_number_4"`
	ArrayNumber5  types.ArrayUint32  `json:"array_number_5"`
	ArrayNumber6  types.ArrayUint64  `json:"array_number_6"`
	ArrayNumber7  types.ArrayFloat32 `json:"array_number_7"`
	ArrayNumber8  types.ArrayFloat64 `json:"array_number_8"`
	ArrayNumber9  types.ArrayInt     `json:"array_number_9"`
	ArrayNumber10 types.ArrayInt     `json:"array_number_10"`

	ArrayString1  types.ArrayKeyword `json:"array_string_1"`
	ArrayString2  types.ArrayKeyword `json:"array_string_2"`
	ArrayString3  types.ArrayKeyword `json:"array_string_3"`
	ArrayString4  types.ArrayKeyword `json:"array_string_4"`
	ArrayString5  types.ArrayKeyword `json:"array_string_5"`
	ArrayString6  types.ArrayKeyword `json:"array_string_6"`
	ArrayString7  types.ArrayKeyword `json:"array_string_7"`
	ArrayString8  types.ArrayKeyword `json:"array_string_8"`
	ArrayString9  types.ArrayKeyword `json:"array_string_9"`
	ArrayString10 types.ArrayKeyword `json:"array_string_10"`

	ArrayString11 types.ArrayString `json:"array_string_11"`
	ArrayString12 types.ArrayString `json:"array_string_12"`
}

const testJsonString = `{
	"array_number_1": 1,
	"array_number_2": [2, 3],
	"array_number_3": "4",
	"array_number_4": ["5", "6"],
	"array_number_5": ["7", 8],
	"array_number_6": "[9,10]",
	"array_number_7": "11,12 ,13",
	"array_number_8": ["14.23, 15.1, 16, 17", "18, 19, 20"],
	"array_number_9": [[21, 22, 23], [24, 25, 26]],
	"array_number_10": {"xx": 27, "cc": {"x": [28, 29]}},

	"array_string_1": "1zcxzv",
	"array_string_2": [2, 3],
	"array_string_3": "4",
	"array_string_4": ["5", "xxx6"],
	"array_string_5": ["7zxx", 8],
	"array_string_6": "[9,10]",
	"array_string_7": "11,12 ,13",
	"array_string_8": ["14.23, 15.1, 16, 17", "18, 19, 20"],
	"array_string_9": [[21, 22, 23], [24, 25, 26]],
	"array_string_10": {"xx": 27, "cc": {"x": [28, 29]}},

	"array_string_11": "最艰难的[asdk\"cee",
	"array_string_12": ["xsssss", "埃松加大\"家乘客都"]
}`

func main() {
	var data ArrayTest
	err := json.Unmarshal([]byte(testJsonString), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", data)
}
```

```
{ArrayNumber1:[1] ArrayNumber2:[2 3] ArrayNumber3:[4] ArrayNumber4:[5 6] ArrayNumber5:[7 8] ArrayNumber6:[9 10] ArrayNumber7:[11 12 13] ArrayNumber8:[14.23 15.1 16 17 18 19 20] ArrayNumber9:[21 22 23 24 25 26] ArrayNumber10:[27 28 29] ArrayString1:[1zcxzv] ArrayString2:[2 3] ArrayString3:[4] ArrayString4:[5 xxx6] ArrayString5:[7zxx 8] ArrayString6:[9 10] ArrayString7:[11 12 13] ArrayString8:[14.23 15.1 16 17 18 19 20] ArrayString9:[21 22 23 24 25 26] ArrayString10:[27 28 29] ArrayString11:[最艰难的[asdk"cee] ArrayString12:[xsssss 埃松加大"家乘客都]}
```
