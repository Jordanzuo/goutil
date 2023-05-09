package runtimeUtil

import "testing"

func TestGetFuncName(t *testing.T) {
	funcName := GetFuncName("GetFuncName", add)
	expectName := "GetFuncName/github.com/Jordanzuo/goutil/runtimeUtil.add"
	if funcName != expectName {
		t.Errorf("Expect to get %s, but get %s instead.", expectName, funcName)
		return
	}

	funcName = GetFuncName("GetFuncName", echo)
	expectName = "GetFuncName/github.com/Jordanzuo/goutil/runtimeUtil.echo"
	if funcName != expectName {
		t.Errorf("Expect to get %s, but get %s instead.", expectName, funcName)
		return
	}
}

func add(a, b int) int {
	return a + b
}

func echo(in string) string {
	return in
}

func BenchmarkGetCurrFuncName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCurrFuncName()
	}
}
