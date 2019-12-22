package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read standard input
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan(){
		// read line
		text := stdin.Text()
		// input value := "4 5 6"
		slice := strings.Split(text, " ")
		sum := 0
		for _, v := range slice {
			a, _ := strconv.Atoi(v)
			sum = sum + a
		}
		if sum >= 22 {
			fmt.Print("bust")
		} else if sum <= 21 {
			fmt.Print("win")
		}
	}
}
