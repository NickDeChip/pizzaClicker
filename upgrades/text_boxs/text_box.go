package textboxs

import (
	//"github.com/NickDeChip/pizzaClicker/state"
	"github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type TextBox struct {
	x          float32
	y          float32
	Rec        rl.Rectangle
	ShowBox    bool
	ShowTextTW bool
	ShowTextAW bool
}

func New() *TextBox {
	if texture == nil {
		tex := rl.LoadTexture("Resources/textbox.png")
		texture = &tex
	}
	textBox := &TextBox{}
	textBox.Setup()
	return textBox
}

func (t *TextBox) Setup() {
	t.x = 2
	t.y = 110
	t.Rec = rl.NewRectangle(t.x, t.y, float32(texture.Width), float32(texture.Height))
	t.ShowBox = false
}

func (t *TextBox) Draw() {
	if t.ShowBox {
		rl.DrawTexture(*texture, int32(t.x), int32(t.y), rl.White)
	}
}

func (t *TextBox) Update() {
	t.ShowBox = false
}

func (t *TextBox) CollisionCheck(mouse rl.Vector2, uRec rl.Rectangle) bool {
	if rl.CheckCollisionPointRec(mouse, uRec) {
		t.ShowBox = true
		return true
	}
	return false
}
