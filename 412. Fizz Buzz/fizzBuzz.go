package main

import (
	"strconv"
)

func fizzBuzz(n int) []string { //The good way
	mySlice := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			mySlice = append(mySlice, "Fizz Buzz")
		} else if i%3 == 0 {
			mySlice = append(mySlice, "Fizz")
		} else if i%5 == 0 {
			mySlice = append(mySlice, "Buzz")
		} else {
			mySlice = append(mySlice, strconv.Itoa(i))
		}
	}
	return mySlice
}
