package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	instructions := "input.txt"
	getTLS(instructions)
}

func getTLS(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0
	tracker := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		right := 0
		wrong := 0

		for i := 1; i < len(line)-2; i++ {

			if line[i] == "[" {
				tracker = false
			}

			if line[i] == "]" {
				tracker = true
			}

			if line[i] == line[i+1] {
				if line[i-1] == line[i+2] {
					if tracker == true {
						right++
					}
					if tracker == false {
						wrong++
					}
				}
			}
		}
		if right == 1 && wrong < 1 {
			counter++
		}
	}
	fmt.Println(counter)
}
