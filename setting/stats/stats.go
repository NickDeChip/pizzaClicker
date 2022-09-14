package stats

import (
	"fmt"

	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/gen2brain/raylib-go/raylib"
)

var texture *rl.Texture2D

type Stats struct {
	x        int32
	y        int32
	Rec      rl.Rectangle
	IconRec  rl.Rectangle
	tex      *rl.Texture2D
	ShowText bool
}

func New(US *rl.Texture2D) *Stats {
	if texture == nil {
		tex := rl.LoadTexture("Resources/powerupbackground.png")
		texture = &tex
	}
	stats := &Stats{}
	stats.tex = US
	stats.SetUp()
	return stats
}

func (s *Stats) SetUp() {
	s.x = 10
	s.y = 66
	s.Rec = rl.NewRectangle(float32(s.x), float32(s.y), float32(texture.Width), float32(texture.Height))
	s.IconRec = rl.NewRectangle(0, 256, float32(s.tex.Height/10), float32(s.tex.Width/10))
	s.ShowText = false
}

func (s *Stats) Draw(state *state.State, p pizza.Pizza) {
	rl.DrawTexture(*texture, s.x, s.y, rl.White)
	rl.DrawTextureRec(*s.tex, s.IconRec, rl.NewVector2(float32(s.x-18), float32(s.y-11)), rl.White)
	if s.ShowText {
		rl.DrawText("Stats", 30, 120, 30, rl.LightGray)
		rl.DrawText(fmt.Sprintf("Total Pizzas Made %.0f", state.TotalPizzaCount), 30, 160, 18, rl.LightGray)
		rl.DrawText(fmt.Sprintf("Total Pizzas Clicked %.0f", p.ClickCount*p.PizzaClickMultiplyer), 30, 190, 18, rl.LightGray)
		rl.DrawText(fmt.Sprintf("Total Times Pizza Clicked %.0f", p.ClickCount), 30, 220, 18, rl.LightGray)
	}

}
