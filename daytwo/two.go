package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Credential struct {
	gorm.Model
	Email    string
	Password string
}

func main() {
	connDB()
	// emptyInterface()
	// interfaceAssert()
	// interfaceSwitch()

	// var s myString
	// fmt.Printf("%q\n", s)
}

func connDB() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Credential{})

	// Create
	db.Create(&Credential{Email: "admin@admin.com", Password: "1234"})

	// Read
	var cred Credential
	// db.First(&cred, 1) // find cred with integer primary key
	db.First(&cred, "email = ?", "admin@admin.com") // find cred with code D42

	fmt.Printf("%#v\n", cred)

	// Update - update cred's price to 200
	// db.Model(&cred).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&cred).Updates(Credential{Email: "ad@ad.com", Password: "F42"}) // non-zero fields
	// db.Model(&cred).Updates(map[string]interface{}{"Email": "ad@ad.com", "Password": "F42"})

	// // Delete - delete product
	// db.Delete(&cred, 1)
}

func emptyInterface() {
	var i interface{}

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = struct {
		number int
		text   string
	}{
		number: 10,
		text:   "ten",
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v\n", i, i)
}

/* interface type assert */
func interfaceAssert() {
	var i interface{}

	if i == nil {
		fmt.Println("nillll")
	}

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)

	if s, ok := i.(string); ok {
		fmt.Printf("%T %v\n", s, s)
	}

	i = struct {
		number int
		text   string
	}{
		number: 10,
		text:   "ten",
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v\n", i, i)
}

/* interface type assert switch */
func interfaceSwitch() {
	var i interface{}

	if i == nil {
		fmt.Println("nillll")
	}

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = "ten"
	fmt.Printf("type is %T, value is %v\n", i, i)

	switch s := i.(type) {
	case int:
		fmt.Printf("int %T %v\n", s, s)
	case string:
		fmt.Printf("string %T %v\n", s, s)
	default:
		fmt.Printf("no not type\n")
	}

	i = struct {
		number int
		text   string
	}{
		number: 10,
		text:   "ten",
	}
	fmt.Printf("type is %T, value is %v\n", i, i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v\n", i, i)
}

type myString string

func (s myString) String() string {
	return "this is my string implementation"
}
