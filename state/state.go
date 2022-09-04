package state

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	PizzaCount float64
	FPScap     bool
	Background rl.Texture2D
	DT         float32
	Timer      float32
}
