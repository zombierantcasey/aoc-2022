package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	open, _ := os.Open("input.txt")
	txt := bufio.NewScanner(open)
	part1_2(txt, 14)

}

//part 1 and part 2
func part1_2(t *bufio.Scanner, counter int) {

	var split []string
	var global_couter = 0
	var reduce bool
	copy_counter := counter

	for t.Scan() {

		n := t.Text()
		split = strings.Split(n, "")
	}
loop:
	for n := range split {
		for i := 1; i < copy_counter; i++ {
			if n+i >= len(split)-1 {
				break
			}
			if split[n] == split[n+i] {
				global_couter = 0
				reduce = false
				copy_counter = counter
				break
			}
			if i == copy_counter-1 {
				reduce = true
				global_couter++
			}
		}
		if reduce {
			copy_counter--
		}
		if global_couter+1 == counter {
			fmt.Println(n + 2)
			break loop
		}

	}

}
