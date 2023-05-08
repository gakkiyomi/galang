// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// IntToString .
func (*GalangTransform) IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString .
func (*GalangTransform) Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// StringToInt64 .
func (*GalangTransform) StringToInt64(i string) (int64, error) {
	return strconv.ParseInt(i, 10, 64)
}

// StringToInt .
func (*GalangTransform) StringToInt(i string) (int, error) {
	return strconv.Atoi(i)
}

// Float64ToInt64 .
func (*GalangTransform) Float64ToInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// Float64ToString .
func (*GalangTransform) Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float32ToString .
func (*GalangTransform) Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

// Int64ToFloat64 .
func (*GalangTransform) Int64ToFloat64(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// BoolToString bool2string
func (*GalangTransform) BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// BytesToString byte2string
func (*GalangTransform) BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// StringToBytes string2bytes
func (*GalangTransform) StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

// AnyToString return a string of any type
func (gl *GalangTransform) AnyToString(any interface{}) string {
	switch val := any.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(any)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return gl.BoolToString(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return gl.Int64ToString(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return gl.Float64ToString(v.Float())
	case reflect.Float32:
		return gl.Float32ToString(float32(v.Float()))
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", any)
}
