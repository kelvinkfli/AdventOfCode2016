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

	counter := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		var continueCount int

		for i, v := range line {
			counter++
			fmt.Println(continueCount)

			if continueCount > 0 {
				continueCount--
				continue
			}
			if v == "(" {
				check := false
				count := 1
				for !check {
					if line[i+count] != ")" {
						count++
					}
					if line[i+count] == ")" {
						check = true
					}
				}

				instruct := strings.Join(line[i:i+count+1], "")
				letters, _ := strconv.Atoi(instruct[1:strings.Index(instruct, "x")])
				amount, _ := strconv.Atoi(instruct[strings.Index(instruct, "x")+1 : len(instruct)-1])
				counter += letters * amount
				counter -= letters + len(instruct)
				continueCount = letters + 4
			}
		}
	}
	fmt.Println(counter)
}
