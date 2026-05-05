package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Standard downward fall
func FallNormal(e *Entity, dt float32) {
	e.Position.Y += e.Velocity.Y * dt
}

// Side-to-side swaying
func ZigZag(e *Entity, dt float32) {
	// Use math.Sin to create a wave based on time
	offset := float64(rl.GetRandomValue(0, 1000))
	width := 2.0 + float32(rl.GetRandomValue(1, 3))
	e.Position.X += float32(math.Sin(float64(rl.GetTime()*3.0)+offset)) * width
}

func CreateZigZag() func(*Entity, float32) {
	offset := float64(rl.GetRandomValue(0, 1000))
	width := 2.0 + float32(rl.GetRandomValue(1, 3))

	return func(e *Entity, dt float32) {
		e.Position.X += float32(math.Sin(float64(rl.GetTime())*2.0+offset)) * width
	}
}

// Drifts toward the center of the screen
func SeekCenter(e *Entity, dt float32) {
	if e.Position.X < 400 {
		e.Position.X += 20 * dt
	} else {
		e.Position.X -= 20 * dt
	}
}

func RotateTowardMouse(e *Entity, dt float32) {
	mousePos := rl.GetMousePosition()

	// Calculate the difference between mouse and entity
	dy := mousePos.Y - e.Position.Y
	dx := mousePos.X - e.Position.X

	// Atan2 returns the angle in radians
	angleRad := math.Atan2(float64(dy), float64(dx))

	// Convert to degrees for Raylib
	// We add 90 because most sprites point "Up" (0 degrees in Raylib is "Right")
	e.Rotation = float32(angleRad * (180 / math.Pi))
}

func MoveForward(e *Entity, dt float32) {
	e.Position.X += e.Velocity.X * dt
	e.Position.Y += e.Velocity.Y * dt
}

func GetWaveManifest(waveNum int) []func(*GameScene, float32, float32) *Entity {
	manifest := []func(*GameScene, float32, float32) *Entity{}

	for i := 0; i < 3+waveNum; i++ {
		manifest = append(manifest, NewBasicAlien)
	}

	return manifest
}

func CreateManifestSpawner(s *GameScene) func(*Entity, float32) {
	var timer float32 = 5
	var waveCount int = 1

	return func(e *Entity, dt float32) {
		timer += dt

		// New wave ever 10 seconds
		if timer >= 10.0 {

			waveItems := GetWaveManifest(waveCount)
			for _, spawnFunc := range waveItems {
				x := float32(rl.GetRandomValue(50, screenWidth-50))
				y := float32(rl.GetRandomValue(-1000, -50))

				s.EnemyEntities = append(s.EnemyEntities, spawnFunc(s, x, y))
			}

			timer = 0
			waveCount++
		}
	}
}
