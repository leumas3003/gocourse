package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "https://derivco.com/"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(benchSite(ctx, url))

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	bid := bidOn(ctx, "http://goo.gl/2")
	fmt.Println(bid)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	bid = bidOn(ctx, "http://goo.gl/2")
	fmt.Println(bid)

	ch := make(chan int, 1)
	ch <- 3 // not blocking
	// ch <- 3 // blocking

	fmt.Println("DONE")

}

func bidOn(ctx context.Context, url string) Bid {
	// Use buffered channel to avoid goroutine leak
	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()
	/* Solution for goroutine leak without buffered channel
	go func() {
		select {
		case ch <- bestBid(url):
		case <-ctx.Done():
		}
	}()
	*/

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

var defaultBid = Bid{
	AdURL: "http://adsRus.com/default",
	Price: 0.02,
}

type Bid struct {
	AdURL string
	Price float64 // in USD
}

// Written by algo team, time to completion varies
func bestBid(url string) Bid {
	// Simulate work
	time.Sleep(time.Duration(len(url)) * 3 * time.Millisecond)

	return Bid{
		AdURL: "http://adsRus.com/ad42",
		Price: 0.07,
	}
}

func benchSite(ctx context.Context, url string) (time.Duration, error) {
	start := time.Now()
	// resp, err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
