package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	p := Payment{
		From:   "Wile. E. Coyote",
		To:     "ACME",
		Amount: 107.3,
	}

	now := time.Now()
	p.Pay(now)
	p.Pay(now)

	var wg sync.WaitGroup
	//var m sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100_000; j++ {
				/* option 1: use a lock
				m.Lock()
				counter++
				m.Unlock()
				*/
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

var counter int64

func (p *Payment) Pay(t time.Time) {
	fn := func() {
		p.pay(t)
	}
	p.once.Do(fn)
}

func (p *Payment) pay(t time.Time) {
	ts := t.Format("2006-01002T15:04:05")
	fmt.Printf("%s %s -> [$%.2f] -> %s\n", ts, p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64

	once sync.Once
}
