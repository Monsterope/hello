package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/* custom type int */
type Int int

/* custom type string */
type String string

func main() {

	/* array and slice example */
	// arrayExam()

	// result := "abcdef"
	// strs := []rune(result)

	// var item []string

	// c := 1
	// cl := len(result)

	// for _, v := range strs {
	// 	cc := c + 3
	// 	var ir []string
	// 	for i := c; i < cc; i++ {
	// 		ch := string(v)
	// 		ir = append(ir, ch)
	// 	}
	// 	item = append(item, ir)

	// }
	// fmt.Println(item)

	// fmt.Println(byte strs)
	// c := len(result)
	// cc := len

	// for pos, char := range result {

	// }

	/* slice */
	// fmt.Println("%#v", trippleChar("abcdef"))

	/* read file and use map */
	// readCSVfile()
	// csvFile2()

	/* pointer example */
	// n := 5
	// p := &n
	// setPointer(p)
	// fmt.Println(n)

	/* type : custom type */
	// var i Int = 12
	// fmt.Printf("%q\n", i.toString())
	// var s String = "10"
	// fmt.Printf("%d\n", s.toInt())

	/* set get pointer */
	// s.set("yyy")
	// fmt.Println(s.get())

	/* defer */
	// defer myPrint(6)
	// defer myPrint(8)
	// myPrint(4)

	/* defer try catch */
	// catchMe()

	// goroutineEx()

	// gochannelEx()

	// chro()

	// pp()

	// ma()

	// log.Fatal("test")
	// log.Panic("test")

}

func ma() {
	slice := []int{1, 2, 3, 4, 5}
	variadic(slice...)
}

func variadic(v ...int) {
	for _, val := range v {
		fmt.Println(val)
	}
}

var str string
var mux sync.Mutex

func pp() {
	for i := 0; i < 100; i++ {
		go writeString("test")
		go readString()
	}
}

func writeString(s string) {
	mux.Lock()
	str = s
	mux.Unlock()
}
func readString() string {
	mux.Lock()
	defer mux.Unlock()
	return str
}

func chro() {
	ch := make(chan int)
	chStop := make(chan struct{})
	go fibocEx(ch, chStop)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	chStop <- struct{}{}
	fmt.Println("finish")

}

func fibocEx(ch chan int, chStop <-chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-chStop:
			fmt.Println("agreacful")
			break
		}
	}

}

func gochannelEx() {
	total := 10
	ch := make(chan int, total)
	for i := total; i > 0; i-- {
		ch <- i
	}
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}

}

var wg = sync.WaitGroup{}

func goroutineEx() {
	total := 10
	now := time.Now()

	wg.Add(10)
	for i := 0; i < total; i++ {
		go printout(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	fmt.Println(i)
	wg.Done()
	panic("aa")
}

/* array and slice */
func arrayExam() {
	var array [5]int
	for i, v := range array {
		fmt.Println(i, v)
	}
	for _, v := range array {
		fmt.Println(v)
	}
	for i, v := range array {
		_ = i
		fmt.Println(v)
	}

	fmt.Println("slice")
	var slice []int
	primes := [...]int{2, 3, 5, 7, 11, 13}
	slice = primes[1:4]

	for i := range slice {
		fmt.Println(slice[i])
	}

	var s []int
	if s == nil {
		fmt.Println("it's nil")
	}

	// s := make([]int, 5)
	// s := []int{1, 2, 3, 4, 5}
	// s = append(s, 6, 7, 8)
}

/* slice example */
func trippleChar(s string) []string {
	count := 3 - (len(s) % 3)
	s += strings.Repeat("*", count)

	round := len(s) / 3

	result := []string{}

	for i := 0; i < round; i++ {
		p := s[:3]
		s = s[3:]
		result = append(result, p)
	}

	return result

}

/* use map example and read file csv */
func readCSVfile() {
	b, err := ioutil.ReadFile("./oscar_age_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(b))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nameAndCount := map[string]int{}
	for _, record := range records {
		nameAndCount[record[3]]++
	}

	// fmt.Println(nameAndCount)

	for name, count := range nameAndCount {
		if count > 1 {
			fmt.Println(name, count)
		}
	}

}

/* use map example and read file csv */
func csvFile2() {
	csvfile, err := os.Open("oscar_age_male.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s\n", record[3])
	}
}

/* pointer example */
func setPointer(n *int) {
	fmt.Println(*n)
	*n = *n + *n
}

/* custom type example */
func (i Int) toString() string {
	return strconv.Itoa((int(i)))
}

/* custom type example */
func toStrings(i Int) string {
	return strconv.Itoa((int(i)))
}

/* custom type example */
func (s String) toInt() (n int) {
	n, _ = strconv.Atoi(string(s))
	return
}

// set pointer
func (s *String) set(ns String) {
	*s = ns
}

// get pointer
func (s String) get() String {
	return s
}

/* struct example */
func structExample() {
	jsonString := `{
		"email": "admin@admin.com",
		"password": "123456"
	}`

	type credential struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}

	var cred credential
	err := json.Unmarshal([]byte(jsonString), &cred)

	if err != nil {
		log.Fatal(err)
	}

}

/* defer example */
func myPrint(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()
	n = n + n
	fmt.Println(n)
}

/* defer example try catch */
func catchMe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	s := []int{}
	fmt.Println(s[1])
}
