package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	m := parseCraneCreateMaps("crane.txt")

	open, _ := os.Open("input.txt")
	txt := bufio.NewScanner(open)
	part2(txt, m)

}

func parseCraneCreateMaps(file string) [][]string {

	a := "abcdefghijklmnopqrstuvwxyz"
	s_a := strings.Split(a, "")
	open, _ := os.Open(file)
	txt := bufio.NewScanner(open)

	controller := [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	for txt.Scan() {
		n := txt.Text()
		s := strings.Split(n, "")
		for v := range s {
			for letter := range s_a {
				u := strings.ToUpper(s_a[letter])
				if s[v] == u {
					switch v {
					case 1:
						controller[0] = append(controller[0], u)
					case 5:
						controller[1] = append(controller[1], u)
					case 9:
						controller[2] = append(controller[2], u)
					case 13:
						controller[3] = append(controller[3], u)
					case 17:
						controller[4] = append(controller[4], u)
					case 21:
						controller[5] = append(controller[5], u)
					case 25:
						controller[6] = append(controller[6], u)
					case 29:
						controller[7] = append(controller[7], u)
					case 33:
						controller[8] = append(controller[8], u)
					}
				}
			}

		}
	}

	return controller

}

func part1(t *bufio.Scanner, crates [][]string) {

	for t.Scan() {
		n := t.Text()

		re := regexp.MustCompile("[0-9]+")
		s := re.FindAllString(n, -1)
		n_of_crates_to_move, _ := strconv.Atoi(s[0])
		sending_stack, _ := strconv.Atoi(s[1])
		receiving_stack, _ := strconv.Atoi(s[2])

		for i := 0; i < n_of_crates_to_move; i++ {
			crates[receiving_stack-1] = append([]string{crates[sending_stack-1][0]}, crates[receiving_stack-1]...)
			crates[sending_stack-1] = crates[sending_stack-1][1:]
		}

	}

	var top_crates string

	for v := range crates {
		top_crates = top_crates + crates[v][0]

	}

	fmt.Println(top_crates)
}

func part2(t *bufio.Scanner, crates [][]string) {
	for t.Scan() {

		n := t.Text()

		re := regexp.MustCompile("[0-9]+")
		s := re.FindAllString(n, -1)
		n_of_crates_to_move, _ := strconv.Atoi(s[0])
		sending_stack, _ := strconv.Atoi(s[1])
		receiving_stack, _ := strconv.Atoi(s[2])

		adjustable_crates_to_move := n_of_crates_to_move

		if n_of_crates_to_move == 1 {
			crates[receiving_stack-1] = append([]string{crates[sending_stack-1][0]}, crates[receiving_stack-1]...)
			crates[sending_stack-1] = crates[sending_stack-1][1:]
		} else if n_of_crates_to_move > 1 {
			for {
				if adjustable_crates_to_move == 0 {
					break
				}
				crates[receiving_stack-1] = append([]string{crates[sending_stack-1][adjustable_crates_to_move-1]}, crates[receiving_stack-1]...)
				adjustable_crates_to_move--
			}
			crates[sending_stack-1] = crates[sending_stack-1][n_of_crates_to_move:]

		}

	}

	var top_crates string

	for v := range crates {
		top_crates = top_crates + crates[v][0]

	}

	fmt.Println(top_crates)
}
