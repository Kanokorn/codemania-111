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
	var wg sync.WaitGroup

	wg.Add(1)
	wCh := cat("testdata/a.txt", &wg)

	go func() {
		for {
			fmt.Println(<-wCh) // HL
		}
	}()

	wg.Wait()
}

func cat(filename string, wg *sync.WaitGroup) chan string {
	wCh := make(chan string) // HL

	go func() {
		defer wg.Done()
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(string(b))
		for _, word := range words {
			wCh <- word // HL
		}
	}()

	return wCh
}
