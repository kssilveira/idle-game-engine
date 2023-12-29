package data

type Resource struct {
	Name     string  `json:",omitempty"`
	Type     string  `json:",omitempty"`
	IsHidden bool    `json:",omitempty"`
	Count    float64 `json:",omitempty"`
	Cap      float64 `json:",omitempty"`

	Producers []Resource `json:",omitempty"`
	// Count += Producer.Count * Factor * elapsedTime
	Factor float64 `json:",omitempty"`
	// Count += floor(Producer.Count) * Factor * elapsedTime
	ProductionFloor bool `json:",omitempty"`
	// Count += (Producer.Count > 0 ? 1 : 0) * Factor * elapsedTime
	ProductionBoolean bool `json:",omitempty"`
	// Count = StartCount + Producer.Count * Factor
	StartCount float64 `json:",omitempty"`
	// Count = Producer.Count * Factor
	StartCountFromZero bool `json:",omitempty"`
	// Count = StartCount + (Producer.Count * Factor) % ProductionModulus
	ProductionModulus int `json:",omitempty"`
	// Count = StartCount if (Producer.Count * Factor) % ProductionModulus == ProductionModulusEquals else 0
	ProductionModulusEquals int `json:",omitempty"`

	// Production *= 1 + sum(Bonus)
	Bonus []Resource `json:",omitempty"`
	// Production *= sum(Bonus)
	BonusStartsFromZero bool `json:",omitempty"`
	// Production *= 1 + product(Bonus)
	BonusIsMultiplicative bool `json:",omitempty"`

	// negative production reduces consumers
	ProductionOnGone bool       `json:",omitempty"`
	OnGone           []Resource `json:",omitempty"`

	// Cap = CapResource.Count
	CapResource string `json:",omitempty"`
	// Count = ResetResource.Count
	ResetResource string `json:",omitempty"`

	// Cost = Count * pow(CostExponentBase, add.Count)
	CostExponentBase float64 `json:",omitempty"`

	ProducerAction string `json:",omitempty"`

	// generated production formula
	Formula string `json:",omitempty"`
}

type Action struct {
	Name       string     `json:",omitempty"`
	Type       string     `json:",omitempty"`
	UnlockedBy string     `json:",omitempty"`
	LockedBy   string     `json:",omitempty"`
	Costs      []Resource `json:",omitempty"`
	Adds       []Resource `json:",omitempty"`
	IsHidden   bool       `json:",omitempty"`

	// Cost = Count * pow(CostExponentBase, add.Count)
	CostExponentBase float64 `json:",omitempty"`
	// Cost = Count * pow(rate(CostExponentBaseResource), add.Count)
	CostExponentBaseResource Resource `json:",omitempty"`

	// extra fields for convenience
	Producers     []Resource `json:",omitempty"`
	Bonus         []Resource `json:",omitempty"`
	ResetResource string     `json:",omitempty"`
}

type ParsedInput struct {
	IsSkip   bool
	IsCreate bool
	IsMax    bool
	IsReset  bool
	Index    int
	Action   Action
}

func (r *Resource) Add(add Resource) {
	r.Cap += add.Cap
	r.Count += add.Count
	if r.Count > r.Cap && r.Cap >= 0 {
		r.Count = r.Cap
	}
	if r.Count < 0 {
		r.Count = 0
	}
}
