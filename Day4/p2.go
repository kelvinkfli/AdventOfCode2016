package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	fmt.Println(secIDSum(filename))
}

func secIDSum(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line, "")
		checkSum := list[len(list)-6 : len(list)-1]
		value := list[len(list)-10 : len(list)-7]
		// fmt.Println(value)
		encrypted := list[0 : len(list)-11]

		letterTracker := map[string]int{}

		for _, letter := range checkSum {
			for _, encryptLetter := range encrypted {
				if letter == encryptLetter {
					letterTracker[letter]++
				}
			}
		}

		n := map[int][]string{}

		max := 0
		for k, v := range letterTracker {
			n[v] = append(n[v], k)
			if v > max {
				max = v
			}
		}

		for _, letters := range n {
			sort.Strings(letters)
		}

		finalstr := ""
		for p := max; p > 0; p-- {
			finalstr += strings.Join(n[p], "")
		}

		original := strings.Join(checkSum, "")
		if finalstr == original {
			increment, err := strconv.Atoi(strings.Join(value, ""))
			if err != nil {
				log.Fatal(err)
			}
			total += increment
			shiftNum := increment % 26
			newList := shift(encrypted, shiftNum)
			newString := strings.Join(newList, "")
			// fmt.Println(newString)
			if strings.Contains(newString, "north") {
				return increment
			}
		}
	}
	return total
}

func shift(letters []string, num int) []string {
	newstr := []string{}
	for _, v := range letters {
		s := []byte(v)
		for _, byte := range s {
			for i := 1; i <= num; i++ {
				byte = byte + 1
				if byte > 122 {
					byte = 97
				}
			}
			newAlpha := string(byte)
			newstr = append(newstr, newAlpha)
		}
	}
	return newstr
}
