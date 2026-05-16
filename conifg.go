package main

type Scene interface {
	Enter()
	Update() string
	Draw()
	Exit()
}

const (
	screenWidth   = 1920
	screenHeight  = 1080
	title         = "Defender Game"
	assetBasePath = "assets"
)
