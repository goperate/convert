package dict

import (
	"fmt"
	"reflect"
)

func StructList2Map(data interface{}, keyName interface{}) interface{} {
	value := reflect.ValueOf(data)
	typ := value.Type()
	keyNameValue := reflect.ValueOf(keyName)
	if value.Kind() != reflect.Slice {
		panic("data is not array")
	}
	elemType := typ.Elem()
	elemType2 := elemType
	isPtr := false
	if elemType2.Kind() == reflect.Ptr {
		isPtr = true
		elemType2 = elemType2.Elem()
	}
	var keyType reflect.Type
	isMap := false
	switch elemType2.Kind() {
	case reflect.Struct:
		key, ok := elemType2.FieldByName(keyNameValue.String())
		if !ok {
			panic(keyNameValue.String() + " is not in struct field")
		}
		keyType = key.Type
	case reflect.Map:
		isMap = true
		keyType = elemType2.Elem()
	default:
		panic("elem is not struct or map")
	}
	fmt.Println(elemType2)

	resValue := reflect.MakeMap(reflect.MapOf(keyType, elemType))
	for i := 0; i < value.Len(); i++ {
		v := value.Index(i)
		vv := v
		if isPtr {
			vv = v.Elem()
		}
		var key reflect.Value
		if isMap {
			key = vv.MapIndex(keyNameValue)
		} else {
			key = vv.FieldByName(keyNameValue.String())
		}
		resValue.SetMapIndex(key, v)
	}
	return resValue.Interface()
}
