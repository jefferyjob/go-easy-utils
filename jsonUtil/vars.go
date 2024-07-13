package jsonUtil

import (
	"errors"
	"strconv"
)

var (
	// ErrSyntax indicates that a value does not have the right syntax for the target type
	// 指示值不具有目标类型的正确语法
	ErrSyntax = strconv.ErrSyntax

	// ErrType indicates that a value does not have the right syntax for the target type
	// 指示值不具有目标类型的正确语法
	ErrType = errors.New("unsupported type")
)
