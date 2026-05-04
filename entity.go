package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	ID       int
	Position rl.Vector2
	Velocity rl.Vector2
	Radius   float32
	Active   bool

	Shield     *ShieldData
	AmmoEffect AmmoType

	// Active Behaviors (Movement, Shooting)
	Behaviors []func(e *Entity, dt float32)

	// Reactive Hooks (On hit, On Death)
	OnHitHooks   []func(attacker *Entity, damage int) int
	OnDeathHooks []func(e *Entity)
}

type ShieldData struct {
	Active bool
	Radius float32 // Pixels beyond entity radius
	Health int
}

type AmmoType string

const (
	AmmoKinetic AmmoType = "bullet"
	AmmoFlame   AmmoType = "flame"
	AmmoPhaser  AmmoType = "phaser"
)
