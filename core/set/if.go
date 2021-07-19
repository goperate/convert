package set

import "reflect"

func IF(val interface{}, condition bool, tv, fv interface{}) {
	v1 := reflect.ValueOf(val).Elem()
	if condition {
		v1.Set(reflect.ValueOf(tv))
	} else {
		v1.Set(reflect.ValueOf(fv))
	}
}
