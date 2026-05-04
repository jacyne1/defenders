package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Alien struct {
	id int
}

type Gun struct {
	id int
}

type GameScene struct {
	aliens []Alien
	guns   []Gun
}

func (s *GameScene) Update() string {
	return "game"
}

func (s *GameScene) Draw() {
	rl.ClearBackground(rl.DarkBlue)
}
