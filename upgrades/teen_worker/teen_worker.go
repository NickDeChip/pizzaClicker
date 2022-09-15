package teen_worker

import (
	"fmt"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/NickDeChip/pizzaClicker/upgrades/aprons"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Worker struct {
	Count     int
	Gain      float64
	TotalGain float64
	Cost      float64
	x         float32
	y         float32
	Rec       rl.Rectangle
	iconRec   rl.Rectangle
	tex       *rl.Texture2D
	ShowText  bool
}

func New(US *rl.Texture2D) *Worker {
	if texture == nil {
		tex := rl.LoadTexture("Resources/upgradebox.png")
		texture = &tex
	}
	worker := &Worker{}
	worker.tex = US
	worker.Setup()
	return worker
}

func (w *Worker) Setup() {
	w.Count = 0
	w.Gain = 0.1
	w.TotalGain = 0
	w.Cost = 20
	w.x = 437
	w.y = 0
	w.Rec = rl.NewRectangle(w.x, w.y, float32(texture.Width), float32(texture.Height))
	w.iconRec = rl.NewRectangle(0, 0, float32(w.tex.Width/10), float32(w.tex.Height/10))
	w.ShowText = false
}
func (w *Worker) Draw() {
	rl.DrawTexture(*texture, int32(w.x), int32(w.y), rl.White)
	rl.DrawTextureRec(*w.tex, w.iconRec, rl.NewVector2(w.x+8, w.y+13), rl.White)
	rl.DrawText("Teen Workers", int32(w.x+100), int32(w.y+10), 24, rl.White)
	rl.DrawText(fmt.Sprintf("Cost: %.2f", w.Cost), int32(w.x+100), int32(w.y+40), 20, rl.White)
	rl.DrawText(fmt.Sprintf("Amount: %d", w.Count), int32(w.x+100), int32(w.y+62), 20, rl.White)
	if w.ShowText {
		rl.DrawText("Teen Worker", 30, 120, 30, rl.LightGray)
		rl.DrawText("This is a cheap worker,\nbut not the most efficient", 30, 160, 20, rl.LightGray)
		rl.DrawText(fmt.Sprintf("Creates %.1f PizzasPerSecond", w.Gain), 30, 220, 20, rl.LightGray)
		if w.Count >= 1 {
			rl.DrawText(fmt.Sprintf("You Gain %.1f PPS from this worker", w.Gain*float64(w.Count)), 30, 260, 18, rl.LightGray)
			rl.DrawText(fmt.Sprintf("You have made a total of %.1f \npizzas From this worker", w.TotalGain), 30, 290, 18, rl.LightGray)
		}
	}
}

func (w *Worker) Update(state *state.State, a *aprons.Aprons) {
	if w.Count >= 1 {
		state.PizzaCount += w.Gain * float64(w.Count)
		state.TotalPizzaCount += w.Gain * float64(w.Count)
		w.TotalGain += w.Gain * float64(w.Count)
	}
	if a.IsApronBought && w.Gain == 0.1 {
		w.Gain = w.Gain * 2
	} else if a.IsSABought && w.Gain == 0.2 {
		w.Gain = w.Gain * 2
	}
}

func (w *Worker) GetCost() float64 {
	return w.Cost
}

func (w *Worker) IncrementCost() {
	w.Cost *= 1.15
}

func (w *Worker) IncrementCount() {
	w.Count += 1
}
