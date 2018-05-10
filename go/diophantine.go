package _go
import (
	"fmt"
	"math"
)

// one must be negative. no need to check two zeros.

func main() {
	var found bool
	var a float64;

	for !found  {
		if int(a) % 1000000 == 0 {
			fmt.Printf("a: %f", a)
		}
		for b := a; b <= 2300000000; b++ {
			if found {
				break
			}
			for c := b; c <= 2300000000; c++ {
				if int(c) % 1000000 == 0 {
					fmt.Printf("c: %f\n", c)
				}
				if math.Pow(a,3) + math.Pow(b,3) + math.Pow(c,3) == 30 {
					found = true
					fmt.Println(a,b,c)
					break
				}
			}
		}
	}
}

