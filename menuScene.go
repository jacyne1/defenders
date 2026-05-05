package main

import rl "github.com/gen2brain/raylib-go/raylib"

type MenuScene struct{}

func (s *MenuScene) Update() string {
	if rl.IsKeyPressed(rl.KeyEnter) {
		return "game"
	}
	return "menu"
}

func (s *MenuScene) Draw() {
	rl.ClearBackground(rl.Black)
	rl.DrawText("Defender: Main Menu", 100, 100, 20, rl.White)
	rl.DrawText("Press Enter to Start", 100, 150, 10, rl.Gray)
}

func (s *MenuScene) Enter() {

}

func (s *MenuScene) Exit() {}
