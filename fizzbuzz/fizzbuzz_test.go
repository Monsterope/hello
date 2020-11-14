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
func TestFizzBuzzGivenSix(t *testing.T) {
	given := 6
	want := "Fizz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenSeven(t *testing.T) {
	given := 7
	want := "7"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenEight(t *testing.T) {
	given := 8
	want := "8"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenNine(t *testing.T) {
	given := 9
	want := "Fizz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenTen(t *testing.T) {
	given := 10
	want := "Buzz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenTwly(t *testing.T) {
	given := 12
	want := "Fizz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenFiveteen(t *testing.T) {
	given := 15
	want := "FizzBuzz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenEightteen(t *testing.T) {
	given := 18
	want := "Fizz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenTwyy(t *testing.T) {
	given := 20
	want := "Buzz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
func TestFizzBuzzGivenThreely(t *testing.T) {
	given := 30
	want := "FizzBuzz"

	get := fizzbuzz.String(given)

	if want != get {
		t.Errorf("given %d want %q but got %q\n", given, want, get)
	}

}
