package mouse

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var cursorTex *rl.Texture2D

type Mouse struct {
	Position rl.Vector2
}

func New() *Mouse {
	if cursorTex == nil {
		tex := rl.LoadTexture("Resources/cursor.png")
		cursorTex = &tex
	}
	pizza := &Mouse{}
	pizza.Setup()
	return pizza
}

func (m *Mouse) Setup() {
	m.Position = rl.GetMousePosition()
}

func (m *Mouse) Draw() {
	rl.DrawTextureV(*cursorTex, m.Position, rl.White)
}

func (m *Mouse) Update() {
	m.Position = rl.GetMousePosition()
}
