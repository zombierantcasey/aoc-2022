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
	part1_part2(txt, false)
}

//I understand that this is a silly solution, but it seemed like the easiest in my tiny little head. i can't help but wince when I look at it.
//create two slices, iterate over the start and finish and then compare the longer slice to the shorter one. Part 1 and Part 2 are in a single function
func returnGroups(s1, s2, s3, s4 int) ([]int, []int) {
	var group_1 []int
	var group_2 []int

	for i := s1; i <= s2; i++ {
		group_1 = append(group_1, i)
	}

	for i := s3; i <= s4; i++ {
		group_2 = append(group_2, i)
	}
	return group_1, group_2
}

func part1_part2(t *bufio.Scanner, part bool) {
	var part_1_count int
	var part_2_count int
	var matched bool
	for t.Scan() {
		n := t.Text()
		split := strings.Split(n, ",")
		p1 := strings.Split(split[0], "-")
		p2 := strings.Split(split[1], "-")

		_s1, _ := strconv.Atoi(p1[0])
		_s2, _ := strconv.Atoi(p1[1])
		_s3, _ := strconv.Atoi(p2[0])
		_s4, _ := strconv.Atoi(p2[1])

		group_1, group_2 := returnGroups(_s1, _s2, _s3, _s4)
		var c int
		if len(group_1) > len(group_2) || len(group_1) == len(group_2) {
			for v := range group_1 {
				for v2 := range group_2 {
					if group_1[v] == group_2[v2] {
						matched = true
						c++
					}
				}
			}

			if len(group_2) == c {
				part_1_count++
			}
		} else {
			for v := range group_2 {
				for v2 := range group_1 {
					if group_2[v] == group_1[v2] {
						matched = true
						c++
					}
				}
			}
			if len(group_1) == c {
				part_1_count++
			}
		}
		if matched {
			part_2_count++
		}
		matched = false
	}

	if part {
		fmt.Println(part_1_count)
	} else {
		fmt.Println(part_2_count)
	}
}
