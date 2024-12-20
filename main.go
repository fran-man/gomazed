package main

import (
	"bytes"
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
	gridSize = 5
)

var (
	mazeTemplateImage *e.Image
	mazeImages        []*e.Image
)

type Game struct {
	frame int32
}

func (g *Game) Update() error {
	g.frame++
	return nil
}

func (g *Game) Draw(screen *e.Image) {
	topLeft := mazeImages[0]
	topRow := mazeImages[3]
	leftCol := mazeImages[1]
	other := mazeImages[4]
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			opt := &e.DrawImageOptions{}
			opt.GeoM.Translate(10.0, 10.0)
			if x == 0 && y == 0 {
				screen.DrawImage(topLeft, opt)
			} else if x == 0 {
				opt.GeoM.Translate(0, float64(y*cellSize))
				screen.DrawImage(leftCol, opt)
			} else if y == 0 {
				opt.GeoM.Translate(float64(x*cellSize), 0)
				screen.DrawImage(topRow, opt)
			} else {
				opt.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
				screen.DrawImage(other, opt)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	b, err := os.ReadFile("Maze2.png")
	if err != nil {
		log.Fatal("Could not load the file!")
	}
	img, _, err2 := image.Decode(bytes.NewReader(b))
	if err2 != nil {
		log.Fatalf("Could not decode the file: %v", err2)
	}
	mazeTemplateImage = e.NewImageFromImage(img)
	mazeImages = populateMazeImagesSlice(mazeTemplateImage)

	e.SetWindowSize(screenW, screenH)
	e.SetWindowTitle("GOMAZED")

	m := initMaze(gridSize, gridSize)
	m.GenerateDFS()

	if err := e.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func populateMazeImagesSlice(*e.Image) []*e.Image {
	imgs := make([]*e.Image, 7)
	for i := 0; i < 7; i++ {
		si := mazeTemplateImage.SubImage(image.Rect(cellSize*i, 0, cellSize+(cellSize*i), cellSize))
		imgs[i] = si.(*e.Image)
	}
	return imgs
}
