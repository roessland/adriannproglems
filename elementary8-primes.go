package main

import (
	"fmt"
	"github.com/cznic/mathutil"
	"math/big"
)

func main() {
	one := big.NewInt(1)
	remainder := big.NewInt(0)

	for n := big.NewInt(2); ; n = n.Add(n, one) {
		sqrtn := mathutil.SqrtBig(n)
		divisible := false
		for divisor := big.NewInt(2); divisor.Cmp(sqrtn) != 1; divisor = divisor.Add(divisor, one) {
			remainder = remainder.Mod(n, divisor)
			if remainder.Cmp(big.NewInt(0)) == 0 {
				divisible = true
				break
			}
		}
		if divisible {
			//fmt.Printf("%v is not prime\n", n.String())
		} else {
			fmt.Printf("%v is prime\n", n.String())
		}
	}
}
