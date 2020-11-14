package fizzbuzz_test

import (
	"hello/fizzbuzz"
	"testing"
)

func TestFizzBuzzGivenOne(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got % %q\n", given, want, get)
	}

}
