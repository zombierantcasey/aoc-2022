package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	open, _ := os.Open("input.txt")
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

	var head_positions []string
	tail_map := make(map[string]int)

	head_horizontal := 0
	head_lateral := 0

	tail_horizontal := 0
	tail_lateral := 0

	for v := range positions {
		s := strings.Split(positions[v][0], " ")
		direction := s[0]
		movement, _ := strconv.Atoi(s[1])

		switch direction {
		case "R":
			head_lateral = head_lateral + movement

		case "U":
			head_horizontal = head_horizontal + movement
			head_positions = append(head_positions, fmt.Sprintf("%d:%d", head_horizontal, head_lateral))
		case "D":
			head_horizontal = head_horizontal - movement
			head_positions = append(head_positions, fmt.Sprintf("%d:%d", head_horizontal, head_lateral))
		case "L":
			head_lateral = head_lateral - movement
			if head_lateral < 0 {
				if head_lateral+2 == tail_lateral {
					for i := 0; i >= head_lateral; i-- {
						fmt.Println(i)
						tail_map[fmt.Sprintf("%d:%d", tail_horizontal, head_lateral+1)] = head_lateral + 1
					}
				}
			} else if head_lateral-2 == tail_lateral {
				tail_map[fmt.Sprintf("%d:%d", tail_horizontal, tail_lateral)] = head_lateral
			}
		}

	}

	fmt.Println(tail_map)

}
