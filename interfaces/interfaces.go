package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func sha1Sig(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}

	sig := sha1.New()
	_, err = io.Copy(sig, r)
	if err != nil {
		return "", err
	}

	hex := fmt.Sprintf("%x", sig.Sum(nil))
	return hex, nil
}

func main() {
	sig, err := sha1Sig("~/WS/gocourse/interfaces/httpd.log.gz")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(sig)
	}
	var loc1 Location
	fmt.Println(loc1)
	fmt.Printf("v: %v\n", loc1)
	fmt.Printf("+v: %+v\n", loc1)
	fmt.Printf("#v: %#v\n", loc1)

	loc1.X = 100
	fmt.Printf("loc1: %#v\n", loc1)

	loc2 := Location{100, 200}
	fmt.Printf("loc2: %#v\n", loc2)

	loc3 := Location{
		Y: 300,
		//		X: 400,
	}
	fmt.Printf("loc3: %#v\n", loc3)

	fmt.Println(NewLocation(-1, 7))
	fmt.Println(NewLocation(1, 7))

	loc4 := Location{-1, 20003}
	fmt.Println("loc4", loc4)

	loc4.Move(21, 42)
	fmt.Println("loc4 move", loc4)
	fmt.Println("loc4 kind", loc4.Kind())

	var loc5 *Location
	// You can call methods on nil values
	fmt.Println("loc5 kind", loc5, loc5.Kind())

	p1 := Player{
		//		X:        1.2,
		Name:     "Parzival",
		Location: Location{10, 20},
	}
	fmt.Println("p1.X", p1.X)
	fmt.Println("p1.Location.X", p1.Location.X)
	p1.Move(100, 200)
	fmt.Printf("p1: %#v\n", p1)

	p2, err := NewPlayer("Art3mis", 400, 300)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("p2: %#v\n", p2)
	}

	items := []mover{
		&p1,
		p2,
		&loc1,
	}
	moveAll(maxX/2, maxY/2, items)
	fmt.Println(p1, p2, loc1)
}

/* Interface rules
- Use small interfaces (1-2 methods)
- Accept interfaces, return structs
*/
type mover interface {
	Move(int, int)
}

func moveAll(x, y int, items []mover) {
	for _, m := range items {
		m.Move(x, y)
	}
}

// Write a NewPlayer function, validate the name is not empty
// Make sure x, y are valid as well
func NewPlayer(name string, x, y int) (*Player, error) {
	loc, err := NewLocation(x, y)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, fmt.Errorf("empty name")
	}

	p := Player{name, loc}
	// Go's compiler does escape analysis, p is going to be allocated on the heap
	return &p, nil
}

type Player struct {
	//X        float64
	Name     string
	Location // Player embeds Location
}

func (*Location) Kind() string {
	return "Location"
}

// l is called "receiver"
func (l *Location) Move(x, y int) {
	l.X = x
	l.Y = y
}

func NewLocation(x, y int) (Location, error) {
	if x < minX || x > maxX || y < minY || y > maxY {
		return Location{}, fmt.Errorf("%d/%d out of bounds %d,%d", x, y, minX, minY)
	}

	loc := Location{
		X: x,
		Y: y,
	}
	return loc, nil
}

const (
	minX = 0
	maxX = 1000
	minY = 0
	maxY = 1000
)

type Location struct {
	X int
	Y int
}

/*
Interface design, the Go one (bottom) allows for buffer re-use = performance

type Reader interface {
  Read(n int) ([]byte, err error)
}

type Reader interface {
  Read(p []byte) (n int, err error)
}
*/
