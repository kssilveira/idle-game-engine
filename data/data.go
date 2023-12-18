package data

type Resource struct {
	Name     string  `json:",omitempty"`
	Type     string  `json:",omitempty"`
	IsHidden bool    `json:",omitempty"`
	Count    float64 `json:",omitempty"`
	Capacity float64 `json:",omitempty"`

	Producers []Resource `json:",omitempty"`
	// Count += Producer.Count * Factor * elapsedTime
	Factor float64 `json:",omitempty"`
	// Count += floor(Producer.Count) * Factor * elapsedTime
	ProductionFloor bool `json:",omitempty"`
	// Count += (Producer.Count > 0 ? 1 : 0) * Factor * elapsedTime
	ProductionBoolean bool `json:",omitempty"`
	// Count = StartCount + Producer.Count * Factor
	StartCount float64 `json:",omitempty"`
	// Count = StartCount + (Producer.Count * Factor) % ProductionModulus
	ProductionModulus int `json:",omitempty"`
	// Count = StartCount if (Producer.Count * Factor) % ProductionModulus == ProductionModulusEquals else 0
	ProductionModulusEquals int `json:",omitempty"`

	// production *= 1 + sum(bonus)
	Bonus []Resource `json:",omitempty"`
	// production *= sum(bonus)
	BonusStartsFromZero bool `json:",omitempty"`

	// negative production reduces consumers
	ProductionOnGone bool       `json:",omitempty"`
	OnGone           []Resource `json:",omitempty"`

	CapacityProducers []Resource `json:",omitempty"`
	// Capacity = StartCapacity + CapacityProducer.Count * Factor
	StartCapacity float64 `json:",omitempty"`

	// cost = Count * pow(CostExponentBase, add.Count)
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
	IsHidden  bool       `json:",omitempty"`
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
	r.Count += add.Count
	if r.Count > r.Capacity && r.Capacity >= 0 {
		r.Count = r.Capacity
	}
	if r.Count < 0 {
		r.Count = 0
	}
}
