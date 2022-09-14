package pizzaoven

import (
	"fmt"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type PizzaOven struct {
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

func New(US *rl.Texture2D) *PizzaOven {
	if texture == nil {
		tex := rl.LoadTexture("Resources/powerupbackground.png")
		texture = &tex
	}
	pizzaOven := &PizzaOven{}
	pizzaOven.tex = US
	pizzaOven.Setup()
	return pizzaOven
}

func (po *PizzaOven) Setup() {
	po.Cost = 100
	po.x = 388
	po.y = 0
	po.iconRec = rl.NewRectangle(0, 128, float32(po.tex.Width/10), float32(po.tex.Height/10))
	po.Rec = rl.NewRectangle(po.x, po.y, float32(texture.Width), float32(texture.Height))
	po.IsBought = false
	po.DisplayUpgrade = false
	po.ShowText = false
}

func (po *PizzaOven) Draw() {
	if po.DisplayUpgrade {
		rl.DrawTexture(*texture, int32(po.x), int32(po.y), rl.White)
		rl.DrawTextureRec(*po.tex, po.iconRec, rl.NewVector2(po.x-18, po.y-11), rl.White)
		if po.ShowText {
			rl.DrawText("Pizza Oven", 30, 120, 30, rl.LightGray)
			rl.DrawText("Better pizza basses in\nhalf the time\nDoubles amount gain from clicking!", 30, 160, 20, rl.LightGray)
			rl.DrawText(fmt.Sprintf("Costs %.0f pizzas", po.Cost), 30, 250, 20, rl.LightGray)
		}
	}
}

func (po *PizzaOven) Update(state *state.State, mouse rl.Vector2) {
	if !po.IsBought && po.DisplayUpgrade {
		if rl.CheckCollisionPointRec(mouse, po.Rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && state.PizzaCount >= po.Cost {
				state.PizzaCount -= po.Cost
				po.IsBought = true
				po.DisplayUpgrade = false
			}
		}
	}
	if !po.DisplayUpgrade && !po.IsBought {
		if state.TotalPizzaCount >= 70 {
			po.DisplayUpgrade = true
		}
	}
}
