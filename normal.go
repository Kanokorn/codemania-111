package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filenames := []string{"testdata/a.txt", "testdata/b.txt"}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Fields(string(b))
		for _, word := range words {
			fmt.Println(word)
		}
	}
}
