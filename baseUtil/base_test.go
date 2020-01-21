package baseUtil

import (
	"testing"
)

func TestNew(t *testing.T) {
	elements := ""
	_, err := New(elements)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	elements = "00"
	_, err = New(elements)
	if err == nil {
		t.Errorf("There should be an error, but now there isn't.")
		return
	}

	elements = "01"
	_, err = New(elements)
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestNewBase2(t *testing.T) {
	_, err := NewBase2()
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestNewBase8(t *testing.T) {
	_, err := NewBase8()
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestNewBase16(t *testing.T) {
	_, err := NewBase16()
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestNewBase26(t *testing.T) {
	_, err := NewBase26()
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestNewBase36(t *testing.T) {
	_, err := NewBase36()
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestNewBase62(t *testing.T) {
	_, err := NewBase62()
	if err != nil {
		t.Errorf("There should be no error, but now there is one:%v.", err)
		return
	}
}

func TestTransform(t *testing.T) {
	base2, _ := NewBase2()
	base8, _ := NewBase8()
	base16, _ := NewBase16()
	base26, _ := NewBase26()
	base36, _ := NewBase36()
	base62, _ := NewBase62()

	var source uint64 = 0
	expected := "0"
	got := base2.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "0"
	got = base8.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "0"
	got = base16.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "a"
	got = base26.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "0"
	got = base36.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "0"
	got = base62.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	source = 1
	expected = "1"
	got = base2.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "1"
	got = base8.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "1"
	got = base16.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "b"
	got = base26.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "1"
	got = base36.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "1"
	got = base62.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	source = 2
	expected = "10"
	got = base2.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "2"
	got = base8.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "2"
	got = base16.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "c"
	got = base26.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "2"
	got = base36.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "2"
	got = base62.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	source = 100
	expected = "1100100"
	got = base2.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "144"
	got = base8.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "64"
	got = base16.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "dw"
	got = base26.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "2s"
	got = base36.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}

	expected = "1C"
	got = base62.Transform(source)
	if got != expected {
		t.Errorf("Expected to get %s, but got %s", expected, got)
		return
	}
}

func TestParse(t *testing.T) {
	base2, _ := NewBase2()
	base8, _ := NewBase8()
	base16, _ := NewBase16()
	base26, _ := NewBase26()
	base36, _ := NewBase36()
	base62, _ := NewBase62()

	expected := uint64(0)
	got := base2.Parse("0")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base8.Parse("0")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base16.Parse("0")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base26.Parse("a")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base36.Parse("0")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base62.Parse("0")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}

	expected = uint64(1)
	got = base2.Parse("1")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base8.Parse("1")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base16.Parse("1")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base26.Parse("b")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base36.Parse("1")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base62.Parse("1")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}

	expected = uint64(2)
	got = base2.Parse("10")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base8.Parse("2")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base16.Parse("2")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base26.Parse("c")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base36.Parse("2")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base62.Parse("2")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}

	expected = uint64(100)
	got = base2.Parse("1100100")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base8.Parse("144")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base16.Parse("64")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base26.Parse("dw")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base36.Parse("2s")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
	got = base62.Parse("1C")
	if got != expected {
		t.Errorf("Expected to get %d, but got %d", expected, got)
		return
	}
}
