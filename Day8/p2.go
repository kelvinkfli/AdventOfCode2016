// same as p1 lul

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
	numLights(file)
}

func numLights(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0

	row1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	row2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	row3 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	row4 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	row5 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	row6 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	keyPAD := [][]int{row1, row2, row3, row4, row5, row6}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		// fmt.Println(line)

		if line[3] == "t" {
			widthValue, _ := strconv.Atoi(strings.Join(line[5:strings.Index(scanner.Text(), "x")], ""))
			lengthValue, _ := strconv.Atoi(line[len(line)-1])
			for l := 0; l < lengthValue; l++ {
				for w := 0; w < widthValue; w++ {
					keyPAD[l][w] = 1
				}
			}
		} else if line[7] == "r" {
			rowNum, _ := strconv.Atoi(line[13])
			rowValue, _ := strconv.Atoi(strings.Join(line[18:len(line)], ""))
			newRow := make([]int, 50)

			for i, v := range keyPAD[rowNum] {
				newpos := (i + rowValue) % 50
				newRow[newpos] = v
			}

			keyPAD[rowNum] = newRow
		} else if line[7] == "c" {
			// fmt.Println(line)
			colNum, _ := strconv.Atoi(strings.Join(line[16:strings.Index(scanner.Text(), "b")-1], ""))
			colValue, _ := strconv.Atoi(line[len(line)-1])
			oldCol := make([]int, 6)
			newCol := make([]int, 6)
			// fmt.Println(colNum, colValue)
			for i := 0; i < 6; i++ {
				currentVal := keyPAD[i][colNum]
				oldCol = append(oldCol, currentVal)
				for i, v := range oldCol {
					newpos := (i + colValue) % 6
					newCol[newpos] = v
				}
			}
			for i := 0; i < 6; i++ {
				keyPAD[i][colNum] = newCol[i]
			}
		}
	}
	fmt.Println(keyPAD)
	for i := 0; i < 6; i++ {
		for _, v := range keyPAD[i] {
			if v == 1 {
				counter++
			}
		}
	}
	fmt.Println(counter)
}
