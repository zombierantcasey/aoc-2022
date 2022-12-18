package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	open, _ := os.Open("input.txt")
	txt := bufio.NewScanner(open)
	part1_and_2(txt)
}

//this is generally a bad solution because you can't guarantee when a routine will write/end, which means that while most of the time
//the answers to either part are correct, there's a small chance that both part 1 and part 2 may be incorrect
//it's not a concern for a problem like this, but is clearly unacceptable for normal applications.

func part1_and_2(t *bufio.Scanner) {
	x := 1
	clock_tick := 1
	next_tick := make(chan bool)
	var interval_values []int
	interval_cycles := []int{20, 60, 100, 140, 180, 220}
	wg := new(sync.WaitGroup)
	var position int

	for t.Scan() {
		n := t.Text()
		go executeProgramLine(next_tick, n, &clock_tick, &x, &position, wg, &interval_values, interval_cycles)
		wg.Add(1)
		next_tick <- true
	}

	go func() {
		wg.Wait()
		var sum int
		for v := range interval_values {
			sum = interval_values[v] + sum
		}
		fmt.Println(sum) //part 1
		os.Exit(1)       //bad thread management. channel deadlock will occur here unless I force an exit. additionally, 1 in 10 runs will generate the wrong total for part 1.
	}()

	for {
		next_tick <- true
	}
}

func executeProgramLine(receive chan bool, values string, cycle, x, position *int, wg *sync.WaitGroup, interval_values *[]int, interval_cycles []int) {
	defer wg.Done()
	s := strings.Split(values, " ")
	var cycles int
	var increase_value int

	if len(s) > 1 {
		if *cycle == 1 {
			draw(position, x)
		}
		*cycle = *cycle + 2
		cycles = 2
		increase_value, _ = strconv.Atoi(s[1])
		draw(position, x)
		for v := range interval_cycles {
			if interval_cycles[v] == *cycle-1 {
				c := *cycle - 1
				*interval_values = append(*interval_values, *x*c)
				break
			}
		}
	} else {
		cycles = 1
		*cycle++
	}
	cycles--
	t := <-receive
	if t && cycles != 0 {
		cycles--
	} else {
		draw(position, x)
		for v := range interval_cycles {
			if interval_cycles[v] == *cycle {
				c := *cycle
				*interval_values = append(*interval_values, *x*c)
				return
			}
		}
		return
	}
	if cycles == 0 {
		*x = *x + increase_value
		for v := range interval_cycles {
			if interval_cycles[v] == *cycle {
				c := *cycle
				*interval_values = append(*interval_values, *x*c)
				break
			}
		}
		draw(position, x)
	}
}

func draw(position, cycle *int) { //part2
	if *position >= 40 {
		*position = 0
		fmt.Println()
	}
	if *position+1 == *cycle || *position-1 == *cycle || *position == *cycle {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	*position++
}
