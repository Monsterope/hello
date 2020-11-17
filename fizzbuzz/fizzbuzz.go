package fizzbuzz

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

// func Handler(c *gin.Context) {
// 	number := c.Param("number")
// 	n, err := strconv.Atoi(number)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	c.String(http.StatusOK, String(n))
// }
func Handler(c echo.Context) error {
	number := c.Param("number")
	n, err := strconv.Atoi(number)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, String(n))
}

type Intner interface {
	Intn(n int) int
}

func generateCode(r Intner) string {
	// src := rand.NewSource(time.Now().Unix())
	// r := rand.New(src)

	n1, n2, n3, n4 := r.Intn(100)+1, r.Intn(100)+1, r.Intn(100)+1, r.Intn(100)+1

	return fmt.Sprintf("%s %s %s %s", String(n1), String(n2), String(n3), String(n4))

}
