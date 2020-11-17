package fizzbuzz

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type fakeIntner struct {
	// n1, n2, n3, n4, count int
	count int
	n     []int
}

/* method on function */
type IntnFunc func(int) int

/* method on function */
func (fn IntnFunc) Intn(n int) int {
	return fn(n)
}

/* method on function */
func fakeIntn(n1, n2, n3, n4 int) IntnFunc {
	var count int
	n := []int{n1, n2, n3, n4}
	return func(int) int {
		defer func() { count++ }()
		return n[count]
	}
}

func (i *fakeIntner) Intn(int) int {
	defer func() { i.count++ }()
	return i.n[i.count]
}

func TestGenerateCode(t *testing.T) {
	// src := rand.NewSource(time.Now().Unix())
	// r := rand.New(src)
	// fake := fakeIntner(3)
	given := []int{2, 3, 4, 5}
	// fake := fakeIntner{0, given}
	want := "Fizz 4 Buzz Fizz"

	// get := generateCode(&fake)

	/* method on function */
	get := generateCode(fakeIntn(2, 3, 4, 5))

	if want != get {
		t.Errorf("it should be %q but got %q\n", want, given)
	}

	// fmt.Println(generateCode())
}

func TestHOF(t *testing.T) {

	// fmt.Println(fibonacciCall(10))
	// fmt.Println(fibonacciRef(10))

	// var fn = func(s string) string {
	// 	return ">>>" + s
	// }

	// hof(fn)

	fn := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}

}

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, a+b }()
		return a
	}
}

func fibonacciCall(n int) string {
	ss := "0,1"
	be := 0
	af := 1
	for i := 0; i < n; i++ {
		cc := be + af
		ss = ss + "," + strconv.Itoa(cc)
		be = af
		af = cc
	}

	return ss
}

func fibonacciRef(n int) string {
	a, b := 0, 1
	series := []string{strconv.Itoa(a), strconv.Itoa(b)}
	for i := 0; i < n; i++ {
		a, b = b, a+b
		series = append(series, strconv.Itoa(b))
	}

	return strings.Join(series, ",")
}

type echoStringFunc func(string) string

func hof(fn echoStringFunc) {
	s := fn("test")
	fmt.Println(s + "^^^")
}
