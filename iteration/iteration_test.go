package iteration

import (
	"errors"
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {

		repeated, _ := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("prpeat b 5 times", func(t *testing.T) {
		repeated, _ := Repeat("b", 5)
		expected := "bbbbb"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeat 9 times", func(t *testing.T) {
		repeated, _ := Repeat("a", 9)
		expected := "aaaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("refuses negative numbers", func(t *testing.T) {
		repeated, err := Repeat("a", -1)
		expectedVal := ""
		expectedErr := errors.New("cannot repeat negative times")
		assertCorrectMessage(t, repeated, expectedVal)
		assertCorrectError(t, err, expectedErr)
	})
	t.Run("repeat 0 times returns empty string", func(t *testing.T) {
		repeated, _ := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t *testing.T, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func assertCorrectError(t *testing.T, repeated error, expected error) {
	t.Helper()
	if repeated.Error() != expected.Error() {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated, _ := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Repeat("a", 5)
		if err != nil {
			b.Errorf("error: %v", err)
		}
	}
}
