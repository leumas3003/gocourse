package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	v := []int{1, 2, 3}
	doubleAt(v, 1)
	fmt.Println(v)

	v = appendVal(v, 4)
	fmt.Println(v)

	n := 7
	double(n)
	fmt.Println("double", n)

	doublePtr(&n)
	fmt.Println("doublePtr", n)

	if d, err := div(7, 2); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(d)
	}

	if d, err := div(7, 0); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(d)
	}

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum())

	// On windows use -n
	args := []string{"ping", "-c", "3", "google.com"}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	/*
		if err := cmd.Run(); err != nil {
			fmt.Println("ERROR:", err)
		}
	*/

	inc := func(n int) int {
		return n + 1
	}
	fmt.Println("inc", inc(9))

	fn := dispatch["mul"]
	fmt.Println("mul", fn(3, 7))

	add7 := makeAdder(7)
	fmt.Println("add7", add7(20))

	if size, err := fileSize("functions.go"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("size:", size)
	}

	deferOrder()
}

func deferOrder() {
	for i := 0; i < 3; i++ {
		defer func(n int) {
			fmt.Println("closing", n)
		}(i)

		/* Option 2 for fixing the bug
		i := i // new i will shadow i from line 70
		defer func() {
			fmt.Println("closing", i)
		}()
		*/

		/* BUG
		defer func() {
			fmt.Println("closing", i)
		}()
		*/
	}
	fmt.Println("for loop done")
}

func fileSize(fileName string) (int, error) {
	// Common idiom: Acquire resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(info.Size()), nil
}

func makeAdder(n int) func(int) int {
	return func(val int) int {
		return val + n // n comes from function closure
	}
}

var dispatch = map[string]func(int, int) int{
	"add": func(a, b int) int { return a + b },
	"mul": func(a, b int) int { return a * b },
}

func sum(nums ...int) int {
	// Here nums is a []int
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func doublePtr(n *int) {
	*n *= 2
}

func double(n int) {
	n *= 2
}

func appendVal(v []int, i int) []int {
	v = append(v, i)
	return v
}

func doubleAt(v []int, i int) {
	v[i] *= 2
}
