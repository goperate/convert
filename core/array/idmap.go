package array

import (
	"reflect"
)

type ObjArray struct {
	value        reflect.Value
	typ          reflect.Type
	keyNameValue reflect.Value
	keyType      reflect.Type
	elemType     reflect.Type
	isPtr        bool
	isMap        bool
}

func NewObjArray(data interface{}, keyName interface{}) (res *ObjArray) {
	res = new(ObjArray)
	res.value = reflect.ValueOf(data)
	res.typ = res.value.Type()
	res.keyNameValue = reflect.ValueOf(keyName)
	if res.value.Kind() != reflect.Slice {
		panic("data is not array")
	}
	res.elemType = res.typ.Elem()
	elemType2 := res.elemType
	if elemType2.Kind() == reflect.Ptr {
		res.isPtr = true
		elemType2 = elemType2.Elem()
	}
	switch elemType2.Kind() {
	case reflect.Struct:
		key, ok := elemType2.FieldByName(res.keyNameValue.String())
		if !ok {
			panic(res.keyNameValue.String() + " is not in struct field")
		}
		res.keyType = key.Type
	case reflect.Map:
		res.isMap = true
		res.keyType = elemType2.Elem()
	default:
		panic("elem is not struct or map")
	}
	return
}

func (t *ObjArray) elemGetKey(v reflect.Value) (key reflect.Value) {
	if t.isPtr {
		v = v.Elem()
	}
	if t.isMap {
		key = v.MapIndex(t.keyNameValue)
	} else {
		key = v.FieldByName(t.keyNameValue.String())
	}
	return
}

func (t *ObjArray) ToIdMap() interface{} {
	resValue := reflect.MakeMap(reflect.MapOf(t.keyType, t.elemType))
	for i := 0; i < t.value.Len(); i++ {
		v := t.value.Index(i)
		resValue.SetMapIndex(t.elemGetKey(v), v)
	}
	return resValue.Interface()
}

func (t *ObjArray) ToIdMapArray() interface{} {
	resValue := reflect.MakeMap(reflect.MapOf(t.keyType, reflect.SliceOf(t.elemType)))
	for i := 0; i < t.value.Len(); i++ {
		v := t.value.Index(i)
		key := t.elemGetKey(v)
		if resValue.MapIndex(key).Kind() == reflect.Invalid {
			resValue.SetMapIndex(key, reflect.MakeSlice(reflect.SliceOf(t.elemType), 0, 0))
		}
		resValue.SetMapIndex(key, reflect.Append(resValue.MapIndex(key), v))
	}
	return resValue.Interface()
}
