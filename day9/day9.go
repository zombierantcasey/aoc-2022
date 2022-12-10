package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	open, _ := os.Open("test.txt")
	txt := bufio.NewScanner(open)
	part1(txt)

}

func part1(t *bufio.Scanner) {

	var positions [][]string

	for t.Scan() {
		n := t.Text()
		o := []string{n}
		positions = append(positions, o)
	}

	for v := range positions {
		s := strings.Split(positions[v][0], " ")
		direction := s[0]
		movement, _ := strconv.Atoi(s[1])
		fmt.Println(direction, movement)
	}

}
