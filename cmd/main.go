package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	instructs, ints := ReadInts(strings.NewReader(string(body)))

	distance := determineDist(instructs, ints)

	fmt.Println(distance)

}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]string, []int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var instructs []string
	var ints []int

	for scanner.Scan() {
		x := scanner.Text()
		num, err := strconv.Atoi(x)
		if err != nil {
			instructs = append(instructs, x)
		}
		if num != 0 {
			ints = append(ints, num)
		}
	}
	return instructs, ints
}

func determineDist(instructs []string, ints []int) int {
	var horizontal, depth int

	fmt.Println(instructs)
	fmt.Println(ints)
	for i := 0; i < len(instructs); i++ {
		if instructs[i] == "forward" {
			horizontal += ints[i]
		}
		if instructs[i] == "up" {
			depth -= ints[i]
		}
		if instructs[i] == "down" {
			depth += ints[i]
		}
	}

	fmt.Println("horizontal:", horizontal)
	fmt.Println("depth:", depth)
	return (horizontal * depth)
}
