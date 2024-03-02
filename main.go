package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var time float32 = 0.0
	var objectsList []interface{}
	objectsList = append(objectsList, rl.NewRectangle(400, 200, 100, 100))

	for !rl.WindowShouldClose() {
		moveByArrow(&objectsList)
		clickToCreateRect(&objectsList)
		draw(&time, objectsList)
	}
}

func draw(time *float32, objectsList []interface{}) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	timeText := fmt.Sprintf("%02.02f seconds", *time)
	rl.DrawText("cur time: "+timeText, 650, 1, 5, rl.Green)

	for _, obj := range objectsList {
		rl.DrawRectangleRec(obj.(rl.Rectangle), rl.Red)
	}

	rl.EndDrawing()

	*time += rl.GetFrameTime()
}

func moveByArrow(objectsList *[]interface{}) {
	// get input
	if rl.IsKeyDown(rl.KeyRight) {
		for i, obj := range *objectsList {
			(*objectsList)[i] = rl.NewRectangle(obj.(rl.Rectangle).X+1, obj.(rl.Rectangle).Y, obj.(rl.Rectangle).Width, obj.(rl.Rectangle).Height)
		}
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		for i, obj := range *objectsList {
			(*objectsList)[i] = rl.NewRectangle(obj.(rl.Rectangle).X-1, obj.(rl.Rectangle).Y, obj.(rl.Rectangle).Width, obj.(rl.Rectangle).Height)
		}
	}
	if rl.IsKeyDown(rl.KeyUp) {
		for i, obj := range *objectsList {
			(*objectsList)[i] = rl.NewRectangle(obj.(rl.Rectangle).X, obj.(rl.Rectangle).Y-1, obj.(rl.Rectangle).Width, obj.(rl.Rectangle).Height)
		}
	}
	if rl.IsKeyDown(rl.KeyDown) {
		for i, obj := range *objectsList {
			(*objectsList)[i] = rl.NewRectangle(obj.(rl.Rectangle).X, obj.(rl.Rectangle).Y+1, obj.(rl.Rectangle).Width, obj.(rl.Rectangle).Height)
		}
	}

}

func clickToCreateRect(objectsList *[]interface{}) {
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		mouseXFloat32 := float32(mouseX)
		mouseYFloat32 := float32(mouseY)
		*objectsList = append(*objectsList, rl.NewRectangle(mouseXFloat32, mouseYFloat32, 1, 1))
	}
}
