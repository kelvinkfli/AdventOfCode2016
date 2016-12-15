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
	defer file.Close()

	counter := 0

	isOutside := true

	scanner := bufio.NewScanner(file)
Top:
	for scanner.Scan() {
		aba := []string{}
		bab := []string{}

		line := strings.Split(scanner.Text(), "")

		for i := 0; i < len(line)-2; i++ {
			if line[i] == "[" {
				isOutside = false
				continue
			}

			if line[i] == "]" {
				isOutside = true
				continue
			}

			if line[i] == line[i+2] && line[i+1] != line[i] {
				firstLetter := line[i]
				secondLetter := line[i+1]
				pair := firstLetter + secondLetter
				if isOutside {
					aba = append(aba, pair)
				} else {
					bab = append(bab, pair)
				}
			}
		}

		for _, a := range aba {
			aba1, aba2 := a[0], a[1]
			for _, b := range bab {
				bab1, bab2 := b[0], b[1]
				if aba1 == bab2 && aba2 == bab1 {
					counter++
					continue Top
				}
			}
		}
	}

	fmt.Println(counter)
}
