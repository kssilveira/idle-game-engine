package data

type Resource struct {
	Name     string
	Type     string
	IsHidden bool
	Quantity float64
	Capacity float64

	Producers []Resource

	// quantity += producer.Quantity * ProductionFactor * elapsedTime
	ProductionFactor float64
	// quantity += producer.Quantity * ProductionFactor * elapsedTime * ProductionResourceFactor.Quantity
	ProductionResourceFactor string
	// quantity += floor(producer.Quantity) * ProductionFactor * elapsedTime
	ProductionFloor bool
	// quantity = StartQuantity + producer.Quantity * ProductionFactor
	StartQuantity float64
	// quantity = StartQuantity + (producer.Quantity * ProductionFactor) % ProductionModulus
	ProductionModulus int
	// quantity = StartQuantity if (producer.Quantity * ProductionFactor) % ProductionModulus == ProductionModulusEquals else 0
	ProductionModulusEquals int

	// production *= 1 + bonus
	ProductionBonus []Resource

	OnGone []Resource

	// cost = Quantity * pow(CostExponentBase, add.Quantity)
	CostExponentBase float64
}

func (r *Resource) Add(add Resource) {
	r.Capacity += add.Capacity
	r.Quantity += add.Quantity
	if r.Quantity > r.Capacity && r.Capacity >= 0 {
		r.Quantity = r.Capacity
	}
	if r.Quantity < 0 {
		r.Quantity = 0
	}
}
