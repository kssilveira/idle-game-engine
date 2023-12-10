package data

type Resource struct {
	Name     string
	Type     string
	IsHidden bool
	Quantity float64
	Capacity float64

	Producers []Resource
	// Quantity += Producer.Quantity * ProductionFactor * elapsedTime
	ProductionFactor float64
	// Quantity += Producer.Quantity * ProductionFactor * elapsedTime * ProductionResourceFactor.Quantity
	ProductionResourceFactor string
	// Quantity += floor(Producer.Quantity) * ProductionFactor * elapsedTime
	ProductionFloor bool
	// Quantity += (Producer.Quantity > 0 ? 1 : 0) * ProductionFactor * elapsedTime
	ProductionBoolean bool
	// Quantity = StartQuantity + Producer.Quantity * ProductionFactor
	StartQuantity float64
	// Quantity = StartQuantity + (Producer.Quantity * ProductionFactor) % ProductionModulus
	ProductionModulus int
	// Quantity = StartQuantity if (Producer.Quantity * ProductionFactor) % ProductionModulus == ProductionModulusEquals else 0
	ProductionModulusEquals int

	// production *= 1 + bonus
	ProductionBonus []Resource

	// negative production reduces consumers
	ProductionOnGone bool
	OnGone           []Resource

	CapacityProducers []Resource
	// Capacity = StartCapacity + CapacityProducer.Quantity * ProductionFactor
	StartCapacity float64

	// cost = Quantity * pow(CostExponentBase, add.Quantity)
	CostExponentBase float64

	ProducerAction string
}

type Action struct {
	Name       string
	Type       string
	UnlockedBy string
	LockedBy   string
	Costs      []Resource
	Adds       []Resource
}

type ParsedInput struct {
	IsSkip bool
	IsMake bool
	Action Action
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
