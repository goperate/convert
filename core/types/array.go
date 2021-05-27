package types

import (
	"encoding/json"
	"github.com/goperate/convert/core/array"
	"strings"
)

type ArrayInt []int
type ArrayInt32 []int32
type ArrayInt64 []int64
type ArrayUint []uint
type ArrayUint32 []uint32
type ArrayUint64 []uint64
type ArrayFloat32 []float32
type ArrayFloat64 []float64

type ArrayKeyword []string
type ArrayString []string

func toElementArray(b []byte) (res []string) {
	ss := strings.Split(string(b), ",")
	for _, v := range ss {
		vv := strings.Split(v, ":")
		v = strings.Trim(vv[len(vv)-1], "[ ]\"}\n")
		if v != "" {
			res = append(res, v)
		}
	}
	return
}

func (t *ArrayInt) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToIntArray()...)
	return any.Error
}

func (t *ArrayInt32) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToInt32Array()...)
	return any.Error
}

func (t *ArrayInt64) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToInt64Array()...)
	return any.Error
}

func (t *ArrayUint) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToUintArray()...)
	return any.Error
}

func (t *ArrayUint32) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToUint32Array()...)
	return any.Error
}

func (t *ArrayUint64) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToUint64Array()...)
	return any.Error
}

func (t *ArrayFloat32) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToFloat32Array()...)
	return any.Error
}

func (t *ArrayFloat64) UnmarshalJSON(b []byte) (err error) {
	any := array.NewArray(toElementArray(b))
	*t = append(*t, any.ToFloat64Array()...)
	return any.Error
}

func (t *ArrayKeyword) UnmarshalJSON(b []byte) (err error) {
	*t = append(*t, toElementArray(b)...)
	return
}

func (t *ArrayString) UnmarshalJSON(b []byte) (err error) {
	n := len(b)
	if n == 0 {
		return
	}
	if b[0] == '[' {
		var val []string
		err = json.Unmarshal(b, &val)
		*t = append(*t, val...)
	} else {
		var val string
		err = json.Unmarshal(b, &val)
		*t = append(*t, val)
	}
	return
}
