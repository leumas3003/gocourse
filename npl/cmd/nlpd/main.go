package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	nlp "github.com/leumas3003/npl"
)

var (
	logger  *zap.Logger
	version string
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version & exit")
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer logger.Sync()

	http.HandleFunc("/tokenize", tokenizeHandler)
	http.HandleFunc("/health", healthHandler)

	addr := ":8080"
	// log.Printf("server ready on %s", addr)
	logger.Info(
		"server ready",
		zap.String("address", addr),
	)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// log.Printf("error: tokenize: %s", err)
		logger.Error(
			"tokenize",
			zap.Error(err),
		)
		http.Error(w, "can't read input", http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(string(data))
	resp := map[string]interface{}{
		"tokens": tokens,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error(
			"tokenize",
			zap.Error(err),
		)
	}
}
