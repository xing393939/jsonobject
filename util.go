package jsonobject

import (
	"runtime"
	"testing"
)

func assertEqual(t testing.TB, got interface{}, want interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok && got != want {
		t.Errorf("\r\n\t%s:%d: unexpected value obtained; got %#v; want %#v", file, line, got, want)
	}
}