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
	part1(txt)
}

func part1(t *bufio.Scanner) {
	x := 1
	clock_tick := 1
	var process_counter int
	next_tick := make(chan bool)
	var interval_values []int
	interval_cycles := []int{20, 60, 100, 140, 180, 220}

	wg := new(sync.WaitGroup)

	for t.Scan() {
		n := t.Text()
		go executeProgramLine(next_tick, n, &clock_tick, &x, &process_counter, wg, &interval_values, interval_cycles)
		wg.Add(1)
		next_tick <- true

	}

	go func() {
		wg.Wait()
		fmt.Println(interval_values)
		var sum int
		for v := range interval_values {
			sum = interval_values[v] + sum
		}
		fmt.Println(sum)
	}()

	for {
		next_tick <- true
	}
}

func executeProgramLine(receive chan bool, values string, cycle, x, process_counter *int, wg *sync.WaitGroup, interval_values *[]int, interval_cycles []int) {

	defer wg.Done()
	s := strings.Split(values, " ")
	var cycles int
	var increase_value int

	if len(s) > 1 {
		*cycle = *cycle + 2
		cycles = 2
		increase_value, _ = strconv.Atoi(s[1])
		for v := range interval_cycles {
			if interval_cycles[v] == *cycle-1 {
				c := *cycle - 1

				*interval_values = append(*interval_values, *x*c)
			}
		}
	} else {
		cycles = 1
		*cycle++
		increase_value = 0
	}
	cycles--
	for {
		t := <-receive
		if t && cycles != 0 {
			cycles--
		} else {
			for v := range interval_cycles {
				if interval_cycles[v] == *cycle {
					c := *cycle
					*interval_values = append(*interval_values, *x*c)
				}
			}
			break
		}
		if cycles == 0 {
			*x = *x + increase_value
			for v := range interval_cycles {
				if interval_cycles[v] == *cycle {
					c := *cycle
					*interval_values = append(*interval_values, *x*c)
				}
			}
			break
		}

	}

}
