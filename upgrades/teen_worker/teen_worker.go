package teen_worker

import (
	"fmt"

	"github.com/NickDeChip/pizzaClicker/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Worker struct {
	Count int
	Gain  float64
	Cost  float64
	x     float32
	y     float32
	Rec   rl.Rectangle
}

func New() *Worker {
	if texture == nil {
		tex := rl.LoadTexture("Resources/teenworker.png")
		texture = &tex
	}
	worker := &Worker{}
	worker.Setup()
	return worker
}

func (w *Worker) Setup() {
	w.Count = 0
	w.Gain = 0.1
	w.Cost = 20
	w.x = 437
	w.y = 0
	w.Rec = rl.NewRectangle(w.x, w.y, float32(texture.Width), float32(texture.Height))
}

func (w *Worker) Draw() {
	rl.DrawTexture(*texture, int32(w.x), int32(w.y), rl.White)
	rl.DrawText("Teen Workers", int32(w.x+100), int32(w.y+10), 24, rl.White)
	rl.DrawText(fmt.Sprintf("Cost: %.2f", w.Cost), int32(w.x+100), int32(w.y+40), 20, rl.White)
	rl.DrawText(fmt.Sprintf("Amount: %d", w.Count), int32(w.x+100), int32(w.y+65), 20, rl.White)
}

func (w *Worker) Update(state *state.State) {
	if w.Count >= 1 {
		state.PizzaCount += w.Gain * float64(w.Count)
		state.Timer = 0
	}
}

func (w *Worker) GetCost() float64 {
	return w.Cost
}

func (w *Worker) IncrementCost() {
	w.Cost *= 1.1
}

func (w *Worker) IncrementCount() {
	w.Count += 1
}
