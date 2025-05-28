package sand_process

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"FallingSand-GO-Raylib/config"
)

func GetInput(g *[][]int) {
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
				if dx*dx+dy*dy <= config.MOUSE_RADIOUS*config.MOUSE_RADIOUS { //raio do circulo
					if nx >= 0 && ny >= 0 && nx < config.COLS && ny < config.ROWS {
						(*g)[ny][nx] = 1 //(*g) disrrefrencia po
					}
				}
			}
		}
	}
}

func Update(g *[][]int) {
	for i, _ := range *g {
		for j, _ := range (*g)[i] {
			if (*g)[i][j] == 1 {
				moved := false

				if i+1 < config.ROWS && (*g)[i+1][j] == 0 { //verifica bloco abaixo
					(*g)[i+1][j] = 1
					(*g)[i][j] = 0
					moved = true
				}

				if !moved && i+1 < config.ROWS && j > 0 && (*g)[i+1][j-1] == 0 { //verifica a esquerda
					(*g)[i+1][j-1] = 1
					(*g)[i][j] = 0
				}

				if !moved && i+1 < config.ROWS && j+1 < config.COLS && (*g)[i+1][j+1] == 0 { //verifica a direita
					(*g)[i+1][j+1] = 1
					(*g)[i][j] = 0

				}

			}
		}
	}
}
