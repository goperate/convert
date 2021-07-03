package set

import "reflect"

func Default(field, _default interface{}) {
	v1 := reflect.ValueOf(field).Elem()
	v2 := reflect.ValueOf(_default)
	if v1.IsZero() {
		v1.Set(v2)
	}
}
