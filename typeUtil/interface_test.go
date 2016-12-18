package typeUtil

import (
	"testing"
)

// 转换为int测试
func TestToInt(t *testing.T) {
	_, isSuccess := Int(1)
	if isSuccess == false {
		t.Error("int=>int error")
		return
	}

	_, isSuccess = Int(1.1)
	if isSuccess == false {
		t.Error("float=>int error")
		return
	}

	_, isSuccess = Int(1.1)
	if isSuccess == false {
		t.Error("float=>int error")
		return
	}

	_, isSuccess = Int("1")
	if isSuccess == false {
		t.Error("int string=>int error")
		return
	}

	_, isSuccess = Int("1.1")
	if isSuccess == false {
		t.Error("float string=>int error")
		return
	}
}

// 转换为Float测试
func TestToFloat(t *testing.T) {
	_, isSuccess := Float64(1)
	if isSuccess == false {
		t.Error("int=>float error")
		return
	}

	_, isSuccess = Float64(1.1)
	if isSuccess == false {
		t.Error("float=>float error")
		return
	}

	_, isSuccess = Float64(1.1)
	if isSuccess == false {
		t.Error("float=>float error")
		return
	}

	_, isSuccess = Float64("1")
	if isSuccess == false {
		t.Error("int string=>float error")
		return
	}

	_, isSuccess = Float64("1.1")
	if isSuccess == false {
		t.Error("float string=>float error")
		return
	}
}

// 转换为String测试
func TestToString(t *testing.T) {
	_, isSuccess := String(1)
	if isSuccess == false {
		t.Error("int=>String error")
		return
	}

	_, isSuccess = String(1.1)
	if isSuccess == false {
		t.Error("float=>String error")
		return
	}

	_, isSuccess = String(1.1)
	if isSuccess == false {
		t.Error("float=>String error")
		return
	}

	_, isSuccess = String("1")
	if isSuccess == false {
		t.Error("int string=>String error")
		return
	}

	_, isSuccess = String("1.1")
	if isSuccess == false {
		t.Error("float string=>String error")
		return
	}
}
