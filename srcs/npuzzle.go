package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type State struct {
	board    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index  int // The index of the item in the heap.
	parent *State
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want the lowest priority so we use smaller than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	fmt.Println(x)
	state := x.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	state.index = -1 // for safety
	*pq = old[0 : n-1]
	return state
}

func getNewState(index, indexToMove int, currentState *State, chanState chan<- *State) {
	// calculer nouveaux states
}

func play(e Env) {
	getFinalState(&e)
	indexToMove := getIndexToMove(e.initState)
	fmt.Println(indexToMove)
	openList := make(PriorityQueue, 1)
	openList[0] = &State{
		board:    e.initState,
		priority: -1,
		parent:   nil,
	}
	heap.Init(&openList)
	// new := &State{
	// 	board:    e.finalState,
	// 	priority: 0,
	// 	parent:   nil,
	// }
	// heap.Push(&openList, new)
	// heuristic := heuristic(e, new)
	// fmt.Println(heuristic)

	chanState := make(chan *State)
	for i := 0; i < 4; i++ {
		go getNewState(i, indexToMove, openList[0], chanState)
	}
	var closedList []*State
	for i := 0; i < 4; i++ {
		ngbState := <-chanState
		//check if the neighbour is not in the closed list
		if !findInClosedList(ngbState, closedList) {
			//check if the neighbour is in the open list
			index := findInOpenList(ngbState, openList)
			if index != -1 {
				//modify priority if it is higher (== worse) in the open list
				if openList[i].priority > ngbState.priority {
					openList[i].priority = ngbState.priority
				}
			} else {
				//push neighbour to open list
				heap.Push(&openList, ngbState)
			}
		}
	}
	//if heapqueue is empty there is no solution => the alhorithm stops here
	if len(openList) != 0 {
		//sort the open list
		sort.Sort(&openList)
		bestState := openList[0]
	}
	// all states were reviewed
}
