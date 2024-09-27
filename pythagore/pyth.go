package main

import "fmt"


func main() {


	var a float64
	var b float64
	var c float64

   	fmt.Println("Enter a value")
	fmt.Scan(&a)
   	fmt.Println("Enter b value")
	fmt.Scan(&b)
   	fmt.Println("Enter c value")
	fmt.Scan(&c)
   	fmt.Println("a^2 + b^2 = ", a*a + b*b)
   	fmt.Println("c^2 = ", c*c)
   	if a*a + b*b == c*c {
   		fmt.Println("a, b, and c form a Pythagorean triple")
   	} else {
   		fmt.Println("a, b, and c do not form a Pythagorean triple")
   	}



}