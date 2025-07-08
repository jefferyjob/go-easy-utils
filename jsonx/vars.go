package jsonx

import (
	"errors"
	"strconv"
)

var (
	// ErrPoint 不是指针类型
	ErrPoint = errors.New("the argument to Result must be a non-nil pointer")
	// ErrNotMap 不是Map类型
	ErrNotMap = errors.New("cannot parse map, value is not a map")
	// ErrNotSlice 不是Slice类型
	ErrNotSlice = errors.New("cannot parse slice, value is not a slice")
	// ErrSyntax 指示值不具有目标类型的正确语法
	ErrSyntax = strconv.ErrSyntax
	// ErrType 指示值不具有目标类型的正确语法
	ErrType = errors.New("unsupported type")
)
