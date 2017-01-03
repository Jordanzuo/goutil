package mysqlUtil

import (
	"testing"
)

func TestConvertConnectionStringFromCSharpToGo(t *testing.T) {
	csharp := "DataSource=10.162.2.205;port=3306;UserId=admin;Password=MOQIkaka$#@!1234;Database=s201_dzz_log;Allow Zero Datetime=true;charset=utf8;pooling=true;MinimumPoolSize=20;maximumpoolsize=200;command timeout=60;AllowUserVariables=True;"
	expected := "admin:MOQIkaka$#@!1234@tcp(10.162.2.205:3306)/s201_dzz_log?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=200||MaxIdleConns=20"

	if goConn := ConvertConnectionStringFromCSharpToGo(csharp); goConn != expected {
		t.Errorf("Expected %s, but got %s", expected, goConn)
	}

	csharp = "DataSource=10.162.2.205;port=3306;UserId=admin;Password=MOQIkaka$#@!1234;Database=s201_dzz_log;Allow Zero Datetime=true;charset=utf8;pooling=false;MinimumPoolSize=20;maximumpoolsize=200;command timeout=60;AllowUserVariables=True;"
	expected = "admin:MOQIkaka$#@!1234@tcp(10.162.2.205:3306)/s201_dzz_log?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=0||MaxIdleConns=0"

	if goConn := ConvertConnectionStringFromCSharpToGo(csharp); goConn != expected {
		t.Errorf("Expected %s, but got %s", expected, goConn)
	}
}
