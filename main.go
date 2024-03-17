package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	snakeLength  = 256
	squareSize   = 31
	screenWidth  = 1200
	screenHeight = 800
)

type Snake struct {
	Position rl.Vector2
	Size     rl.Vector2
	Speed    rl.Vector2
	Color    rl.Color
}

type Score struct {
	Position rl.Vector2
	Size     rl.Vector2
	Active   bool
	Color    rl.Color
}

var (
	framesCounter = 0
	gameOver      = false
	pause         = false
	score         = Score{}
	snake         [snakeLength]Snake
	snakePosition [snakeLength]rl.Vector2
	allowMove     = false
	offset        = rl.Vector2{}
	counterTail   = 1
)

func InitGame() {
	framesCounter = 0
	gameOver = false
	pause = false

	counterTail = 1
	allowMove = false

	offset.X = float32(screenWidth % squareSize)
	offset.Y = float32(screenHeight % squareSize)

	for i := 0; i < snakeLength; i++ {
		snake[i].Position = rl.NewVector2(offset.X/2, offset.Y/2)
		snake[i].Size = rl.NewVector2(squareSize, squareSize)
		snake[i].Speed = rl.NewVector2(squareSize, 0)
		if i == 0 {
			snake[i].Color = rl.DarkBlue
		} else {
			snake[i].Color = rl.Blue
		}
	}

	for i := 0; i < snakeLength; i++ {
		snakePosition[i] = rl.NewVector2(0.0, 0.0)
	}

	score.Size = rl.NewVector2(squareSize, squareSize)
	score.Color = rl.SkyBlue
	score.Active = false
}

func UpdateGame() {
	if !gameOver {
		if rl.IsKeyPressed(rl.KeyP) {
			pause = !pause
		}

		if !pause {
			if rl.IsKeyPressed(rl.KeyRight) && (snake[0].Speed.X == 0) && allowMove {
				snake[0].Speed = rl.NewVector2(squareSize, 0)
				allowMove = false
			}
			if rl.IsKeyPressed(rl.KeyLeft) && (snake[0].Speed.X == 0) && allowMove {
				snake[0].Speed = rl.NewVector2(-squareSize, 0)
				allowMove = false
			}
			if rl.IsKeyPressed(rl.KeyUp) && (snake[0].Speed.Y == 0) && allowMove {
				snake[0].Speed = rl.NewVector2(0, -squareSize)
				allowMove = false
			}
			if rl.IsKeyPressed(rl.KeyDown) && (snake[0].Speed.Y == 0) && allowMove {
				snake[0].Speed = rl.NewVector2(0, squareSize)
				allowMove = false
			}

			for i := 0; i < counterTail; i++ {
				snakePosition[i] = snake[i].Position
			}

			if (framesCounter % 5) == 0 {
				for i := 0; i < counterTail; i++ {
					if i == 0 {
						snake[0].Position.X += snake[0].Speed.X
						snake[0].Position.Y += snake[0].Speed.Y
						allowMove = true
					} else {
						snake[i].Position = snakePosition[i-1]
					}
				}
			}

			if (snake[0].Position.X > (screenWidth - offset.X)) ||
				(snake[0].Position.Y > (screenHeight - offset.Y)) ||
				(snake[0].Position.X < 0) || (snake[0].Position.Y < 0) {
				gameOver = true
			}

			for i := 1; i < counterTail; i++ {
				if (snake[0].Position.X == snake[i].Position.X) && (snake[0].Position.Y == snake[i].Position.Y) {
					gameOver = true
				}
			}

			if !score.Active {
				score.Active = true
				score.Position = rl.NewVector2(float32(rl.GetRandomValue(0, (screenWidth/squareSize)-1)*squareSize)+offset.X/2, float32(rl.GetRandomValue(0, (screenHeight/squareSize)-1)*squareSize)+offset.Y/2)

				for i := 0; i < counterTail; i++ {
					for (score.Position.X == snake[i].Position.X) && (score.Position.Y == snake[i].Position.Y) {
						score.Position = rl.NewVector2(float32(rl.GetRandomValue(0, (screenWidth/squareSize)-1)*squareSize)+offset.X/2, float32(rl.GetRandomValue(0, (screenHeight/squareSize)-1)*squareSize)+offset.Y/2)
						i = 0
					}
				}
			}

			if (snake[0].Position.X < (score.Position.X+score.Size.X) && (snake[0].Position.X+snake[0].Size.X) > score.Position.X) &&
				(snake[0].Position.Y < (score.Position.Y+score.Size.Y) && (snake[0].Position.Y+snake[0].Size.Y) > score.Position.Y) {
				snake[counterTail].Position = snakePosition[counterTail-1]
				counterTail++
				score.Active = false
			}

			framesCounter++
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			InitGame()
			gameOver = false
		}
	}
}

func DrawGame() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	if !gameOver {
		for i := 0; i < counterTail; i++ {
			rl.DrawRectangleV(snake[i].Position, snake[i].Size, snake[i].Color)
		}

		if score.Active {
			rl.DrawRectangleV(score.Position, score.Size, score.Color)
		}

		if pause {
			rl.DrawText("PAUSED", screenWidth/2-rl.MeasureText("GAME PAUSED", 40)/2, screenHeight/2-40, 40, rl.Gray)
		}
	} else {
		rl.DrawText("PRESS [ENTER] TO PLAY", screenWidth/2-rl.MeasureText("PRESS [ENTER] TO PLAY", 20)/2, screenHeight/2-50, 20, rl.Gray)
	}

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "snake game in raylib-go")

	InitGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		UpdateGame()
		DrawGame()
	}

	rl.CloseWindow()
}
