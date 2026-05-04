package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	scenes := map[string]Scene{
		"menu": &MenuScene{},
		"game": &GameScene{},
	}

	currentScene := "menu"

	for !rl.WindowShouldClose() {
		currentScene = scenes[currentScene].Update()

		rl.BeginDrawing()
		scenes[currentScene].Draw()
		rl.EndDrawing()
	}
}
