package main

import (
	"fmt"
	"Users/dayoadepoju/Documents/Go/word_quote/quote"
	"Users/dayoadepoju/Documents/Go/word_quote/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k,v := range word.UseCount(quote.SunAlso){
		fmt.Println(v,k)
	}
}
