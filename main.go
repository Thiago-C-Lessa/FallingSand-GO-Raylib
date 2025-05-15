package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	
	"FallingSand-GO-Raylib/sand_process"
	"FallingSand-GO-Raylib/config"
)


func main() {
	rl.InitWindow(config.SCREEN_WIDTH, config.SCREEN_HEIGHT, "Falling Sand")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	grid := sand_process.Intit() //y a esquerda e x a direita!!

	for !rl.WindowShouldClose() {
		

		sand_process.Update(&grid)

		rl.BeginDrawing()

		sand_process.Render(grid)

		rl.EndDrawing()
	}

	sand_process.Deinit()
}
