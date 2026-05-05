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

	currentKey := "menu"
	activeScene := scenes[currentKey]
	activeScene.Enter()

	for !rl.WindowShouldClose() {
		nextKey := activeScene.Update()

		if nextKey != currentKey {
			activeScene.Exit()
			activeScene = scenes[nextKey]
			activeScene.Enter()
			currentKey = nextKey
		}

		rl.BeginDrawing()
		activeScene.Draw()
		rl.EndDrawing()
	}
}
