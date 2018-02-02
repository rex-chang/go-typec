package typec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inFloat32 = float32(1234.56)
var inFloat64 = float64(1234.56)

var inStruct = struct {
	t string
}{t: "123"}

func TestInt2String(t *testing.T) {
	//in test
	var in = 123
	var inBool = true
	var inInt8 = int8(127)
	var inUint8 = uint8(244)
	var inString = "123a"
	out, _ := In(in).String()
	outFloat32Ret, _ := In(inFloat32).String()
	outFloat64Ret, _ := In(inFloat64).String()
	outBool, _ := In(inBool).String()
	outInt8, _ := In(inInt8).String()
	outUint8, _ := In(inUint8).String()
	outString, _ := In(inString).String()
	_, err := In(inStruct).String()
	assert.Equal(t, "123", out, "convert error ")
	assert.Equal(t, "1234.56", outFloat32Ret, " float32 convert error")
	assert.Equal(t, "1234.56", outFloat64Ret, " float64 convert error")
	assert.Equal(t, "true", outBool, " bool convert error")
	assert.Equal(t, "127", outInt8, " int convert error")
	assert.Equal(t, "244", outUint8, " uint convert error")
	assert.Equal(t, "123a", outString, " string convert error")
	assert.Error(t, err)
}

func TestAny2Int(t *testing.T) {
	//in test
	var in uint = 123
	var inInt8 = int8(123)
	var inString = "123"
	//expect type expeced
	var expect = int(123)

	outUint, _ := In(in).Int()
	outInt8, _ := In(inInt8).Int()
	outStr, _ := In(inString).Int()
	outFloat32Ret, _ := In(inFloat32).Int()
	outFloat64Ret, _ := In(inFloat64).Int()
	_, err := In(inStruct).Int()
	assert.Equal(t, expect, outUint, "convert error ")
	assert.Equal(t, expect, outStr, "string convert error")
	assert.Equal(t, expect, outInt8, "string convert error")
	assert.Equal(t, int(1234), outFloat32Ret, "string convert error")
	assert.Equal(t, int(1234), outFloat64Ret, "string convert error")
	assert.Error(t, err)

}
