package main

type Scene interface {
	Enter()
	Update() string
	Draw()
	Exit()
}

const (
	screenWidth   = 1200
	screenHeight  = 720
	title         = "Defender Game"
	assetBasePath = "assets"
)
