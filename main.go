package main

import (
	"fmt"
	"GO/name"
)

// variable
// Int int8 int32 int64
// uint uint8
// Float float32 float64
// byte
// rune
// complex

var hobby string
// var height, weight int = 169, 69
var height int = 169
var weight = "69.0"

func main() {
	hobby = "play game"
	thisYear := 2020 
	fname, lname := name.Name()

	// operation + - * / %
	fmt.Println("hello",fname ,lname)
	fmt.Println("I am ", name.Age(1997, thisYear))
	fmt.Println("my hobby is " + hobby)
	fmt.Println("my height ", height , "weight ", weight)

	// IF ELSE
	if h, _ := body(); h > 170  {
		fmt.Println("I was born in Udonthani")
	} else if h == 170 {
		fmt.Println("I am tall")
	} else {
		fmt.Println("I am short")
	}
}

func body() (int,int) {
	return 170, 69
}

