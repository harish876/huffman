package main

import (
	"fmt"
	"testing"
)

func TestHuffmanEncode(t *testing.T) {
	hf := NewHuffmanTree()
	word := "abbcdbccdaabbeeebeab"
	hf.Encode(word)
	fmt.Println("Code for b is", hf.Codes[byte('d')])
}
