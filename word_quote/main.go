package main

import (
	"fmt"
	"github.com/adepoju/Golang/tree/master/word_quote/quote"
	"github.com/adepoju/Golang/tree/master/word_quote/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k,v := range word.UseCount(quote.SunAlso){
		fmt.Println(v,k)
	}
}
