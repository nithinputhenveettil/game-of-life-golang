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

func drawScreen(gol *golGame) {
	rl.ClearBackground(rl.White)

	for i := 0; i < gol.size[0]; i = i + 25 {
		rl.DrawLine(int32(i), int32(50), int32(i), int32(gol.size[1]-100), rl.Black)
	}
	for i := 50; i <= gol.size[1]-100; i = i + 25 {
		rl.DrawLine(int32(0), int32(i), int32(gol.size[0]), int32(i), rl.Black)
	}

	rl.DrawText(gol.name, int32(gol.size[0]/2-60), int32(18), 25, rl.Black)

	instr := "Mouse click to select initial configuration.Press 's' to start/stop simulation or 'r' to reset the screen."
	rl.DrawText(instr, int32(50), int32(gol.size[1]-85), 19, rl.Black)

	rl.DrawText("Generations : "+strconv.Itoa(gol.generations), int32(50), int32(gol.size[1]-40), 20, rl.Black)
	rl.DrawText("Alive : "+strconv.Itoa(gol.alive), int32(gol.size[0]/2), int32(gol.size[1]-40), 20, rl.Black)
	rl.DrawText("Dead : "+strconv.Itoa(gol.dead), int32(gol.size[0]-140), int32(gol.size[1]-40), 20, rl.Black)
}

func main() {
	gol := new(golGame)
	reset(gol)

	rl.InitWindow(int32(gol.size[0]), int32(gol.size[1]), gol.name)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		drawScreen(gol)
		rl.EndDrawing()
	}
	rl.CloseWindow()

}
