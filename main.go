package main

import (
	"hello/fizzbuzz"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func main() {
	/* call package file example */
	// prime.Prime()

	/* call gin framework exmaple */
	// ginGonic()

	/* call echo framework exmaple */
	// echoFra()
}

/* gin framework */
func ginGonic() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/fizzbuzz/:number", func(c *gin.Context) {
		number := c.Param("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusOK, fizzbuzz.String(n))
	})

	router.POST("/credentials", CredentialsHandler)

	router.Run(":8080")
}

/* call struct exmaple */
type credential struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

/* func jwt and response */
func CredentialsHandler(c *gin.Context) {
	var cred credential
	err := c.Bind(&cred)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Wrap(err, "bind credential").Error(),
		})
		return
	}
	mySigningKey := []byte("P@ssw0rd")

	type MyCustomClaims struct {
		Email    string `json: "email"`
		Password string `json: "password"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		cred.Email,
		cred.Password,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	// c.JSON(http.StatusOK, gin.H{
	// 	"token:": ss,
	// })
	c.JSON(http.StatusOK, map[string]string{
		"token:": ss,
	})
}

/* echo framework */
func echoFra() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.GET("/fizzbuzz/:number", func(c echo.Context) error {
		number := c.Param("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.String(http.StatusOK, fizzbuzz.String(n))
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
