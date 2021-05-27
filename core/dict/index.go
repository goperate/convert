package dict

import (
	"errors"
	"fmt"
	"github.com/goperate/convert/core/element"
	"reflect"
)

type Dict struct {
	Data  interface{}
	Error error
	Value reflect.Value
}

func NewDict(m interface{}) (res *Dict) {
	res = &Dict{
		Data:  m,
		Value: reflect.ValueOf(m),
	}
	if m == nil && res.Value.Kind() != reflect.Invalid {
		res.Value = reflect.MakeMap(res.Value.Type())
		res.Data = res.Value.Interface()
	}
	return
}

func (t *Dict) initArray(val reflect.Value) (res reflect.Value) {
	res = reflect.MakeSlice(reflect.SliceOf(val.Type()), 1, 1)
	res.Index(0).Set(val)
	return
}

func (t *Dict) a2b(a, b reflect.Value) reflect.Value {
	iter := a.MapRange()
	for iter.Next() {
		b.SetMapIndex(iter.Key(), iter.Value())
	}
	return b
}

func (t *Dict) toInterfaceSlice(old reflect.Value) (res reflect.Value) {
	res = reflect.ValueOf([]interface{}{})
	for i := 0; i < old.Len(); i++ {
		res = reflect.Append(res, old.Index(i))
	}
	return
}

func (t *Dict) append(old, new reflect.Value) reflect.Value {
	switch old.Kind() {
	case reflect.Slice:
		ote := old.Type().Elem()
		if new.Kind() == reflect.Slice {
			if ote.Kind() != reflect.Interface && old.Type() != new.Type() {
				old = t.toInterfaceSlice(old)
				new = t.toInterfaceSlice(new)
			}
			return reflect.AppendSlice(old, new)
		}
		if ote.Kind() != reflect.Interface && ote != new.Type() {
			old = t.toInterfaceSlice(old)
		}
		return reflect.Append(old, new)
	case reflect.Interface:
		return t.append(old.Elem(), new)
	}
	return new
}

func (t *Dict) setMap(out, key, val reflect.Value) reflect.Value {
	switch out.Kind() {
	case reflect.Invalid:
		out = t.setMap(reflect.MakeMap(reflect.MapOf(key.Type(), val.Type())), key, val)
	case reflect.Interface:
		out = t.setMap(out.Elem(), key, val)
	case reflect.Map:
		newOut := out
		reset := false
		keyType := out.Type().Key()
		elemType := out.Type().Elem()
		b := reflect.TypeOf(make(map[interface{}]interface{}))
		if keyType.Kind() != reflect.Interface && keyType != key.Type() {
			newOut = reflect.MakeMap(reflect.MapOf(b.Key(), newOut.Type().Elem()))
			reset = true
		}
		if elemType.Kind() != reflect.Interface && elemType != val.Type() {
			newOut = reflect.MakeMap(reflect.MapOf(newOut.Type().Key(), b.Elem()))
			reset = true
		}
		if reset {
			out = t.a2b(out, newOut)
		}
		out.SetMapIndex(key, t.append(out.MapIndex(key), val))
	}
	return out
}

func (t *Dict) initMap(val reflect.Value, key reflect.Value, args ...interface{}) (res reflect.Value) {
	for i := len(args) - 1; i >= 0; i-- {
		key := reflect.ValueOf(args[i])
		val = t.setMap(res, key, val)
	}
	res = t.setMap(res, key, val)
	return
}

func (t *Dict) set(out, val reflect.Value, isArray bool, key reflect.Value, args ...interface{}) (res reflect.Value) {
	if isArray {
		val = t.initArray(val)
	}
	switch out.Kind() {
	case reflect.Interface:
		res = t.set(out.Elem(), val, false, key, args...)
	case reflect.Map:
		if len(args) == 0 {
			res = t.setMap(out, key, val)
			return
		}
		newOut := out
		keyType := out.Type().Key()
		b := reflect.TypeOf(make(map[interface{}]interface{}))
		if keyType.Kind() != reflect.Interface && keyType != key.Type() {
			newOut = reflect.MakeMap(reflect.MapOf(b.Key(), newOut.Type().Elem()))
			out = t.a2b(out, newOut)
		}
		next := out.MapIndex(key)
		res = t.setMap(out, key, t.set(next, val, false, reflect.ValueOf(args[0]), args[1:]...))
	case reflect.Invalid:
		res = t.initMap(val, key, args...)
	}
	return
}

func (t *Dict) Set(val interface{}, isArray bool, key interface{}, args ...interface{}) {
	t.Value = t.set(t.Value, reflect.ValueOf(val), isArray, reflect.ValueOf(key), args...)
	t.Data = t.Value.Interface()
}

func (t *Dict) get(value, key reflect.Value, args ...interface{}) (res reflect.Value) {
	switch value.Kind() {
	case reflect.Map:
		if len(args) == 0 {
			return value.MapIndex(key)
		}
		return t.get(value.MapIndex(key), reflect.ValueOf(args[0]), args[1:]...)
	case reflect.Interface:
		return t.get(value.Elem(), key, args...)
	case reflect.Slice:
		if len(args) == 0 {
			return value.Index(int(key.Int()))
		}
		return t.get(value.Index(int(key.Int())), reflect.ValueOf(args[0]), args[1:]...)
	}
	t.Error = errors.New("不支持的数据类型")
	return
}

func (t *Dict) Get(key interface{}, args ...interface{}) interface{} {
	defer func() {
		t.Error = errors.New(fmt.Sprintf("%v", recover()))
	}()
	out := t.get(t.Value, reflect.ValueOf(key), args...)
	if out.Kind() == reflect.Invalid {
		return nil
	}
	return out.Interface()
}

func (t *Dict) GetElement(key interface{}, args ...interface{}) *element.Element {
	return element.NewElement(t.Get(key, args...))
}
