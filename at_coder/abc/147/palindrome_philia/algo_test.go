package main

import (
	"fmt"
	"testing"
)

func Test_Algo(t *testing.T) {
	count := 0
	text := "redcoder"
	n1, n2 := Split(text)
	if n1 == Reverse(n2) {
		fmt.Println(0)
		return
	}
	n1Rune := []rune(n1)
	fmt.Println(n1Rune)
	n2ReversedRune := []rune(Reverse(n2))
	fmt.Println(n2ReversedRune)
	for i:=0;i<len(n1);i++ {
		if n1Rune[i] != n2ReversedRune[i] {
			count = count + 1
		}
	}
	fmt.Println(count)
	return
}