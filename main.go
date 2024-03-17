package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

var windowWidth int32 = 1280
var windowHeight int32 = 800
var fps int32 = 60
var grids []Grid
var chosenGridPlayerOne *Grid
var chosenGridPlayerTwo *Grid
var turn int32 = 1

type Grid struct {
	Width int32
	x     int32
	y     int32
	Unit  *Unit
}

func initGrids() {
	for i := int32(200); i < 1100; i += 100 {
		for j := int32(200); j < 700; j += 100 {
			grid := Grid{Width: 100, x: i, y: j}
			grids = append(grids, grid)
		}
	}
}

type Unit struct {
	Image rl.Texture2D
	side  int32
}

func initUnit() {
	image := rl.LoadImage("picture/01.png")
	texture := rl.LoadTextureFromImage(image)
	unit := Unit{Image: texture, side: 1}
	grids[0].Unit = &unit
}

func main() {
	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	rl.InitWindow(windowWidth, windowHeight, "my little game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	initGrids()
	initUnit()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		drawGrid()
		drawChosenGrid()
		detectClickOnGrid()

		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}

func drawGrid() {
	for _, grid := range grids {
		rl.DrawRectangleLines(grid.x, grid.y, grid.Width, grid.Width, rl.White)
		if grid.Unit != nil {
			rl.DrawTexture(grid.Unit.Image, grid.x, grid.y, rl.White)
		}
	}
}

func drawChosenGrid() {
	if chosenGridPlayerOne != nil {
		rl.DrawRectangleLines(chosenGridPlayerOne.x, chosenGridPlayerOne.y, chosenGridPlayerOne.Width, chosenGridPlayerOne.Width, rl.Red)
	}
	if chosenGridPlayerTwo != nil {
		rl.DrawRectangleLines(chosenGridPlayerTwo.x, chosenGridPlayerTwo.y, chosenGridPlayerTwo.Width, chosenGridPlayerTwo.Width, rl.Blue)
	}
}

func detectClickOnGrid() {
	if turn != 1 {
		return
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePos := rl.GetMousePosition()
		for i := range grids {
			grid := &grids[i]
			if mousePos.X > float32(grid.x) && mousePos.X < float32(grid.x+grid.Width) && mousePos.Y > float32(grid.y) && mousePos.Y < float32(grid.y+grid.Width) {
				chosenGridPlayerOne = grid
				break
			}
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		mousePos := rl.GetMousePosition()
		for i := range grids {
			grid := &grids[i]
			if mousePos.X > float32(grid.x) &&
				mousePos.X < float32(grid.x+grid.Width) &&
				mousePos.Y > float32(grid.y) &&
				mousePos.Y < float32(grid.y+grid.Width) {

				if chosenGridPlayerOne != nil && chosenGridPlayerOne.Unit != nil {
					grid.Unit = chosenGridPlayerOne.Unit
					chosenGridPlayerOne.Unit = nil
					chosenGridPlayerOne = grid
					break
				}
			}
		}
	}
}
