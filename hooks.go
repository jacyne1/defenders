package main

import "fmt"

type DamageHook func(attacker *Entity, currentDamage int) int

func ArmorHook(attacker *Entity, damage int) int {
	return damage / 2 // Passive 50% damage reduction
}

func ExplosiveDeathHook(e *Entity) {
	// Logic to spawn an explosion entity at e.Position
	fmt.Println("BOOM at", e.Position)
}
