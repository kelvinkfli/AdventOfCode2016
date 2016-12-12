package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	valTri()
}

func valTri() {
	counter := 0

	file, _ := os.Open("AOC_input_d3.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stringline := strings.Fields(line)

		// string to int conversion
		intTest := []int{}
		for i := range stringline {
			text := stringline[i]
			number, _ := strconv.Atoi(text)
			intTest = append(intTest, number)
		}

		sort.Ints(intTest)
		fmt.Println(intTest)

		if intTest[0]+intTest[1] <= intTest[2] {
			continue
		}

		counter++
	}
	fmt.Println(counter)
}
