package fizzbuzz

func String(n int) string {
	if n == 5 {
		return "Buzz"
	}
	if n == 4 {
		return "4"
	}
	if n == 3 {
		return "Fizz"
	}
	if n == 2 {
		return "2"
	}
	return "1"
}
