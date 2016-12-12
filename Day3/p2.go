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
	valTriangle(filename)
}

func valTriangle(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	obj := [][]string{}
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		newline := strings.Fields(line)
		obj = append(obj, newline)
		if len(obj) == 3 {
			validnum := evalDimensions(obj)
			counter += validnum
			obj = nil
		}
	}
	fmt.Println(counter)
}

func evalDimensions(object [][]string) int {
	counter := 0

	triangles := []string{}
	for i := 0; i < 3; i++ {
		triangles = append(triangles, object[0][i], object[1][i], object[2][i])
	}

	intList := []int{}
	for _, v := range triangles {
		num, _ := strconv.Atoi(v)
		intList = append(intList, num)
	}

	testList := []int{}

	for i := 0; i <= 6; i++ {
		if i%3 != 0 {
			continue
		}
		testList := append(testList, intList[i], intList[i+1], intList[i+2])
		sort.Ints(testList)
		if testList[0]+testList[1] <= testList[2] {
			continue
		}
		counter++
	}
	return counter
}
