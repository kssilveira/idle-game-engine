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
			Name: "Catnip Field", ProductionFactor: 0.125 * 5 * (1 + 0.50), ProductionResourceFactor: "Spring",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.125 * 5, ProductionResourceFactor: "Summer",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.125 * 5, ProductionResourceFactor: "Autumn",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.125 * 5 * (1 - 0.75), ProductionResourceFactor: "Winter",
		}, {
			Name: "kitten", ProductionFactor: -4.25, ProductionFloor: true, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, {
			Name: "woodcutter", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, {
			Name: "scholar", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, {
			Name: "farmer", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, {
			Name: "hunter", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, {
			Name: "miner", ProductionFactor: -4.25, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, {
			Name: "farmer", ProductionFactor: 5, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mineral Hoes", ProductionFactor: 0.5,
			}, {
				Name: "Iron Hoes", ProductionFactor: 0.3,
			}},
		}, {
			Name: "Brewery", ProductionFactor: -1 * 5,
		}},
		ProductionBonus: []data.Resource{{
			Name: "Aqueduct", ProductionFactor: 0.03,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 5000,
		}, {
			Name: "Harbour", ProductionFactor: 2500,
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
			Name: "Active Smelter", ProductionFactor: -0.05 * 5,
		}},
		ProductionBonus: []data.Resource{{
			Name: "Lumber Mill", ProductionFactor: 0.10,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 200,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 150,
		}, {
			Name: "Harbour", ProductionFactor: 700,
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
		}, {
			Name: "Observatory", ProductionFactor: 0.25,
		}, {
			Name: "Bio Lab", ProductionFactor: 0.35,
		}, {
			Name: "Data Center", ProductionFactor: 0.10,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Library", ProductionFactor: 250,
		}, {
			Name: "Academy", ProductionFactor: 500,
		}, {
			Name: "Observatory", ProductionFactor: 1000,
		}, {
			Name: "Bio Lab", ProductionFactor: 1500,
		}, {
			Name: "Data Center", ProductionFactor: 750,
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
		}, {
			Name: "Mint", ProductionFactor: -0.75 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Hut", ProductionFactor: 75,
		}, {
			Name: "Log House", ProductionFactor: 50,
		}, {
			Name: "Mansion", ProductionFactor: 50,
		}},
	}, {
		Name: "minerals", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "miner", ProductionFactor: 0.25, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mine", ProductionFactor: 0.2,
			}, {
				Name: "Quarry", ProductionFactor: 0.35,
			}},
		}, {
			Name: "Active Smelter", ProductionFactor: -0.1 * 5,
		}, {
			Name: "Calciner", ProductionFactor: -1.5 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 250,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 200,
		}, {
			Name: "Harbour", ProductionFactor: 950,
		}},
	}, {
		Name: "iron", Type: "Resource", StartCapacity: 50,
		Producers: []data.Resource{{
			Name: "Active Smelter", ProductionFactor: 0.02 * 5,
		}, {
			Name: "Calciner", ProductionFactor: 0.15 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 50,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 25,
		}, {
			Name: "Harbour", ProductionFactor: 150,
		}},
	}, {
		Name: "coal", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Quarry", ProductionFactor: 0.015 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 60,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 30,
		}, {
			Name: "Harbour", ProductionFactor: 100,
		}},
	}, {
		Name: "gold", Type: "Resource", StartCapacity: 20,
		Producers: []data.Resource{{
			Name: "Mint", ProductionFactor: -0.005 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 10,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 5,
		}, {
			Name: "Harbour", ProductionFactor: 25,
		}, {
			Name: "Mint", ProductionFactor: 100,
		}},
	}, {
		Name: "titanium", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Accelerator", ProductionFactor: -0.015 * 5,
		}, {
			Name: "Calciner", ProductionFactor: 0.0005 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 2,
			ProductionBonus: []data.Resource{{
				Name: "Expanded Barns", ProductionFactor: 0.75,
			}, {
				Name: "Reinforced Barns", ProductionFactor: 0.80,
			}},
		}, {
			Name: "Warehouse", ProductionFactor: 10,
		}, {
			Name: "Harbour", ProductionFactor: 50,
		}},
	}, {
		Name: "oil", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Oil Well", ProductionFactor: 0.02 * 5,
		}, {
			Name: "Magneto", ProductionFactor: -0.05 * 5,
		}, {
			Name: "Calciner", ProductionFactor: -0.024 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Oil Well", ProductionFactor: 1500,
		}},
	}, {
		Name: "uranium", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Accelerator", ProductionFactor: 0.0025 * 5,
		}, {
			Name: "Reactor", ProductionFactor: -0.001 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Reactor", ProductionFactor: 250,
		}},
	}, {
		Name: "unobtainium", Type: "Resource", StartCapacity: 1,
	}, {
		Name: "time crystal", Type: "Resource", StartCapacity: 1,
	}, {
		Name: "antimatter", Type: "Resource", StartCapacity: 1,
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
			ProductionBonus: []data.Resource{{
				Name: "Tradepost", ProductionFactor: -0.04,
			}},
		}},
	}, {
		Name: "ivory", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.035,
			ProductionBonus: []data.Resource{{
				Name: "Tradepost", ProductionFactor: -0.04,
			}},
		}},
	}, {
		Name: "spice", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.005,
			ProductionBonus: []data.Resource{{
				Name: "Tradepost", ProductionFactor: -0.04,
			}},
		}, {
			Name: "Brewery", ProductionFactor: -0.1 * 5,
		}},
	}, {
		Name: "unicorns", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "Unic. Pasture", ProductionFactor: 0.001 * 5,
		}},
	}, {
		Name: "culture", Type: "Resource", StartCapacity: 575,
		Producers: []data.Resource{{
			Name: "Amphitheatre", ProductionFactor: 0.005 * 5,
		}, {
			Name: "Chapel", ProductionFactor: 0.05 * 5,
		}, {
			Name: "Temple", ProductionFactor: 0.1 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Library", ProductionFactor: 10,
			ProductionBonus: []data.Resource{{Name: "Ziggurat", ProductionFactor: 0.08}},
		}, {
			Name: "Academy", ProductionFactor: 25,
			ProductionBonus: []data.Resource{{Name: "Ziggurat", ProductionFactor: 0.08}},
		}, {
			Name: "Amphitheatre", ProductionFactor: 50,
			ProductionBonus: []data.Resource{{Name: "Ziggurat", ProductionFactor: 0.08}},
		}, {
			Name: "Chapel", ProductionFactor: 200,
			ProductionBonus: []data.Resource{{Name: "Ziggurat", ProductionFactor: 0.08}},
		}, {
			Name: "Data Center", ProductionFactor: 250,
			ProductionBonus: []data.Resource{{Name: "Ziggurat", ProductionFactor: 0.08}},
		}},
	}, {
		Name: "faith", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Chapel", ProductionFactor: 0.005 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Temple", ProductionFactor: 100,
		}},
	}, {
		Name: "beam", Type: "Resource", Capacity: -1,
		ProducerAction: "@beam",
	}, {
		Name: "slab", Type: "Resource", Capacity: -1,
		ProducerAction: "@slab",
	}, {
		Name: "plate", Type: "Resource", Capacity: -1,
		ProducerAction: "@plate",
	}, {
		Name: "steel", Type: "Resource", Capacity: -1,
	}, {
		Name: "gear", Type: "Resource", Capacity: -1,
		ProducerAction: "@gear",
	}, {
		Name: "concrete", Type: "Resource", Capacity: -1,
	}, {
		Name: "scaffold", Type: "Resource", Capacity: -1,
		ProducerAction: "@scaffold",
	}, {
		Name: "alloy", Type: "Resource", Capacity: -1,
	}, {
		Name: "parchment", Type: "Resource", Capacity: -1,
	}, {
		Name: "manuscript", Type: "Resource", Capacity: -1,
		ProducerAction: "@manuscript",
	}, {
		Name: "blueprint", Type: "Resource", Capacity: -1,
	}, {
		Name: "megalith", Type: "Resource", Capacity: -1,
		ProducerAction: "@megalith",
	}, {
		Name: "gigaflops", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "AI Core", ProductionFactor: 0.02 * 5,
		}},
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
		Name: "Aqueduct", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Mansion", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Observatory", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Bio Lab", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Harbour", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Quarry", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Lumber Mill", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Oil Well", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Accelerator", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Steamworks", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Magneto", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Calciner", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Factory", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Reactor", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Amphitheatre", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Chapel", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Temple", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Tradepost", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Mint", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Brewery", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Ziggurat", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Chronosphere", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "AI Core", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Solar Farm", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Hydro Plant", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Data Center", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Broadcast Tower", Type: "Bonfire", IsHidden: true, Capacity: -1,
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
		}, {
			Name: "Amphitheatre", ProductionFactor: -0.048,
		}, {
			Name: "Broadcast Tower", ProductionFactor: -0.75,
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
		Name: "Architecture", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Astronomy", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Biology", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Navigation", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Geology", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Chemistry", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Particle Physics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Machinery", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Electricity", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mechanization", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Nuclear Fission", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Writing", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Acoustics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Philosophy", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Chronophysics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Artificial Intelligence", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Ecology", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Robotics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Electronics", Type: "Science", IsHidden: true, Capacity: 1,
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
	g.AddActions([]data.Action{{
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
		Name: "Aqueduct", Type: "Bonfire", UnlockedBy: "Engineering",
		Costs: []data.Resource{{Name: "minerals", Quantity: 75, CostExponentBase: 1.12}},
		Adds:  []data.Resource{{Name: "Aqueduct", Quantity: 1}},
	}, {
		Name: "Mansion", Type: "Bonfire", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "slab", Quantity: 185, CostExponentBase: 1.15,
		}, {
			Name: "steel", Quantity: 75, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{
			Name: "Mansion", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "Observatory", Type: "Bonfire", UnlockedBy: "Astronomy",
		Costs: []data.Resource{{
			Name: "scaffold", Quantity: 50, CostExponentBase: 1.1,
		}, {
			Name: "slab", Quantity: 35, CostExponentBase: 1.1,
		}, {
			Name: "iron", Quantity: 750, CostExponentBase: 1.1,
		}, {
			Name: "science", Quantity: 1000, CostExponentBase: 1.1,
		}},
		Adds: []data.Resource{{Name: "Observatory", Quantity: 1}},
	}, {
		Name: "Bio Lab", Type: "Bonfire", UnlockedBy: "Biology",
		Costs: []data.Resource{{
			Name: "slab", Quantity: 100, CostExponentBase: 1.1,
		}, {
			Name: "alloy", Quantity: 25, CostExponentBase: 1.1,
		}, {
			Name: "science", Quantity: 1500, CostExponentBase: 1.1,
		}},
		Adds: []data.Resource{{Name: "Bio Lab", Quantity: 1}},
	}, {
		Name: "Harbour", Type: "Bonfire", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "scaffold", Quantity: 5, CostExponentBase: 1.15,
		}, {
			Name: "slab", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 75, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Harbour", Quantity: 1}},
	}, {
		Name: "Quarry", Type: "Bonfire", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "scaffold", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "steel", Quantity: 125, CostExponentBase: 1.15,
		}, {
			Name: "slab", Quantity: 1000, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Quarry", Quantity: 1}},
	}, {
		Name: "Lumber Mill", Type: "Bonfire", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "iron", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Lumber Mill", Quantity: 1}},
	}, {
		Name: "Oil Well", Type: "Bonfire", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "gear", Quantity: 25, CostExponentBase: 1.15,
		}, {
			Name: "scaffold", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Oil Well", Quantity: 1}},
	}, {
		Name: "Accelerator", Type: "Bonfire", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 7500, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Quantity: 125, CostExponentBase: 1.15,
		}, {
			Name: "uranium", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Accelerator", Quantity: 1}},
	}, {
		Name: "Steamworks", Type: "Bonfire", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 65, CostExponentBase: 1.25,
		}, {
			Name: "gear", Quantity: 20, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.25,
		}},
		Adds: []data.Resource{{Name: "Steamworks", Quantity: 1}},
	}, {
		Name: "Magneto", Type: "Bonfire", UnlockedBy: "Electricity",
		Costs: []data.Resource{{
			Name: "alloy", Quantity: 10, CostExponentBase: 1.25,
		}, {
			Name: "gear", Quantity: 5, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.25,
		}},
		Adds: []data.Resource{{Name: "Magneto", Quantity: 1}},
	}, {
		Name: "Calciner", Type: "Bonfire", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Quantity: 15, CostExponentBase: 1.15,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Calciner", Quantity: 1}},
	}, {
		Name: "Factory", Type: "Bonfire", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 2000, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 25000, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Quantity: 15, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Factory", Quantity: 1}},
	}, {
		Name: "Reactor", Type: "Bonfire", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 3500, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 5000, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "blueprint", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Reactor", Quantity: 1}},
	}, {
		Name: "Amphitheatre", Type: "Bonfire", UnlockedBy: "Writing",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 1200, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Quantity: 3, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Amphitheatre", Quantity: 1}},
	}, {
		Name: "Chapel", Type: "Bonfire", UnlockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 2000, CostExponentBase: 1.15,
		}, {
			Name: "culture", Quantity: 250, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Quantity: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Chapel", Quantity: 1}},
	}, {
		Name: "Temple", Type: "Bonfire", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "slab", Quantity: 25, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 15, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "manuscript", Quantity: 10, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Temple", Quantity: 1}},
	}, {
		Name: "Tradepost", Type: "Bonfire", UnlockedBy: "Currency",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 500, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 10, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Tradepost", Quantity: 1}},
	}, {
		Name: "Mint", Type: "Bonfire", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 5000, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 500, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Mint", Quantity: 1}},
	}, {
		Name: "Brewery", Type: "Bonfire", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 1000, CostExponentBase: 1.5,
		}, {
			Name: "culture", Quantity: 750, CostExponentBase: 1.5,
		}, {
			Name: "spice", Quantity: 5, CostExponentBase: 1.5,
		}, {
			Name: "parchment", Quantity: 375, CostExponentBase: 1.5,
		}},
		Adds: []data.Resource{{Name: "Brewery", Quantity: 1}},
	}, {
		Name: "Ziggurat", Type: "Bonfire", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "megalith", Quantity: 50, CostExponentBase: 1.25,
		}, {
			Name: "scaffold", Quantity: 50, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.25,
		}},
		Adds: []data.Resource{{Name: "Ziggurat", Quantity: 1}},
	}, {
		Name: "Chronosphere", Type: "Bonfire", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "unobtainium", Quantity: 2500, CostExponentBase: 1.25,
		}, {
			Name: "time crystal", Quantity: 1, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 100, CostExponentBase: 1.25,
		}, {
			Name: "science", Quantity: 250000, CostExponentBase: 1.25,
		}},
		Adds: []data.Resource{{Name: "Chronosphere", Quantity: 1}},
	}, {
		Name: "AI Core", Type: "Bonfire", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "antimatter", Quantity: 125, CostExponentBase: 1.15,
		}, {
			Name: "science", Quantity: 500000, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "AI Core", Quantity: 1}},
	}, {
		Name: "Solar Farm", Type: "Bonfire", UnlockedBy: "Ecology",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Solar Farm", Quantity: 1}},
	}, {
		Name: "Hydro Plant", Type: "Bonfire", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "concrete", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Quantity: 2500, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Hydro Plant", Quantity: 1}},
	}, {
		Name: "Data Center", Type: "Bonfire", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "concrete", Quantity: 10, CostExponentBase: 1.15,
		}, {
			Name: "steel", Quantity: 100, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "Data Center", Quantity: 1}},
	}, {
		Name: "Broadcast Tower", Type: "Bonfire", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 1250, CostExponentBase: 1.18,
		}, {
			Name: "titanium", Quantity: 75, CostExponentBase: 1.18,
		}},
		Adds: []data.Resource{{Name: "Broadcast Tower", Quantity: 1}},
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
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}},
	}, {
		Name: "@slab", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "minerals", Quantity: 250}},
		Adds: []data.Resource{{
			Name: "slab", Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}},
	}, {
		Name: "@plate", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "iron", Quantity: 125}},
		Adds: []data.Resource{{
			Name: "plate", Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}},
	}, {
		Name: "@gear", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "steel", Quantity: 15}},
		Adds: []data.Resource{{
			Name: "gear", Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}},
	}, {
		Name: "@scaffold", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "beam", Quantity: 50}},
		Adds: []data.Resource{{
			Name: "scaffold", Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}},
	}, {
		Name: "@manuscript", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "culture", Quantity: 400,
		}, {
			Name: "parchment", Quantity: 25,
		}},
		Adds: []data.Resource{{
			Name: "manuscript", Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}},
	}, {
		Name: "@megalith", Type: "Craft", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "beam", Quantity: 25,
		}, {
			Name: "slab", Quantity: 50,
		}, {
			Name: "plate", Quantity: 5,
		}},
		Adds: []data.Resource{{
			Name: "megalith", Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
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
	aqueduct
	mansion
	observatory
	biolab
	harbour
	quarry
	lumbermill
	oilwell
	accelerator
	steamworks
	magneto
	calciner
	factory
	reactor
	amphitheatre
	chapel
	temple
	tradepost
	mint
	brewery
	ziggurat
	chronosphere
	aicore
	solarfarm
	hydroplant
	datacenter
	broadcasttower
	woodcutter
	scholar
	farmer
	hunter
	miner
	sendhunters
	beam
	slab
	plate
	gear
	scaffold
	manuscript
	megalith
	lizards
	calendar
	agriculture
	archery
	mining
	animalhusbandry
	metalworking
	civilservice
	mathematics
	construction
	currency
	celestialmechanics
	engineering
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
		{[]int{s + field, field}, 55},
		{[]int{s + refine, refine}, 5},
		{[]int{hut, s + woodcutter, woodcutter}, 1},
		{[]int{s + library, library, s + scholar, scholar}, 1},
		{[]int{s + library, library}, 15},
		{[]int{s + calendar, calendar}, 1},
		{[]int{s + agriculture, agriculture}, 1},

		{[]int{s + barn, barn}, 10},
		{[]int{s + field, field}, 25},
		{[]int{s + library, library}, 15},
		{[]int{s + hut, hut, s + farmer, farmer}, 9},

		{[]int{s + archery, archery}, 1},
		{[]int{s + hunter, hunter}, 1}, // hut

		{[]int{s + animalhusbandry, animalhusbandry}, 1},
		{[]int{s + pasture, pasture}, 40},
		{[]int{s + sendhunters, sendhunters}, 40},
		{[]int{unicpasture}, 1},
		{[]int{s + unicpasture, unicpasture}, 10},
		{[]int{s + civilservice, civilservice}, 1},

		{[]int{s + mathematics, mathematics}, 1},
		{[]int{s + celestialmechanics, celestialmechanics}, 1},

		{[]int{s + construction, construction}, 1},
		{[]int{s + engineering, engineering}, 1},
		{[]int{s + currency, currency}, 1},
		{[]int{s + catnipenrichment, catnipenrichment}, 1},

		{[]int{s + mining, mining}, 1},
		{[]int{s + mine, mine}, 20},
		{[]int{miner}, 1}, // hut

		{[]int{s + academy, academy}, 25},
		{[]int{s + workshop, workshop}, 15},
		{[]int{s + mineralhoes, mineralhoes}, 1},
		{[]int{s + mineralaxe, mineralaxe}, 1},
		{[]int{s + bolas, bolas}, 1},
		{[]int{woodcutter}, 1}, // hut
		{[]int{s + loghouse, loghouse, s + woodcutter, woodcutter}, 7},
		{[]int{s + loghouse, loghouse, s + miner, miner}, 1},
		{[]int{s + loghouse, loghouse, s + farmer, farmer}, 10},

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

		{[]int{s + barn, barn}, 5},
		{[]int{s + field, field}, 5},
		{[]int{s + hut, hut, s + farmer, farmer}, 5},
		{[]int{s + library, library}, 20},
		{[]int{s + mine, mine}, 20},
		{[]int{s + workshop, workshop}, 20},
		{[]int{s + smelter, smelter}, 20},
		{[]int{s + pasture, pasture}, 50},
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
