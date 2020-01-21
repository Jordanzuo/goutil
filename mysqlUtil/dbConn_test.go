package mysqlUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestOpenMysqlConnection(t *testing.T) {
	connectionString := "root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=10||MaxIdleConns=5"

	if _, err := OpenMysqlConnection(connectionString); err != nil {
		t.Errorf("there should be no err, but now has：%s", err)
	}
}

func TestOpenMysqlConnection2(t *testing.T) {
	connectionString := "root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s"
	maxOpenConns := 10
	maxIdleConns := 5

	if _, err := OpenMysqlConnection2(connectionString, maxOpenConns, maxIdleConns); err != nil {
		t.Errorf("there should be no err, but now has：%s", err)
	}
}

func TestOpenMysqlConnection3(t *testing.T) {
	dbConfigObj := NewDBConfig("root:moqikaka3306@tcp(10.1.0.10:3306)/sdkcenter?charset=utf8&parseTime=true&loc=Local&timeout=10s", 5, 2)
	if _, err := OpenMysqlConnection3(dbConfigObj); err != nil {
		t.Errorf("there should be no err, but now has：%s", err)
	}
}

func TestTestConnection(t *testing.T) {
	dbConfigObj := NewDBConfig("root:moqikaka3306@tcp(10.1.0.10:3306)/sdkcenter?charset=utf8&parseTime=true&loc=Local&timeout=10s", 5, 2)
	dbObj, err := OpenMysqlConnection3(dbConfigObj)
	if err != nil {
		t.Errorf("there should be no err, but now has：%s", err)
	}

	succeedCount := 0
	expectedCount := 5
	for i := 0; i < expectedCount; i++ {
		if err := TestConnection(dbObj); err != nil {
			fmt.Printf("%s:%s\n", time.Now(), err)
		} else {
			succeedCount += 1
			fmt.Printf("%s:%s\n", time.Now(), "ok")
		}
		time.Sleep(time.Second * 3)
	}

	if succeedCount != expectedCount {
		t.Errorf("ExecptedCount:%d, but got %d", expectedCount, succeedCount)
	}
}
