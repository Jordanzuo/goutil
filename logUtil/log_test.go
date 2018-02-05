package logUtil

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
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

	Log("test Info", Info, true)
	Log("test Debug", Debug, true)
	Log("test Error", Error, true)
	Log("test Fatal", Fatal, true)
}

func TestLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	Log("a test log message", Info, true)
	Log("another test log message", Debug, false)
}

func TestInfoLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	InfoLog("info记录")
	InfoLog("info记录2:%v", time.Now())
}

func TestErrorLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	ErrorLog("Error记录")
	ErrorLog("Error记录2:%v", time.Now())
}

func TestDebugLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	DebugLog("info记录")
	DebugLog("info记录2:%v", time.Now())
}

func TestWarnLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	WarnLog("Warn记录")
	WarnLog("Warn记录2:%v", time.Now())
}

func TestFatalLog(t *testing.T) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	logPath := filepath.Dir(path)

	SetLogPath(logPath)

	FatalLog("Fatal记录")
	FatalLog("Fatal记录2:%v", time.Now())
}
