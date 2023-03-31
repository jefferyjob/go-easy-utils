package sliceUtil

import "reflect"

// UniqueSlice 移除slice中重复的值
func UniqueSlice(slice interface{}) interface{} {
	if reflect.ValueOf(slice).Kind() != reflect.Slice {
		return slice
	}

	s := reflect.ValueOf(slice)
	if s.Len() == 0 {
		return slice
	}
	t := reflect.TypeOf(slice).Elem()
	m := make(map[interface{}]bool)
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		if _, ok := m[v.Interface()]; ok {
			s = reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len()))
			i--
		} else {
			m[v.Interface()] = true
		}
	}
	result := reflect.MakeSlice(reflect.SliceOf(t), s.Len(), s.Len())
	reflect.Copy(result, s)
	return result.Interface()
}

//func UniqueStr(v []string) []string {
//	m := make(map[string]bool)
//	var result []string
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueInt(v []int) []int {
//	m := make(map[int]bool)
//	var result []int
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueInt8(v []int8) []int8 {
//	m := make(map[int8]bool)
//	var result []int8
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueInt16(v []int16) []int16 {
//	m := make(map[int16]bool)
//	var result []int16
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueInt32(v []int32) []int32 {
//	m := make(map[int32]bool)
//	var result []int32
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueInt64(v []int64) []int64 {
//	m := make(map[int64]bool)
//	var result []int64
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueFloat32(v []float32) []float32 {
//	m := make(map[float32]bool)
//	var result []float32
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueFloat64(v []float64) []float64 {
//	m := make(map[float64]bool)
//	var result []float64
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueUint(v []uint) []uint {
//	m := make(map[uint]bool)
//	var result []uint
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueUint8(v []uint8) []uint8 {
//	m := make(map[uint8]bool)
//	var result []uint8
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueUint16(v []uint16) []uint16 {
//	m := make(map[uint16]bool)
//	var result []uint16
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueUint32(v []uint32) []uint32 {
//	m := make(map[uint32]bool)
//	var result []uint32
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
//
//func UniqueUint64(v []uint64) []uint64 {
//	m := make(map[uint64]bool)
//	var result []uint64
//	for _, v := range v {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			result = append(result, v)
//		}
//	}
//	return result
//}
