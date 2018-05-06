package main

import "fmt"

func collatz(n int16, i int16) bool {
	i++
	if n == 1 {
		fmt.Println(i)
		return true
	}
	if n % 2 == 0 {
		return collatz(n/2, i)
	} else {
		return collatz(3*n+1, i)
	}
}

func main () {
	var n int16;
	for n = 2; n < 1000; n++ {
		fmt.Println(n)
		collatz(n, 0)
	}
}
