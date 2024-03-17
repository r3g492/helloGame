package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

var windowWidth int32 = 1280
var windowHeight int32 = 800
var fps int32 = 60
var nodes []Node
var player Player

type Node struct {
	Images    []rl.Texture2D
	x         int32
	y         int32
	direction int32
}

func createNode(x int32, y int32, direction int32) {
	images := make([]rl.Texture2D, 4)
	for i := 0; i < 4; i++ {
		image := rl.LoadImage("picture/0" + fmt.Sprintf("%d", i+1) + ".png")
		if image.Width == 0 || image.Height == 0 {
			fmt.Println("Failed to load the image.")
		} else {
			fmt.Println("Image loaded successfully.")
		}
		texture := rl.LoadTextureFromImage(image)
		rl.UnloadImage(image)
		images[i] = texture
	}
	node := Node{Images: images, x: x, y: y, direction: direction}
	nodes = append(nodes, node)
}

type Player struct {
	Images    []rl.Texture2D
	x         int32
	y         int32
	direction int32
}

func initPlayer(x int32, y int32, direction int32) {
	images := make([]rl.Texture2D, 4)
	for i := 0; i < 4; i++ {
		image := rl.LoadImage("picture/0" + fmt.Sprintf("%d", i+1) + ".png")
		if image.Width == 0 || image.Height == 0 {
			fmt.Println("Failed to load the image.")
		} else {
			fmt.Println("Image loaded successfully.")
		}
		texture := rl.LoadTextureFromImage(image)
		rl.UnloadImage(image)
		images[i] = texture
	}
	player = Player{Images: images, x: x, y: y, direction: direction}
}

func main() {
	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	rl.InitWindow(windowWidth, windowHeight, "my little game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	createNode(100, 100, 0)
	initPlayer(200, 200, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		decideNodeDirection()
		drawNodes()

		decidePlayerDirection()
		drawPlayer()

		rl.DrawText("My little game", 10, 10, 20, rl.White)
		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}
}

func decideNodeDirection() {
	for i, _ := range nodes {
		random := rl.GetRandomValue(0, 3)
		nodes[i].direction = int32(random)

		if random == 0 {
			nodes[i].y += 1
		} else if random == 1 {
			nodes[i].x += 1
		} else if random == 2 {
			nodes[i].y -= 1
		} else if random == 3 {
			nodes[i].x -= 1
		}
	}
}

func decidePlayerDirection() {
	input := rl.GetKeyPressed()
	if input == rl.KeyS {
		player.y -= 1
		player.direction = 2
	} else if input == rl.KeyW {
		player.y += 1
		player.direction = 0
	} else if input == rl.KeyA {
		player.x -= 1
		player.direction = 3
	} else if input == rl.KeyD {
		player.x += 1
		player.direction = 1
	}
}

func drawNodes() {
	for _, n := range nodes {
		rl.DrawTexture(n.Images[n.direction], n.x, n.y, rl.White)
	}
}

func drawPlayer() {
	rl.DrawTexture(player.Images[player.direction], player.x, player.y, rl.White)
}
