package main

import (
	"github.com/NickDeChip/pizzaClicker/enity"
	"github.com/NickDeChip/pizzaClicker/mouse"
	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/NickDeChip/pizzaClicker/upgrades"
	"github.com/NickDeChip/pizzaClicker/upgrades/adult_worker"
	"github.com/NickDeChip/pizzaClicker/upgrades/necronomicon"
	"github.com/NickDeChip/pizzaClicker/upgrades/teen_worker"
	textboxs "github.com/NickDeChip/pizzaClicker/upgrades/text_boxs"
	"github.com/NickDeChip/pizzaClicker/upgrades/zombe_aw"
	"github.com/NickDeChip/pizzaClicker/upgrades/zombe_tw"
	"github.com/gen2brain/raylib-go/raylib"
)

const ScreenX = 800
const ScreenY = 450

func main() {
	rl.InitWindow(ScreenX, ScreenY, "PizzaClicker")
	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
	upgradeSheet := rl.LoadTexture("Resources/upgradespritesheet.png")

	state := &state.State{
		PizzaCount:      1000,
		TotalPizzaCount: 1000,
		FPScap:          true,
		Background:      rl.LoadTexture("Resources/PizzaClickerBackground.png"),
		DT:              rl.GetFrameTime(),
		Timer:           0,
	}

	enity := enity.Enity{
		BigPizza: pizza.New(),
		Mouse:    mouse.New(),
		TW:       teen_worker.New(&upgradeSheet),
		AW:       adultworker.New(&upgradeSheet),
		ZTW:      zombetw.New(&upgradeSheet),
		ZAW:      zombeaw.New(&upgradeSheet),
		Necro:    necronomicon.New(&upgradeSheet),
		TextBox:  textboxs.New(),
	}

	rl.HideCursor()

	for !rl.WindowShouldClose() {
		state.DT = rl.GetFrameTime()
		state.Timer += state.DT

		if rl.IsKeyPressed(rl.KeyF) {
			state.FPScap = !state.FPScap
			if state.FPScap {
				rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
			} else {
				rl.SetTargetFPS(0)
			}
		}

		upgradeColision(enity.TW.Rec, enity.Mouse.Position, enity.TW, state)
		upgradeColision(enity.AW.Rec, enity.Mouse.Position, enity.AW, state)

		enity.TextBox.Update()

		enity.TW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.TW.Rec)
		enity.AW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.AW.Rec)

		if enity.Necro.IsBought {
			upgradeColision(enity.ZTW.Rec, enity.Mouse.Position, enity.ZTW, state)
			upgradeColision(enity.ZAW.Rec, enity.Mouse.Position, enity.ZAW, state)
			enity.ZTW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.ZTW.Rec)
			enity.ZAW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.ZAW.Rec)
		}

		enity.Necro.Update(state, enity.Mouse.Position)

		if state.Timer >= 1 {
			enity.TW.Update(state)
			enity.AW.Update(state)
			enity.ZTW.Update(state)
			enity.ZAW.Update(state)

			state.Timer = 0
		}

		enity.Mouse.Update()
		pizzaColision(enity.BigPizza.Crec, enity.Mouse.Position, enity.BigPizza, state)
		enity.BigPizza.Animation()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(state.Background, 0, 0, rl.White)
		enity.BigPizza.Draw(state.PizzaCount)

		enity.TextBox.Draw()
		enity.TW.Draw()
		enity.AW.Draw()
		enity.Necro.Draw()

		if enity.Necro.IsBought {
			enity.ZTW.Draw()
			enity.ZAW.Draw()
		}

		enity.Mouse.Draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func pizzaColision(prec rl.Rectangle, mouse rl.Vector2, p *pizza.Pizza, state *state.State) {
	if rl.CheckCollisionPointRec(mouse, prec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			state.PizzaCount += 1
			state.TotalPizzaCount += 1
			p.IsPizzaClicked = true
		}
	}
}

func upgradeColision(urec rl.Rectangle, mouse rl.Vector2, u upgrades.Upgradable, state *state.State) {
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
