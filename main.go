package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Env struct {
	initState      []int
	finalState     []int
	file           string
	boardSize      int
	heuristic      int
	timeComplexity int
	sizeComplexity int
	moves          int
}

func main() {
	var arg = os.Args[1:]
	if len(arg) != 1 {
		printError("Wrong number of arguments : need 1")
	}
	e := Env{}
	file, err := ioutil.ReadFile(arg[0])
	if err != nil {
		msg := arg[0] + " is not a valid file"
		printError(msg)
	}
	e.file = string(file)
	tab, size := parseFile(string(e.file))
	e.initState = tab
	e.boardSize = size
	e.heuristic = 2
	getFinalState(&e)
	if sameArrays(e.initState, e.finalState) {
		fmt.Println("Puzzle already solved")
		return
	}
	checkSolvability(e)
	play(&e)
	fmt.Println("States selected in the openList: ", e.timeComplexity)
	fmt.Println("Maximum number of states in memory: ", e.sizeComplexity)
	fmt.Println("Number of moves: ", e.moves)
}
