package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "Go ♡"
	fmt.Println("len", len(text))
	// rune
	c := text[0]
	// byte is an alias to uint8
	fmt.Printf("%c (%T)\n", c, c)
	for i := range text {
		fmt.Println(i)
	}
	// rune is an alias to int32
	for i, c := range text {
		// fmt.Println(i, c)
		fmt.Printf("%d: %c (%T)\n", i, c, c)
	}

	fmt.Println("# runes:", utf8.RuneCountInString(text))
	// text[0] = 'g' // strings are immutable
	// Hello hello

	name, age := "Daffy", 83
	s := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(s)

	fmt.Printf("%-20s is %04d year old\n", name, age)

	v1, v2 := 1, "1"
	fmt.Printf("%v == %v\n", v1, v2)
	fmt.Printf("%#v == %#v\n", v1, v2)

	fmt.Println("euler 14:", euler14())

	s1 := fmt.Sprint(78)
	fmt.Println(s1) // N (not "78")

	//Slices
	vec1 := []int{1, 2, 3}
	fmt.Println("vec1", vec1)
	vec1[1] = 200
	fmt.Println("vec1", vec1)
	vec1 = append(vec1, 4)
	fmt.Println("vec1", vec1)

	nums := []int{1, 2, 3}

	if d, err := dot(nums, nums); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("dot", d)
	}

	// fmt.Println(nums[7]) // panic

	// slicing
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	s3 := s2[2:5] // slices are half-open range, in math - [)
	fmt.Println(s3)
	s3[0] = 100
	fmt.Println(s2, s3)
	s2 = append(s2, 8)
	s3[0] = 200
	fmt.Println(s2, s3)

	s4 := []int{1, 2, 3}
	s4 = appendInt(s4, 4)
	fmt.Println(s4)
	s4 = appendInt(s4, 5)
	fmt.Println(s4)

	values := []float64{1, 3, 2}
	if m, err := median(values); err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println(values, "->", m)
	}

	//Maps
	var stocks map[string]float64
	fmt.Printf("stocks: len=%d, type=%T\n", len(stocks), stocks)
	goog := stocks["GOOG"]
	fmt.Println("goog", goog)

	goog, ok := stocks["GOOG"] // "comma, ok"
	if !ok {
		fmt.Println("GOOG not in stocks")
	} else {
		fmt.Println("goog", goog)
	}

	file, err := os.Open("road.txt")
	
	if err != nil {
		log.Fatal(err)
	}
	freq := frequency(file)
	fmt.Println(freq)
}

func frequency(r io.Reader) map[string]int {
	freq := make(map[string]int)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := strings.ToLower(s.Text())
		freq[word]++
	}
	return freq
}

/* median
- if empty return error
- sort values (sort.Float64)
- if odd number of values return middle
- if even number of values return average of middle
*/
func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}
	return (nums[i-1] + nums[i]) / 2, nil
}

func appendInt(s []int, v int) []int {
	size := len(s)
	if size < cap(s) { // There space in the underlying array
		fmt.Println("DEBUG: enough space")
		s = s[:size+1]
	} else { // need to re-allocate & copy
		fmt.Println("DEBUG: re-allocate")
		newSize := size * 2
		arr := make([]int, newSize)
		copy(arr, s)
		s = arr[:size+1]
	}
	s[size] = v
	return s
}

// dot product v1·v2 = v1[0]ⅹv2[0] + v1[1]ⅹv2[1] ...
func dot(v1, v2 []int) (int, error) {
	if len(v1) != len(v2) {
		return 0, fmt.Errorf("length mismatch")
	}

	d := 0
	for i, v := range v1 {
		d += v * v2[i]
	}

	return d, nil
}

/*
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is
	9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.

Answer:  906609
*/
func euler14() int {
	maxPali := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if m := i * j; isPalindrome(m) && m > maxPali {
				maxPali = m
			}
		}
	}
	return maxPali
}

func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n) // strconv.Iota
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
