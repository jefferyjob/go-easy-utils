package go_easy_utils

import (
	"errors"
	"strconv"
)

// ErrSyntax indicates that a value does not have the right syntax for the target type
// 指示值不具有目标类型的正确语法
var ErrSyntax = strconv.ErrSyntax

// ErrType indicates that a value does not have the right syntax for the target type
// 指示值不具有目标类型的正确语法
var ErrType = errors.New("unsupported type")

// ErrValOut Indicates that the range corresponding to the type is exceeded
// 表示超出了类型对应的范围
var ErrValOut = errors.New("value out of range")

// ErrUnsignedInt The identity type is negative, so it cannot be converted to an unsigned integer
// 标识类型为负数，所以无法转为无符号的整型
var ErrUnsignedInt = errors.New("cannot convert negative value to unsigned integer")
