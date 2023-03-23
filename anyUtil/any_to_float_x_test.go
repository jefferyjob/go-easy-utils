package anyUtil

import (
	"testing"
)

func TestAnyToFloat32(t *testing.T) {
	// 测试类型转换为float32
	f, err := AnyToFloat32(float32(3.14))
	if err != nil {
		t.Errorf("AnyToFloat32 failed: %v", err)
	}
	if f != float32(3.14) {
		t.Errorf("AnyToFloat32 failed: expected %v, but got %v", float32(3.14), f)
	}

	// 测试类型转换为float64
	f, err = AnyToFloat32(float64(3.14))
	if err != nil {
		t.Errorf("AnyToFloat32 failed: %v", err)
	}
	if f != float32(3.14) {
		t.Errorf("AnyToFloat32 failed: expected %v, but got %v", float32(3.14), f)
	}

	// 测试类型转换为int
	f, err = AnyToFloat32(int(3))
	if err != nil {
		t.Errorf("AnyToFloat32 failed: %v", err)
	}
	if f != float32(3) {
		t.Errorf("AnyToFloat32 failed: expected %v, but got %v", float32(3), f)
	}

	// 测试类型转换为string
	f, err = AnyToFloat32("3.14")
	if err != nil {
		t.Errorf("AnyToFloat32 failed: %v", err)
	}
	if f != float32(3.14) {
		t.Errorf("AnyToFloat32 failed: expected %v, but got %v", float32(3.14), f)
	}

	// 测试错误的值
	_, err = AnyToFloat32("invalid")
	if err == nil {
		t.Errorf("AnyToFloat32 failed: expected an error, but got none")
	}
}

func TestAnyToFloat64(t *testing.T) {
	// 测试类型转换为float64
	f, err := AnyToFloat64(float64(3.14))
	if err != nil {
		t.Errorf("AnyToFloat64 failed: %v", err)
	}
	if f != float64(3.14) {
		t.Errorf("AnyToFloat64 failed: expected %v, but got %v", float64(3.14), f)
	}

	//// 测试类型转换为float32
	//f, err = AnyToFloat64(float32(3.1))
	//if err != nil {
	//	t.Errorf("AnyToFloat64 failed: %v", err)
	//}
	//if f != float64(3.1) {
	//	t.Errorf("AnyToFloat64 failed: expected %v, but got %v", float64(3.14), f)
	//}

	// 测试类型转换为int
	f, err = AnyToFloat64(int(3))
	if err != nil {
		t.Errorf("AnyToFloat64 failed: %v", err)
	}
	if f != float64(3) {
		t.Errorf("AnyToFloat64 failed:expected %v, but got %v", float64(3), f)
	}

	// 测试类型转换为string
	f, err = AnyToFloat64("3.14")
	if err != nil {
		t.Errorf("AnyToFloat64 failed: %v", err)
	}
	if f != float64(3.14) {
		t.Errorf("AnyToFloat64 failed: expected %v, but got %v", float64(3.14), f)
	}

	// 测试错误的值
	_, err = AnyToFloat64("invalid")
	if err == nil {
		t.Errorf("AnyToFloat64 failed: expected an error, but got none")
	}
}
