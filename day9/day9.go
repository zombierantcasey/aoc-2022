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

type Position struct {
	x, y int
}

type Rope struct {
	Head         Position
	Tail         Position
	head_visited map[string]string
	tail_visited map[string]string
}

func (r *Rope) absoluteNum(n int) int {

	if n < 0 {
		return -n
	}
	return n

}

func (r *Rope) moveHead(direction string) {
	switch direction {
	case "R":
		r.Head.x++
	case "L":
		r.Head.x--
	case "U":
		r.Head.y++
	case "D":
		r.Head.y--
	}
	r.head_visited[fmt.Sprintf("%d:%d", r.Head.x, r.Head.y)] = direction
}

func (r *Rope) moveTail(direction string) {

	diff_1 := r.absoluteNum(r.Head.x - r.Tail.x)
	diff_2 := r.absoluteNum(r.Head.y - r.Tail.y)

	if diff_1+diff_2 == 3 {
		if r.Head.x == r.Tail.x-2 {
			r.Tail.x--
			a := r.Head.y - r.Tail.y
			r.Tail.y = r.Tail.y + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
		if r.Head.x == r.Tail.x+2 {
			r.Tail.x++
			a := r.Head.y - r.Tail.y
			r.Tail.y = r.Tail.y + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
		if r.Head.y == r.Tail.y-2 {
			r.Tail.y--
			a := r.Head.x - r.Tail.x
			r.Tail.x = r.Tail.x + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
		if r.Head.y == r.Tail.y+2 {
			r.Tail.y++
			a := r.Head.x - r.Tail.x
			r.Tail.x = r.Tail.x + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}

	} else if diff_1 > 1 || diff_2 > 1 {
		if r.Head.x == r.Tail.x-2 {
			r.Tail.x--
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
		if r.Head.x == r.Tail.x+2 {
			r.Tail.x++
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
		if r.Head.y == r.Tail.y-2 {
			r.Tail.y--
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
		if r.Head.y == r.Tail.y+2 {
			r.Tail.y++
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = direction
		}
	}
}

func part1(t *bufio.Scanner) {

	var positions [][]string

	for t.Scan() {
		n := t.Text()
		o := []string{n}
		positions = append(positions, o)
	}

	headPosition := Position{
		x: 0,
		y: 0,
	}

	tailPosition := Position{
		x: 0,
		y: 0,
	}

	head_visited := make(map[string]string)
	tail_visited := make(map[string]string)

	rope := Rope{
		Head:         headPosition,
		Tail:         tailPosition,
		head_visited: head_visited,
		tail_visited: tail_visited,
	}

	for v := range positions {
		if v == 0 {
			head_visited["0:0"] = "S"
			tail_visited["0:0"] = "S"
		}
		s := strings.Split(positions[v][0], " ")
		movement, _ := strconv.Atoi(s[1])
		for i := 1; i <= movement; i++ {
			rope.moveHead(s[0])
			rope.moveTail(s[0])
		}

	}

	fmt.Println(len(rope.tail_visited))
}
