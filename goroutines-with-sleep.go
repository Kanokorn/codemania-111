// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	go func() { // HL
		b, err := ioutil.ReadFile("testdata/a.txt")
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(string(b))
		for _, word := range words {
			fmt.Println(word)
		}
	}() // HL

	time.Sleep(1 * time.Second)
}
