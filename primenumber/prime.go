package main

import "fmt"


func main() {

	var testprime int
	
   	fmt.Println("Enter a value")
	fmt.Scan(&testprime)

	//if testprime == 2 {
	//	fmt.Println(testprime, "is a prime number")
	//	return
	//} else {
		for i:=2; i<testprime; i++ {
			if testprime%i == 0 {
				fmt.Println(testprime, "is not a prime number")
				return
			} else {
				fmt.Println(testprime, "is a prime number")
				return
			}
		}
	//}



}