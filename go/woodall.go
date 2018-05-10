package _go

import (
	"math/big"
	"fmt"
)

// https://www.youtube.com/watch?v=PQDvEJFdY1U&t=0s&list=PL0D0BD149128BB06F&index=5
// https://en.wikipedia.org/wiki/Woodall_number

// n*2^n -1

func main() {
	n := big.NewInt(17016602)
	var i = new(big.Int).Exp(big.NewInt(2), n, nil)
	var m = new(big.Int).Mul(n, i)
	var p = new(big.Int).Sub(big.NewInt(1), m)

	fmt.Println(p.ProbablyPrime(10))
}