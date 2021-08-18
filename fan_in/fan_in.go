package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1, ch2 := producer("p1", 5), producer("p2", 3)
	for msg := range fanIn(ch1, ch2) {
		fmt.Println(msg)
	}
	fmt.Println("DONE")
}

func fanIn(chans ...chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	go func() {
		for _, ch := range chans {
			// wg.Add(1)
			ch := ch
			go func() {
				defer wg.Done()
				for msg := range ch {
					out <- msg
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	/*
		go func() {
			wg.Wait()
			close(out)
		}()
	*/

	return out
}

func producer(name string, size int) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < size; i++ {
			msg := fmt.Sprintf("%s: %d", name, i)
			ch <- msg
		}
		close(ch)
	}()
	return ch
}
