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

func prioritize(letters []string) int {
	a := "abcdefghijklmnopqrstuvwxyz"
	a_s := strings.Split(a, "")
	var sum int
	for v := range a_s {
		for l := range letters {
			if letters[l] == a_s[v] {
				sum = sum + v + 1
			} else {
				c := strings.ToUpper(letters[l])
				f := strings.ToUpper(a_s[v])
				if c == f {
					sum = sum + v + 26 + 1
				}
			}
		}
	}
	return sum
}

func part1(t *bufio.Scanner) {
	var letters []string
	for t.Scan() {
		n := t.Text()
		s := len(n) / 2
		split := strings.Split(n, "")
	loop:
		for v := range split {
			for i := s; len(split) > i; i++ {
				if split[v] == split[i] {
					letters = append(letters, split[v])
					break loop
				}
			}
		}
	}
	fmt.Println(prioritize(letters))
}

func part2(t *bufio.Scanner) {
	var lines [][][]string
	var set [][]string
	var _new_ [][]string
	cc := 0
	for t.Scan() {
		n := t.Text()
		split := strings.Split(n, "")
		cc++
		set = append(set, split)
		if cc == 3 {
			lines = append(lines, set)
			set = _new_
			cc = 0
		}

	}
	var items []string
	for group := range lines {
	next_group:
		for sack := range lines[group] {
			for item_1 := range lines[group][sack] {
				for item_2 := range lines[group][sack+1] {
					if lines[group][sack][item_1] == lines[group][sack+1][item_2] {
						for item_3 := range lines[group][sack+2] {
							if lines[group][sack][item_1] == lines[group][sack+2][item_3] {
								items = append(items, lines[group][sack+2][item_3])
								break next_group
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(prioritize(items))
}
