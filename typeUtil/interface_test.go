package typeUtil

import (
	"reflect"
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	var val interface{}
	expected := int(100)

	// Test with value == nil
	got, err := Int(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Int(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Int(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Int(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Int(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Int(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Int(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Int(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Int(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Int(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Int(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Int(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Int(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Int(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Int(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Int(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestIntArray(t *testing.T) {
	got, err := IntArray(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []int{1, 2}

	got, err = IntArray(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestInt8(t *testing.T) {
	var val interface{}
	expected := int8(100)

	// Test with value == nil
	got, err := Int8(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Int8(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Int8(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Int8(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Int8(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Int8(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Int8(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Int8(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Int8(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Int8(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Int8(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Int8(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Int8(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Int8(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Int8(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Int8(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestInt8Array(t *testing.T) {
	got, err := Int8Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []int8{1, 2}

	got, err = Int8Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestInt16(t *testing.T) {
	var val interface{}
	expected := int16(100)

	// Test with value == nil
	got, err := Int16(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Int16(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Int16(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Int16(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Int16(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Int16(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Int16(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Int16(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Int16(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Int16(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Int16(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Int16(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Int16(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Int16(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Int16(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Int16(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestInt16Array(t *testing.T) {
	got, err := Int16Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []int16{1, 2}

	got, err = Int16Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestInt32(t *testing.T) {
	var val interface{}
	expected := int32(100)

	// Test with value == nil
	got, err := Int32(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Int32(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Int32(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Int32(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Int32(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Int32(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Int32(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Int32(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Int32(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Int32(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Int32(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Int32(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Int32(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Int32(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Int32(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Int32(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestInt32Array(t *testing.T) {
	got, err := Int32Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []int32{1, 2}

	got, err = Int32Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestInt64(t *testing.T) {
	var val interface{}
	expected := int64(100)

	// Test with value == nil
	got, err := Int64(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Int64(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Int64(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Int64(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Int64(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Int64(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Int64(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Int64(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Int64(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Int64(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Int64(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Int64(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Int64(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Int64(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Int64(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Int64(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestInt64Array(t *testing.T) {
	got, err := Int64Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []int64{1, 2}

	got, err = Int64Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestUint(t *testing.T) {
	var val interface{}
	expected := uint(100)

	// Test with value == nil
	got, err := Uint(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Uint(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Uint(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Uint(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Uint(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Uint(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Uint(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Uint(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Uint(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Uint(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Uint(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Uint(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Uint(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Uint(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Uint(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Uint(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestUintArray(t *testing.T) {
	got, err := UintArray(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []uint{1, 2}

	got, err = UintArray(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestUint8(t *testing.T) {
	var val interface{}
	expected := uint8(100)

	// Test with value == nil
	got, err := Uint8(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Uint8(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Uint8(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Uint8(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Uint8(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Uint8(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Uint8(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Uint8(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Uint8(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Uint8(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Uint8(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Uint8(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Uint8(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Uint8(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Uint8(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Uint8(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestUint8Array(t *testing.T) {
	got, err := Uint8Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []uint8{1, 2}

	got, err = Uint8Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestByte(t *testing.T) {
	var val interface{}
	expected := byte(100)

	// Test with value == nil
	got, err := Byte(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Byte(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Byte(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Byte(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Byte(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Byte(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Byte(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Byte(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Byte(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Byte(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Byte(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Byte(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Byte(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Byte(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Byte(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Byte(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestByteArray(t *testing.T) {
	got, err := ByteArray(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []byte{1, 2}

	got, err = ByteArray(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestUint16(t *testing.T) {
	var val interface{}
	expected := uint16(100)

	// Test with value == nil
	got, err := Uint16(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Uint16(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Uint16(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Uint16(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Uint16(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Uint16(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Uint16(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Uint16(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Uint16(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Uint16(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Uint16(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Uint16(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Uint16(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Uint16(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Uint16(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Uint16(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestUint16Array(t *testing.T) {
	got, err := Uint16Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []uint16{1, 2}

	got, err = Uint16Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestUint32(t *testing.T) {
	var val interface{}
	expected := uint32(100)

	// Test with value == nil
	got, err := Uint32(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Uint32(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Uint32(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Uint32(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Uint32(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Uint32(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Uint32(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Uint32(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Uint32(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Uint32(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Uint32(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Uint32(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Uint32(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Uint32(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Uint32(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Uint32(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestUint32Array(t *testing.T) {
	got, err := Uint32Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []uint32{1, 2}

	got, err = Uint32Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestUint64(t *testing.T) {
	var val interface{}
	expected := uint64(100)

	// Test with value == nil
	got, err := Uint64(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Uint64(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Uint64(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Uint64(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Uint64(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Uint64(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Uint64(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Uint64(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Uint64(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Uint64(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Uint64(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Uint64(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Uint64(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Uint64(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Uint64(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Uint64(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestUint64Array(t *testing.T) {
	got, err := Uint64Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []uint64{1, 2}

	got, err = Uint64Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestFloat32(t *testing.T) {
	var val interface{}
	expected := float32(100.0)

	// Test with value == nil
	got, err := Float32(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Float32(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Float32(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Float32(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Float32(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Float32(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Float32(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Float32(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Float32(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Float32(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Float32(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Float32(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Float32(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Float32(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Float32(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Float32(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}
}

func TestFloat32Array(t *testing.T) {
	got, err := Float32Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []float32{1.0, 2.0}

	got, err = Float32Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestFloat64(t *testing.T) {
	var val interface{}
	expected := float64(100.0)

	// Test with value == nil
	got, err := Float64(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Float64(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Float64(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Float64(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Float64(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Float64(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Float64(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Float64(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Float64(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Float64(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Float64(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Float64(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Float64(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Float64(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Float64(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Float64(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}
}

func TestFloat64Array(t *testing.T) {
	got, err := Float64Array(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []float64{1.0, 2.0}

	got, err = Float64Array(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestBool(t *testing.T) {
	var val interface{}
	expected := true

	// Test with value == nil
	got, err := Bool(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = Bool(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = Bool(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = Bool(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = Bool(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = Bool(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = Bool(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = Bool(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = Bool(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = Bool(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = Bool(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = Bool(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100.0)
	got, err = Bool(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100.0)
	got, err = Bool(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as string type
	val_string1 := "abc"
	got, err = Bool(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 := "100"
	got, err = Bool(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with false result
	expected = false

	// Test with value as int type
	val_int = int(0)
	got, err = Bool(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 = int8(0)
	got, err = Bool(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 = int16(0)
	got, err = Bool(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 = int32(0)
	got, err = Bool(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 = int64(0)
	got, err = Bool(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint type
	val_uint = uint(0)
	got, err = Bool(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 = uint8(0)
	got, err = Bool(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 = uint16(0)
	got, err = Bool(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 = uint32(0)
	got, err = Bool(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 = uint64(0)
	got, err = Bool(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 = float32(0.0)
	got, err = Bool(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 = float64(0.0)
	got, err = Bool(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}

	// Test with value as string type
	val_string1 = "abc"
	got, err = Bool(val_string1)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't")
		return
	}

	val_string2 = "0"
	got, err = Bool(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}
}

func TestBoolArray(t *testing.T) {
	got, err := BoolArray(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 0)
	expected := []bool{true, false}

	got, err = BoolArray(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestString(t *testing.T) {
	var val interface{}
	expected := "100"

	// Test with value == nil
	got, err := String(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	got, err = String(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as int type
	val_int := int(100)
	got, err = String(val_int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as int8 type
	val_int8 := int8(100)
	got, err = String(val_int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as int16 type
	val_int16 := int16(100)
	got, err = String(val_int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as int32 type
	val_int32 := int32(100)
	got, err = String(val_int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as int64 type
	val_int64 := int64(100)
	got, err = String(val_int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as uint type
	val_uint := uint(100)
	got, err = String(val_uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as uint8 type
	val_uint8 := uint8(100)
	got, err = String(val_uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as uint16 type
	val_uint16 := uint16(100)
	got, err = String(val_uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as uint32 type
	val_uint32 := uint32(100)
	got, err = String(val_uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as uint64 type
	val_uint64 := uint64(100)
	got, err = String(val_uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as float32 type
	val_float32 := float32(100)
	got, err = String(val_float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as float64 type
	val_float64 := float64(100)
	got, err = String(val_float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test with value as string type
	val_string2 := "100"
	got, err = String(val_string2)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}
}

func TestStringArray(t *testing.T) {
	got, err := StringArray(nil)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	valArray := make([]interface{}, 0, 2)
	valArray = append(valArray, 1)
	valArray = append(valArray, 2)
	expected := []string{"1", "2"}

	got, err = StringArray(valArray)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}

	if len(expected) != len(got) || expected[0] != got[0] || expected[1] != got[1] {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestDateTime(t *testing.T) {
	var val interface{}
	expected := time.Date(2017, time.February, 14, 5, 20, 0, 0, time.Local)

	// Test with value == nil
	_, err := DateTime(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test with value as non basic type
	val = make([]int, 0, 32)
	_, err = DateTime(val)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	timeVal := "2017-02-14 05:20:00"
	got, err := DateTime(timeVal)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}

	got, err = DateTime(got)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}

	got, err = DateTime(got.Unix())
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

func TestDateTimeArray(t *testing.T) {

}

func TestDateTimeByFormat(t *testing.T) {
	// var val interface{}
	// expected := time.Date(2017, time.February, 14, 5, 20, 0, 0, time.Local)

	// // Test with value == nil
	// _, err := DateTimeByFormat(val, "2006-01-02 15:04:05")
	// if err == nil {
	// 	t.Errorf("There should be an error, but now there isn't.")
	// 	return
	// }

	// // Test with value as non basic type
	// val = make([]int, 0, 32)
	// _, err = DateTimeByFormat(val, "2006-01-02 15:04:05")
	// if err == nil {
	// 	t.Errorf("There should be an error, but now there isn't.")
	// 	return
	// }

	// timeVal := "2017-02-14 05:20:00"
	// got, err := DateTimeByFormat(timeVal, "2006-01-02 15:04:05")
	// if err != nil {
	// 	t.Errorf("There should be no error, but now there is one:%s", err)
	// 	return
	// }
	// if got != expected {
	// 	t.Errorf("Expected %v, but got %v", expected, got)
	// 	return
	// }

	// got, err = DateTimeByFormat(got, "2006-01-02 15:04:05")
	// if err != nil {
	// 	t.Errorf("There should be no error, but now there is one:%s", err)
	// 	return
	// }
	// if got != expected {
	// 	t.Errorf("Expected %v, but got %v", expected, got)
	// 	return
	// }

	// got, err = DateTimeByFormat(got.Unix(), "2006-01-02 15:04:05")
	// if err != nil {
	// 	t.Errorf("There should be no error, but now there is one:%s", err)
	// 	return
	// }
	// if got != expected {
	// 	t.Errorf("Expected %v, but got %v", expected, got)
	// 	return
	// }
}

func TestDateTimeArrayByFormat(t *testing.T) {

}

func TestConvert(t *testing.T) {
	result, err := Convert([]int{}, reflect.Map)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test Convert with type int
	expected_int := 1
	result, err = Convert(1, reflect.Int)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(int); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_int {
		t.Errorf("Expected %d, but got %d", expected_int, got)
		return
	}

	// Test Convert with type int8
	expected_int8 := int8(1)
	result, err = Convert(1, reflect.Int8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(int8); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_int8 {
		t.Errorf("Expected %d, but got %d", expected_int8, got)
		return
	}

	// Test Convert with type int16
	expected_int16 := int16(1)
	result, err = Convert(1, reflect.Int16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(int16); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_int16 {
		t.Errorf("Expected %d, but got %d", expected_int16, got)
		return
	}

	// Test Convert with type int32
	expected_int32 := int32(1)
	result, err = Convert(1, reflect.Int32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(int32); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_int32 {
		t.Errorf("Expected %d, but got %d", expected_int32, got)
		return
	}

	// Test Convert with type int64
	expected_int64 := int64(1)
	result, err = Convert(1, reflect.Int64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(int64); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_int64 {
		t.Errorf("Expected %d, but got %d", expected_int64, got)
		return
	}

	// Test Convert with type uint
	expected_uint := uint(1)
	result, err = Convert(1, reflect.Uint)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(uint); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_uint {
		t.Errorf("Expected %d, but got %d", expected_uint, got)
		return
	}

	// Test Convert with type int8
	expected_uint8 := uint8(1)
	result, err = Convert(1, reflect.Uint8)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(uint8); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_uint8 {
		t.Errorf("Expected %d, but got %d", expected_uint8, got)
		return
	}

	// Test Convert with type int16
	expected_uint16 := uint16(1)
	result, err = Convert(1, reflect.Uint16)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(uint16); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_uint16 {
		t.Errorf("Expected %d, but got %d", expected_uint16, got)
		return
	}

	// Test Convert with type uint32
	expected_uint32 := uint32(1)
	result, err = Convert(1, reflect.Uint32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(uint32); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_uint32 {
		t.Errorf("Expected %d, but got %d", expected_uint32, got)
		return
	}

	// Test Convert with type uint64
	expected_uint64 := uint64(1)
	result, err = Convert(1, reflect.Uint64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(uint64); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_uint64 {
		t.Errorf("Expected %d, but got %d", expected_uint64, got)
		return
	}

	// Test Convert with type float32
	expected_float32 := float32(1.0)
	result, err = Convert(1, reflect.Float32)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(float32); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_float32 {
		t.Errorf("Expected %f, but got %f", expected_float32, got)
		return
	}

	// Test Convert with type float64
	expected_float64 := float64(1.0)
	result, err = Convert(1, reflect.Float64)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(float64); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_float64 {
		t.Errorf("Expected %f, but got %f", expected_float64, got)
		return
	}

	// Test Convert with type bool
	expected_bool := true
	result, err = Convert(true, reflect.Bool)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(bool); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_bool {
		t.Errorf("Expected %t, but got %t", expected_bool, got)
		return
	}

	// Test Convert with type string
	expected_string := "hello"
	result, err = Convert("hello", reflect.String)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%s", err)
		return
	}
	if got, ok := result.(string); !ok {
		t.Errorf("The type should be int, but now it's not.")
		return
	} else if got != expected_string {
		t.Errorf("Expected %s, but got %s", expected_string, got)
		return
	}
}
