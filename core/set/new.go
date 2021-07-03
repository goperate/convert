package set

import (
	"fmt"
	"reflect"
)

func new_(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Map:
		return reflect.MakeMap(t)
	case reflect.Slice:
		return reflect.MakeSlice(t, 0, 0)
	case reflect.Struct:
		return reflect.New(t).Elem()
	case reflect.Ptr:
		return new_(t.Elem()).Addr()
	default:
		return reflect.New(t).Elem()
	}
}

//New 初始化"零值"
func New(val interface{}) {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr {
		panic("val参数必须传入指针")
	}
	t := v.Type().Elem()
	v = v.Elem()
	if !v.CanSet() {
		panic(fmt.Sprintf("%v不支持设置值", v.Kind()))
	}
	if !v.IsZero() {
		return
	}
	v.Set(new_(t))
}
