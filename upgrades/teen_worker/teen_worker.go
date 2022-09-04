package teen_worker

import (
	"github.com/NickDeChip/pizzaClicker/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Worker struct {
	Count float64
	Gain  float64
	cost  float64
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
	w.Gain = 1
	w.cost = 20
	w.x = 439
	w.y = 0
	w.Rec = rl.NewRectangle(w.x, w.y, float32(texture.Width), float32(texture.Width))
}

func (w *Worker) Draw() {
	rl.DrawTexture(*texture, int32(w.x), int32(w.y), rl.White)
	rl.DrawText("Teen Workers", int32(w.x+80), int32(w.y), 24, rl.White)
}

func (w *Worker) Update(state *state.State) {
	if w.Count >= 1 && state.Timer >= 1 {
		state.PizzaCount += w.Gain
		state.DT = 0
	}
}
