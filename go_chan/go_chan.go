package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go fmt.Println("in a goroutine")
	fmt.Println("in main")

	a := 2.2
	{
		a := "hello"
		fmt.Println(a)
	}
	fmt.Println(a)

	for i := 0; i < 4; i++ {
		// i := i
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	ch := make(chan int)
	go func() {
		ch <- 24 // send
	}()
	val := <-ch // receive
	fmt.Println("val", val)

	time.Sleep(10 * time.Millisecond)

	nums := []int{15, 8, 4, 16, 42, 23}
	fmt.Println(nums, "->", sleepSort(nums))

	ch2 := make(chan int)
	close(ch2)
	// receive from a closed channel will return the zero value - non blocking
	val2 := <-ch2
	fmt.Println("val2", val2)
	if val3, ok := <-ch2; ok {
		fmt.Println("val3", val3)
	} else {
		fmt.Println("ch2 closed")
	}
	// ch2 <- 3 // panic
	// close(ch2) // panic

	ch3 := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	for n := range ch3 {
		fmt.Println("ch3", n)
	}
	/* The above is like
	for {
		n, ok := <- ch3
		if !ok {
			break
		}
		// user code
	}
	*/

	queue := make(chan string)
	for i := 0; i < 3; i++ {
		go producer(i, queue)
		go consumer(i, queue)
	}

	time.Sleep(time.Second)
}

func consumer(id int, ch <-chan string) {
	for {
		randSleep() // simulate work
		msg := <-ch
		fmt.Printf("%s -> %d\n", msg, id)
	}
}

func producer(id int, ch chan<- string) {
	n := 0
	for {
		randSleep() // simulate work
		n++
		msg := fmt.Sprintf("%d -> {%d}", id, n)
		ch <- msg
	}
}

func randSleep() {
	n := rand.Intn(50)
	time.Sleep(time.Duration(n) * time.Millisecond)

}

/* For every n in nums, spin a gorouting that will
- sleep n milliseconds
- send n over a channel

The function will collect the number from the channel
In bash:
	for n in $@; do
		(sleep $n && echo $n)&
	done
	wait
*/
func sleepSort(nums []int) []int {
	// TODO: Write the code
	ch := make(chan int)
	for _, n := range nums {
		go func(i int) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			ch <- i
		}(n)
	}

	sorted := make([]int, len(nums))
	for i := range nums {
		n := <-ch
		sorted[i] = n
	}
	return sorted
}
