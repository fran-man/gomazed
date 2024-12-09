package main

type Cell struct {
	visited bool
	nbors   []Neighbor
}

// coords: coordinates of the neighbor
// wall: Is the neighbor connected with a wall?
type Neighbor struct {
	coords [2]int
	wall   bool
}

// grid is a slice of cells. One dimensional slice, index is coordinates like [2,1]
type MazeData struct {
	width  int16
	height int16
	grid   map[[2]int]Cell
}

func initMaze(w int, h int) *MazeData {
	g := make(map[[2]int]Cell, w*h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			nc := neighborCount(x, y, w, h)
			idx := 0
			neighbors := make([]Neighbor, nc)
			if x > 0 {
				neighbors[idx] = Neighbor{
					coords: [2]int{x - 1, y},
					wall:   true,
				}
				idx++
			}
			if x < w-1 {
				neighbors[idx] = Neighbor{
					coords: [2]int{x + 1, y},
					wall:   true,
				}
				idx++
			}
			if y > 0 {
				neighbors[idx] = Neighbor{
					coords: [2]int{x, y - 1},
					wall:   true,
				}
				idx++
			}
			if y < h-1 {
				neighbors[idx] = Neighbor{
					coords: [2]int{x, y + 1},
					wall:   true,
				}
				idx++
			}
			g[[2]int{x, y}] = Cell{
				visited: false,
				nbors:   neighbors,
			}
		}
	}

	m := MazeData{
		width:  int16(w),
		height: int16(h),
		grid:   g,
	}
	return &m
}

func neighborCount(x, y, w, h int) int {
	c := 2
	if x > 0 && x < w-1 {
		c++
	}
	if y > 0 && y < h-1 {
		c++
	}
	return c
}
