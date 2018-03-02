package main

import (
	"fmt"

	"github.com/siwally/golang/coins"
)

func main() {
	f := coins.PermsFn()

	for i := 1; i < 50; i++ {
		perms := f()

		fmt.Printf("p(%d) = %d\n", i, perms)
	}
}
