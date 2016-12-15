package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	getLength(file)
}

func getLength(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		fmt.Println(recursive(line))
	}
}

func recursive(line []string) int {
	total := 0
	for i := 0; i < len(line); i++ {
		if line[i] == "(" {
			markerLength := 0
			check := false
			for !check {
				if line[i+markerLength] != ")" {
					markerLength++
				}
				if line[i+markerLength] == ")" {
					check = true
				}
			}

			instruct := strings.Join(line[i:i+markerLength+1], "")
			letters, _ := strconv.Atoi(instruct[1:strings.Index(instruct, "x")])
			amount, _ := strconv.Atoi(instruct[strings.Index(instruct, "x")+1 : len(instruct)-1])
			newLine := line[i+markerLength+1 : i+markerLength+letters+1]

			total += amount * recursive(newLine)
			i += markerLength + letters

		} else {
			total++
		}
	}
	return total
}
