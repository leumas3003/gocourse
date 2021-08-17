package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	file, err := os.OpenFile("kill.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Print(err)
	} else {
		log.SetOutput(file)
	}

	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		log.Printf("error: %+v", err)
	}

	/* PANIC
	v := div(1, 0)
	fmt.Println(v)
	*/
	v, err := safeDiv(1, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("safeDiv:", v)
	}
	fmt.Println("DONE")

	q, r := divmod(7, 3)
	fmt.Println(q, r)
}

// Miki don't like this style
func divmod(a, b int) (q int, r int) {
	if b == 0 {
		return
	}
	q = a / b
	r = a % b
	return
}

func safeDiv(a, b int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil { // "e" is interface{}
			//fmt.Println("error:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return div(a, b), nil
}

func div(a, b int) int {
	return a / b
}

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		// return fmt.Errorf("%s: %w", pidFile, err)
		return errors.Wrapf(err, "%s: bad pid", pidFile)
	}

	fmt.Printf("killing %d\n", pid) // simulate kill
	os.Remove(pidFile)              // TODO: What do do with error here?
	return nil
}
