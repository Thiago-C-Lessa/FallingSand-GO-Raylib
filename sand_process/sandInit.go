package sand_process

import (
	"FallingSand-GO-Raylib/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// initialize the audio and generete the grid tha hold the cells
func Intit() [][]int {
	rl.InitAudioDevice()
	return CreateGrid(0)
}

func Deinit() {
	rl.CloseAudioDevice()
}

func CreateGrid(defaultValue int) [][]int {

	grid := make([][]int, config.ROWS)
	for y := 0; y < config.ROWS; y++ {
		grid[y] = make([]int, config.COLS)
		for x := 0; x < config.COLS; x++ {
			grid[y][x] = defaultValue
		}
	}

	return grid
}
