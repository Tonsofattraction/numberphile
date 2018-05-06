package main

import "fmt"

func collatz(n int16, i int16) {
	fmt.Printf("%d", n)
	i++
	if n == 1 {
		fmt.Println(i)
		return
	}
	if n % 2 == 0 {
		collatz(n/2, i)
	} else {
		collatz(3*n+1, i)
	}
}

func main () {
	var n int16;
	for n = 2; n < 1000; n++ {
		collatz(n, 1)
	}
}
