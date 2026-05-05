package main

import rl "github.com/gen2brain/raylib-go/raylib"

func NewBasicAlien(s *GameScene, x, y float32) *Entity {
	texName := "alien"
	tex := s.Textures[texName]
	radius := float32(tex.Width) / 2.0
	if tex.Height > tex.Width {
		radius = float32(tex.Height) / 2.0
	}

	return &Entity{
		ID:        1,
		TextureID: texName,
		Radius:    radius,
		Position:  rl.NewVector2(x, y),

		Velocity:  rl.NewVector2(0, 80),
		Active:    true,
		Health:    20,
		MaxHealth: 20,

		Class: ClassAlien,

		// MIX: It falls AND it zig-zags
		Behaviors: []func(*Entity, float32){
			FallNormal,
			CreateZigZag(),
		},
	}
}
