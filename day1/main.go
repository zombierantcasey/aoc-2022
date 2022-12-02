package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	open, _ := os.Open("input.txt")
	txt := bufio.NewScanner(open)
	part2(txt)
}

func part1(t *bufio.Scanner) {

	c := 0
	m := 0

	for t.Scan() {
		n := t.Text()
		if len(n) == 0 {
			if c > m {
				m = c
			}
			c = 0
			continue
		}
		c_i, _ := strconv.Atoi(n)
		c = c_i + c

	}

	fmt.Println(m)

}

func part2(t *bufio.Scanner) {

	l := make(map[int]int)
	tt := 0
	ll := 1

	for t.Scan() {
		n := t.Text()
		if len(n) == 0 {
			if len(l) >= 3 {
				for k, v := range l {
					if tt > v {
						l[k] = tt
						tt = 0
						break
					} else {
						continue
					}
				}
			} else {
				l[ll] = tt
				ll++
			}
			tt = 0
		}
		c_i, _ := strconv.Atoi(n)
		tt = c_i + tt
	}

	fmt.Println(l[1] + l[2] + l[3])

}
