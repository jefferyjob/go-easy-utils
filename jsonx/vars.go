package jsonx

import (
	"errors"
	"strconv"
)

var (
	// ErrPoint 不是指针类型
	// ErrPoint indicates that the argument is not a non-nil pointer to a struct
	ErrPoint = errors.New("the argument to ToStruct must be a non-nil pointer to a struct")

	// ErrNotBool 不是Bool类型
	// ErrNotBool = errors.New("cannot parse bool, value is not a bool")
	// ErrNotMap 不是Map类型
	// ErrNotMap = errors.New("cannot parse map, value is not a map")
	// ErrNotSlice 不是Slice类型
	// ErrNotSlice = errors.New("cannot parse slice, value is not a slice")

	// ErrSyntax 指示值不具有目标类型的正确语法
	ErrSyntax = strconv.ErrSyntax
	// ErrUnsupported 不支持的类型
	ErrUnsupported = errors.New("unsupported type")

	// ErrValOut 指示值超出目标类型可表示的范围
	// ErrValOut indicates that the value is out of range for the target type
	ErrValOut = errors.New("value out of range")
)
