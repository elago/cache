package cache

import (
	// "github.com/gogather/com/log"
	"testing"
)

func TestSet(t *testing.T) {
	Set("region", "key", "hello world")

	var str string
	Get("region", "key", &str)

	if str != "hello world" {
		t.Error("get content is not put content")
	}
}
