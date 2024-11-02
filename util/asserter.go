package util

import "testing"

func AssertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf(`got "%d" want "%d"`, got, want)
	}
}
