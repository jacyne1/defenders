package main

import (
	"fmt"
	"math"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	PlayerEntities []*Entity
	EnemyEntities  []*Entity
	Textures       map[string]rl.Texture2D
}

func (s *GameScene) Update() string {
	dt := rl.GetFrameTime()

	for _, e := range s.PlayerEntities {
		s.calculateBehaviors(e, dt)
	}

	for _, e := range s.EnemyEntities {
		s.calculateBehaviors(e, dt)
	}

	s.CollisionSystem()

	return "game"
}

func (s *GameScene) Draw() {
	rl.ClearBackground(rl.DarkBlue)

	for _, e := range s.PlayerEntities {
		s.drawEntity(e)
	}

	for _, e := range s.EnemyEntities {
		s.drawEntity(e)
	}
}

func (s *GameScene) Enter() {
	s.Textures = make(map[string]rl.Texture2D)

	s.loadAsset("alien", filepath.Join(assetBasePath, "alien4.png"))
	s.loadAsset("cannon", filepath.Join(assetBasePath, "cannon2.png"))
	s.loadAsset("projectile", filepath.Join(assetBasePath, "projectile.png"))

	waveManager := &Entity{
		Active: true,
		Class:  ClassNone,
		Behaviors: []func(*Entity, float32){
			CreateManifestSpawner(s),
		},
	}

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

		for _, e := range s.EnemyEntities {
			// Only check if the alient is actually active
			if !e.Active || e.Class != ClassAlien {
				continue
			}

			// Perform circle collision check
			if rl.CheckCollisionCircles(p.Position, p.Radius, e.Position, e.Radius) {
				fmt.Println("Ding")

				// 1. Deactivate the bullet
				p.Active = false

				// 2. Handle damage / hooks for the alien
				e.TakeDamage(p, p.AttackDamage)

			}
		}
	}
}
