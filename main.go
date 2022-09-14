package main

import (
	"github.com/NickDeChip/pizzaClicker/enity"
	"github.com/NickDeChip/pizzaClicker/mouse"
	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/NickDeChip/pizzaClicker/setting/stats"
	"github.com/NickDeChip/pizzaClicker/state"
	"github.com/NickDeChip/pizzaClicker/upgrades"
	"github.com/NickDeChip/pizzaClicker/upgrades/adult_worker"
	"github.com/NickDeChip/pizzaClicker/upgrades/aprons"
	"github.com/NickDeChip/pizzaClicker/upgrades/necronomicon"
	"github.com/NickDeChip/pizzaClicker/upgrades/pizza_oven"
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
		PizzaCount:      10000,
		TotalPizzaCount: 10000,
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
		PO:       pizzaoven.New(&upgradeSheet),
		Aprons:   aprons.New(&upgradeSheet),
		TextBox:  textboxs.New(),
		Stats:    stats.New(&upgradeSheet),
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

		upgrades.UpgradeColision(enity.TW.Rec, enity.Mouse.Position, enity.TW, state)
		upgrades.UpgradeColision(enity.AW.Rec, enity.Mouse.Position, enity.AW, state)

		enity.TextBox.Update()

		enity.Stats.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.Stats.Rec)

		enity.TW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.TW.Rec)
		enity.AW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.AW.Rec)

		if enity.Necro.IsBought {
			upgrades.UpgradeColision(enity.ZTW.Rec, enity.Mouse.Position, enity.ZTW, state)
			upgrades.UpgradeColision(enity.ZAW.Rec, enity.Mouse.Position, enity.ZAW, state)
			enity.ZTW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.ZTW.Rec)
			enity.ZAW.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.ZAW.Rec)
		}

		if !enity.Necro.IsBought && enity.Necro.DisplayUpgrade {
			enity.Necro.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.Necro.Rec)
		}

		if !enity.PO.IsBought && enity.PO.DisplayUpgrade {
			enity.PO.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.PO.Rec)
		}

		if !enity.Aprons.IsApronBought && enity.Aprons.DisplayUpgradeApron {
			enity.Aprons.ShowText = enity.TextBox.CollisionCheck(enity.Mouse.Position, enity.Aprons.Rec)
		}

		enity.PO.Update(state, enity.Mouse.Position)
		enity.Aprons.Update(state, enity.Mouse.Position)
		enity.Necro.Update(state, enity.Mouse.Position)
		enity.BigPizza.Update(enity.PO)

		if state.Timer >= 1 {
			enity.TW.Update(state, enity.Aprons)
			enity.AW.Update(state, enity.Aprons)
			enity.ZTW.Update(state)
			enity.ZAW.Update(state)

			state.Timer = 0
		}

		enity.Mouse.Update()
		enity.BigPizza.PizzaColision(enity.BigPizza.Crec, enity.Mouse.Position, state)
		enity.BigPizza.Animation()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(state.Background, 0, 0, rl.White)
		enity.BigPizza.Draw(state.PizzaCount)

		enity.TextBox.Draw()
		enity.Stats.Draw(state, *enity.BigPizza)
		enity.TW.Draw()
		enity.AW.Draw()

		enity.PO.Draw()
		enity.Aprons.Draw()
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
