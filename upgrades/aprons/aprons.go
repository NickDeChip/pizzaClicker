package aprons

import (
	"fmt"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Aprons struct {
	Cost                float64
	x                   float32
	y                   float32
	Rec                 rl.Rectangle
	iconRec             rl.Rectangle
	IsApronBought       bool
	IsSABought          bool
	DisplayUpgradeApron bool
	DisplayUpgradeSA    bool
	tex                 *rl.Texture2D
	ShowText            bool
}

func New(US *rl.Texture2D) *Aprons {
	if texture == nil {
		tex := rl.LoadTexture("Resources/powerupbackground.png")
		texture = &tex
	}
	aprons := &Aprons{}
	aprons.tex = US
	aprons.Setup()
	return aprons
}

func (a *Aprons) Setup() {
	a.Cost = 250
	a.x = 388
	a.y = 45
	a.iconRec = rl.NewRectangle(0, 192, float32(a.tex.Width/10), float32(a.tex.Height/10))
	a.Rec = rl.NewRectangle(a.x, a.y, float32(texture.Width), float32(texture.Height))
	a.IsApronBought = false
	a.DisplayUpgradeApron = false
	a.DisplayUpgradeSA = false
	a.ShowText = false
}

func (a *Aprons) Draw() {
	if a.DisplayUpgradeApron {
		rl.DrawTexture(*texture, int32(a.x), int32(a.y), rl.White)
		rl.DrawTextureRec(*a.tex, a.iconRec, rl.NewVector2(a.x-18, a.y-11), rl.White)
		if a.ShowText {
			rl.DrawText("Aprons", 30, 120, 30, rl.LightGray)
			rl.DrawText("Better looking workers\nmakes them work harder\nDoubles efficiency Of workers!", 30, 160, 20, rl.LightGray)
			rl.DrawText(fmt.Sprintf("Costs %.0f pizzas", a.Cost), 30, 250, 20, rl.LightGray)
		}
	}

	if a.DisplayUpgradeSA {
		rl.DrawTexture(*texture, int32(a.x), int32(a.y), rl.White)
		rl.DrawTextureRec(*a.tex, a.iconRec, rl.NewVector2(a.x-18, a.y-11), rl.White)
		if a.ShowText {
			rl.DrawText("Silver Aprons", 30, 120, 30, rl.LightGray)
			rl.DrawText("Richer workers work harder\nfor the same price\nDoubles efficiency Of workers!", 30, 160, 20, rl.LightGray)
			rl.DrawText(fmt.Sprintf("Costs %.0f pizzas", a.Cost), 30, 250, 20, rl.LightGray)
		}
	}
}

func (a *Aprons) Update(state *state.State, mouse rl.Vector2) {
	if !a.IsApronBought && a.DisplayUpgradeApron {
		if rl.CheckCollisionPointRec(mouse, a.Rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && state.PizzaCount >= a.Cost {
				state.PizzaCount -= a.Cost
				a.IsApronBought = true
				a.DisplayUpgradeApron = false
			}
		}
	}
	if !a.DisplayUpgradeApron && !a.IsApronBought {
		if state.TotalPizzaCount >= 100 {
			a.DisplayUpgradeApron = true
		}
	}
	if !a.DisplayUpgradeSA && !a.IsSABought && a.IsApronBought {
		if state.TotalPizzaCount >= 1000 {
			a.DisplayUpgradeSA = true
			a.iconRec = rl.NewRectangle(80, 192, float32(a.tex.Width/10), float32(a.tex.Height/10))
		}
	}
}
