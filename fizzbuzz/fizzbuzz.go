package fizzbuzz

import (
	"strconv"
)

func String(n int) string {

	var returnResult string = ""

	if n%3 == 0 {
		returnResult = "Fizz"
	}
	if n%5 == 0 {
		returnResult = returnResult + "Buzz"
	}
	if n%3 != 0 && n%5 != 0 {
		returnResult = strconv.Itoa(n)
	}

	return returnResult
}
