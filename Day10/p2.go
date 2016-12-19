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

	for scanner.Scan() {
		line := scanner.Text()
		instruct := strings.Split(scanner.Text(), " ")
		if instruct[0] == "value" {
			valueInstruct = append(valueInstruct, line)
		} else {
			currentBot, _ := strconv.Atoi(instruct[1])
			lowBotNum, _ := strconv.Atoi(instruct[6])
			highBotNum, _ := strconv.Atoi(instruct[11])
			lowBotObj := instruct[5]
			highBotObj := instruct[10]

			if myBots[lowBotNum] == nil {
				myBots[lowBotNum] = &bot{
					currentNum: lowBotNum,
				}
			}
			if myBots[highBotNum] == nil {
				myBots[highBotNum] = &bot{
					currentNum: highBotNum,
				}
			}

			if myBots[currentBot] == nil {
				myBots[currentBot] = &bot{
					currentNum: currentBot,
					sendLowTo:  myBots[lowBotNum],
					sendHighTo: myBots[highBotNum],
					lowObj:     lowBotObj,
					highObj:   highBotObj,
				}
			} else {
				myBots[currentBot].sendHighTo = myBots[highBotNum]
				myBots[currentBot].sendLowTo = myBots[lowBotNum]
                myBots[currentBot].lowObj = lowBotObj
                myBots[currentBot].highObj = highBotObj
			}
		}
	}
	for _, v := range valueInstruct {
		line := strings.Split(v, " ")
		number, _ := strconv.Atoi(line[1])
        fmt.Println(number)
		bot, _ := strconv.Atoi(line[5])
		myBots[bot].obtainValue(number)
	}
}

type bot struct {
	values     []int
	currentNum int
	sendLowTo  *bot
	sendHighTo *bot
	lowObj     string
	highObj    string
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

        if b.sendLowTo.currentNum == 0 && b.lowObj == "output" {
            fmt.Println(b, "send to output:",b.sendLowTo.currentNum)
            fmt.Println("low:",lowerNum)
        }
        if b.sendLowTo.currentNum == 1 && b.lowObj == "output" {
            fmt.Println(b, "send to output:",b.sendLowTo.currentNum)
            fmt.Println("low:",lowerNum)
        }
        if b.sendLowTo.currentNum == 2 && b.lowObj == "output" {
            fmt.Println(b, "send to output:",b.sendLowTo.currentNum)
            fmt.Println("low:",lowerNum)
        }
	}
}
