// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func main() {
	done := make(chan struct{}) // HL
	defer close(done)           // HL

	out := <-fanIn(done, cat(done, "testdata/a.txt"), cat(done, "testdata/b.txt")) // HL
	fmt.Println(out)
	fmt.Println("Done.")
}

func fanIn(done <-chan struct{}, inputs ...<-chan string) <-chan string { // HL
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(input <-chan string) {
			defer wg.Done()
			for i := range input {
				select {
				case c <- i:
				case <-done: // HL
					return // HL
				}
			}
		}(input)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func cat(done <-chan struct{}, filename string) <-chan string { // HL
	wCh := make(chan string)

	go func() {
		defer close(wCh)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(string(b))
		for _, word := range words {
			select {
			case wCh <- word:
			case <-done: // HL
				return // HL
			}
		}
	}()

	return wCh
}
