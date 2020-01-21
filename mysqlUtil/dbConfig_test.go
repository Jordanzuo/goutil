package mysqlUtil

import (
	"testing"
)

func TestNewDBConfig(t *testing.T) {
	connectionString := "root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s"
	maxOpenConns := 10
	maxIdleConns := 5

	dbConfigObj := NewDBConfig(connectionString, maxOpenConns, maxIdleConns)
	if connectionString != dbConfigObj.ConnectionString {
		t.Errorf("Expected %s, but now %s", connectionString, dbConfigObj.ConnectionString)
	}

	if maxOpenConns != dbConfigObj.MaxOpenConns {
		t.Errorf("Expected %d, but now %d", maxOpenConns, dbConfigObj.MaxOpenConns)
	}

	if maxIdleConns != dbConfigObj.MaxIdleConns {
		t.Errorf("Expected %d, but now %d", maxIdleConns, dbConfigObj.MaxIdleConns)
	}
}

func TestNewDBConfig2(t *testing.T) {
	dbConfigStr := "root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=10||MaxIdleConns2=5"

	if _, err := NewDBConfig2(dbConfigStr); err == nil {
		t.Errorf("there should be err, but now not.")
	}

	dbConfigStr = "root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=10||MaxIdleConns=5"
	dbConfigObj, err := NewDBConfig2(dbConfigStr)
	if err != nil {
		t.Errorf("there should be no err, but now has.")
	}

	connectionString := "root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s"
	maxOpenConns := 10
	maxIdleConns := 5

	if connectionString != dbConfigObj.ConnectionString {
		t.Errorf("Expected %s, but now %s", connectionString, dbConfigObj.ConnectionString)
	}

	if maxOpenConns != dbConfigObj.MaxOpenConns {
		t.Errorf("Expected %d, but now %d", maxOpenConns, dbConfigObj.MaxOpenConns)
	}

	if maxIdleConns != dbConfigObj.MaxIdleConns {
		t.Errorf("Expected %d, but now %d", maxIdleConns, dbConfigObj.MaxIdleConns)
	}
}
