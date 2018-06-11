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
	var wg sync.WaitGroup // HL

	wg.Add(1) // HL
	go func() {
		defer wg.Done() // HL
		b, err := ioutil.ReadFile("testdata/a.txt")
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(string(b))
		for _, word := range words {
			fmt.Println(word)
		}
	}()

	wg.Wait() // HL
}
