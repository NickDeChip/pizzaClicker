package pizza

import (
	"fmt"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/NickDeChip/pizzaClicker/upgrades/pizza_oven"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pizzaTexture *rl.Texture2D

type Pizza struct {
	x                    float32
	y                    float32
	width                float32
	height               float32
	PizzaClickMultiplyer float64
	ClickCount           float64
	Rec                  rl.Rectangle
	Crec                 rl.Rectangle
	MainPizza            float64
	animationspeed       float32
	currentPizzaFrame    int
	IsPizzaClicked       bool
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
	p.PizzaClickMultiplyer = 1
	p.ClickCount = 0
	p.Rec = rl.NewRectangle(0, 0, p.width/4, p.height)
	p.Crec = rl.NewRectangle(p.x, p.y, p.width/4, p.height)
	p.animationspeed = 0
	p.IsPizzaClicked = false
}

func (p *Pizza) Draw(pizzacount float64) {
	rl.DrawTextureRec(*pizzaTexture, p.Rec, rl.NewVector2(p.x, p.y), rl.White)
	rl.DrawText(fmt.Sprintf("Pizzas: %d", int64(pizzacount)), 40, 10, 50, rl.White)
}

func (p *Pizza) Update(po *pizzaoven.PizzaOven) {
	if po.IsBought && p.PizzaClickMultiplyer == 1 {
		p.PizzaClickMultiplyer = 2
	}
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

func (p *Pizza) PizzaColision(prec rl.Rectangle, mouse rl.Vector2, state *state.State) {
	if rl.CheckCollisionPointRec(mouse, prec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			state.PizzaCount += 1 * p.PizzaClickMultiplyer
			state.TotalPizzaCount += 1 * p.PizzaClickMultiplyer
			p.ClickCount += 1
			p.IsPizzaClicked = true
		}
	}
}
