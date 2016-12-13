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
	string := getMessage(instructions)
	fmt.Println(string)
}

func getMessage(input string) []string {

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	letter0 := map[string]int{}
	letter1 := map[string]int{}
	letter2 := map[string]int{}
	letter3 := map[string]int{}
	letter4 := map[string]int{}
	letter5 := map[string]int{}
	letter6 := map[string]int{}
	letter7 := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(reflect.TypeOf(line))
		sep := strings.Split(line, "")

		letter0[sep[0]]++
		letter1[sep[1]]++
		letter2[sep[2]]++
		letter3[sep[3]]++
		letter4[sep[4]]++
		letter5[sep[5]]++
		letter6[sep[6]]++
		letter7[sep[7]]++
	}

	final := []string{getLetter(letter0), getLetter(letter1), getLetter(letter2), getLetter(letter3), getLetter(letter4), getLetter(letter5), getLetter(letter6), getLetter(letter7)}
	return final
}

func getLetter(myMap map[string]int) string {
	for k, v := range myMap {
		if v == 26 {
			return k
		}
	}
	return ""
}
