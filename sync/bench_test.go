package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	count     int64
	m         sync.Mutex
	countChan chan int64
)

func BenchmarkMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		count++
		m.Unlock()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&count, 1)
	}
}

func BenchmarkChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countChan <- 1
	}
}

func init() {
	countChan = make(chan int64)
	go func() {
		for {
			val := <-countChan
			count += val
		}
	}()
}
