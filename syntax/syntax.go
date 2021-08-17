package main

import (
	"fmt"
	"time"
)

func main() {
	var i1 int
	fmt.Println("i1", i1) // Go assigns the zero value for un-intialized values
	/* Number types in Go
	- int8, int16, int32, int64, int
	- uint8, uint16, uint32, uint64, uint
	- float32, float64
	- complex64, complex128
	- big.Int, big.Rat (e.g. ½), big.Float
	*/

	var i2 = 2
	fmt.Println("i2", i2)

	i3 := 4 // type inference
	fmt.Printf("i3: %d (%T)\n", i3, i3)

	var (
		i4 int64 = 7
		i5       = int64(8)
	)
	fmt.Println("i4", i4, "i5", i5)
	// fmt.Println(i4 == i3) // compilation error
	fmt.Println("eq?", int(i4) == i3)

	i6 := i3 / 3 // integer division
	fmt.Println("i6", i6)

	f1 := 3.0
	f2 := float64(i3) / f1
	fmt.Println("f2", f2)

	// constants type is defined at usage location
	fmt.Println("div", i3/3.0)

	// timeout := 10 // compilation error
	const timeout = 10
	time.Sleep(timeout * time.Millisecond)

	a := 0
	{ // opens a new scope
		a := 1 // a shadows the a on line 47
		fmt.Println("inner a", a)
	}
	fmt.Println("outer a", a)

	// Conditions
	n := 10
	if n > 5 {
		fmt.Println("n is big")
	}

	if n > 10 {
		fmt.Println("n is big")
	} else {
		fmt.Println("n is small")
	}

	if n > 0 && n < 100 {
		fmt.Println("n in range")
	}

	if m := n * 10; m < 10_000 {
		fmt.Println("m", m)
	}
	// fmt.Println(m) // m's scope is the if (and else) only

	switch {
	case n < 10:
		fmt.Println("n is small")
	case n < 100:
		fmt.Println("n is medium")
	default:
		fmt.Println("n is huge")
	}

	switch n % 2 {
	case 0:
		fmt.Println("n is even")
	case 1:
		fmt.Println("n is odd")
	}

	// Numbers Formats
	fmt.Println(1_000_000)
	fmt.Println(0x10)
	fmt.Println(0o10)
	fmt.Println(0b10)
	fmt.Println(1e9)

	fmt.Println(collatz(6)) // 3
	fmt.Println(collatz(7)) // 22

	fmt.Println(2 ^ 10) // ^ is binary XOR, there's no power operator

	// Looping - there is only a "for" loop
	for i := 0; i < 3; i++ {
		fmt.Println("for", i)
	}

	// while
	i := 3
	for i > 0 {
		fmt.Println("while", i)
		i--
	}

	// while true
	i = 0
	for {
		if i > 2 {
			break
		}
		fmt.Println("while true", i)
		i++
	}

	// there's also for ... range, we'll talk about it later

	fmt.Println(collatzLen(12)) // 10

	// Find the first two single digit number that their multiplication is above 12
loop: // loop is label (and yes, there's a goto statement)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if m := i * j; m > 12 {
				fmt.Println(i, j)
				// break // breaks the inner loop
				break loop // breaks the outer loop
			}
		}
	}
}

// exercise: write collatzLen(n int) int
// how many collatz steps does it take to get from n to 1?
func collatzLen(n int) int {
	size := 1
	for n != 1 {
		n = collatz(n)
		size++
	}
	return size
}

/* collatz n
- n/2 if n is even
- n*3 + 1 if n is odd
*/
func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}
