package enity

import (
	"github.com/NickDeChip/pizzaClicker/mouse"
	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/NickDeChip/pizzaClicker/setting/stats"
	"github.com/NickDeChip/pizzaClicker/upgrades/adult_worker"
	"github.com/NickDeChip/pizzaClicker/upgrades/necronomicon"
	"github.com/NickDeChip/pizzaClicker/upgrades/pizza_oven"
	"github.com/NickDeChip/pizzaClicker/upgrades/teen_worker"
	"github.com/NickDeChip/pizzaClicker/upgrades/text_boxs"
	"github.com/NickDeChip/pizzaClicker/upgrades/zombe_aw"
	"github.com/NickDeChip/pizzaClicker/upgrades/zombe_tw"
)

type Enity struct {
	BigPizza *pizza.Pizza
	Mouse    *mouse.Mouse
	TW       *teen_worker.Worker
	AW       *adultworker.Worker
	ZTW      *zombetw.Worker
	ZAW      *zombeaw.Worker
	Necro    *necronomicon.Necro
	PO       *pizzaoven.PizzaOven
	TextBox  *textboxs.TextBox
	Stats    *stats.Stats
}
