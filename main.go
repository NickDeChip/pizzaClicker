package main

import (
	"github.com/NickDeChip/pizzaClicker/enity"
	"github.com/NickDeChip/pizzaClicker/mouse"
	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/NickDeChip/pizzaClicker/upgrades/teen_worker"
	"github.com/gen2brain/raylib-go/raylib"
)

const ScreenX = 800
const ScreenY = 450

func main() {
	rl.InitWindow(ScreenX, ScreenY, "PizzaClicker")

	state := &state.State{
		PizzaCount: 0,
		FPScap:     true,
		Background: rl.LoadTexture("Resources/PizzaClickerBackground.png"),
		DT:         rl.GetFrameTime(),
		Timer:      0,
	}

	enity := enity.Enity{
		BigPizza: *pizza.New(),
		Mouse:    *mouse.New(),
		TW:       *teen_worker.New(),
	}
	rl.HideCursor()

	for !rl.WindowShouldClose() {
		state.Timer += state.DT

		if rl.IsKeyPressed(rl.KeyF) {
			state.FPScap = !state.FPScap
		}
		if state.FPScap {
			rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
		} else {
			rl.SetTargetFPS(0)
		}

		upgradeColision(enity.TW.Rec, enity.Mouse.Position, &enity.TW)
		enity.TW.Update(state, enity)

		enity.Mouse.Update()
		pizzaColision(enity.BigPizza.Crec, enity.Mouse.Position, &enity.BigPizza, state)

		enity.BigPizza.Animation()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(state.Background, 0, 0, rl.White)
		enity.BigPizza.Draw(state.PizzaCount)

		enity.TW.Draw()

		enity.Mouse.Draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func pizzaColision(prec rl.Rectangle, mouse rl.Vector2, p *pizza.Pizza, state *state.State) {
	if rl.CheckCollisionPointRec(mouse, prec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			state.PizzaCount += .1
			p.IsPizzaClicked = true
		}
	}
}

func upgradeColision(urec rl.Rectangle, mouse rl.Vector2, u *teenworker.Worker) {
	if rl.CheckCollisionPointRec(mouse, urec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			u.Count += 1
		}
	}
}
