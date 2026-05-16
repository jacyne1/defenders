package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *GameScene) NewBasicAlien(x, y float32) {
	texName := "alien"
	tex := s.Textures[texName]
	radius := float32(tex.Width) / 2.0
	if tex.Height > tex.Width {
		radius = float32(tex.Height) / 2.0
	}
	fmt.Println("Spawned Alien")
	for i := range s.EnemyEntities {
		if !s.EnemyEntities[i].Active {
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
			}
			return
		}
	}
}
