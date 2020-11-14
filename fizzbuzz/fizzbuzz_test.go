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
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenTwo(t *testing.T) {
	given := 2
	want := "2"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenThere(t *testing.T) {
	given := 3
	want := "Fizz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}

func TestFizzBuzzGivenFour(t *testing.T) {
	given := 4
	want := "4"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}

func TestFizzBuzzGivenFive(t *testing.T) {
	given := 5
	want := "Buzz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
