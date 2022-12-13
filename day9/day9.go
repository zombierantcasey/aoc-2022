package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Had some trouble with this problem. Re-wrote it in OOP design because I originally wrote it like a degenerate idiot, and it was very hard to
//read when troubleshooting.

func main() {

	open, _ := os.Open("input.txt")
	txt := bufio.NewScanner(open)
	part1_and_2(txt)

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
		r.Head.y--
	case "D":
		r.Head.y++
	}
	r.head_visited[fmt.Sprintf("%d:%d", r.Head.x, r.Head.y)] = direction
}

func (r *Rope) moveTail(direction, movement string) {

	diff_1 := r.absoluteNum(r.Head.x - r.Tail.x)
	diff_2 := r.absoluteNum(r.Head.y - r.Tail.y)

	if diff_1+diff_2 == 3 {
		if r.Head.x == r.Tail.x-2 {
			r.Tail.x--
			a := r.Head.y - r.Tail.y
			r.Tail.y = r.Tail.y + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
		if r.Head.x == r.Tail.x+2 {
			r.Tail.x++
			a := r.Head.y - r.Tail.y
			r.Tail.y = r.Tail.y + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
		if r.Head.y == r.Tail.y-2 {
			r.Tail.y--
			a := r.Head.x - r.Tail.x
			r.Tail.x = r.Tail.x + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
		if r.Head.y == r.Tail.y+2 {
			r.Tail.y++
			a := r.Head.x - r.Tail.x
			r.Tail.x = r.Tail.x + a
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}

	} else if diff_1 > 1 || diff_2 > 1 {
		if r.Head.x == r.Tail.x-2 {
			r.Tail.x--
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
		if r.Head.x == r.Tail.x+2 {
			r.Tail.x++
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
		if r.Head.y == r.Tail.y-2 {
			r.Tail.y--
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
		if r.Head.y == r.Tail.y+2 {
			r.Tail.y++
			r.tail_visited[fmt.Sprintf("%d:%d", r.Tail.x, r.Tail.y)] = fmt.Sprintf(" %s%s", direction, movement)
		}
	}
}

func part1_and_2(t *bufio.Scanner) {

	var positions [][]string

	for t.Scan() {
		n := t.Text()
		o := []string{n}
		positions = append(positions, o)
	}

	var ropeKnots []Rope

	for i := 1; i < 10; i++ {
		r := Rope{Head: Position{x: 0, y: 0},
			Tail:         Position{x: 0, y: 0},
			head_visited: make(map[string]string),
			tail_visited: make(map[string]string)}
		ropeKnots = append(ropeKnots, r)
	}

	v_2_tail_visited := make(map[string]string)

	for v := range positions {
		if v == 0 {
			for r := range ropeKnots {
				ropeKnots[r].head_visited["0:0"] = "S"
				ropeKnots[r].tail_visited["0:0"] = "S"
			}
		}
		s := strings.Split(positions[v][0], " ")
		movement, _ := strconv.Atoi(s[1])

		for i := 1; i <= movement; i++ {
			for r := range ropeKnots {
				if r == 0 {
					ropeKnots[r].moveHead(s[0])
					ropeKnots[r].moveTail(s[0], s[1])
					continue
				}
				ropeKnots[r].Head = ropeKnots[r-1].Tail
				ropeKnots[r].moveTail(s[0], s[1])
				if r == 8 {
					v_2_tail_visited[fmt.Sprintf("%d:%d", ropeKnots[r].Tail.x, ropeKnots[r].Tail.y)] = fmt.Sprintf(" %s%d", s[0], movement)
				}

			}

		}
	}

	fmt.Println(len(ropeKnots[0].tail_visited)) //part 1
	fmt.Println(len(v_2_tail_visited))          //part 2

}
