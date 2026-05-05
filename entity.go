package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityClass int

const (
	ClassNone EntityClass = iota
	ClassPlayerTurret
	ClassTurret
	ClassBullet
	ClassAlien
	ClassBase
)

type Entity struct {
	ID        int
	TextureID string
	Position  rl.Vector2
	Velocity  rl.Vector2
	Radius    float32
	Active    bool
	Rotation  float32
	Class     EntityClass

	Health       int
	MaxHealth    int
	AttackDamage int

	Shield     *ShieldData
	AmmoEffect AmmoType

	// Active Behaviors (Movement, Shooting)
	Behaviors []func(e *Entity, dt float32)

	// Reactive Hooks (On hit, On Death)
	DefensiveHooks  []DamageHook
	MultiplierHooks []DamageHook
	OnDeathHooks    []func(e *Entity)
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

func NewBasicCannon() *Entity {
	return &Entity{
		ID:        2,
		TextureID: "cannon",
		Position:  rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)-100),
		Active:    true,

		Behaviors: []func(*Entity, float32){
			RotateTowardMouse,
		},
	}
}

func FilterActive(entities []*Entity) []*Entity {
	active := entities[:0]
	for _, e := range entities {
		// Also deactivate if it's way off screen
		if e.Position.Y < -200 || e.Position.Y > screenHeight+100 || e.Position.X < -200 || e.Position.X > screenWidth+200 {
			e.Active = false
		}

		if e.Active {
			active = append(active, e)
		}
	}
	return active
}

func (e *Entity) TakeDamage(attacker *Entity, rawDamage int) {
	finalDamage := rawDamage

	// 1. Process Reductions / DefensiveHooks
	for _, hook := range e.DefensiveHooks {
		finalDamage = hook(attacker, finalDamage)
	}

	// 2. Process MultiplierHooks
	for _, hook := range e.MultiplierHooks {
		finalDamage = hook(attacker, finalDamage)
	}

	e.Health -= finalDamage

	if e.Health <= 0 {
		e.Health = 0
		e.Active = false
		for _, deathHook := range e.OnDeathHooks {
			deathHook(e)
		}
	}
}
