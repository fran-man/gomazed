package main

import (
	"fmt"
	"math/rand"
)

func (m *MazeData) GenerateDFS() {

	// Set first cell to visited
	c := m.grid[[2]int{0, 0}]
	c.visited = true
	m.grid[[2]int{0, 0}] = c
	unusedNeighbors := c.unusedNeighbors()
	nextNeighborIdx := rand.Intn(len(unusedNeighbors))
	nextNeighbor := c.nbors[nextNeighborIdx]
	nextNeighbor.wall = false
	c.nbors[nextNeighborIdx] = nextNeighbor
	m.grid[[2]int{0, 0}] = c
	// remove wall from the corresponding neighbor
	m.removeWall(nextNeighbor.coords, [2]int{0, 0})
	fmt.Printf("MAZE %v", m)
}

func (c Cell) unusedNeighbors() []int {
	unused := make([]int, 0)
	for i, n := range c.nbors {
		if n.wall {
			unused = append(unused, i)
		}
	}
	fmt.Printf("Unused Neighbors: %v", unused)
	return unused
}

func (m *MazeData) removeWall(cellIdx [2]int, wallToRemove [2]int) {
	c := m.grid[cellIdx]
	c.visited = true
	for i, n := range c.nbors {
		if n.coords == wallToRemove {
			n.wall = false
			c.nbors[i] = n
			m.grid[cellIdx] = c
		}
	}
}
