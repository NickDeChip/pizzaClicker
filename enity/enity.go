package enity

import (
	"github.com/NickDeChip/pizzaClicker/mouse"
	"github.com/NickDeChip/pizzaClicker/pizza"
	"github.com/NickDeChip/pizzaClicker/upgrades/adult_worker"
	"github.com/NickDeChip/pizzaClicker/upgrades/teen_worker"
)

type Enity struct {
	BigPizza *pizza.Pizza
	Mouse    *mouse.Mouse
	TW       *teen_worker.Worker
	AW       *adultworker.Worker
}
