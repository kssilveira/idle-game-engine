package data

type Resource struct {
	Name     string  `json:",omitempty"`
	Type     string  `json:",omitempty"`
	IsHidden bool    `json:",omitempty"`
	Quantity float64 `json:",omitempty"`
	Capacity float64 `json:",omitempty"`

	Producers []Resource `json:",omitempty"`
	// Quantity += Producer.Quantity * Factor * elapsedTime
	Factor float64 `json:",omitempty"`
	// Quantity += floor(Producer.Quantity) * Factor * elapsedTime
	ProductionFloor bool `json:",omitempty"`
	// Quantity += (Producer.Quantity > 0 ? 1 : 0) * Factor * elapsedTime
	ProductionBoolean bool `json:",omitempty"`
	// Quantity = StartQuantity + Producer.Quantity * Factor
	StartQuantity float64 `json:",omitempty"`
	// Quantity = StartQuantity + (Producer.Quantity * Factor) % ProductionModulus
	ProductionModulus int `json:",omitempty"`
	// Quantity = StartQuantity if (Producer.Quantity * Factor) % ProductionModulus == ProductionModulusEquals else 0
	ProductionModulusEquals int `json:",omitempty"`

	// production *= 1 + sum(bonus)
	Bonus []Resource `json:",omitempty"`
	// production *= sum(bonus)
	BonusStartsFromZero bool `json:",omitempty"`

	// negative production reduces consumers
	ProductionOnGone bool       `json:",omitempty"`
	OnGone           []Resource `json:",omitempty"`

	CapacityProducers []Resource `json:",omitempty"`
	// Capacity = StartCapacity + CapacityProducer.Quantity * Factor
	StartCapacity float64 `json:",omitempty"`

	// cost = Quantity * pow(CostExponentBase, add.Quantity)
	CostExponentBase float64 `json:",omitempty"`

	ProducerAction string `json:",omitempty"`
}

type Action struct {
	Name       string     `json:",omitempty"`
	Type       string     `json:",omitempty"`
	UnlockedBy string     `json:",omitempty"`
	LockedBy   string     `json:",omitempty"`
	Costs      []Resource `json:",omitempty"`
	Adds       []Resource `json:",omitempty"`
	// Producers of the corresponding Resource
	Producers []Resource `json:",omitempty"`
}

type ParsedInput struct {
	IsSkip   bool
	IsCreate bool
	IsMax    bool
	Index    int
	Action   Action
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
