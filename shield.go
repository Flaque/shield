package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flaque/shield/tokenizer"
)

func main() {
	b, _ := ioutil.ReadFile(os.Args[1])

	toko := tokenizer.NewTokenizer(string(b))
	tokens := toko.Run()

	fmt.Println(string(b))
	fmt.Println(tokens)
}
