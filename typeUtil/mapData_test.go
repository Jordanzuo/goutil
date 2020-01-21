package typeUtil

import (
	"testing"
	"time"
)

func TestMapDataByte(t *testing.T) {
	TestMapDataUint8(t)
}

func TestMapDataInt(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected int = 0

	// Test when key doesn't exist
	got, err := mapData.Int(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Int(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Int(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataInt8(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected int8 = 0

	// Test when key doesn't exist
	got, err := mapData.Int8(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Int8(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Int8(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataInt16(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected int16 = 0

	// Test when key doesn't exist
	got, err := mapData.Int16(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Int16(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Int16(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataInt32(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected int32 = 0

	// Test when key doesn't exist
	got, err := mapData.Int32(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Int32(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Int32(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataInt64(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected int64 = 0

	// Test when key doesn't exist
	got, err := mapData.Int64(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Int64(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Int64(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataUint(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected uint = 0

	// Test when key doesn't exist
	got, err := mapData.Uint(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Uint(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Uint(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataUint8(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected uint8 = 0

	// Test when key doesn't exist
	got, err := mapData.Uint8(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Uint8(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Uint8(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataUint16(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected uint16 = 0

	// Test when key doesn't exist
	got, err := mapData.Uint16(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Uint16(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Uint16(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataUint32(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected uint32 = 0

	// Test when key doesn't exist
	got, err := mapData.Uint32(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Uint32(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Uint32(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataUint64(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected uint64 = 0

	// Test when key doesn't exist
	got, err := mapData.Uint64(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Uint64(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Uint64(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %d, but got %d", expected, got)
		return
	}
}

func TestMapDataFloat32(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected float32 = 0

	// Test when key doesn't exist
	got, err := mapData.Float32(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Float32(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Float32(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}
}

func TestMapDataFloat64(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected float64 = 0

	// Test when key doesn't exist
	got, err := mapData.Float64(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Float64(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = 1
	expected = 1
	got, err = mapData.Float64(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %f, but got %f", expected, got)
		return
	}
}

func TestMapDataBool(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected bool = true

	// Test when key doesn't exist
	got, err := mapData.Bool(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = "abc"
	got, err = mapData.Bool(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = true
	expected = true
	got, err = mapData.Bool(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %t, but got %t", expected, got)
		return
	}
}

func TestMapDataString(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected string = ""

	// Test when key doesn't exist
	got, err := mapData.String(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist, but type doesn't match
	mapData[key] = 123
	expected = "123"
	got, err = mapData.String(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}

	// Test when key exist and value matches
	mapData[key] = "hello"
	expected = "hello"
	got, err = mapData.String(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}
}

// Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
func TestMapDataDateTime(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected time.Time = time.Date(2019, time.December, 25, 12, 0, 0, 0, time.UTC)

	// Test when key doesn't exist
	got, err := mapData.DateTime(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist but value doesn't match
	mapData[key] = "123"
	expected = time.Date(2019, time.December, 25, 12, 0, 0, 0, time.UTC)
	got, err = mapData.DateTime(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = time.Date(2019, time.December, 25, 12, 0, 0, 0, time.UTC)
	expected = time.Date(2019, time.December, 25, 12, 0, 0, 0, time.UTC)
	got, err = mapData.DateTime(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
		return
	}
}

func TestMapDataInterface(t *testing.T) {
	data := make(map[string]interface{})
	mapData := NewMapData(data)
	key := "key"
	var expected *Person

	// Test when key doesn't exist
	got, err := mapData.Interface(key)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	// Test when key exist and value matches
	mapData[key] = NewPerson("Jordan", 34)
	expected = NewPerson("Jordan", 34)
	got, err = mapData.Interface(key)
	if err != nil {
		t.Errorf("There should be no error, but now there is:%s", err)
		return
	}
	if gotPersonObj, ok := got.(*Person); !ok {
		t.Errorf("Expected type *Person")
	} else if gotPersonObj.SameAs(expected) == false {
		t.Errorf("Expected %v, but got %v", expected, got)
		return
	}
}

type Person struct {
	Name string
	Age  int
}

func (this *Person) SameAs(other *Person) bool {
	return this.Name == other.Name && this.Age == other.Age
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
