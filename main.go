package main

import (
	"github.com/NickDeChip/pizzaClicker/mouse"
	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/gen2brain/raylib-go/raylib"
)

const ScreenX = 800
const ScreenY = 450

var pizzaCount float64 = 0

func main() {
	rl.InitWindow(ScreenX, ScreenY, "PizzaClicker")

	fpscap := true
	Background := rl.LoadTexture("Resources/PizzaClickerBackground.png")
	bigPizza := pizza.New()
	mouse := mouse.New()
	rl.HideCursor()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyF) {
			fpscap = !fpscap
		}
		if fpscap {
			rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
		} else {
			rl.SetTargetFPS(0)
		}

		mouse.Update()
		colision(bigPizza.Crec, mouse.Position, bigPizza)

		bigPizza.Animation()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(Background, 0, 0, rl.White)
		bigPizza.Draw(pizzaCount)
		mouse.Draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func colision(prec rl.Rectangle, mouse rl.Vector2, p *pizza.Pizza) {
	if rl.CheckCollisionPointRec(mouse, prec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			pizzaCount += .1
			p.IsPizzaClicked = true
		}
	}
}
