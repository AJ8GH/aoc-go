package util

import (
	"fmt"
	"testing"
)

func Assert[T comparable](t testing.TB, got, want T) {
	t.Helper()
	verb := verb(got)
	if got != want {
		t.Errorf(fmt.Sprintf("got %v want %v", verb, verb), got, want)
	}
}

func verb(got interface{}) string {
	switch got.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return "%d"
	case float32, float64:
		return "%f"
	default:
		return "%q"
	}
}
