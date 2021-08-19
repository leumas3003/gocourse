package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

/* Exercise:
https://api.stocktwits.com/api/2/streams/symbol/aapl.json

Possible additions:
- Command line processing with "flag"
- Limit size of incoming data with io.LimitReader
*/

func getRelatedStocks(symbol string) (map[string]int, error) {
	return nil, fmt.Errorf("not implemented")
}

func main() {
	file, err := os.Open("aapl.json")
	if err != nil {
		log.Fatal(err)
	}

	counts, err := relatedStocks(file)
	if err != nil {
		log.Fatal(err)
	}
	for sym, count := range counts {
		fmt.Printf("%-10s -> %02d\n", sym, count)
	}
}

// realtedStocks return a map of stocks mentioned with count for each
func relatedStocks(r io.Reader) (map[string]int, error) {
	var reply struct {
		Messages []struct {
			Symbols []struct {
				Name string `json:"symbol"`
			}
		}
	}

	if err := json.NewDecoder(r).Decode(&reply); err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	for _, m := range reply.Messages {
		for _, symbol := range m.Symbols {
			counts[symbol.Name]++
		}
	}

	return counts, nil
}
