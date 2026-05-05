package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SpawnBasicMissle(s *GameScene, pos rl.Vector2, rotation float32, texID string) *Entity {
	rad := (rotation) * (math.Pi / 180)
	tex := s.Textures[texID]
	radius := float32(tex.Width) / 2.0
	if tex.Height > tex.Width {
		radius = float32(tex.Height) / 2.0
	}

	speed := float32(400)
	vx := float32(math.Cos(float64(rad))) * speed
	vy := float32(math.Sin(float64(rad))) * speed

	return &Entity{
		Position:     pos,
		Velocity:     rl.NewVector2(vx, vy),
		Rotation:     rotation,
		Radius:       radius,
		Active:       true,
		TextureID:    texID,
		AttackDamage: 10,
		Class:        ClassBullet,
		Behaviors: []func(*Entity, float32){
			MoveForward,
		},
	}
}
