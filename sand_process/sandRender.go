package sand_process

import (
	"FallingSand-GO-Raylib/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render(g [][]int) {
	ANTIQUE_WHITE := rl.Color{R: 250, G: 235, B: 215, A: 255}

	for y := 0; y < config.ROWS; y++ {
		for x := 0; x < config.COLS; x++ {
			var color rl.Color

			switch g[y][x] {
			case 0:
				color = ANTIQUE_WHITE
			default:
				color = rl.Black
			}

			rl.DrawRectangle(int32(x*config.GRID_SIZE), int32(y*config.GRID_SIZE), int32(config.GRID_SIZE), int32(config.GRID_SIZE), color)
			//rl.DrawRectangleLines(int32(x*config.GRID_SIZE), int32(y*config.GRID_SIZE), int32(config.GRID_SIZE), int32(config.GRID_SIZE), rl.Black)
		}
	}
}
