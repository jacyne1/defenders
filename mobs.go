package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MobBasic = iota
	MobFast
	MobZigZag
)

func (s *GameScene) NewBasicAlien(x, y float32) {
	texName := "alien"
	tex := s.Textures[texName]
	radius := float32(tex.Width) / 2.0
	if tex.Height > tex.Width {
		radius = float32(tex.Height) / 2.0
	}

	for i := range s.EnemyEntities {
		if !s.EnemyEntities[i].Active {
			fmt.Println("Spawning Alien")
			s.EnemyEntities[i].Active = true
			s.EnemyEntities[i].ID = 1
			s.EnemyEntities[i].TextureID = texName
			s.EnemyEntities[i].Radius = radius
			s.EnemyEntities[i].Position = rl.NewVector2(x, y)

			s.EnemyEntities[i].Velocity = rl.NewVector2(0, 80)
			s.EnemyEntities[i].Health = 20
			s.EnemyEntities[i].MaxHealth = 20
			s.EnemyEntities[i].PointValue = 100

			s.EnemyEntities[i].Class = ClassAlien

			s.EnemyEntities[i].Behaviors = []func(*Entity, float32){
				FallNormal,
				SeekCenter,
			}
			return
		}
	}
}

func (s *GameScene) NewBasicZigZagAlien(x, y float32) {
	texName := "alien"
	tex := s.Textures[texName]
	radius := float32(tex.Width) / 2.0
	if tex.Height > tex.Width {
		radius = float32(tex.Height) / 2.0
	}

	for i := range s.EnemyEntities {
		if !s.EnemyEntities[i].Active {
			fmt.Println("Spawning Alien")
			s.EnemyEntities[i].Active = true
			s.EnemyEntities[i].ID = 1
			s.EnemyEntities[i].TextureID = texName
			s.EnemyEntities[i].Radius = radius
			s.EnemyEntities[i].Position = rl.NewVector2(x, y)

			s.EnemyEntities[i].Velocity = rl.NewVector2(0, 80)
			s.EnemyEntities[i].Health = 20
			s.EnemyEntities[i].MaxHealth = 20
			s.EnemyEntities[i].PointValue = 100

			s.EnemyEntities[i].Class = ClassAlien

			// MIX: It falls AND it zig-zags
			s.EnemyEntities[i].Behaviors = []func(*Entity, float32){
				FallNormal,
				CreateZigZag(),
				SeekCenter,
			}
			return
		}
	}
}

func (s *GameScene) SpawnAlien(mobType int, x, y float32) {
	for i := range s.EnemyEntities {
		if !s.EnemyEntities[i].Active {
			e := &s.EnemyEntities[i]

			// Common to all aliens
			e.Active = true
			e.Class = ClassAlien
			e.Position = rl.NewVector2(x, y)

			// Default values that get overridden below if needed
			var texName string = "alien"
			var baseSpeed float32 = 80

			switch mobType {
			case MobBasic:
				e.ID = 1
				e.Health = 20
				e.MaxHealth = 20
				e.PointValue = 100

				e.Behaviors = []func(*Entity, float32){
					FallNormal,
					SeekCenter,
				}
			case MobZigZag:
				e.ID = 2
				e.Health = 20
				e.MaxHealth = 20
				e.PointValue = 200
				e.Behaviors = []func(*Entity, float32){
					FallNormal,
					CreateZigZag(),
					SeekCenter,
				}
			}

			// Texture and Radius Setup
			e.TextureID = texName
			tex := s.Textures[texName]
			radius := float32(tex.Width) / 2.0
			if tex.Height > tex.Width {
				radius = float32(tex.Height) / 2.0
			}
			e.Radius = radius

			e.Velocity = rl.NewVector2(0, baseSpeed)
			fmt.Println("Spawning Alien ID:", e.ID)
			return // Successfully spawned, exit function

		}
	}
}
