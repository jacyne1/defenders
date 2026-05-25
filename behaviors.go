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
	if e.Position.X < screenWidth/2 {
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

func GetWaveManifest(s *GameScene, waveNum int) []func(float32, float32) {
	manifest := make([]func(float32, float32), 0, 3+waveNum)

	for i := 0; i < 3+waveNum; i++ {
		manifest = append(manifest, s.NewBasicAlien)
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
			spawnCount := 3 + waveCount

			for i := 0; i < spawnCount; i++ {
				x := float32(rl.GetRandomValue(150, screenWidth-50))
				y := float32(rl.GetRandomValue(-1000, -50))

				s.SpawnAlien(MobBasic, x, y)
			}

			timer = 0
			waveCount++
		}
	}
}

func CreateShootBehavior(s *GameScene) func(*Entity, float32) {
	return func(e *Entity, dt float32) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rad := float64(e.Rotation) * (math.Pi / 180)
			barrelLength := float32(60)
			spawnX := e.Position.X + float32(math.Cos(rad))*barrelLength
			spawnY := e.Position.Y + float32(math.Sin(rad))*barrelLength
			s.SpawnBasicMissle(rl.NewVector2(spawnX, spawnY), e.Rotation, "projectile")
		}
	}
}

func CreateCampaignSpawner(s *GameScene) func(*Entity, float32) {
	var currentLevel int = 0
	var currentWave int = 0

	var spawnTimer float32 = 0
	var intermissionTimer float32 = 0
	var totalSpawnedInWave int = 0
	var isWaitingForNextWave bool = false

	return func(e *Entity, dt float32) {
		// Safety check: if player beat the final level, stop processing
		if currentLevel >= len(GameCampaign) {
			// Optional: Trigger game win screen here
			return
		}

		levelData := GameCampaign[currentLevel]
		waveData := levelData.Waves[currentWave]

		// ----------------------------------------------------
		// PHASE 1: SPAWNING THE WAVE
		// ----------------------------------------------------

		if totalSpawnedInWave < waveData.TotalToSpawn {
			spawnTimer += dt
			if spawnTimer >= waveData.SpawnRate {
				spawnTimer = 0
				totalSpawnedInWave++

				// Coordinate Logic
				x := float32(rl.GetRandomValue(150, screenWidth-50))
				y := float32(rl.GetRandomValue(-200, -50))

				chosenMobType := GetRandomMobType(waveData.MobPool)

				s.SpawnAlien(chosenMobType, x, y)
			}
			return
		}

		// ----------------------------------------------------
		// PHASE 2: ACTIVE COMBAT (Waiting for board clearance)
		// ----------------------------------------------------

		// Count how many are active right now
		activeCount := 0
		for i := range s.EnemyEntities {
			if s.EnemyEntities[i].Active {
				activeCount++
			}
		}

		if activeCount == 0 && !isWaitingForNextWave {
			isWaitingForNextWave = true
			intermissionTimer = 0
		}

		// ----------------------------------------------------
		// PHASE 3: INTERMISSION / LEVEL ADVANCEMENT
		// ----------------------------------------------------

		if isWaitingForNextWave {
			intermissionTimer += dt

			// Visual feedback using standard Raylib text functions
			// e.g., rl.DrawText("WAVE COMPLETE!", X, Y, 20, rl.RayWhite)

			if intermissionTimer >= waveData.IntermissionDelay {
				// Reset internal counters for upcoming wave
				isWaitingForNextWave = false
				totalSpawnedInWave = 0
				spawnTimer = 0

				if currentWave < len(levelData.Waves)-1 {
					currentWave++
				} else {
					currentLevel++
					currentWave = 0
				}
			}
		}
	}
}

func GetRandomMobType(pool []WaveMobConfig) int {
	totalWeight := 0
	for i := range pool {
		totalWeight += pool[i].Weight
	}

	randomValue := int(rl.GetRandomValue(0, int32(totalWeight-1)))

	currentSum := 0

	for i := range pool {
		currentSum += pool[i].Weight
		if randomValue < currentSum {
			return pool[i].MobType
		}
	}
	return pool[0].MobType
}
