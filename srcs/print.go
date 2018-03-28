package main

import (
	"fmt"
	"os"
)

func printError(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func printState(e Env, state State) {
	fmt.Println("State Index : ", state.index)
	fmt.Println("State Parent : ", state.parent)
	fmt.Println("State Priority : ", state.priority)
	fmt.Println("State Iteration : ", state.iteration)
	fmt.Println("State Heuristic : ", state.heuristic)
	if state.board != nil {
		fmt.Println("State Board : ")
		for i := 0; i < e.boardSize; i++ {
			for j := 0; j < e.boardSize; j++ {
				fmt.Printf("%d\t", state.board[i*e.boardSize+j])
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("State Board : ", state.board)
	}
}
