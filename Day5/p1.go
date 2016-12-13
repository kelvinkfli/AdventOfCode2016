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
	fmt.Println(getPass())
}

func getPass() []string {
	doorID := []string{"a", "b", "b", "h", "d", "w", "s", "y"}
	// doorID := []string{"a", "b", "c"}
	passcode := []string{}

	charCount := 0
	counter := 0

	myHash := md5.New()

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
			digit := string(hexStr[5])
			passcode = append(passcode, digit)
			charCount++
			fmt.Println(hexStr)
		}

		counter++
		myHash.Reset()
	}
	return passcode
}
