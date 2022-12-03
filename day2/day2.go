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
	part2(txt)

}

func part1(t *bufio.Scanner) {

	a := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	var total int

	for t.Scan() {
		n := t.Text()
		s := strings.Split(n, " ")
		v1 := a[s[0]]
		v2 := a[s[1]]
		if v2 == v1 {
			total = total + v2 + 3
			continue
		}
		switch v2 {
		case 1:
			if v1 == 3 {
				total = total + v2 + 6
			} else {
				total = total + v2
			}
		case 2:
			if v1 == 1 {
				total = total + v2 + 6
			} else {
				total = total + v2
			}
		case 3:
			if v1 == 2 {
				total = total + v2 + 6
			} else {
				total = total + v2

			}
		}
	}
	fmt.Println(total)
}

func part2(t *bufio.Scanner) {
	b := map[string]int{"A": 1, "B": 2, "C": 3}
	c := map[string]string{"A": "C", "B": "A", "C": "B"}
	var total int
	for t.Scan() {
		n := t.Text()
		s := strings.Split(n, " ")
		switch s[1] {
		case "Y":
			total = total + b[s[0]] + 3
		case "Z":
			for k := range c {
				if c[k] == s[0] {
					total = total + b[k] + 6
				}
			}
		case "X":
			for k := range c {
				if k == s[0] {
					total = total + b[c[k]]
				}
			}

		}
	}

	fmt.Println(total)

}
