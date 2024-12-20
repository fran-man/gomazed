package main

import (
	"fmt"
	"math/rand"
)

func (m *MazeData) GenerateDFS() {

	// Set first cell to visited
	coords := [2]int{0, 0}
	c := m.grid[coords]
	for {
		fmt.Printf("Currently at Cell: %v\n", coords)
		c.visited = true
		m.grid[coords] = c
		unusedNeighbors := m.unusedNeighbors(coords)
		if len(unusedNeighbors) == 0 {
			break
		}
		nextNeighborIdx := rand.Intn(len(unusedNeighbors))
		nextNeighbor := c.nbors[unusedNeighbors[nextNeighborIdx]]
		nextNeighbor.wall = false
		c.nbors[unusedNeighbors[nextNeighborIdx]] = nextNeighbor
		m.grid[coords] = c
		// remove wall from the corresponding neighbor
		m.removeWall(nextNeighbor.coords, coords)
		coords = nextNeighbor.coords
		c = m.grid[coords]
	}
	// fmt.Printf("MAZE %v", m)
}

func (m *MazeData) unusedNeighbors(coords [2]int) []int {
	c := m.grid[coords]
	unused := make([]int, 0)
	for i, n := range c.nbors {
		if n.wall && !m.grid[n.coords].visited {
			unused = append(unused, i)
		}
	}
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
