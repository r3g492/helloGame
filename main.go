package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var windowWidth int32 = 1280
var windowHeight int32 = 800
var fps int32 = 60

func main() {
	rl.InitWindow(windowWidth, windowHeight, "my little game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	cubePosition := rl.NewVector3(0.0, 0.0, 0.0)

	rl.DisableCursor()

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		if rl.IsKeyDown('Z') {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)
		rl.DrawGrid(10, 1.0)

		rl.DrawCube(rl.NewVector3(-2.0, 0.0, 0.0), 1.0, 2.0, 1.0, rl.Green)
		rl.DrawCubeWires(rl.NewVector3(-2.0, 0.0, 0.0), 1.0, 2.0, 1.0, rl.Lime)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)

		rl.EndDrawing()
	}
}

func draw(time *float32, objectsList *[]interface{}) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	timeText := fmt.Sprintf("%02.02f seconds", *time)

	var windowWidthFloat = float32(windowWidth)
	var windowWidth90 = windowWidthFloat * 0.9
	var rightUpper = rl.NewVector2(windowWidth90, 1)
	rl.DrawText("cur time: "+timeText, int32(rightUpper.X), int32(rightUpper.Y), 5, rl.Green)

	for _, obj := range *objectsList {
		rl.DrawCircle(int32(obj.(rl.Rectangle).X), int32(obj.(rl.Rectangle).Y), 5, rl.Blue)
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

func clickToDelete(objectsList *[]interface{}) {
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		mouseXFloat32 := float32(mouseX)
		mouseYFloat32 := float32(mouseY)
		for i, obj := range *objectsList {
			if rl.CheckCollisionPointRec(rl.NewVector2(mouseXFloat32, mouseYFloat32), obj.(rl.Rectangle)) {
				*objectsList = append((*objectsList)[:i], (*objectsList)[i+1:]...)
				break
			}
		}
	}
}

func clickToCreate(objectsList *[]interface{}) {
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		mouseXFloat32 := float32(mouseX)
		mouseYFloat32 := float32(mouseY)
		*objectsList = append(*objectsList, rl.NewRectangle(mouseXFloat32, mouseYFloat32, 10, 10))
	}
}
