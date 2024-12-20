package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	e "github.com/hajimehoshi/ebiten/v2"
)

const (
	screenW  = 220
	screenH  = 220
	cellSize = 20
	gridSize = 10
)

var (
	mazeTemplateImage *e.Image
	mazeImages        []*e.Image
	mazeData          *MazeData
)

type Game struct {
	frame int32
}

func (g *Game) Update() error {
	g.frame++
	return nil
}

func (g *Game) Draw(screen *e.Image) {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			opt := &e.DrawImageOptions{}
			opt.GeoM.Translate(10.0, 10.0)
			if x == 0 && y == 0 {
				drawTopLeftCell(screen, opt, x, y)
			} else if x == 0 {
				opt.GeoM.Translate(0, float64(y*cellSize))
				drawLeftColumnCell(screen, opt, x, y)
			} else if y == 0 {
				opt.GeoM.Translate(float64(x*cellSize), 0)
				drawTopRowCell(screen, opt, x, y)
			} else {
				opt.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
				drawOtherCell(screen, opt, x, y)
			}
		}
	}
}

func drawTopLeftCell(screen *e.Image, o *e.DrawImageOptions, x int, y int) {
	c := mazeData.grid[[2]int{x, y}]
	var nborR, nborB Neighbor
	for _, n := range c.nbors {
		if n.coords == [2]int{x + 1, y} {
			nborR = n
		} else if n.coords == [2]int{x, y + 1} {
			nborB = n
		}
	}

	if nborB.wall && nborR.wall {
		screen.DrawImage(mazeImages[0], o)
	} else if nborB.wall {
		screen.DrawImage(mazeImages[2], o)
	} else if nborR.wall {
		screen.DrawImage(mazeImages[3], o)
	} else {
		screen.DrawImage(mazeImages[6], o)
	}
}

func drawTopRowCell(screen *e.Image, o *e.DrawImageOptions, x int, y int) {
	c := mazeData.grid[[2]int{x, y}]
	var nborR, nborB Neighbor
	for _, n := range c.nbors {
		if n.coords == [2]int{x + 1, y} {
			nborR = n
		} else if n.coords == [2]int{x, y + 1} {
			nborB = n
		}
	}

	hasWallOnRight := nborR.wall || x == gridSize-1
	hasWallOnBottom := nborB.wall || y == gridSize-1

	if hasWallOnBottom && hasWallOnRight {
		screen.DrawImage(mazeImages[4], o)
	} else if hasWallOnRight {
		screen.DrawImage(mazeImages[7], o)
	} else if hasWallOnBottom {
		screen.DrawImage(mazeImages[9], o)
	} else {
		screen.DrawImage(mazeImages[11], o)
	}
}

func drawLeftColumnCell(screen *e.Image, o *e.DrawImageOptions, x int, y int) {
	c := mazeData.grid[[2]int{x, y}]
	var nborR, nborB Neighbor
	for _, n := range c.nbors {
		if n.coords == [2]int{x + 1, y} {
			nborR = n
		} else if n.coords == [2]int{x, y + 1} {
			nborB = n
		}
	}

	hasWallOnRight := nborR.wall || x == gridSize-1
	hasWallOnBottom := nborB.wall || y == gridSize-1

	if hasWallOnBottom && hasWallOnRight {
		screen.DrawImage(mazeImages[1], o)
	} else if hasWallOnRight {
		screen.DrawImage(mazeImages[10], o)
	} else if hasWallOnBottom {
		screen.DrawImage(mazeImages[5], o)
	} else {
		screen.DrawImage(mazeImages[14], o)
	}
}

func drawOtherCell(screen *e.Image, o *e.DrawImageOptions, x int, y int) {
	c := mazeData.grid[[2]int{x, y}]
	var nborR, nborB Neighbor
	for _, n := range c.nbors {
		if n.coords == [2]int{x + 1, y} {
			nborR = n
		} else if n.coords == [2]int{x, y + 1} {
			nborB = n
		}
	}

	hasWallOnRight := nborR.wall || x == gridSize-1
	hasWallOnBottom := nborB.wall || y == gridSize-1

	if hasWallOnBottom && hasWallOnRight {
		screen.DrawImage(mazeImages[8], o)
	} else if hasWallOnRight {
		screen.DrawImage(mazeImages[12], o)
	} else if hasWallOnBottom {
		screen.DrawImage(mazeImages[13], o)
	} else {
		screen.DrawImage(mazeImages[15], o)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	var n Neighbor
	fmt.Printf("%v", n)
	b, err := os.ReadFile("Maze_all.png")
	if err != nil {
		log.Fatal("Could not load the file!")
	}
	img, _, err2 := image.Decode(bytes.NewReader(b))
	if err2 != nil {
		log.Fatalf("Could not decode the file: %v", err2)
	}
	mazeTemplateImage = e.NewImageFromImage(img)
	mazeImages = populateMazeImagesSlice(mazeTemplateImage)

	mazeData = initMaze(gridSize, gridSize)
	mazeData.GenerateDFS()

	e.SetWindowSize(screenW, screenH)
	e.SetWindowTitle("GOMAZED")

	if err := e.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func populateMazeImagesSlice(*e.Image) []*e.Image {
	imgs := make([]*e.Image, 16)
	for i := 0; i < 16; i++ {
		si := mazeTemplateImage.SubImage(image.Rect(cellSize*i, 0, cellSize+(cellSize*i), cellSize))
		imgs[i] = si.(*e.Image)
	}
	return imgs
}
