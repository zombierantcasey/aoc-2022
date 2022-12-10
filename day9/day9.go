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

	h_map := make(map[string]int)
	t_map := make(map[string]int)

	h_h := 0
	h_l := 0

	t_h := 0
	t_l := 0

	h_map[fmt.Sprintf("%d:%d", h_h, h_l)] = t_h
	t_map[fmt.Sprintf("%d:%d", t_h, t_l)] = t_h

	for v := range positions {
		s := strings.Split(positions[v][0], " ")
		direction := s[0]
		movement, _ := strconv.Atoi(s[1])
		h_h, h_l, t_h, t_l = addToMap(direction, movement, h_map, t_map, h_h, h_l, t_h, t_l)
		fmt.Println("--")
	}

	fmt.Println(h_map)
	fmt.Println(len(h_map))

	fmt.Println(t_map)
	fmt.Println(len(t_map))

}

func addToMap(direction string, movement int, m map[string]int, t map[string]int, h_h, h_l, t_h, t_l int) (int, int, int, int) {

	fmt.Printf("HEAD -- Horizontal: %d, Lateral: %d\n", h_h, h_l)
	fmt.Printf("TAIL -- Horizontal: %d, Lateral: %d\n", t_h, t_l)

	switch direction {
	case "R":
		for i := 1; i <= movement; i++ {
			m[fmt.Sprintf("%d:%d", h_h, h_l+1)] = i
			h_l = h_l + 1

			if h_l-t_l == 2 {
				t_l = h_l - 1
				t_h = h_h
				fmt.Println(t_h, t_l)
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			} else if h_l+t_l == 2 {
				t_l = h_l + 1
				t_h = h_h
				fmt.Println(t_h, t_l)
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			}
		}
	case "L":
		for i := movement; i >= 1; i-- {
			m[fmt.Sprintf("%d:%d", h_h, h_l-1)] = i
			h_l = h_l - 1

			if t_l-h_l == 2 {
				t_l = h_l + 1
				t_h = h_h
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			} else if t_l+h_l == 2 {
				t_l = h_l - 1
				t_h = h_h
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			}
		}
	case "U":
		for i := 1; i <= movement; i++ {
			m[fmt.Sprintf("%d:%d", h_h+1, h_l)] = i
			h_h = h_h + 1

			if h_h-t_h == 2 {
				t_h = h_h - 1
				t_l = h_l
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			} else if h_h+t_h == 2 {
				t_h = h_h + 1
				t_l = h_l
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			}
		}
	case "D":
		for i := movement; i >= 1; i-- {
			m[fmt.Sprintf("%d:%d", h_h-1, h_l)] = i
			h_h = h_h - 1

			if t_h-h_h == 2 {
				t_h = t_h + 1
				t_l = h_l
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			} else if t_h+h_h == 2 {
				t_h = t_h - 1
				t_l = h_l
				t[fmt.Sprintf("%d:%d", t_h, t_l)] = i
			}
		}

	}

	fmt.Printf("LEAVING HEAD -- Horizontal: %d, Lateral: %d\n", h_h, h_l)
	fmt.Printf("LEAVINGTAIL -- Horizontal: %d, Lateral: %d\n", t_h, t_l)
	return h_h, h_l, t_h, t_l
}
