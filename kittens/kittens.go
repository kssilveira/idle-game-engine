package kittens

import (
	"fmt"
	"log"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

func NewGame(now game.Now) *game.Game {
	g := game.NewGame(now())
	g.AddResources([]data.Resource{{
		Name: "day", Type: "Calendar", IsHidden: true, Quantity: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.5,
		}},
	}, {
		Name: "year", Type: "Calendar", Quantity: 1, StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.0025, ProductionFloor: true,
		}},
	}, {
		Name: "Spring", Type: "Calendar", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Summer", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Autumn", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 2, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Winter", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 3, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "day_of_year", Type: "Calendar", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 400, ProductionModulusEquals: -1,
		Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 1, ProductionFloor: true,
		}},
	}, {
		Name: "catnip", Type: "Resource", StartCapacity: 5000,
		Producers: []data.Resource{{
			Name: "Catnip Field", ProductionFactor: 0.63 * (1 + 0.50), ProductionResourceFactor: "Spring",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63, ProductionResourceFactor: "Summer",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63, ProductionResourceFactor: "Autumn",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63 * (1 - 0.75), ProductionResourceFactor: "Winter",
		}, {
			Name: "kitten", ProductionFactor: -4.25, ProductionFloor: true, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.015,
			}},
		}, {
			Name: "woodcutter", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.015,
			}},
		}, {
			Name: "scholar", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.015,
			}},
		}, {
			Name: "farmer", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.015,
			}},
		}, {
			Name: "hunter", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.015,
			}},
		}, {
			Name: "miner", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.015,
			}},
		}, {
			Name: "farmer", ProductionFactor: 5, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mineral Hoes", ProductionFactor: 0.5,
			}, {
				Name: "Iron Hoes", ProductionFactor: 0.3,
			}},
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 5000,
		}},
	}, {
		Name: "wood", Type: "Resource", StartCapacity: 200,
		Producers: []data.Resource{{
			Name: "woodcutter", ProductionFactor: 0.09, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mineral Axe", ProductionFactor: 0.7,
			}, {
				Name: "Iron Axe", ProductionFactor: 0.5,
			}},
		}, {
			Name: "Active Smelter", ProductionFactor: -0.25,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 200,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 262.5,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}},
	}, {
		Name: "science", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "scholar", ProductionFactor: 0.175, ProductionResourceFactor: "happiness",
		}},
		ProductionBonus: []data.Resource{{
			Name: "Library", ProductionFactor: 0.1,
		}, {
			Name: "Academy", ProductionFactor: 0.2,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Library", ProductionFactor: 250,
		}, {
			Name: "Academy", ProductionFactor: 500,
		}},
	}, {
		Name: "catpower", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "hunter", ProductionFactor: 0.3, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Bolas", ProductionFactor: 1,
			}, {
				Name: "Hunting Armor", ProductionFactor: 2,
			}, {
				Name: "Composite Bow", ProductionFactor: 0.5,
			}},
		}},
		CapacityProducers: []data.Resource{{
			Name: "Hut", ProductionFactor: 75,
		}, {
			Name: "Log House", ProductionFactor: 50,
		}},
	}, {
		Name: "minerals", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "miner", ProductionFactor: 0.25, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mine", ProductionFactor: 0.2,
			}},
		}, {
			Name: "Active Smelter", ProductionFactor: -0.5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 250,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 350,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}},
	}, {
		Name: "iron", Type: "Resource", StartCapacity: 50,
		Producers: []data.Resource{{
			Name: "Active Smelter", ProductionFactor: 0.1,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 50,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 43.75,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}},
	}, {
		Name: "gold", Type: "Resource", Capacity: 20,
	}, {
		Name: "kitten", Type: "Resource", Capacity: 0,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.05,
		}},
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}},
	}, {
		Name: "all kittens", Type: "Resource", IsHidden: true, Capacity: -1, StartQuantity: 1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: -1,
		}, {
			Name: "kitten", ProductionFactor: 1, ProductionFloor: true,
		}, {
			Name: "woodcutter", ProductionFactor: 1,
		}, {
			Name: "scholar", ProductionFactor: 1,
		}, {
			Name: "farmer", ProductionFactor: 1,
		}, {
			Name: "hunter", ProductionFactor: 1,
		}, {
			Name: "miner", ProductionFactor: 1,
		}},
	}, {
		Name: "furs", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.05,
		}},
	}, {
		Name: "ivory", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.035,
		}},
	}, {
		Name: "spice", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.005,
		}},
	}, {
		Name: "beam", Type: "Resource", Capacity: -1,
		ProducerAction: "@beam",
	}, {
		Name: "slab", Type: "Resource", Capacity: -1,
		ProducerAction: "@slab",
	}, {
		Name: "unicorns", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "Unic. Pasture", ProductionFactor: 0.005,
		}},
	}, {
		Name: "blueprint", Type: "Resource", Capacity: -1,
	}, {
		Name: "gone kitten", Type: "Resource", Capacity: -1,
	}, {
		Name: "Catnip Field", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Hut", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Library", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Barn", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Mine", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Workshop", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Smelter", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Idle Smelter", Type: "Bonfire", Capacity: -1, StartQuantity: 1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: -1,
		}, {
			Name: "Smelter", ProductionFactor: 1,
		}, {
			Name: "Active Smelter", ProductionFactor: -1,
		}},
	}, {
		Name: "Active Smelter", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Pasture", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Unic. Pasture", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Academy", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Warehouse", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Log House", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "woodcutter", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "scholar", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "farmer", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "hunter", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "miner", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "happiness", Type: "Village", StartQuantity: 1.1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.02,
		}, {
			Name: "ivory", ProductionFactor: 0.1, ProductionBoolean: true,
		}, {
			Name: "furs", ProductionFactor: 0.1, ProductionBoolean: true,
		}, {
			Name: "spice", ProductionFactor: 0.1, ProductionBoolean: true,
		}, {
			Name: "unicorns", ProductionFactor: 0.1, ProductionBoolean: true,
		}},
	}, {
		Name: "Calendar", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Agriculture", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Archery", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mining", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Animal Husbandry", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Metal Working", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Civil Service", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mathematics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Construction", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Currency", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Celestial Mechanics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Engineering", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mineral Hoes", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Iron Hoes", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mineral Axe", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Iron Axe", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Expanded Barns", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Reinforced Barns", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Bolas", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Hunting Armor", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Reinforced Saw", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Composite Bow", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Catnip Enrichment", Type: "Workshop", IsHidden: true, Capacity: 1,
	}})
	g.AddActions([]game.Action{{
		Name: "Gather catnip", Type: "Bonfire", LockedBy: "Catnip Field",
		Adds: []data.Resource{{Name: "catnip", Quantity: 1}},
	}, {
		Name: "Refine catnip", Type: "Bonfire", UnlockedBy: "catnip", LockedBy: "woodcutter",
		Costs: []data.Resource{{Name: "catnip", Quantity: 100}},
		Adds:  []data.Resource{{Name: "wood", Quantity: 1}},
	}, {
		Name: "Catnip Field", Type: "Bonfire", UnlockedBy: "catnip",
		Costs: []data.Resource{{Name: "catnip", Quantity: 10, CostExponentBase: 1.12}},
		Adds:  []data.Resource{{Name: "Catnip Field", Quantity: 1}},
	}, {
		Name: "Hut", Type: "Bonfire", UnlockedBy: "wood",
		Costs: []data.Resource{{Name: "wood", Quantity: 5, CostExponentBase: 2.5}},
		Adds: []data.Resource{{
			Name: "Hut", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 2,
		}},
	}, {
		Name: "Library", Type: "Bonfire", UnlockedBy: "wood",
		Costs: []data.Resource{{Name: "wood", Quantity: 25, CostExponentBase: 1.15}},
		Adds:  []data.Resource{{Name: "Library", Quantity: 1}},
	}, {
		Name: "Barn", Type: "Bonfire", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "wood", Quantity: 50, CostExponentBase: 1.75}},
		Adds:  []data.Resource{{Name: "Barn", Quantity: 1}},
	}, {
		Name: "Mine", Type: "Bonfire", UnlockedBy: "Mining",
		Costs: []data.Resource{{Name: "wood", Quantity: 100, CostExponentBase: 1.15}},
		Adds:  []data.Resource{{Name: "Mine", Quantity: 1}},
	}, {
		Name: "Workshop", Type: "Bonfire", UnlockedBy: "Mining",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 400, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Workshop", Quantity: 1}},
	}, {
		Name: "Smelter", Type: "Bonfire", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{Name: "minerals", Quantity: 200, CostExponentBase: 1.15}},
		Adds:  []data.Resource{{Name: "Smelter", Quantity: 1}},
	}, {
		Name: "Active Smelter", Type: "Bonfire", UnlockedBy: "Smelter",
		Costs: []data.Resource{{Name: "Idle Smelter", Quantity: 1}},
		Adds:  []data.Resource{{Name: "Active Smelter", Quantity: 1}},
	}, {
		Name: "Pasture", Type: "Bonfire", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "wood", Quantity: 10, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Pasture", Quantity: 1}},
	}, {
		Name: "Unic. Pasture", Type: "Bonfire", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "unicorns", Quantity: 2, CostExponentBase: 1.75}},
		Adds:  []data.Resource{{Name: "Unic. Pasture", Quantity: 1}},
	}, {
		Name: "Academy", Type: "Bonfire", UnlockedBy: "Mathematics",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 70, CostExponentBase: 1.15,
		}, {
			Name: "science", Quantity: 100, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Academy", Quantity: 1}},
	}, {
		Name: "Warehouse", Type: "Bonfire", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "beam", Quantity: 1.5, CostExponentBase: 1.15,
		}, {
			Name: "slab", Quantity: 2, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Warehouse", Quantity: 1}},
	}, {
		Name: "Log House", Type: "Bonfire", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{
			Name: "Log House", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "woodcutter", Type: "Village", UnlockedBy: "Hut",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "woodcutter", Quantity: 1}},
	}, {
		Name: "scholar", Type: "Village", UnlockedBy: "Library",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "scholar", Quantity: 1}},
	}, {
		Name: "farmer", Type: "Village", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "farmer", Quantity: 1}},
	}, {
		Name: "hunter", Type: "Village", UnlockedBy: "Archery",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "hunter", Quantity: 1}},
	}, {
		Name: "miner", Type: "Village", UnlockedBy: "Mine",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "miner", Quantity: 1}},
	}, {
		Name: "Send hunters", Type: "Village", UnlockedBy: "Archery",
		Costs: []data.Resource{{Name: "catpower", Quantity: 100}},
		Adds: []data.Resource{{
			Name: "furs", Quantity: 39.5,
		}, {
			Name: "ivory", Quantity: 10.78,
		}, {
			Name: "unicorns", Quantity: 0.05,
		}},
	}, {
		Name: "@beam", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "wood", Quantity: 175}},
		Adds: []data.Resource{{
			Name: "beam", Quantity: 1,
			ProductionBonus: []data.Resource{{Name: "Workshop", ProductionFactor: 0.06}},
		}},
	}, {
		Name: "@slab", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "minerals", Quantity: 250}},
		Adds: []data.Resource{{
			Name: "slab", Quantity: 1,
			ProductionBonus: []data.Resource{{Name: "Workshop", ProductionFactor: 0.06}},
		}},
	}, {
		Name: "Lizards", Type: "Trade", UnlockedBy: "Archery",
		Costs: []data.Resource{{
			Name: "catpower", Quantity: 50,
		}, {
			Name: "gold", Quantity: 15,
		}, {
			Name: "minerals", Quantity: 1000,
		}},
		Adds: []data.Resource{{
			Name: "wood", Quantity: 500,
		}, {
			Name: "blueprint", Quantity: 0.1,
		}, {
			Name: "spice", Quantity: 8.75,
		}},
	}, {
		Name: "Calendar", Type: "Science", UnlockedBy: "Library", LockedBy: "Calendar",
		Costs: []data.Resource{{Name: "science", Quantity: 30}},
		Adds:  []data.Resource{{Name: "Calendar", Quantity: 1}},
	}, {
		Name: "Agriculture", Type: "Science", UnlockedBy: "Calendar", LockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "science", Quantity: 100}},
		Adds:  []data.Resource{{Name: "Agriculture", Quantity: 1}},
	}, {
		Name: "Archery", Type: "Science", UnlockedBy: "Calendar", LockedBy: "Archery",
		Costs: []data.Resource{{Name: "science", Quantity: 300}},
		Adds:  []data.Resource{{Name: "Archery", Quantity: 1}},
	}, {
		Name: "Mining", Type: "Science", UnlockedBy: "Calendar", LockedBy: "Mining",
		Costs: []data.Resource{{Name: "science", Quantity: 500}},
		Adds:  []data.Resource{{Name: "Mining", Quantity: 1}},
	}, {
		Name: "Animal Husbandry", Type: "Science", UnlockedBy: "Archery", LockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Quantity: 500}},
		Adds:  []data.Resource{{Name: "Animal Husbandry", Quantity: 1}},
	}, {
		Name: "Metal Working", Type: "Science", UnlockedBy: "Mining", LockedBy: "Metal Working",
		Costs: []data.Resource{{Name: "science", Quantity: 900}},
		Adds:  []data.Resource{{Name: "Metal Working", Quantity: 1}},
	}, {
		Name: "Civil Service", Type: "Science", UnlockedBy: "Animal Husbandry", LockedBy: "Civil Service",
		Costs: []data.Resource{{Name: "science", Quantity: 1500}},
		Adds:  []data.Resource{{Name: "Civil Service", Quantity: 1}},
	}, {
		Name: "Mathematics", Type: "Science", UnlockedBy: "Animal Husbandry", LockedBy: "Mathematics",
		Costs: []data.Resource{{Name: "science", Quantity: 1000}},
		Adds:  []data.Resource{{Name: "Mathematics", Quantity: 1}},
	}, {
		Name: "Construction", Type: "Science", UnlockedBy: "Animal Husbandry", LockedBy: "Construction",
		Costs: []data.Resource{{Name: "science", Quantity: 1300}},
		Adds:  []data.Resource{{Name: "Construction", Quantity: 1}},
	}, {
		Name: "Currency", Type: "Science", UnlockedBy: "Animal Husbandry", LockedBy: "Currency",
		Costs: []data.Resource{{Name: "science", Quantity: 2200}},
		Adds:  []data.Resource{{Name: "Currency", Quantity: 1}},
	}, {
		Name: "Celestial Mechanics", Type: "Science", UnlockedBy: "Mathematics", LockedBy: "Celestial Mechanics",
		Costs: []data.Resource{{Name: "science", Quantity: 250}},
		Adds:  []data.Resource{{Name: "Celestial Mechanics", Quantity: 1}},
	}, {
		Name: "Engineering", Type: "Science", UnlockedBy: "Construction", LockedBy: "Engineering",
		Costs: []data.Resource{{Name: "science", Quantity: 1500}},
		Adds:  []data.Resource{{Name: "Engineering", Quantity: 1}},
	}, {
		Name: "Mineral Hoes", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Mineral Hoes",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 275,
		}, {
			Name: "science", Quantity: 100,
		}},
		Adds: []data.Resource{{Name: "Mineral Hoes", Quantity: 1}},
	}, {
		Name: "Iron Hoes", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Iron Hoes",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 25,
		}, {
			Name: "science", Quantity: 200,
		}},
		Adds: []data.Resource{{Name: "Iron Hoes", Quantity: 1}},
	}, {
		Name: "Mineral Axe", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Mineral Axe",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 500,
		}, {
			Name: "science", Quantity: 100,
		}},
		Adds: []data.Resource{{Name: "Mineral Axe", Quantity: 1}},
	}, {
		Name: "Iron Axe", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Iron Axe",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 50,
		}, {
			Name: "science", Quantity: 200,
		}},
		Adds: []data.Resource{{Name: "Iron Axe", Quantity: 1}},
	}, {
		Name: "Expanded Barns", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Expanded Barns",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 1000,
		}, {
			Name: "minerals", Quantity: 750,
		}, {
			Name: "iron", Quantity: 50,
		}, {
			Name: "science", Quantity: 500,
		}},
		Adds: []data.Resource{{Name: "Expanded Barns", Quantity: 1}},
	}, {
		Name: "Reinforced Barns", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Reinforced Barns",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 100,
		}, {
			Name: "science", Quantity: 800,
		}, {
			Name: "beam", Quantity: 25,
		}, {
			Name: "slab", Quantity: 10,
		}},
		Adds: []data.Resource{{Name: "Reinforced Barns", Quantity: 1}},
	}, {
		Name: "Bolas", Type: "Workshop", UnlockedBy: "Workshop", LockedBy: "Bolas",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 50,
		}, {
			Name: "minerals", Quantity: 250,
		}, {
			Name: "science", Quantity: 1000,
		}},
		Adds: []data.Resource{{Name: "Bolas", Quantity: 1}},
	}, {
		Name: "Hunting Armor", Type: "Workshop", UnlockedBy: "Metal Working", LockedBy: "Hunting Armor",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 750,
		}, {
			Name: "science", Quantity: 2000,
		}},
		Adds: []data.Resource{{Name: "Hunting Armor", Quantity: 1}},
	}, {
		Name: "Reinforced Saw", Type: "Workshop", UnlockedBy: "Construction", LockedBy: "Reinforced Saw",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 1000,
		}, {
			Name: "science", Quantity: 2500,
		}},
		Adds: []data.Resource{{Name: "Reinforced Saw", Quantity: 1}},
	}, {
		Name: "Composite Bow", Type: "Workshop", UnlockedBy: "Construction", LockedBy: "Composite Bow",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 200,
		}, {
			Name: "iron", Quantity: 100,
		}, {
			Name: "science", Quantity: 500,
		}},
		Adds: []data.Resource{{Name: "Composite Bow", Quantity: 1}},
	}, {
		Name: "Catnip Enrichment", Type: "Workshop", UnlockedBy: "Construction", LockedBy: "Catnip Enrichment",
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 5000,
		}, {
			Name: "science", Quantity: 500,
		}},
		Adds: []data.Resource{{Name: "Catnip Enrichment", Quantity: 1}},
	}})
	return g
}

const (
	gather = iota
	refine
	field
	hut
	library
	barn
	mine
	workshop
	smelter
	activesmelter
	pasture
	unicpasture
	academy
	warehouse
	loghouse
	woodcutter
	scholar
	farmer
	hunter
	miner
	sendhunters
	beam
	slab
	lizards
	calendar
	agriculture
	archery
	mining
	animalhusbandry
	metalworking
	civilservice
	mathematics
	celestialmechanics
	engineering
	construction
	currency
	mineralhoes
	ironhoes
	mineralaxe
	ironaxe
	expandedbarns
	reinforcedbarns
	bolas
	huntingarmor
	reinforcedsaw
	compositebow
	catnipenrichment
)

const (
	delta = 100
)

const (
	_ = iota * delta
	s
	m
)

func Solve(input chan string, sleepMS int) {
	for _, one := range []struct {
		cmds  []int
		count int
	}{
		{[]int{gather}, 10},
		{[]int{field}, 1},
		{[]int{s + field, field}, 58},
		{[]int{s + refine, refine}, 5},
		{[]int{hut, s + woodcutter, woodcutter}, 1},
		{[]int{s + library, library, s + scholar, scholar}, 1},
		{[]int{s + library, library}, 14},
		{[]int{s + calendar, calendar}, 1},
		{[]int{s + agriculture, agriculture}, 1},

		{[]int{s + barn, barn}, 6},
		{[]int{s + field, field}, 25},
		{[]int{s + library, library}, 15},
		{[]int{s + hut, hut, s + farmer, farmer}, 4},
		{[]int{s + farmer, farmer}, 2},

		{[]int{s + archery, archery}, 1},
		{[]int{hunter}, 1}, // hut

		{[]int{s + animalhusbandry, animalhusbandry}, 1},
		{[]int{s + pasture, pasture}, 40},
		{[]int{s + sendhunters, sendhunters}, 40},
		{[]int{unicpasture}, 1},
		{[]int{s + unicpasture, unicpasture}, 10},
		{[]int{s + civilservice, civilservice}, 1},

		{[]int{s + mathematics, mathematics}, 1},
		{[]int{s + celestialmechanics, celestialmechanics}, 1},
		{[]int{s + engineering, engineering}, 1},

		{[]int{s + construction, construction}, 1},
		{[]int{s + currency, currency}, 1},
		{[]int{s + catnipenrichment, catnipenrichment}, 1},

		{[]int{s + mining, mining}, 1},
		{[]int{s + mine, mine}, 20},
		{[]int{miner}, 1}, // hut

		{[]int{s + academy, academy}, 30},
		{[]int{s + workshop, workshop}, 20},
		{[]int{s + mineralhoes, mineralhoes}, 1},
		{[]int{s + mineralaxe, mineralaxe}, 1},
		{[]int{s + bolas, bolas}, 1},
		{[]int{woodcutter}, 1}, // hut
		{[]int{s + loghouse, loghouse, s + woodcutter, woodcutter}, 5},
		{[]int{s + loghouse, loghouse, s + farmer, farmer}, 20},

		{[]int{s + metalworking, metalworking}, 1},
		{[]int{s + smelter, smelter}, 20},
		{[]int{activesmelter}, 1},
		{[]int{s + ironhoes, ironhoes}, 1},
		{[]int{s + ironaxe, ironaxe}, 1},
		{[]int{s + compositebow, compositebow}, 1},
		{[]int{s + expandedbarns, expandedbarns}, 1},
		{[]int{
			s + reinforcedbarns, m + reinforcedbarns,
			s + reinforcedbarns, m + reinforcedbarns, reinforcedbarns}, 1},
		{[]int{s + warehouse, m + warehouse, warehouse}, 11},

		{[]int{s + barn, barn}, 10},
		{[]int{s + field, field}, 10},
		{[]int{s + hut, hut, s + farmer, farmer}, 10},
		{[]int{s + library, library}, 20},
		{[]int{s + mine, mine}, 20},
		{[]int{s + workshop, workshop}, 20},
		{[]int{s + smelter, smelter}, 20},
		{[]int{s + pasture, pasture}, 20},
		{[]int{s + academy, academy}, 20},
		{[]int{s + loghouse, loghouse, s + farmer, farmer}, 20},

		{[]int{s + huntingarmor, huntingarmor}, 1},
		{[]int{s + reinforcedsaw, m + reinforcedsaw, reinforcedsaw}, 1},
	} {
		for i := 0; i < one.count; i++ {
			for _, cmd := range one.cmds {
				input <- toInput(cmd)
				time.Sleep(time.Second * time.Duration(sleepMS) / 1000.)
			}
		}
	}
}

func toInput(cmd int) string {
	prefix := ""
	if cmd >= m {
		prefix = "m"
		cmd -= m
	}
	if cmd >= s {
		prefix = "s"
		cmd -= s
	}
	return fmt.Sprintf("%s%d", prefix, cmd)
}

func Graph(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf("digraph {\n")
	typeToShape := map[string]string{
		"Resource": "cylinder",
		"Bonfire":  "box3d",
		"Village":  "house",
		"Science":  "diamond",
		"Workshop": "hexagon",
		"Craft":    "cds",
		"Trade":    "cds",
	}
	nodes := map[string]bool{}
	edges := map[string]bool{}
	edgefn := func(from, to, color string) {
		edge(logger, nodes, edges, colors, from, to, color)
	}
	for _, r := range g.Resources {
		if r.Name == "happiness" {
			continue
		}
		last := ""
		for _, p := range r.Producers {
			if p.Name == "" || p.Name == "day" || p.Name == last {
				continue
			}
			last = p.Name
			if p.ProductionFactor < 0 {
				edgefn(r.Name, p.Name, "red")
			} else {
				edgefn(p.Name, r.Name, "green")
			}
			for _, b := range p.ProductionBonus {
				if p.ProductionFactor < 0 {
					edgefn(b.Name, p.Name, "red")
				} else {
					edgefn(b.Name, p.Name, "green")
				}
			}
		}
		for _, b := range r.ProductionBonus {
			edgefn(b.Name, r.Name, "green")
		}
		for _, p := range r.CapacityProducers {
			edgefn(p.Name, r.Name, "limegreen")
			for _, b := range p.ProductionBonus {
				edgefn(b.Name, p.Name, "green")
			}
		}
	}
	for _, a := range g.Actions {
		for _, c := range a.Costs {
			edgefn(c.Name, a.Name, "orange")
		}
		for _, add := range a.Adds {
			if a.Name == add.Name {
				continue
			}
			edgefn(a.Name, add.Name, "limegreen")
		}
		if a.UnlockedBy != "" {
			edgefn(a.UnlockedBy, a.Name, "blue")
		}
	}
	for _, r := range g.Resources {
		if r.Type == "Calendar" || r.Name == "gone kitten" || r.Name == "happiness" {
			continue
		}
		if !nodes[r.Name] {
			continue
		}
		logger.Printf(`  "%s" [shape="%s"];`+"\n", r.Name, typeToShape[r.Type])
	}
	for _, a := range g.Actions {
		if !nodes[a.Name] {
			continue
		}
		logger.Printf(`  "%s" [shape="%s"];`+"\n", a.Name, typeToShape[a.Type])
	}
	logger.Printf("}\n")
}

func GraphEdges(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf(`
digraph {
  node [label="" width=0 style=invis];
  { rank="same"; n0; n1; n2; n3; n4; n5; n6; n7; n8; n9; }
  n0 -> n1 [color="red" label="feeds"];
  n2 -> n3 [color="green" label="produces"];
  n4 -> n5 [color="orange" label="buys"];
  n6 -> n7 [color="limegreen" label="adds"];
  n8 -> n9 [color="blue" label="unlocks"];
}
`)
}

func GraphNodes(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf(`
digraph {
  "Resource" [shape="cylinder"];
  "Bonfire" [shape="box3d"];
  "Village" [shape="house"];
  "Science" [shape="diamond"];
  "Workshop" [shape="hexagon"];
  "Craft" [shape="cds"];
  "Trade" [shape="cds"];
}
`)
}

func edge(logger *log.Logger, nodes map[string]bool, edges map[string]bool, colors map[string]bool, from, to, color string) {
	if len(colors) > 0 && !colors[color] {
		return
	}
	key := fmt.Sprintf("%s+%s+%s", from, to, color)
	if edges[key] {
		return
	}
	edges[key] = true
	nodes[from] = true
	nodes[to] = true
	logger.Printf(`  "%s" -> "%s" [color="%s"];`+"\n", from, to, color)
}
