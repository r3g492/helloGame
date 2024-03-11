package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

var windowWidth int32 = 1280
var windowHeight int32 = 800
var fps int32 = 60

func main() {
	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	rl.InitWindow(windowWidth, windowHeight, "my little game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	image := rl.LoadImage("picture/01.png") // Make sure the path is correct
	if image.Width == 0 || image.Height == 0 {
		fmt.Println("Failed to load the image.")
	} else {
		fmt.Println("Image loaded successfully.")
	}
	fmt.Println(image.Width, image.Height)
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawRectangle(0, 0, windowWidth, windowHeight, rl.Black)
		rl.DrawTexture(texture, 0, 0, rl.White)
		rl.DrawText("My little game", 10, 10, 20, rl.White)

		rl.EndDrawing()
	}
}
