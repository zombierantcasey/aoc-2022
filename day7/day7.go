package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fp := "input.txt"
	open, _ := os.Open(fp)
	txt := bufio.NewScanner(open)
	part1_and_2(txt, fp)

}

//My solution involves counting the levels when we cd into a dir. In the main function, iterate over the input, and when "cd xx" is found,
//call returnSizeOfDirectory with its position in the txt file and the dir name. Iterate to the position, increase levels when cd xx occurs,
//increase the size whenever a file is found, reduce count when cd .. is found. When levels == 0, we're back to the root of that directory,
//and can return the size of the dir. Main loop does this for every directory and returned dir sizes are stored in a slice. Note that I'm using two
//seperate iterators for input file. The main loop uses a single instance, and returnSizeOfDirectory uses a second that is reset everytime it's called. Sharing
//the iterator would mean having to move back and forth with the last and next cd position. Sharing a single iterator is kinda fine, but adds a layer of complexity
//that is out of the scope of simply getting the right answer
func returnSizeOfDirectory(dir string, pos int, fp string) int {

	var size int
	var in_folder bool
	var levels int
	counter := 0

	open_2, _ := os.Open(fp)
	te := bufio.NewScanner(open_2)

	for te.Scan() {
		counter++
		if counter < pos {
			continue
		}
		n := te.Text()
		s := strings.Split(n, " ")
		if n == fmt.Sprintf("$ cd %s", dir) {
			in_folder = true
			levels++
			continue
		}
		if in_folder {
			if s[1] == "cd" {
				if s[2] == ".." {
					levels--
				} else {
					levels++
				}
			} else if s[0] == "dir" {
				continue
			} else if s[1] == "ls" {
				continue
			} else {
				si, _ := strconv.Atoi(s[0])
				size = size + si
			}

			if levels == 0 {
				break
			}
		}
	}
	return size
}

func part1_and_2(t *bufio.Scanner, fp string) {

	var results []int
	var sum int
	counter := 0

	for t.Scan() {
		n := t.Text()
		split := strings.Split(n, " ")
		counter++
		if split[1] == "cd" {
			if split[2] == ".." {
				continue
			}
			directory_size := returnSizeOfDirectory(split[2], counter, fp)
			results = append(results, directory_size)
		}
	}

	for _, v := range results {
		if v < 100000 {
			sum = sum + v
		}
	}

	fmt.Println(sum) //part 1
	unused_space := 70000000 - results[0]
	space_needed := 30000000 - unused_space
	var directories_worth_deleting []int

	for _, v := range results { //part 2 could be a single loop
		if v > space_needed {
			directories_worth_deleting = append(directories_worth_deleting, v)
		}
	}

	var last_value int

	for v := range directories_worth_deleting {
		if v == 0 {
			last_value = directories_worth_deleting[v]
		}
		if directories_worth_deleting[v] < last_value {
			last_value = directories_worth_deleting[v]
		}
	}

	fmt.Println(last_value) //part 2
}
