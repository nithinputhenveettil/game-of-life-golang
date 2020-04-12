package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	name   = "Game of Life!!"
	width  = 1000
	height = 650
)

type golGame struct {
	name        string
	generations int
	alive       int
	dead        int
	size        [2]int
	simulation  bool
	game        [20][40]bool
}

func reset(gol *golGame) {
	gol.name = name
	gol.generations = 0
	gol.alive = 0
	gol.dead = 0
	gol.size = [2]int{width, height}
	gol.simulation = false

	for i, a := range gol.game {
		for j := range a {
			gol.game[i][j] = false
		}
	}
}

func litsenKeyboardEvents(gol *golGame) {
	if rl.IsKeyDown(83) {
		if gol.simulation {
			gol.simulation = false
		} else {
			gol.simulation = true
		}
	}
	if rl.IsKeyDown(82) {
		reset(gol)
	}
}

func litsenMouseClick(gol *golGame) {
	if rl.IsMouseButtonPressed(0) {
		points := rl.GetMousePosition()
		if points.Y > float32(gol.size[1]-100) || points.Y < float32(50) {
			return
		}
		i := -1
		j := -1
		x := 0
		y := 50
		for x < gol.size[0] {
			if points.X >= float32(x) && points.X <= float32(x+25) {
				j = x / 25
				break
			}
			x += 25
		}
		for y < gol.size[1]-100 {
			if points.Y >= float32(y) && points.Y <= float32(y+25) {
				i = y/25 - 2
				break
			}
			y += 25
		}
		if i == -1 || j == -1 {
			return
		}
		if gol.game[i][j] {
			gol.game[i][j] = false
		} else {
			gol.game[i][j] = true
		}
	}
}

func drawScreen(gol *golGame) {
	rl.ClearBackground(rl.White)

	rl.DrawText(gol.name, int32(gol.size[0]/2-60), int32(18), 25, rl.Black)

	instr := "Mouse click to select initial configuration. Press 's' to start/stop simulation or 'r' to reset the screen."
	rl.DrawText(instr, int32(50), int32(gol.size[1]-85), 19, rl.Black)

	rl.DrawText("Generations : "+strconv.Itoa(gol.generations), int32(50), int32(gol.size[1]-40), 20, rl.Black)
	rl.DrawText("Alive : "+strconv.Itoa(gol.alive), int32(gol.size[0]/2), int32(gol.size[1]-40), 20, rl.Black)
	rl.DrawText("Dead : "+strconv.Itoa(gol.dead), int32(gol.size[0]-140), int32(gol.size[1]-40), 20, rl.Black)

	for i, a := range gol.game {
		for j := range a {
			if gol.game[i][j] {
				var x, y int
				x = j*25 + 1
				y = i*25 + 50 + 1
				rl.DrawRectangle(int32(x), int32(y), 22, 22, rl.Black)
			}
		}
	}

}

func simulate(gol *golGame) {
	cGame := gol.game
	gol.alive = 0
	gol.dead = 0
	for i, a := range gol.game {
		for j := range a {
			var neighbours = 0
			var k, l, m, n int
			if i == 0 {
				k = 19
			} else {
				k = i - 1
			}
			if i == 19 {
				l = 0
			} else {
				l = i + 1
			}
			if j == 0 {
				m = 39
			} else {
				m = j - 1
			}
			if j == 39 {
				n = 0
			} else {
				n = j + 1
			}

			if gol.game[k][m] {
				neighbours++
			}
			if gol.game[k][j] {
				neighbours++
			}
			if gol.game[k][n] {
				neighbours++
			}
			if gol.game[i][m] {
				neighbours++
			}
			if gol.game[i][n] {
				neighbours++
			}
			if gol.game[l][m] {
				neighbours++
			}
			if gol.game[l][j] {
				neighbours++
			}
			if gol.game[l][n] {
				neighbours++
			}

			if gol.game[i][j] {
				if neighbours < 2 || neighbours > 3 {
					cGame[i][j] = false
					gol.dead++
				} else {
					cGame[i][j] = true
					gol.alive++
				}
			} else {
				if neighbours == 3 {
					cGame[i][j] = true
					gol.alive++
				} else {
					cGame[i][j] = false
					gol.dead++
				}
			}
		}
	}
	gol.game = cGame
	gol.generations++
}

func main() {
	gol := new(golGame)
	reset(gol)

	rl.InitWindow(int32(gol.size[0]), int32(gol.size[1]), gol.name)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		litsenMouseClick(gol)
		litsenKeyboardEvents(gol)
		if gol.simulation {
			simulate(gol)
		}
		rl.BeginDrawing()
		drawScreen(gol)
		rl.EndDrawing()
	}
	rl.CloseWindow()

}
