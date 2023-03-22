package conditions

import "testing"

func TestIsOdd(t *testing.T) {
	given := 10
	want := false

	result := isOdd(given)
	if result != want {
		t.Errorf("isOdd(%d) = %t, want %t", given, result, want)
	}
}
