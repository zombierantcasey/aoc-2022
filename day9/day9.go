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

	head_horizontal := 0
	head_lat := 0

	tail_horizontal := 0
	tail_lat := 0

	tail_map := make(map[string]int)

	for v := range positions {
		s := strings.Split(positions[v][0], " ")
		direction := s[0]
		movement, _ := strconv.Atoi(s[1])

		switch direction {
		case "R":
			head_lat = head_lat + movement
		case "L":
			head_lat = head_lat - movement
		case "U":
			head_horizontal = head_horizontal + movement
		case "D":
			head_horizontal = head_horizontal - movement
		}

		fmt.Println(head_horizontal, head_lat)

		fmt.Println(tail_horizontal, tail_lat)

		if head_lat-tail_lat > 1 {
			for i := 0; i < head_lat; i++ {
				if i == head_lat-1 {
					tail_lat = head_lat - 1
					break
				}
				tail_lat = head_lat + 1
				tail_map[fmt.Sprintf("%d:%d", tail_lat, tail_lat)] = 0
			}
		} else if head_lat-tail_lat < -1 {
			for i := 0; i > head_lat; i-- {
				if i == head_lat+1 {
					break
				}
				tail_lat = tail_lat - 1
				tail_map[fmt.Sprintf("%d:%d", tail_lat, tail_lat)] = 0
			}
		}

		if head_horizontal-tail_horizontal > 1 {
			for i := 0; i < head_horizontal; i++ {
				if i == head_horizontal-1 {
					break
				}
				tail_horizontal = tail_horizontal + 1
				tail_map[fmt.Sprintf("%d:%d", tail_lat, tail_lat)] = 0
			}
		} else if head_horizontal-tail_horizontal < -1 {
			for i := 0; i > head_horizontal; i-- {
				if i == head_horizontal+1 {
					break
				}
				tail_horizontal = tail_horizontal - 1
				tail_map[fmt.Sprintf("%d:%d", tail_lat, tail_lat)] = 0
			}
		}
	}

	fmt.Println(tail_map)

}
