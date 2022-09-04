package upgrades

type Upgradable interface {
	GetCost() float64
	IncrementCost()
	IncrementCount()
}
