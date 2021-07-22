package sorted

import (
	"reflect"
	"sort"
)

func Sort(slice interface{}, ascending ...bool) {
	val := reflect.ValueOf(slice)
	kind := val.Type().Elem().Kind()
	asc := len(ascending) == 0 || ascending[0]
	sort.Slice(slice, func(i, j int) bool {
		a, b := val.Index(i), val.Index(j)
		if !asc {
			b, a = a, b
		}
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return a.Int() < b.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return a.Uint() < b.Uint()
		case reflect.Float32, reflect.Float64:
			return a.Float() < b.Float()
		default:
			return a.String() < b.String()
		}
	})
}
