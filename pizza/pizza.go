package pizza

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var pizzaTexture *rl.Texture2D

type Pizza struct {
	x                 float32
	y                 float32
	width             float32
	height            float32
	clickCount        int
	Rec               rl.Rectangle
	Crec              rl.Rectangle
	MainPizza         float64
	animationspeed    float32
	currentPizzaFrame int
	IsPizzaClicked    bool
}

func New() *Pizza {
	if pizzaTexture == nil {
		tex := rl.LoadTexture("Resources/Pizza.png")
		pizzaTexture = &tex
	}
	pizza := &Pizza{}
	pizza.Setup()
	return pizza
}

func (p *Pizza) Setup() {
	p.x = 80
	p.y = 150
	p.width = float32(pizzaTexture.Width)
	p.height = float32(pizzaTexture.Height)
	p.clickCount = 0
	p.Rec = rl.NewRectangle(0, 0, p.width/4, p.height)
	p.Crec = rl.NewRectangle(p.x, p.y, p.width/4, p.height)
	p.animationspeed = 0
	p.IsPizzaClicked = false
}

func (p *Pizza) Draw(pizzacount float64) {
	rl.DrawTextureRec(*pizzaTexture, p.Rec, rl.NewVector2(p.x, p.y), rl.White)
	rl.DrawText(fmt.Sprintf("Pizzas: %.0f", pizzacount), 40, 10, 50, rl.White)
}

func (p *Pizza) Animation() {
	p.animationspeed += rl.GetFrameTime()

	if p.animationspeed >= 0.025 && p.IsPizzaClicked {
		p.animationspeed = 0
		switch p.currentPizzaFrame {
		case 0:
			p.currentPizzaFrame = 1
		case 1:
			p.currentPizzaFrame = 2
		case 2:
			p.currentPizzaFrame = 3
		case 3:
			p.currentPizzaFrame = 0
			p.IsPizzaClicked = false
		}

		p.Rec.X = ((p.Rec.Width) * float32(p.currentPizzaFrame))
	}
}
