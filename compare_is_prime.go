package main

import "math/big"

// take a relatively large prime i.e. 807*2^100095+1 and compare which method is faster


func _is_prime(n big.Int) {
	_, mod := new(big.Int).DivMod(n, big.NewInt(2), nil)
	if mod.Int64() == 0 {
		
	}

}


func main() {
	t := big.NewInt(2)
	e := new(big.Int).Exp(t, big.NewInt(100095), nil)
	m := new(big.Int).Mul(e, big.NewInt(807))
	a := new(big.Int).Add(m, big.NewInt(1))

}
