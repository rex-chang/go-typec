package typec

import (
	"errors"
	"reflect"
	"strconv"
)

//TypeConv 转换实体
type TypeConv struct {
	in interface{}
}

//In 设置一个值
func In(in interface{}) (tc *TypeConv) {
	tc = &TypeConv{in}
	// tc.in = in
	return tc
}

//Int 返回一个Int类型
func (t *TypeConv) Int() (out int, err error) {
	switch t.in.(type) {
	case int8, int32, int, int64:
		return int(reflect.ValueOf(t.in).Int()), nil
	case uint8, uint16, uint, uint32, uint64:
		return int(reflect.ValueOf(t.in).Uint()), nil
	case float32, float64:
		return int(reflect.ValueOf(t.in).Float()), nil
	case string:
		return strconv.Atoi(t.in.(string))
	}
	return 0, errors.New("convert failed")
}

func (t *TypeConv) String() (out string, err error) {
	if str, ok := (t.in).(string); ok {
		return str, nil
	}
	switch t.in.(type) {
	case int:
		return strconv.Itoa(t.in.(int)), nil
	case int8, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(t.in).Int(), 10), nil
	case uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(t.in).Uint(), 10), nil
	case float32:
		return strconv.FormatFloat(reflect.ValueOf(t.in).Float(), 'g', -1, 32), nil
	case float64:
		return strconv.FormatFloat(reflect.ValueOf(t.in).Float(), 'g', -1, 64), nil
	case bool:
		return strconv.FormatBool(t.in.(bool)), nil
	}
	return "", errors.New("convert failed")
}
