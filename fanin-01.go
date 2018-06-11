// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	c := fanIn(cat("testdata/a.txt"), cat("testdata/b.txt")) // HL
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)  // HL
	}
	fmt.Println("Done.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }() // HL
	go func() { for { c <- <-input2 } }() // HL
	return c
}

func cat(filename string) <-chan string {
	wCh := make(chan string)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Fields(string(b))
	go func() {
		for _, word := range words {
			wCh <- word // HL
		}
		close(wCh)
	}()

	return wCh
}
