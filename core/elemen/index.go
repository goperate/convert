package elemen

import (
	"errors"
	"strconv"
)

type Element struct {
	Val   interface{}
	Error error
}

func NewElement(val interface{}) *Element {
	return &Element{
		Val: val,
	}
}

func (t *Element) ToInt() (res int) {
	switch t.Val.(type) {
	case int:
		return t.Val.(int)
	case int32:
		return int(t.Val.(int32))
	case int64:
		return int(t.Val.(int64))
	case uint:
		return int(t.Val.(uint))
	case uint32:
		return int(t.Val.(uint32))
	case uint64:
		return int(t.Val.(uint64))
	case float32:
		return int(t.Val.(float32))
	case float64:
		return int(t.Val.(float64))
	case string:
		res, t.Error = strconv.Atoi(t.Val.(string))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToInt32() (res int32) {
	switch t.Val.(type) {
	case int:
		return int32(t.Val.(int))
	case int32:
		return t.Val.(int32)
	case int64:
		return int32(t.Val.(int64))
	case uint:
		return int32(t.Val.(uint))
	case uint32:
		return int32(t.Val.(uint32))
	case uint64:
		return int32(t.Val.(uint64))
	case float32:
		return int32(t.Val.(float32))
	case float64:
		return int32(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseInt(t.Val.(string), 10, 32)
		return int32(t.Val.(int64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToInt64() (res int64) {
	switch t.Val.(type) {
	case int:
		return int64(t.Val.(int))
	case int32:
		return int64(t.Val.(int32))
	case int64:
		return t.Val.(int64)
	case uint:
		return int64(t.Val.(uint))
	case uint32:
		return int64(t.Val.(uint32))
	case uint64:
		return int64(t.Val.(uint64))
	case float32:
		return int64(t.Val.(float32))
	case float64:
		return int64(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseInt(t.Val.(string), 10, 64)
		return t.Val.(int64)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToUint() (res uint) {
	switch t.Val.(type) {
	case int:
		return uint(t.Val.(int))
	case int32:
		return uint(t.Val.(int32))
	case int64:
		return uint(t.Val.(int64))
	case uint:
		return t.Val.(uint)
	case uint32:
		return uint(t.Val.(uint32))
	case uint64:
		return uint(t.Val.(uint64))
	case float32:
		return uint(t.Val.(float32))
	case float64:
		return uint(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseUint(t.Val.(string), 10, 32)
		return uint(t.Val.(uint64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToUint32() (res uint32) {
	switch t.Val.(type) {
	case int:
		return uint32(t.Val.(int))
	case int32:
		return uint32(t.Val.(int32))
	case int64:
		return uint32(t.Val.(int64))
	case uint:
		return uint32(t.Val.(uint))
	case uint32:
		return t.Val.(uint32)
	case uint64:
		return uint32(t.Val.(uint64))
	case float32:
		return uint32(t.Val.(float32))
	case float64:
		return uint32(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseUint(t.Val.(string), 10, 32)
		return uint32(t.Val.(uint64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToUint64() (res uint64) {
	switch t.Val.(type) {
	case int:
		return uint64(t.Val.(int))
	case int32:
		return uint64(t.Val.(int32))
	case int64:
		return uint64(t.Val.(int64))
	case uint:
		return uint64(t.Val.(uint))
	case uint32:
		return uint64(t.Val.(uint32))
	case uint64:
		return t.Val.(uint64)
	case float32:
		return uint64(t.Val.(float32))
	case float64:
		return uint64(t.Val.(float64))
	case string:
		res, t.Error = strconv.ParseUint(t.Val.(string), 10, 64)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToFloat32() (res float32) {
	switch t.Val.(type) {
	case int:
		return float32(t.Val.(int))
	case int32:
		return float32(t.Val.(int32))
	case int64:
		return float32(t.Val.(int64))
	case uint:
		return float32(t.Val.(uint))
	case uint32:
		return float32(t.Val.(uint32))
	case uint64:
		return float32(t.Val.(uint64))
	case float32:
		return t.Val.(float32)
	case float64:
		return float32(t.Val.(float64))
	case string:
		t.Val, t.Error = strconv.ParseFloat(t.Val.(string), 32)
		return float32(t.Val.(float64))
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToFloat64() (res float64) {
	switch t.Val.(type) {
	case int:
		return float64(t.Val.(int))
	case int32:
		return float64(t.Val.(int32))
	case int64:
		return float64(t.Val.(int64))
	case uint:
		return float64(t.Val.(uint))
	case uint32:
		return float64(t.Val.(uint32))
	case uint64:
		return float64(t.Val.(uint64))
	case float32:
		return float64(t.Val.(float32))
	case float64:
		return t.Val.(float64)
	case string:
		res, t.Error = strconv.ParseFloat(t.Val.(string), 64)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}

func (t *Element) ToString() (res string) {
	switch t.Val.(type) {
	case int:
		return strconv.Itoa(t.Val.(int))
	case int32:
		return strconv.FormatInt(int64(t.Val.(int32)), 10)
	case int64:
		return strconv.FormatInt(t.Val.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(t.Val.(uint)), 10)
	case uint32:
		return strconv.FormatUint(uint64(t.Val.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(t.Val.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(t.Val.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t.Val.(float64), 'f', -1, 64)
	case string:
		return t.Val.(string)
	default:
		t.Error = errors.New("不支持的数据类型")
	}
	return
}
