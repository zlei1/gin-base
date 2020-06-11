package redis

import (
	"testing"
	"time"
)

func TestSetupRedis(t *testing.T) {
	TestSetup()

	err := Client.Ping().Err()
	if err != nil {
		t.Error("ping redis server err: ", err)
		return
	}
	t.Log("ping redis server pass")
}

func TestRedisUse(t *testing.T) {
	TestSetup()

	var setGetKey = "test-key"
	var setGetValue = "test-value"
	Client.Set(setGetKey, setGetValue, time.Second*100)

	expectValue := Client.Get(setGetKey).Val()
	if setGetValue != expectValue {
		t.Log("original value: ", setGetValue)
		t.Log("expect value: ", expectValue)
		return
	}

	t.Log("redis set get test pass")
}
