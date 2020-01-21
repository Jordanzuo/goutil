package mysqlUtil

import (
	"testing"
)

func TestConvertConnectionStringFromCSharpToGo(t *testing.T) {
	csharp := "DataSource=10.66.195.134;port=3306;UserId=admin;Password=MOQIkaka$#@!1234;Database=s9501_sd_log;Allow Zero Datetime=true;charset=utf8;pooling=false;command timeout=60;AllowUserVariables=True;"
	expected := "admin:MOQIkaka$#@!1234@tcp(10.66.195.134:3306)/s9501_sd_log?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=0||MaxIdleConns=0"

	if goConn := ConvertConnectionStringFromCSharpToGo(csharp); goConn != expected {
		t.Errorf("Expected %s, but got %s", expected, goConn)
	}

	// csharp = "DataSource=10.162.2.205;port=3306;UserId=admin;Password=MOQIkaka$#@!1234;Database=s201_dzz_log;Allow Zero Datetime=true;charset=utf8;pooling=false;min pool size=20;max pool size=200;command timeout=60;AllowUserVariables=True;"
	// expected = "admin:MOQIkaka$#@!1234@tcp(10.162.2.205:3306)/s201_dzz_log?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=0||MaxIdleConns=0"

	// if goConn := ConvertConnectionStringFromCSharpToGo(csharp); goConn != expected {
	// 	t.Errorf("Expected %s, but got %s", expected, goConn)
	// }
}

func TestIsCSharpStyle(t *testing.T) {
	connString := "DataSource=10.66.195.134;port=3306;UserId=admin;Password=MOQIkaka$#@!1234;Database=s9501_sd_log;Allow Zero Datetime=true;charset=utf8;pooling=false;command timeout=60;AllowUserVariables=True;"
	if !IsCSharpStyle(connString) {
		t.Errorf("it's should be C# style, but now not")
	}

	connString = "admin:MOQIkaka$#@!1234@tcp(10.66.195.134:3306)/s9501_sd_log?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=0||MaxIdleConns=0"
	if IsCSharpStyle(connString) {
		t.Errorf("it's should not be C# style, but now it is")
	}
}

func TestIsGoStyle(t *testing.T) {
	connString := "admin:MOQIkaka$#@!1234@tcp(10.66.195.134:3306)/s9501_sd_log?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=0||MaxIdleConns=0"
	if !IsGoStyle(connString) {
		t.Errorf("it's should be Go style, but now not")
	}

	connString = "DataSource=10.66.195.134;port=3306;UserId=admin;Password=MOQIkaka$#@!1234;Database=s9501_sd_log;Allow Zero Datetime=true;charset=utf8;pooling=false;command timeout=60;AllowUserVariables=True;"
	if IsGoStyle(connString) {
		t.Errorf("it's should not be Go style, but now it is")
	}
}
