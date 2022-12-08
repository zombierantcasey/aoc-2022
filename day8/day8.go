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

	var trees int
	var all_trees [][]int
	var total_scores []int

	for t.Scan() {
		n := t.Text()
		s := strings.Split(n, "")
		var one_row_trees []int
		for v := range s {
			tree, _ := strconv.Atoi(s[v])
			one_row_trees = append(one_row_trees, tree)
		}
		all_trees = append(all_trees, one_row_trees)

	}

	for v := range all_trees {
		if v == 0 {
			trees = trees + len(all_trees[v])
			continue
		} else if v == len(all_trees)-1 {
			trees = trees + len(all_trees[v])
			continue
		}

		for pos := range all_trees[v] {
			if pos == 0 {
				trees++
				continue
			} else if pos == len(all_trees[v])-1 {
				trees++
				continue
			}

			if checkAllOtherTrees(v, pos, all_trees) {
				trees++
			}

			scenic_score := returnTreeScore(v, pos, all_trees)
			total_scores = append(total_scores, scenic_score)

		}
	}

	fmt.Println(trees)

	var current_score int

	for v := range total_scores {
		if v == 0 {
			current_score = total_scores[v]
		}
		if total_scores[v] > current_score {
			current_score = total_scores[v]
		}

	}

	fmt.Println(current_score)
}

func returnTreeScore(row, position int, trees [][]int) int {

	down_row := row
	up_row := row
	left_check := position
	right_check := position

	down_distance := 0
	up_distance := 0
	right_distance := 0
	left_distance := 0

	for {
		down_row--
		if trees[down_row][position] > trees[row][position] {
			down_distance++
			break
		} else if trees[down_row][position] == trees[row][position] {
			down_distance++
			break
		} else {
			down_distance++
		}

		if down_row == 0 {
			break
		}

	}

	for {

		up_row++
		if trees[up_row][position] > trees[row][position] {
			up_distance++
			break
		} else if trees[up_row][position] == trees[row][position] {
			up_distance++
			break
		} else {
			up_distance++
		}

		if up_row == len(trees)-1 {
			break
		}

	}

	for {

		left_check--
		if trees[row][left_check] > trees[row][position] {
			left_distance++
			break
		} else if trees[row][left_check] == trees[row][position] {
			left_distance++
			break
		} else {
			left_distance++
		}

		if left_check == 0 {
			break
		}

	}

	for {

		right_check++
		if trees[row][right_check] > trees[row][position] {
			right_distance++
			break
		} else if trees[row][right_check] == trees[row][position] {
			right_distance++
			break
		} else {
			right_distance++
		}

		if right_check == len(trees[row])-1 {
			break
		}
	}

	return up_distance * down_distance * right_distance * left_distance

}

func checkAllOtherTrees(row, position int, trees [][]int) bool {

	down_row := row
	up_row := row
	left_check := position
	right_check := position

	for {
		down_row--
		if trees[down_row][position] >= trees[row][position] {
			break
		}

		if down_row == 0 {
			return true
		}

	}

	for {

		up_row++
		if trees[up_row][position] >= trees[row][position] {
			break
		}

		if up_row == len(trees)-1 {
			return true
		}

	}

	for {

		left_check--
		if trees[row][left_check] >= trees[row][position] {
			break
		}

		if left_check == 0 {
			return true
		}

	}

	for {

		right_check++
		if trees[row][right_check] >= trees[row][position] {
			break
		}

		if right_check == len(trees[row])-1 {
			return true
		}
	}

	return false

}
