package sand_process

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"FallingSand-GO-Raylib/config"
)

func Update(g *[][]int) {
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mouseX := int(rl.GetMouseX())
		mouseY := int(rl.GetMouseY())

		// Converte posição do mouse para coordenadas do grid
		gridX := mouseX / config.GRID_SIZE
		gridY := mouseY / config.GRID_SIZE

		for dy := -config.MOUSE_RADIOUS; dy <= config.MOUSE_RADIOUS; dy++ {
			for dx := -config.MOUSE_RADIOUS; dx <= config.MOUSE_RADIOUS; dx++ {
				nx := gridX + dx
				ny := gridY + dy

				// Verifica limites
				if dx*dx+dy*dy <= config.MOUSE_RADIOUS * config.MOUSE_RADIOUS { //raio do circulo
					if nx >= 0 && ny >= 0 && nx < config.COLS && ny < config.ROWS {
						(*g)[ny][nx] = 1 //(*g) disrrefrencia po
					}
				}
			}
		}
	}
}
