package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	StatePlaying GameState = iota
	StatePaused
)

type GameScene struct {
	PlayerEntities []*Entity
	EnemyEntities  []Entity
	Textures       map[string]rl.Texture2D
	State          GameState
	PlayerPoints   int
}

func NewGameScene() *GameScene {
	return &GameScene{
		PlayerEntities: make([]*Entity, 0),
		EnemyEntities:  make([]Entity, 100),
	}
}

func (s *GameScene) Update() string {
	dt := rl.GetFrameTime()

	if s.State == StatePaused {
		if rl.IsKeyPressed(rl.KeyP) {
			s.State = StatePlaying
			fmt.Println("State Paused")
			return "game"
		}
	}

	if s.State == StatePlaying {

		if rl.IsKeyPressed(rl.KeyP) {
			s.State = StatePaused
		}

		for _, e := range s.PlayerEntities {
			s.calculateBehaviors(e, dt)
		}

		for i := range s.EnemyEntities {
			s.calculateBehaviors(&s.EnemyEntities[i], dt)
		}

		ManageEnemyActive(s.EnemyEntities)
		ManagePlayerActive(s.PlayerEntities)

		s.CollisionSystem()
	}

	return "game"
}

func (s *GameScene) Draw() {
	rl.ClearBackground(rl.DarkBlue)
	rl.DrawTexture(s.Textures["background"], 0, 0, rl.White)

	for _, e := range s.PlayerEntities {
		s.drawEntity(e)
	}

	for i := range s.EnemyEntities {
		s.drawEntity(&s.EnemyEntities[i])
	}

	DrawPoints(s.PlayerPoints)
}

func (s *GameScene) Enter() {
	s.Textures = make(map[string]rl.Texture2D)

	s.loadAsset("alien", filepath.Join(assetBasePath, "alien4.png"))
	s.loadAsset("cannon", filepath.Join(assetBasePath, "cannon2.png"))
	s.loadAsset("projectile", filepath.Join(assetBasePath, "projectile.png"))
	s.loadAsset("background", filepath.Join(assetBasePath, "game_backgrounds", "BG.png"))

	s.PlayerPoints = 0
	//s.EnemyEntities = make([]*Entity, 100),

	waveManager := &Entity{
		Active: true,
		Class:  ClassNone,
		Behaviors: []func(*Entity, float32){
			CreateManifestSpawner(s),
		},
	}

	s.State = StatePlaying

	cannon := NewBasicCannon()
	cannon.Behaviors = append(cannon.Behaviors, CreateShootBehavior(s))
	s.PlayerEntities = append(s.PlayerEntities, cannon, waveManager)
}

func (s *GameScene) Exit() {
	for _, tex := range s.Textures {
		rl.UnloadTexture(tex)
	}

	s.Textures = nil
}

func (s *GameScene) loadAsset(key string, path string) {
	if _, exists := s.Textures[key]; !exists {
		s.Textures[key] = rl.LoadTexture(path)
	}
}

func (s *GameScene) drawEntity(e *Entity) {

	if !e.Active {

		return
	}

	texture := s.Textures[e.TextureID]

	source := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	dest := rl.NewRectangle(e.Position.X, e.Position.Y, float32(texture.Width), float32(texture.Height))
	origin := rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2)

	rl.DrawTexturePro(texture, source, dest, origin, e.Rotation, rl.White)

}

func (s *GameScene) calculateBehaviors(e *Entity, dt float32) {
	if !e.Active {
		return
	}

	for _, behavior := range e.Behaviors {
		behavior(e, dt)
	}

}

func CreateShootBehavior(s *GameScene) func(*Entity, float32) {
	return func(e *Entity, dt float32) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rad := float64(e.Rotation) * (math.Pi / 180)
			barrelLength := float32(60)
			spawnX := e.Position.X + float32(math.Cos(rad))*barrelLength
			spawnY := e.Position.Y + float32(math.Sin(rad))*barrelLength
			m := SpawnBasicMissle(s, rl.NewVector2(spawnX, spawnY), e.Rotation, "projectile")
			s.PlayerEntities = append(s.PlayerEntities, m)
		}
	}
}

func (s *GameScene) CollisionSystem() {
	// Check every player entity
	for _, p := range s.PlayerEntities {
		if !p.Active || p.Class != ClassBullet {
			continue
		}

		for i := range s.EnemyEntities {

			if !s.EnemyEntities[i].Active || s.EnemyEntities[i].Class != ClassAlien {
				continue
			}

			if rl.CheckCollisionCircles(p.Position, p.Radius, s.EnemyEntities[i].Position, s.EnemyEntities[i].Radius) {

				// 1. Deactivate the bullet
				p.Active = false
				// 2. Handle damage / hooks for the alien
				points := s.EnemyEntities[i].TakeDamage(p, p.AttackDamage)
				if points > 0 {
					s.PlayerPoints += points
				}
			}
		}

	}
}

func DrawPoints(points int) {
	scoreText := "SCORE: " + strconv.Itoa(points)

	var fontSize int32 = 30
	var posX int32 = 20
	var posY int32 = 20

	rl.DrawText(scoreText, posX+2, posY+2, fontSize, rl.NewColor(0, 0, 0, 150))
	rl.DrawText(scoreText, posX, posY, fontSize, rl.RayWhite)
}
