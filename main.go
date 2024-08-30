package main

import (
	"fmt"
	"os"
	"os/exec"
)

var width, height = 41, 11
var leftRocket, rightRocket = height / 2, height / 2
var ballX, ballY = height / 2, width / 2
var leftScore, rightScore = 0, 0


var ballX1, ballY1 = 0, 0
var table = make([][]string, height)

func main() {
	
	createTable()
	
	for {
		clearScreen()
		validateRockets()
		render()

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "w": leftRocket--
		case "s": leftRocket++
		case "i": rightRocket--
		case "k": rightRocket++
		}

		
	}
	
}

func createTable() {
	for i := range table {
		table[i] = make([]string, width)
	}

	for i := range table {
		for j := range table[i] {
			switch {
			case i == 0 : table[i][j] = "_"
			case i == height - 1: table[i][j] = "‾"
			case (i == leftRocket) && j == 0 || (i == rightRocket) && j == width - 1: table[i][j] = "|"
			case (i == ballX) && (j == ballY): table[i][j] = "o"
			default: table[i][j] = " "
			}
		}
	}
}

func render() {
	fmt.Printf("%19d | %d\n", leftScore, rightScore)
	for i := range table {
		for j := range table[i] {
			fmt.Print(table[i][j])
		}
		fmt.Println()
	}
	
}

func validateRockets() {
	switch {
		case leftRocket <= 0: leftRocket = 1
		case leftRocket == height - 1: leftRocket = height - 2
		case rightRocket == 0: rightRocket = 1
		case rightRocket == height - 1: rightRocket = height - 2
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
	cmd.Stdout = os.Stdout
	cmd.Run()
}