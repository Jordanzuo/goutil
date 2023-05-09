package runtimeUtil

import (
	"fmt"
	"reflect"
	"runtime"
)

// 获取当前执行的方法的名称
func GetCurrFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	if len(pc) == 0 {
		return ""
	}

	return runtime.FuncForPC(pc[0]).Name()
}

// 获取任意方法的名称
func GetFuncName(prefix string, i interface{}) string {
	var name string
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Func {
		name = runtime.FuncForPC(value.Pointer()).Name()
	}

	return fmt.Sprintf("%s/%s", prefix, name)
}
