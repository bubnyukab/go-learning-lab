package week1

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 11)
	expected := "aaaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 11)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 11)
	fmt.Println(result)
	// Output: aaaaaaaaaaa
}
