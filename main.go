package main

import (
	"hello/fizzbuzz"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func main() {

	viper.SetDefault("jwt.signature", "P@ssw0rd")
	viper.SetDefault("app.addr", ":1323")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	/* call package file example */
	// prime.Prime()

	/* call gin framework exmaple */
	// ginGonic()

	/* call echo framework exmaple */
	echoFra()
}

/* gin framework */
func ginGonic() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// router.GET("/fizzbuzz/:number", fizzbuzz.Handler)

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
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes

	// e.GET("/fizzbuzz/:number", func(c echo.Context) error {
	// 	number := c.Param("number")
	// 	n, err := strconv.Atoi(number)
	// 	if err != nil {
	// 		return c.String(http.StatusBadRequest, err.Error())
	// 	}

	// 	return c.String(http.StatusOK, fizzbuzz.String(n))
	// })

	protect := e.Group("/bo", middleware)

	protect.GET("/fizzbuzz/:number", fizzbuzz.Handler, middleware)

	// Start server
	// e.Logger.Fatal(e.Start(":1323"))
	e.Logger.Fatal(e.Start(viper.GetString("app.addr")))
}

func Handler(c echo.Context) error {
	number := c.Param("number")
	n, err := strconv.Atoi(number)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fizzbuzz.String(n))
}

func validateToken(token *jwt.Token) (interface{}, error) {
	return []byte(viper.GetString("jwt.signature")), nil

}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		if len(auth) < 7 {
			log.Println("token empty")
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "auth"})
		}

		tokenString := auth[7:]

		_, err := jwt.Parse(tokenString, validateToken)
		// _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	return hmacSampleSecret, nil
		// })

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unkhow",
			})
		}

		return next(c)
	}
}
