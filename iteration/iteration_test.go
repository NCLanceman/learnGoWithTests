package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	assertErrorMessage(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("i", 4)
	fmt.Println(result)
	// Output: iiii
}

func assertErrorMessage(t testing.TB, want, get string) {
	t.Helper()
	if want != get {
		t.Errorf("expected %q but got %q", want, get)
	}
}
