package array

import (
	"reflect"
)

type _field struct {
	NameValue reflect.Value
	Type      reflect.Type
}

type _return struct {
	Active    bool
	Out       reflect.Value
	IsRun     bool
	KeyIndex  int
	KeyIndex2 int
}

type ObjArray struct {
	value         reflect.Value
	typ           reflect.Type
	keys          []_field
	lenKeys       int
	elemType      reflect.Type
	elemType2     reflect.Type
	idMap         _return
	idMapArray    _return
	idMapOne      _return
	idMapOneArray _return
	idArray       _return
}

// NewObjArray keyName 默认为 "Id"
func NewObjArray(data interface{}, keyName interface{}) (res *ObjArray) {
	res = new(ObjArray)
	res.value = reflect.ValueOf(data)
	res.typ = res.value.Type()
	if res.value.Kind() != reflect.Slice {
		panic("data is not array")
	}
	res.elemType = res.typ.Elem()
	res.elemType2 = res.elemType
	if res.elemType.Kind() == reflect.Ptr {
		res.elemType2 = res.elemType.Elem()
	}
	res.getKeyType(keyName)
	return
}

func (t *ObjArray) getKeyType(name interface{}) int {
	t.lenKeys += 1
	var key _field
	key.NameValue = reflect.ValueOf(name)
	switch t.elemType2.Kind() {
	case reflect.Struct:
		field, ok := t.elemType2.FieldByName(key.NameValue.String())
		if !ok {
			panic(key.NameValue.String() + " is not in struct field")
		}
		key.Type = field.Type
	case reflect.Map:
		key.Type = t.elemType2.Elem()
	default:
		panic("elem is not struct or map")
	}
	t.keys = append(t.keys, key)
	return t.lenKeys - 1
}

// IdMap 生成{1: struct}
func (t *ObjArray) IdMap(key ...interface{}) *ObjArray {
	if len(key) > 0 {
		t.idMap.KeyIndex = t.getKeyType(key[0])
	}
	t.idMap.Active = true
	t.idMap.Out = reflect.MakeMap(reflect.MapOf(
		t.keys[t.idMap.KeyIndex].Type,
		t.elemType,
	))
	return t
}

// IdMapArray 生成{1: [struct]}
func (t *ObjArray) IdMapArray(key ...interface{}) *ObjArray {
	if len(key) > 0 {
		t.idMapArray.KeyIndex = t.getKeyType(key[0])
	}
	t.idMapArray.Active = true
	t.idMapArray.Out = reflect.MakeMap(reflect.MapOf(
		t.keys[t.idMapArray.KeyIndex].Type,
		reflect.SliceOf(t.elemType),
	))
	return t
}

// IdMapOne 生成{1: 2}
func (t *ObjArray) IdMapOne(key interface{}, key2 ...interface{}) *ObjArray {
	t.getKeyType(key)
	if len(key2) > 0 {
		t.idMapOne.KeyIndex = t.lenKeys - 1
		t.idMapOne.KeyIndex2 = t.getKeyType(key2[0])
	} else {
		t.idMapOne.KeyIndex2 = t.lenKeys - 1
	}
	t.idMapOne.Active = true
	t.idMapOne.Out = reflect.MakeMap(reflect.MapOf(
		t.keys[t.idMapOne.KeyIndex].Type,
		t.keys[t.idMapOne.KeyIndex2].Type,
	))
	return t
}

// IdMapOneArray 生成{1: [1, 2]}
func (t *ObjArray) IdMapOneArray(key interface{}, key2 ...interface{}) *ObjArray {
	t.getKeyType(key)
	if len(key2) > 0 {
		t.idMapOneArray.KeyIndex = t.lenKeys - 1
		t.idMapOneArray.KeyIndex2 = t.getKeyType(key2[0])
	} else {
		t.idMapOneArray.KeyIndex2 = t.lenKeys - 1
	}
	t.idMapOneArray.Active = true
	t.idMapOneArray.Out = reflect.MakeMap(reflect.MapOf(
		t.keys[t.idMapOneArray.KeyIndex].Type,
		reflect.SliceOf(t.keys[t.idMapOneArray.KeyIndex2].Type),
	))
	return t
}

// IdArray 生成[1, 2, 3]
func (t *ObjArray) IdArray(key ...interface{}) *ObjArray {
	if len(key) > 0 {
		t.idArray.KeyIndex = t.getKeyType(key[0])
	}
	t.idArray.Active = true
	t.idArray.Out = reflect.MakeSlice(reflect.SliceOf(t.keys[t.idArray.KeyIndex].Type), 0, 0)
	return t
}

func (t *ObjArray) elemGetKey(v reflect.Value, index int) (key reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Map {
		key = v.MapIndex(t.keys[index].NameValue)
	} else {
		key = v.FieldByName(t.keys[index].NameValue.String())
	}
	return
}

func (t *ObjArray) run() *ObjArray {
	for i := 0; i < t.value.Len(); i++ {
		v := t.value.Index(i)
		if t.idMap.Active && !t.idMap.IsRun {
			key := t.elemGetKey(v, t.idMap.KeyIndex)
			t.idMap.Out.SetMapIndex(key, v)
		}
		if t.idMapArray.Active && !t.idMapArray.IsRun {
			key := t.elemGetKey(v, t.idMapArray.KeyIndex)
			if t.idMapArray.Out.MapIndex(key).Kind() == reflect.Invalid {
				t.idMapArray.Out.SetMapIndex(key, reflect.MakeSlice(reflect.SliceOf(t.elemType), 0, 0))
			}
			t.idMapArray.Out.SetMapIndex(key, reflect.Append(t.idMapArray.Out.MapIndex(key), v))
		}
		if t.idMapOne.Active && !t.idMapOne.IsRun {
			key := t.elemGetKey(v, t.idMapOne.KeyIndex)
			t.idMapOne.Out.SetMapIndex(key, t.elemGetKey(v, t.idMapOne.KeyIndex2))
		}
		if t.idMapOneArray.Active && !t.idMapOneArray.IsRun {
			key := t.elemGetKey(v, t.idMapOneArray.KeyIndex)
			if t.idMapOneArray.Out.MapIndex(key).Kind() == reflect.Invalid {
				t.idMapOneArray.Out.SetMapIndex(key, reflect.MakeSlice(reflect.SliceOf(t.keys[t.idMapOneArray.KeyIndex2].Type), 0, 0))
			}
			t.idMapOneArray.Out.SetMapIndex(key, reflect.Append(t.idMapOneArray.Out.MapIndex(key), t.elemGetKey(v, t.idMapOneArray.KeyIndex2)))
		}
		if t.idArray.Active && !t.idArray.IsRun {
			key := t.elemGetKey(v, t.idArray.KeyIndex)
			t.idArray.Out = reflect.Append(t.idArray.Out, key)
		}
	}
	t.idMap.IsRun = t.idMap.Active
	t.idMapArray.IsRun = t.idMapArray.Active
	t.idMapOne.IsRun = t.idMapOne.Active
	t.idMapOneArray.IsRun = t.idMapOneArray.Active
	t.idArray.IsRun = t.idArray.Active
	return t
}

func (t *ObjArray) ToIdMap() interface{} {
	if !t.idMap.Active {
		t.IdMap()
	}
	if !t.idMap.IsRun {
		t.run()
	}
	return t.idMap.Out.Interface()
}

func (t *ObjArray) ToIdMapArray() interface{} {
	if !t.idMapArray.Active {
		t.IdMapArray()
	}
	if !t.idMapArray.IsRun {
		t.run()
	}
	return t.idMapArray.Out.Interface()
}

func (t *ObjArray) ToIdMapOne(key ...interface{}) interface{} {
	if !t.idMapOne.Active {
		t.IdMapOne(key[0])
	}
	if !t.idMapOne.IsRun {
		t.run()
	}
	return t.idMapOne.Out.Interface()
}

func (t *ObjArray) ToIdMapOneArray(key ...interface{}) interface{} {
	if !t.idMapOneArray.Active {
		t.IdMapOneArray(key[0])
	}
	if !t.idMapOneArray.IsRun {
		t.run()
	}
	return t.idMapOneArray.Out.Interface()
}

func (t *ObjArray) ToIdArray() interface{} {
	if !t.idArray.Active {
		t.IdArray()
	}
	if !t.idArray.IsRun {
		t.run()
	}
	return t.idArray.Out.Interface()
}
