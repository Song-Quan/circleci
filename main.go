package main

import (
	"circleci/internal/mathutil"
	"fmt"
)

func main() {
	fmt.Println("Hello, Modular Go!")

	sum := mathutil.Add(3, 5)
	fmt.Printf("3 + 5 = %d\n", sum)
}
