package array

import (
	"github.com/goperate/convert/core/element"
	"reflect"
)

type Array struct {
	Data          interface{}
	Error         error
	deduplication bool
	dict          interface{}
	value         reflect.Value
}

func NewArray(val interface{}, ded ...bool) *Array {
	value := reflect.ValueOf(val)
	if value.Kind() != reflect.Slice {
		return nil
	}
	return &Array{
		Data:          val,
		value:         value,
		deduplication: len(ded) > 0 && ded[0],
	}
}

func (t *Array) Deduplication(v bool) *Array {
	t.deduplication = v
	return t
}

func (t *Array) base(f func(v interface{})) {
	for i := 0; i < t.value.Len(); i++ {
		f(t.value.Index(i).Interface())
	}
}

func (t *Array) ToIntArray() (res []int) {
	m := make(map[int]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToInt()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToInt32Array() (res []int32) {
	m := make(map[int32]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToInt32()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToInt64Array() (res []int64) {
	m := make(map[int64]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToInt64()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToUintArray() (res []uint) {
	m := make(map[uint]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToUint()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToUint32Array() (res []uint32) {
	m := make(map[uint32]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToUint32()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToUint64Array() (res []uint64) {
	m := make(map[uint64]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToUint64()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToFloat32Array() (res []float32) {
	m := make(map[float32]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToFloat32()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToFloat64Array() (res []float64) {
	m := make(map[float64]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToFloat64()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToStringArray() (res []string) {
	m := make(map[string]int)
	f := func(v interface{}) {
		any := element.NewElement(v)
		vv := any.ToString()
		if !t.deduplication || m[vv] == 0 {
			res = append(res, vv)
		}
		m[vv] += 1
		if any.Error != nil {
			t.Error = any.Error
		}
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToInterfaceArray() (res []interface{}) {
	m := make(map[interface{}]int)
	f := func(v interface{}) {
		if !t.deduplication || m[v] == 0 {
			res = append(res, v)
		}
		m[v] += 1
	}
	t.base(f)
	t.dict = m
	return
}

func (t *Array) ToIntMap() (res map[int]int) {
	switch t.dict.(type) {
	case map[int]int:
		return t.dict.(map[int]int)
	default:
		t.ToIntArray()
		return t.ToIntMap()
	}
}

func (t *Array) ToInt32Map() (res map[int32]int) {
	switch t.dict.(type) {
	case map[int32]int:
		return t.dict.(map[int32]int)
	default:
		t.ToInt32Array()
		return t.ToInt32Map()
	}
}

func (t *Array) ToInt64Map() (res map[int64]int) {
	switch t.dict.(type) {
	case map[int64]int:
		return t.dict.(map[int64]int)
	default:
		t.ToInt64Array()
		return t.ToInt64Map()
	}
}

func (t *Array) ToUintMap() (res map[uint]int) {
	switch t.dict.(type) {
	case map[uint]int:
		return t.dict.(map[uint]int)
	default:
		t.ToUintArray()
		return t.ToUintMap()
	}
}

func (t *Array) ToUint32Map() (res map[uint32]int) {
	switch t.dict.(type) {
	case map[uint32]int:
		return t.dict.(map[uint32]int)
	default:
		t.ToUint32Array()
		return t.ToUint32Map()
	}
}

func (t *Array) ToUint64Map() (res map[uint64]int) {
	switch t.dict.(type) {
	case map[uint64]int:
		return t.dict.(map[uint64]int)
	default:
		t.ToUint64Array()
		return t.ToUint64Map()
	}
}

func (t *Array) ToFloat32Map() (res map[float32]int) {
	switch t.dict.(type) {
	case map[float32]int:
		return t.dict.(map[float32]int)
	default:
		t.ToFloat32Array()
		return t.ToFloat32Map()
	}
}

func (t *Array) ToFloat64Map() (res map[float64]int) {
	switch t.dict.(type) {
	case map[float64]int:
		return t.dict.(map[float64]int)
	default:
		t.ToFloat64Array()
		return t.ToFloat64Map()
	}
}

func (t *Array) ToStringMap() (res map[string]int) {
	switch t.dict.(type) {
	case map[string]int:
		return t.dict.(map[string]int)
	default:
		t.ToStringArray()
		return t.ToStringMap()
	}
}

func (t *Array) ToInterfaceMap() (res map[interface{}]int) {
	switch t.dict.(type) {
	case map[interface{}]int:
		return t.dict.(map[interface{}]int)
	default:
		t.ToInterfaceArray()
		return t.ToInterfaceMap()
	}
}
