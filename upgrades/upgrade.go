package upgrades

import (
	"github.com/NickDeChip/pizzaClicker/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Upgradable interface {
	GetCost() float64
	IncrementCost()
	IncrementCount()
}

func UpgradeColision(urec rl.Rectangle, mouse rl.Vector2, u Upgradable, state *state.State) {
	if rl.CheckCollisionPointRec(mouse, urec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if state.PizzaCount >= u.GetCost() {
				state.PizzaCount -= u.GetCost()
				u.IncrementCount()
				u.IncrementCost()
			}
		}
	}
}

func PowerUpColision(IsDiplayed bool, IsBought bool, mouse rl.Vector2, rec rl.Rectangle) bool {
	var answer bool
	if !IsBought && IsDiplayed {
		if rl.CheckCollisionPointRec(mouse, rec) {
			answer = true
		} else {
			answer = false
		}
	}
	return answer
}
