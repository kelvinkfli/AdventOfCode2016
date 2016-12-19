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
	getBotNum(file)
}

func getBotNum(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valueInstruct := []string{}
	myBots := map[int]*bot{}
	myOutput := map[int]*bot{}

	for scanner.Scan() {
		line := scanner.Text()
		instruct := strings.Split(scanner.Text(), " ")
		if instruct[0] == "value" {
			valueInstruct = append(valueInstruct, line)
		} else {
			currentBot, _ := strconv.Atoi(instruct[1])
			lowBotNum, _ := strconv.Atoi(instruct[6])
			highBotNum, _ := strconv.Atoi(instruct[11])

			var lowerMap map[int]*bot
			var higherMap map[int]*bot

			if instruct[5] == "bot" {
				lowerMap = myBots
			} else {
				lowerMap = myOutput
			}

			if instruct[10] == "bot" {
				higherMap = myBots
			} else {
				higherMap = myOutput
			}

			if lowerMap[lowBotNum] == nil {
				lowerMap[lowBotNum] = &bot{
					currentNum: lowBotNum,
				}
			}
			if higherMap[highBotNum] == nil {
				higherMap[highBotNum] = &bot{
					currentNum: highBotNum,
				}
			}

			if myBots[currentBot] == nil {
				myBots[currentBot] = &bot{
					currentNum: currentBot,
					sendLowTo:  lowerMap[lowBotNum],
					sendHighTo: higherMap[highBotNum],
				}
			} else {
				myBots[currentBot].sendHighTo = lowerMap[highBotNum]
				myBots[currentBot].sendLowTo = higherMap[lowBotNum]
			}
		}
	}
	for _, v := range valueInstruct {
		line := strings.Split(v, " ")
		number, _ := strconv.Atoi(line[1])
		bot, _ := strconv.Atoi(line[5])
		myBots[bot].obtainValue(number)
	}
}

type bot struct {
	values     []int
	currentNum int
	sendLowTo  *bot
	sendHighTo *bot
}

func (b *bot) obtainValue(num int) {
	b.values = append(b.values, num)
	b.scanChips()
}

func (b *bot) scanChips() {
	if len(b.values) == 2 {
		var lowerNum int
		var higherNum int
		if b.values[0] < b.values[1] {
			lowerNum, higherNum = b.values[0], b.values[1]
		} else {
			lowerNum, higherNum = b.values[1], b.values[0]
		}
		if lowerNum == 17 && higherNum == 61 {
			fmt.Println(b.currentNum)
		}
		b.sendLowTo.obtainValue(lowerNum)
		b.sendHighTo.obtainValue(higherNum)
	}
}
