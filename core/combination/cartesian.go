package combination

import (
	"reflect"
)

// CartesianProduct 从多个数组中每个数组取一个生成一个组合
type CartesianProduct struct {
	index  int
	slice  []reflect.Value
	nums   [][2]int
	num    int // 结果组合数
	length int // 集合数
}

func NewCartesianProduct(val interface{}, args ...interface{}) *CartesianProduct {
	res := new(CartesianProduct)
	res.num = 1
	for _, v := range append([]interface{}{val}, args...) {
		res.length += 1
		vv := reflect.ValueOf(v)
		n := vv.Len()
		if n == 0 {
			panic("不能传入空数组")
		}
		res.slice = append(res.slice, vv)
		res.nums = append(res.nums, [2]int{res.num, n})
		res.num *= n
	}
	return res
}

// Next 每次取一个组合, 取完一遍会从头开始
func (t *CartesianProduct) Next() (res []interface{}) {
	res = make([]interface{}, t.length)
	for i := t.length - 1; i >= 0; i-- {
		nums := t.nums[i]
		v := t.slice[i].Index(t.index / nums[0] % nums[1])
		res[i] = v.Interface()
	}
	t.index = (t.index + 1) % t.num
	return
}
