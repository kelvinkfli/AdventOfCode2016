package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	secIDSum(filename)
}

func secIDSum(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line, "")
		checkSum := list[len(list)-6 : len(list)-1]
		encrypted := list[0 : len(list)-11]

		letterTracker := map[string]int{}

		for _, letter := range checkSum {
			for _, encryptLetter := range encrypted {
				if letter == encryptLetter {
					letterTracker[letter]++
				}
			}
		}
		fmt.Println(letterTracker)
		for k, v := range letterTracker {

		}
	}
}
