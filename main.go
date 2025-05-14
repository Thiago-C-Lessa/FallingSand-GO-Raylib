package main

import (
	//"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
	GRID_SIZE int = 2
	COLS          = SCREEN_WIDTH / GRID_SIZE
	ROWS          = SCREEN_HEIGHT / GRID_SIZE
)

func IntitGame() [][]int {
	rl.InitAudioDevice()
	return CreateGrid(0)
}

func CloseGame() {
	rl.CloseAudioDevice()
}

func GameUpdate() {

}

func GameRender(g [][]int) {
	ANTIQUE_WHITE := rl.Color{R: 250, G: 235, B: 215, A: 255}

	for y := 0; y < ROWS; y++ {
		for x := 0; x < COLS; x++ {
			var color rl.Color

			switch g[y][x] {
			case 0:
				color = ANTIQUE_WHITE
			default:
				color = rl.Black
			}

			rl.DrawRectangle(int32(x*GRID_SIZE), int32(y*GRID_SIZE), int32(GRID_SIZE), int32(GRID_SIZE), color)
			//rl.DrawRectangleLines(int32(x*GRID_SIZE), int32(y*GRID_SIZE), int32(GRID_SIZE), int32(GRID_SIZE), rl.Black)
		}
	}
}

func CreateGrid(defaultValue int) [][]int {

	grid := make([][]int, ROWS)
	for y := 0; y < ROWS; y++ {
		grid[y] = make([]int, COLS)
		for x := 0; x < COLS; x++ {
			grid[y][x] = defaultValue
		}
	}

	return grid
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Falling Sand")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	grid := IntitGame() //y a esquerda e x a direita!!

	grid[34][34] = 1

	for !rl.WindowShouldClose() {

		GameUpdate()

		rl.BeginDrawing()

		GameRender(grid)

		rl.EndDrawing()
	}

	CloseGame()
}
