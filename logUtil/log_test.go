package logUtil

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestSetLogPath(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)
	retPath := GetLogPath()
	if retPath != logPath {
		t.Errorf("设置路径不正确，Expected:%s, Got:%s", logPath, retPath)
	}
}

func TestLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	Log("a test log message", Info, true)
	Log("another test log message", Debug, false)
}
