package redisUtil

import (
	"testing"
	"time"
)

func TestNewRedisConfig(t *testing.T) {
	redisConfigStr := "ConnectionString=10.1.0.21:6379;Password=redis_pwd;Database=3;MaxActive=50;MaxIdle=20;IdleTimeout=300;DialConnectTimeout=10;"
	redisConfig, err := NewRedisConfig(redisConfigStr)
	if err != nil {
		t.Errorf("there should be no err, but now has:%s", err)
	}

	connectionString := "10.1.0.21:6379"
	password := "redis_pwd"
	database := 3
	maxActive := 50
	maxIdle := 20
	idleTimeout := 300 * time.Second
	dialConnectTimeout := 10 * time.Second
	if redisConfig.ConnectionString != connectionString {
		t.Errorf("expected %s, but now got %s", connectionString, redisConfig.ConnectionString)
	}

	if redisConfig.Password != password {
		t.Errorf("expected %s, but now got %s", password, redisConfig.Password)
	}

	if redisConfig.Database != database {
		t.Errorf("expected %d, but now got %d", database, redisConfig.Database)
	}

	if redisConfig.MaxActive != maxActive {
		t.Errorf("expected %d, but now got %d", maxActive, redisConfig.MaxActive)
	}

	if redisConfig.MaxIdle != maxIdle {
		t.Errorf("expected %d, but now got %d", maxIdle, redisConfig.MaxIdle)
	}

	if redisConfig.IdleTimeout != idleTimeout {
		t.Errorf("expected %d, but now got %d", idleTimeout, redisConfig.IdleTimeout)
	}

	if redisConfig.DialConnectTimeout != dialConnectTimeout {
		t.Errorf("expected %d, but now got %d", dialConnectTimeout, redisConfig.DialConnectTimeout)
	}
}

func TestNewRedisConfig2(t *testing.T) {
	connectionString := "10.1.0.21:6379"
	password := "redis_pwd"
	database := 3
	maxActive := 50
	maxIdle := 20
	idleTimeout := 300 * time.Second
	dialConnectTimeout := 10 * time.Second

	redisConfig := NewRedisConfig2(connectionString, password, database, maxActive, maxIdle, idleTimeout, dialConnectTimeout)
	if redisConfig.ConnectionString != connectionString {
		t.Errorf("expected %s, but now got %s", connectionString, redisConfig.ConnectionString)
	}

	if redisConfig.Password != password {
		t.Errorf("expected %s, but now got %s", password, redisConfig.Password)
	}

	if redisConfig.Database != database {
		t.Errorf("expected %d, but now got %d", database, redisConfig.Database)
	}

	if redisConfig.MaxActive != maxActive {
		t.Errorf("expected %d, but now got %d", maxActive, redisConfig.MaxActive)
	}

	if redisConfig.MaxIdle != maxIdle {
		t.Errorf("expected %d, but now got %d", maxIdle, redisConfig.MaxIdle)
	}

	if redisConfig.IdleTimeout != idleTimeout {
		t.Errorf("expected %d, but now got %d", idleTimeout, redisConfig.IdleTimeout)
	}

	if redisConfig.DialConnectTimeout != dialConnectTimeout {
		t.Errorf("expected %d, but now got %d", dialConnectTimeout, redisConfig.DialConnectTimeout)
	}
}
