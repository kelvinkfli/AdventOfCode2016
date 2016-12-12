package main

import (
    "fmt"
    "strings"
	"strconv"
)

func main()  {
    rawInput := "R2, L5, L4, L5, R4, R1, L4, R5, R3, R1, L1, L1, R4, L4, L1, R4, L4, R4, L3, R5, R4, R1, R3, L1, L1, R1, L2, R5, L4, L3, R1, L2, L2, R192, L3, R5, R48, R5, L2, R76, R4, R2, R1, L1, L5, L1, R185, L5, L1, R5, L4, R1, R3, L4, L3, R1, L5, R4, L4, R4, R5, L3, L1, L2, L4, L3, L4, R2, R2, L3, L5, R2, R5, L1, R1, L3, L5, L3, R4, L4, R3, L1, R5, L3, R2, R4, R2, L1, R3, L1, L3, L5, R4, R5, R2, R2, L5, L3, L1, L1, L5, L2, L3, R3, R3, L3, L4, L5, R2, L1, R1, R3, R4, L2, R1, L1, R3, R3, L4, L2, R5, R5, L1, R4, L5, L5, R1, L5, R4, R2, L1, L4, R1, L1, L1, L5, R3, R4, L2, R1, R2, R1, R1, R3, L5, R1, R4"
    fmt.Println(getFinalDistance(rawInput))
}

func getFinalDistance(input string) int {
    instructions := strings.Split(input, ", ")

    horizontalDist := 0
    verticalDist := 0

    // 0 = N, 1 = E, 2 = S, 3 = W
    Position := 0

    for _, v := range instructions {
        direction := string(v[0])
        distance, err := strconv.Atoi(v[1:])
        fmt.Println(distance)
        if err != nil {
            fmt.Println("whatever, fuck errors")
        }


        if direction == "R" {
            Position = (Position + 1)%4
        } else {
            Position = (Position - 1)
        }

        if Position == -1 {
            Position = 3
        }
        
        if Position == 0 {
            verticalDist += distance
        } else if Position == 1 {
            horizontalDist += distance
        } else if Position == 2 {
            verticalDist -= distance
        } else {
            horizontalDist -= distance
        }
    }

    if horizontalDist < 0 {
        horizontalDist = horizontalDist * -1
    }

    if verticalDist < 0 {
        horizontalDist = horizontalDist * -1
    }

    return horizontalDist + verticalDist
}
