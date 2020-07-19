package main

import (
	"fmt"
	"math"

	"github.com/patil-ashwin/go_crash_course/strutil"
)

func main() {

	firstName, lastName := "Ashwin", "Patil"
	var age int32 = 16
	var isActive = false
	salary := 1000.0

	fmt.Println("Welcome to, Golang", firstName, lastName, age, isActive)
	fmt.Printf("%T\n", salary)

	fmt.Println(math.Ceil(2.8))
	fmt.Println(math.Floor(2.8))
	fmt.Println(math.Sqrt(64))
	fmt.Println(strutil.Reverse("niwhsA"))
}
