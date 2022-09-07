package necronomicon

import (
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Necro struct {
	Cost           float64
	x              float32
	y              float32
	Rec            rl.Rectangle
	iconRec        rl.Rectangle
	isBought       bool
	displayUpgrade bool
	tex            *rl.Texture2D
}

func New(US *rl.Texture2D) *Necro {
	if texture == nil {
		tex := rl.LoadTexture("Resources/powerupbackground.png")
		texture = &tex
	}
	necro := &Necro{}
	necro.tex = US
	necro.Setup()
	return necro
}

func (n *Necro) Setup() {
	n.Cost = 1000
	n.x = 388
	n.y = 0
	n.iconRec = rl.NewRectangle(0, 63, float32(n.tex.Width/10), float32(n.tex.Height/10))
	n.Rec = rl.NewRectangle(n.x, n.y, float32(texture.Width), float32(texture.Height))
	n.isBought = false
	n.displayUpgrade = false
}

func (n *Necro) Draw() {
	if n.displayUpgrade {
		rl.DrawTexture(*texture, int32(n.x), int32(n.y), rl.White)
		rl.DrawTextureRec(*n.tex, n.iconRec, rl.NewVector2(n.x-18, n.y-11), rl.White)
	}
}

func (n *Necro) Update(state *state.State, mouse rl.Vector2) {
	if !n.isBought && n.displayUpgrade {
		if rl.CheckCollisionPointRec(mouse, n.Rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && state.PizzaCount >= n.Cost {
				state.PizzaCount -= n.Cost
				n.isBought = true
				n.displayUpgrade = false
			}
		}
	}
	if !n.displayUpgrade && !n.isBought {
		if state.TotalPizzaCount >= 700 {
			n.displayUpgrade = true
		}
	}
}
