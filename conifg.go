package main

type Scene interface {
	Update() string
	Draw()
}

const (
	screenWidth  = 1200
	screenHeight = 720
	title        = "Defender Game"
)
