package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	password := getPass())
	final := []string{}

	for i := 0; i < 8; i++ {
		final = append(final, password[i])
	}
	fmt.Println(final)
}

func getPass() map[int]string {
	doorID := []string{"a", "b", "b", "h", "d", "w", "s", "y"}

	charCount := 0
	counter := 0

	myHash := md5.New()

	myMap := map[int]string{}

	for charCount < 8 {
		int := strconv.Itoa(counter)
		newdoorID := doorID
		newdoorID = append(newdoorID, int)
		hash := strings.Join(newdoorID, "")
		io.WriteString(myHash, hash)
		hexStr := hex.EncodeToString(myHash.Sum(nil))
		// fmt.Println(hexStr)
		proceed := true

		for j := 0; j < 5; j++ {
			letter := string(hexStr[j])
			if letter != "0" {
				proceed = false
				break
			}
		}

		if proceed == true {
			position, _ := strconv.Atoi(string(hexStr[5]))
			digit := string(hexStr[6])
			if position <= 7 {
				if _, ok := myMap[position]; ok {
					continue
				}

				myMap[position] = digit

				charCount++
			}
		}

		counter++
		myHash.Reset()
	}
	return myMap
}
