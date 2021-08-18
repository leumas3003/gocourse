/*
Write a function that gets an index file with names of files and sha256
signatures in the following format
0c4ccc63a912bbd6d45174251415c089522e5c0e75286794ab1f86cb8e2561fd  taxi-01.csv
f427b5880e9164ec1e6cda53aa4b2d1f1e470da973e5b51748c806ea5c57cbdf  taxi-02.csv
4e251e9e98c5cb7be8b34adfcb46cc806a4ef5ec8c95ba9aac5ff81449fc630c  taxi-03.csv
...

You should compute concurrently sha256 signatures of these files and see if
they math the ones in the index file.

- Print the number of processed files
- If there's a mismatch, print the offending file(s) and exit the program with
  non-zero value

Get taxi.zip from the web site and open it. The index file is sha256sum.txt
*/
package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Parse signature file. Return map of path->signature
func parseSigFile(r io.Reader) (map[string]string, error) {
	sigs := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Line example
		// 6c6427da7893932731901035edbb9214  nasa-00.log
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			// TODO: line number
			return nil, fmt.Errorf("bad line: %q", scanner.Text())
		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}

type result struct {
	fileName string
	err      error
	match    bool
}

func sigWorker(fileName, signature string, ch chan<- result) {
	res := result{
		fileName: fileName,
	}

	sig, err := fileSig(fileName)
	if err != nil {
		res.err = err
	} else {
		res.match = sig == signature
	}

	ch <- res
}

/* Process each file in a different goroutine
Think how to:
- Wait for all the goroutines to finish
- know if there's a failure so you can exit with 1
*/
func main() {
	start := time.Now()
	rootDir := "/tmp/taxi"
	file, err := os.Open(path.Join(rootDir, "sha256sum.txt"))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	sigs, err := parseSigFile(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	ch := make(chan result)
	// Part 1: spin workers
	for name, signature := range sigs {
		fileName := path.Join(rootDir, name) + ".bz2"
		go sigWorker(fileName, signature, ch)
	}

	// Part 2: collect results
	ok := true
	for range sigs {
		res := <-ch
		if res.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", res.err)
			ok = false
		}
		if !res.match {
			fmt.Printf("%s: mismatch\n", res.fileName)
			ok = false
		}
	}

	duration := time.Since(start)
	fmt.Printf("processed %d files in %v\n", len(sigs), duration)
	if !ok {
		os.Exit(1)
	}

}
