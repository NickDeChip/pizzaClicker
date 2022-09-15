package necronomicon

import (
	"fmt"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/NickDeChip/pizzaClicker/upgrades"
	"github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Necro struct {
	Cost           float64
	x              float32
	y              float32
	Rec            rl.Rectangle
	iconRec        rl.Rectangle
	IsBought       bool
	DisplayUpgrade bool
	tex            *rl.Texture2D
	ShowText       bool
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
	n.y = 90
	n.iconRec = rl.NewRectangle(0, 64, float32(n.tex.Width/10), float32(n.tex.Height/10))
	n.Rec = rl.NewRectangle(n.x, n.y, float32(texture.Width), float32(texture.Height))
	n.IsBought = false
	n.DisplayUpgrade = false
	n.ShowText = false
}

func (n *Necro) Draw() {
	if n.DisplayUpgrade {
		rl.DrawTexture(*texture, int32(n.x), int32(n.y), rl.White)
		rl.DrawTextureRec(*n.tex, n.iconRec, rl.NewVector2(n.x-18, n.y-11), rl.White)
		if n.ShowText {
			rl.DrawText("Necronicon", 30, 120, 30, rl.LightGray)
			rl.DrawText("A magic book that turns workers\ninto ZOMBES!", 30, 160, 20, rl.LightGray)
			rl.DrawText(fmt.Sprintf("Costs %.0f pizzas", n.Cost), 30, 220, 20, rl.LightGray)
		}
	}
}

func (n *Necro) Update(state *state.State, mouse rl.Vector2) {
	if upgrades.PowerUpColision(n.DisplayUpgrade, n.IsBought, mouse, n.Rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && state.PizzaCount >= n.Cost {
			state.PizzaCount -= n.Cost
			n.IsBought = true
			n.DisplayUpgrade = false
		}
	}

	if !n.DisplayUpgrade && !n.IsBought {
		if state.TotalPizzaCount >= 700 {
			n.DisplayUpgrade = true
		}
	}
}
