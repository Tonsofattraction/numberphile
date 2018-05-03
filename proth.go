package main
// https://www.youtube.com/watch?v=fcVjitaM3LY
// did up to 86000 yet

import (
	"fmt"
	"math/big"
)

func process1000(e int64, c chan int) {
	var n = big.NewInt(21181)
	var found bool
	var exp = big.NewInt(e)


	fmt.Println(c)
	fmt.Println(e)
	start := e

	for e <= (start+1000) {
		exp = big.NewInt(e)
		// отмечаемся каждую тысячу итераций.
		var _, mod = new(big.Int).DivMod(exp, big.NewInt(100), big.NewInt(0))
		if mod.Int64() == 0 {
			fmt.Println(exp)
		}

		var i = new(big.Int).Exp(big.NewInt(2), exp, nil)
		var m = new(big.Int).Mul(n, i)
		var p = new(big.Int).Add(big.NewInt(1), m)

		found = p.ProbablyPrime(1)
		if found {
			fmt.Println("FOUND")
			fmt.Println(exp)
			break
		} else {
			e++
		}
	}
	if found {
		c <- 1
	} else {
		c <- 0
	}
}

func main() {
	c := make(chan int)
	var e int64 = 1

	for i:=0;i<2;i++ {
		e += 1000
		go process1000(e, c)
	}
	for i := 0; i < 2; i++ {
		<- c // wait for one task to complete

	}
}