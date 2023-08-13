package lru

import (
	"reflect"
	"testing"
)

type Integer int32

func (i Integer) Len() int {
	return 4
}

func TestCache_Get(t *testing.T) {
	cache := New(0, nil)
	cache.Add("zls", Integer(21))
	zlsAge, ok := cache.Get("zls")
	if !ok || !reflect.DeepEqual(zlsAge.(Integer), Integer(21)) {
		t.Fail()
	}
}
