// +build ignore

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // HL
	defer cancel()                                          // HL

	out := <-fanIn(ctx, cat(ctx, "testdata/a.txt"), cat(ctx, "testdata/b.txt")) // HL
	fmt.Println(out)
	fmt.Println("Done.")
}

func fanIn(ctx context.Context, inputs ...<-chan string) <-chan string { // HL
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(input <-chan string) {
			defer wg.Done()
			for i := range input {
				select {
				case c <- i:
				case <-ctx.Done(): // HL
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

func cat(ctx context.Context, filename string) <-chan string { // HL
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
			case <-ctx.Done(): // HL
				return // HL
			}
		}
	}()

	return wCh
}
