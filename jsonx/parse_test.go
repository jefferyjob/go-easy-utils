package jsonx

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMapKey(t *testing.T) {
	var (
		tStr   = reflect.TypeOf("")
		tMyStr = reflect.TypeOf(mapKeyStr(""))
		tBool  = reflect.TypeOf(false)
		tInt   = reflect.TypeOf(int(0))
		tI8    = reflect.TypeOf(int8(0))
		tI16   = reflect.TypeOf(int16(0))
		tI32   = reflect.TypeOf(int32(0))
		tI64   = reflect.TypeOf(int64(0))
		tUint  = reflect.TypeOf(uint(0))
		tU8    = reflect.TypeOf(uint8(0))
		tU16   = reflect.TypeOf(uint16(0))
		tU32   = reflect.TypeOf(uint32(0))
		tU64   = reflect.TypeOf(uint64(0))
		tF32   = reflect.TypeOf(float32(0))
		tF64   = reflect.TypeOf(float64(0))
	)

	testCases := []struct {
		name    string
		key     string
		keyType reflect.Type
		want    any
		wantErr string
	}{
		{name: "string", key: "k", keyType: tStr, want: "k"},
		{name: "string命名类型", key: "k", keyType: tMyStr, want: mapKeyStr("k")},
		{name: "bool", key: "true", keyType: tBool, want: true},
		{name: "bool解析失败", key: "yes", keyType: tBool, wantErr: `无法将 "yes" 转为 bool`},
		{name: "int", key: "-7", keyType: tInt, want: -7},
		{name: "int8", key: "42", keyType: tI8, want: int8(42)},
		{name: "int16", key: "42", keyType: tI16, want: int16(42)},
		{name: "int32", key: "42", keyType: tI32, want: int32(42)},
		{name: "int64", key: "42", keyType: tI64, want: int64(42)},
		{name: "int8溢出", key: "300", keyType: tI8, wantErr: `无法将 "300" 转为 int8`},
		{name: "int解析失败", key: "abc", keyType: tInt, wantErr: `无法将 "abc" 转为 int`},
		{name: "uint", key: "42", keyType: tUint, want: uint(42)},
		{name: "uint8", key: "42", keyType: tU8, want: uint8(42)},
		{name: "uint16", key: "42", keyType: tU16, want: uint16(42)},
		{name: "uint32", key: "42", keyType: tU32, want: uint32(42)},
		{name: "uint64", key: "42", keyType: tU64, want: uint64(42)},
		{name: "uint为负数", key: "-1", keyType: tUint, wantErr: `无法将 "-1" 转为 uint`},
		{name: "float32", key: "1.5", keyType: tF32, want: float32(1.5)},
		{name: "float64", key: "1.5", keyType: tF64, want: 1.5},
		{name: "float64解析失败", key: "x", keyType: tF64, wantErr: `无法将 "x" 转为 float64`},
		{name: "不支持的key类型", key: "k", keyType: reflect.TypeOf(complex128(0)), wantErr: "不支持的 map key 类型 complex128"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseMapKey(tc.key, tc.keyType, "path")
			if tc.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.want, got.Interface())
		})
	}
}

func TestAssignValueStructFromRawMessage(t *testing.T) {
	f := reflect.New(reflect.TypeOf(Address{})).Elem()
	err := assignValue(f, json.RawMessage(`{"country": "China", "city": "Beijing"}`), "addr")
	require.NoError(t, err)
	require.Equal(t, Address{Country: "China", City: "Beijing"}, f.Interface().(Address))
}

func TestAssignValueStructMarshalError(t *testing.T) {
	f := reflect.New(reflect.TypeOf(Address{})).Elem()
	err := assignValue(f, make(chan int), "addr")
	require.Error(t, err)
	require.Contains(t, err.Error(), "addr")
}

func TestAssignValueUnsupportedKind(t *testing.T) {
	var ch chan int
	f := reflect.ValueOf(&ch).Elem()
	err := assignValue(f, 1, "ch")
	require.ErrorIs(t, err, ErrUnsupported)
}
