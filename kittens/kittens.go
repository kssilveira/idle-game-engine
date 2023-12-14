package kittens

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

func NewGame(now game.Now) *game.Game {
	BarnProductionBonus := []data.Resource{{
		Name: "Expanded Barns", ProductionFactor: 0.75,
	}, {
		Name: "Reinforced Barns", ProductionFactor: 0.80,
	}}
	CultureCapacityProductionBonus := []data.Resource{{Name: "Ziggurat", ProductionFactor: 0.08}}
	kittenNames := []string{
		"kitten", "woodcutter", "scholar", "farmer", "hunter", "miner", "priest", "geologist",
	}

	g := game.NewGame(now())

	g.AddResources(join([]data.Resource{{
		Name: "day", Type: "Calendar", IsHidden: true, Quantity: 0, Capacity: -1,
		Producers: []data.Resource{{ProductionFactor: 0.5}},
	}, {
		Name: "year", Type: "Calendar", StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{Name: "day", ProductionFactor: 0.0025, ProductionFloor: true}},
	}}, resourceWithModulus(data.Resource{
		Type: "Calendar", StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{Name: "day", ProductionFactor: 0.01, ProductionFloor: true}},
	}, []string{
		"Spring", "Summer", "Autumn", "Winter"}), []data.Resource{{
		Name: "day_of_year", Type: "Calendar", StartQuantity: 1, Capacity: -1,
		ProductionModulus: 400, ProductionModulusEquals: -1,
		Producers: []data.Resource{{Name: "day", ProductionFactor: 1, ProductionFloor: true}},
	}, {
		Name: "catnip", Type: "Resource", StartCapacity: 5000,
		Producers: join([]data.Resource{{
			Name: "Catnip Field", ProductionFactor: 0.125 * 5 * (1 + 0.50), ProductionResourceFactor: "Spring",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.125 * 5, ProductionResourceFactor: "Summer",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.125 * 5, ProductionResourceFactor: "Autumn",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.125 * 5 * (1 - 0.75), ProductionResourceFactor: "Winter",
		}}, resourceWithName(data.Resource{
			ProductionFactor: -4.25, ProductionFloor: true, ProductionOnGone: true,
			ProductionBonus: []data.Resource{{
				Name: "Pasture", ProductionFactor: -0.005,
			}, {
				Name: "Unic. Pasture", ProductionFactor: -0.0015,
			}},
		}, kittenNames), []data.Resource{{
			Name: "farmer", ProductionFactor: 1 * 5, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mineral Hoes", ProductionFactor: 0.5,
			}, {
				Name: "Iron Hoes", ProductionFactor: 0.3,
			}},
		}, {
			Name: "Brewery", ProductionFactor: -1 * 5,
		}}),
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
			Name: "woodcutter", ProductionFactor: 0.0018 * 5, ProductionResourceFactor: "happiness",
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
			Name: "Barn", ProductionFactor: 200, ProductionBonus: BarnProductionBonus,
		}, {
			Name: "Warehouse", ProductionFactor: 150,
		}, {
			Name: "Harbour", ProductionFactor: 700,
		}},
	}, {
		Name: "science", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "scholar", ProductionFactor: 0.035 * 5, ProductionResourceFactor: "happiness",
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
			Name: "hunter", ProductionFactor: 0.06 * 5, ProductionResourceFactor: "happiness",
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
			Name: "miner", ProductionFactor: 0.05 * 5, ProductionResourceFactor: "happiness",
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
			Name: "Barn", ProductionFactor: 250, ProductionBonus: BarnProductionBonus,
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
			Name: "Barn", ProductionFactor: 50, ProductionBonus: BarnProductionBonus,
		}, {
			Name: "Warehouse", ProductionFactor: 25,
		}, {
			Name: "Harbour", ProductionFactor: 150,
		}},
	}, {
		Name: "coal", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "geologist", ProductionFactor: 0.015 * 5,
		}, {
			Name: "Quarry", ProductionFactor: 0.015 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", ProductionFactor: 60, ProductionBonus: BarnProductionBonus,
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
			Name: "Barn", ProductionFactor: 10, ProductionBonus: BarnProductionBonus,
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
			Name: "Barn", ProductionFactor: 2, ProductionBonus: BarnProductionBonus,
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
		Name: "relic", Type: "Resource", StartCapacity: 1,
	}, {
		Name: "void", Type: "Resource", StartCapacity: 1,
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
		Producers: join([]data.Resource{{
			Name: "", ProductionFactor: -1,
		}}, resourceWithName(data.Resource{
			Name: "kitten", ProductionFactor: 1, ProductionFloor: true,
		}, kittenNames)),
	}, {
		Name: "furs", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.05,
			ProductionBonus: []data.Resource{{Name: "Tradepost", ProductionFactor: -0.04}},
		}},
	}, {
		Name: "ivory", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.035,
			ProductionBonus: []data.Resource{{Name: "Tradepost", ProductionFactor: -0.04}},
		}},
	}, {
		Name: "spice", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.005,
			ProductionBonus: []data.Resource{{Name: "Tradepost", ProductionFactor: -0.04}},
		}, {
			Name: "Brewery", ProductionFactor: -0.1 * 5,
		}},
	}, {
		Name: "unicorns", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{Name: "Unic. Pasture", ProductionFactor: 0.001 * 5}},
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
			Name: "Library", ProductionFactor: 10, ProductionBonus: CultureCapacityProductionBonus,
		}, {
			Name: "Academy", ProductionFactor: 25, ProductionBonus: CultureCapacityProductionBonus,
		}, {
			Name: "Amphitheatre", ProductionFactor: 50, ProductionBonus: CultureCapacityProductionBonus,
		}, {
			Name: "Chapel", ProductionFactor: 200, ProductionBonus: CultureCapacityProductionBonus,
		}, {
			Name: "Data Center", ProductionFactor: 250, ProductionBonus: CultureCapacityProductionBonus,
		}},
	}, {
		Name: "faith", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "priest", ProductionFactor: 0.0015 * 5,
		}, {
			Name: "Chapel", ProductionFactor: 0.005 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Temple", ProductionFactor: 100,
		}},
	}, {
		Name: "steel", Type: "Resource", Capacity: -1,
	}, {
		Name: "concrete", Type: "Resource", Capacity: -1,
	}, {
		Name: "alloy", Type: "Resource", Capacity: -1,
	}, {
		Name: "parchment", Type: "Resource", Capacity: -1,
	}, {
		Name: "compendium", Type: "Resource", Capacity: -1,
	}, {
		Name: "blueprint", Type: "Resource", Capacity: -1,
	}, {
		Name: "gigaflops", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "AI Core", ProductionFactor: 0.02 * 5,
		}},
	}, {
		Name: "gone kitten", Type: "Resource", Capacity: -1,
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
	}}))

	g.AddActions([]data.Action{{
		Name: "Gather catnip", Type: "Bonfire", LockedBy: "Catnip Field",
		Adds: []data.Resource{{Name: "catnip", Quantity: 1}},
	}, {
		Name: "Refine catnip", Type: "Bonfire", UnlockedBy: "catnip", LockedBy: "woodcutter",
		Costs: []data.Resource{{Name: "catnip", Quantity: 100}},
		Adds:  []data.Resource{{Name: "wood", Quantity: 1}},
	}})

	addBuildings(g, []data.Action{{
		Name: "Catnip Field", UnlockedBy: "catnip",
		Costs: []data.Resource{{Name: "catnip", Quantity: 10, CostExponentBase: 1.12}},
	}, {
		Name: "Hut", UnlockedBy: "wood",
		Costs: []data.Resource{{Name: "wood", Quantity: 5, CostExponentBase: 2.5}},
		Adds:  []data.Resource{{Name: "kitten", Capacity: 2}},
	}, {
		Name: "Library", UnlockedBy: "wood",
		Costs: []data.Resource{{Name: "wood", Quantity: 25, CostExponentBase: 1.15}},
	}, {
		Name: "Barn", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "wood", Quantity: 50, CostExponentBase: 1.75}},
	}, {
		Name: "Mine", UnlockedBy: "Mining",
		Costs: []data.Resource{{Name: "wood", Quantity: 100, CostExponentBase: 1.15}},
	}, {
		Name: "Workshop", UnlockedBy: "Mining",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 400, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Smelter", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{Name: "minerals", Quantity: 200, CostExponentBase: 1.15}},
	}, {
		Name: "Pasture", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "wood", Quantity: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Unic. Pasture", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "unicorns", Quantity: 2, CostExponentBase: 1.75}},
	}, {
		Name: "Academy", UnlockedBy: "Mathematics",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 70, CostExponentBase: 1.15,
		}, {
			Name: "science", Quantity: 100, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Warehouse", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "beam", Quantity: 1.5, CostExponentBase: 1.15,
		}, {
			Name: "slab", Quantity: 2, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Log House", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "kitten", Capacity: 1}},
	}, {
		Name: "Aqueduct", UnlockedBy: "Engineering",
		Costs: []data.Resource{{Name: "minerals", Quantity: 75, CostExponentBase: 1.12}},
	}, {
		Name: "Mansion", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "slab", Quantity: 185, CostExponentBase: 1.15,
		}, {
			Name: "steel", Quantity: 75, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "kitten", Capacity: 1}},
	}, {
		Name: "Observatory", UnlockedBy: "Astronomy",
		Costs: []data.Resource{{
			Name: "scaffold", Quantity: 50, CostExponentBase: 1.1,
		}, {
			Name: "slab", Quantity: 35, CostExponentBase: 1.1,
		}, {
			Name: "iron", Quantity: 750, CostExponentBase: 1.1,
		}, {
			Name: "science", Quantity: 1000, CostExponentBase: 1.1,
		}},
	}, {
		Name: "Bio Lab", UnlockedBy: "Biology",
		Costs: []data.Resource{{
			Name: "slab", Quantity: 100, CostExponentBase: 1.1,
		}, {
			Name: "alloy", Quantity: 25, CostExponentBase: 1.1,
		}, {
			Name: "science", Quantity: 1500, CostExponentBase: 1.1,
		}},
	}, {
		Name: "Harbour", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "scaffold", Quantity: 5, CostExponentBase: 1.15,
		}, {
			Name: "slab", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 75, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Quarry", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "scaffold", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "steel", Quantity: 125, CostExponentBase: 1.15,
		}, {
			Name: "slab", Quantity: 1000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Lumber Mill", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "iron", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Oil Well", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "gear", Quantity: 25, CostExponentBase: 1.15,
		}, {
			Name: "scaffold", Quantity: 25, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Accelerator", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 7500, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Quantity: 125, CostExponentBase: 1.15,
		}, {
			Name: "uranium", Quantity: 25, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Steamworks", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 65, CostExponentBase: 1.25,
		}, {
			Name: "gear", Quantity: 20, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Magneto", UnlockedBy: "Electricity",
		Costs: []data.Resource{{
			Name: "alloy", Quantity: 10, CostExponentBase: 1.25,
		}, {
			Name: "gear", Quantity: 5, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Calciner", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Quantity: 15, CostExponentBase: 1.15,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Factory", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 2000, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 25000, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Quantity: 15, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Reactor", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 3500, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 5000, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "blueprint", Quantity: 25, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Amphitheatre", UnlockedBy: "Writing",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 1200, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Quantity: 3, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Chapel", UnlockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 2000, CostExponentBase: 1.15,
		}, {
			Name: "culture", Quantity: 250, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Quantity: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Temple", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "slab", Quantity: 25, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 15, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 50, CostExponentBase: 1.15,
		}, {
			Name: "manuscript", Quantity: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Tradepost", UnlockedBy: "Currency",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 500, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Mint", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 5000, CostExponentBase: 1.15,
		}, {
			Name: "plate", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Brewery", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 1000, CostExponentBase: 1.5,
		}, {
			Name: "culture", Quantity: 750, CostExponentBase: 1.5,
		}, {
			Name: "spice", Quantity: 5, CostExponentBase: 1.5,
		}, {
			Name: "parchment", Quantity: 375, CostExponentBase: 1.5,
		}},
	}, {
		Name: "Ziggurat", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "megalith", Quantity: 50, CostExponentBase: 1.25,
		}, {
			Name: "scaffold", Quantity: 50, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Chronosphere", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "unobtainium", Quantity: 2500, CostExponentBase: 1.25,
		}, {
			Name: "time crystal", Quantity: 1, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Quantity: 100, CostExponentBase: 1.25,
		}, {
			Name: "science", Quantity: 250000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "AI Core", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "antimatter", Quantity: 125, CostExponentBase: 1.15,
		}, {
			Name: "science", Quantity: 500000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Solar Farm", UnlockedBy: "Ecology",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Hydro Plant", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "concrete", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Quantity: 2500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Data Center", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "concrete", Quantity: 10, CostExponentBase: 1.15,
		}, {
			Name: "steel", Quantity: 100, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Broadcast Tower", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 1250, CostExponentBase: 1.18,
		}, {
			Name: "titanium", Quantity: 75, CostExponentBase: 1.18,
		}},
	}})

	addJobs(g, []data.Action{{
		Name: "woodcutter", UnlockedBy: "Hut",
	}, {
		Name: "scholar", UnlockedBy: "Library",
	}, {
		Name: "farmer", UnlockedBy: "Agriculture",
	}, {
		Name: "hunter", UnlockedBy: "Archery",
	}, {
		Name: "miner", UnlockedBy: "Mine",
	}, {
		Name: "priest", UnlockedBy: "Theology",
	}, {
		Name: "geologist", UnlockedBy: "Geology",
	}})

	g.AddActions([]data.Action{{
		Name: "Send hunters", Type: "Village", UnlockedBy: "Archery",
		Costs: []data.Resource{{Name: "catpower", Quantity: 100}},
		Adds: []data.Resource{{
			Name: "furs", Quantity: 39.5,
		}, {
			Name: "ivory", Quantity: 10.78,
		}, {
			Name: "unicorns", Quantity: 0.05,
		}},
	}})

	addCrafts(g, []data.Action{{
		Name: "beam", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "wood", Quantity: 175}},
	}, {
		Name: "slab", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "minerals", Quantity: 250}},
	}, {
		Name: "plate", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "iron", Quantity: 125}},
	}, {
		Name: "gear", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "steel", Quantity: 15}},
	}, {
		Name: "scaffold", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "beam", Quantity: 50}},
	}, {
		Name: "manuscript", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "culture", Quantity: 400,
		}, {
			Name: "parchment", Quantity: 25,
		}},
	}, {
		Name: "megalith", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "beam", Quantity: 25,
		}, {
			Name: "slab", Quantity: 50,
		}, {
			Name: "plate", Quantity: 5,
		}},
	}})

	g.AddActions([]data.Action{{
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
	}})

	addSciences(g, []data.Action{{
		Name: "Calendar", UnlockedBy: "Library",
		Costs: []data.Resource{{Name: "science", Quantity: 30}},
	}, {
		Name: "Agriculture", UnlockedBy: "Calendar",
		Costs: []data.Resource{{Name: "science", Quantity: 100}},
	}, {
		Name: "Archery", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "science", Quantity: 300}},
	}, {
		Name: "Mining", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "science", Quantity: 500}},
	}, {
		Name: "Animal Husbandry", UnlockedBy: "Archery",
		Costs: []data.Resource{{Name: "science", Quantity: 500}},
	}, {
		Name: "Metal Working", UnlockedBy: "Mining",
		Costs: []data.Resource{{Name: "science", Quantity: 900}},
	}, {
		Name: "Civil Service", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Quantity: 1500}},
	}, {
		Name: "Mathematics", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Quantity: 1000}},
	}, {
		Name: "Construction", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Quantity: 1300}},
	}, {
		Name: "Currency", UnlockedBy: "Civil Service",
		Costs: []data.Resource{{Name: "science", Quantity: 2200}},
	}, {
		Name: "Celestial Mechanics", UnlockedBy: "Mathematics",
		Costs: []data.Resource{{Name: "science", Quantity: 250}},
	}, {
		Name: "Engineering", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "science", Quantity: 1500}},
	}, {
		Name: "Writing", UnlockedBy: "Engineering",
		Costs: []data.Resource{{Name: "science", Quantity: 3600}},
	}, {
		Name: "Philosophy", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Quantity: 9500}},
	}, {
		Name: "Steel", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Quantity: 12000}},
	}, {
		Name: "Machinery", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Quantity: 15000}},
	}, {
		Name: "Theology", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 20000,
		}, {
			Name: "manuscript", Quantity: 35,
		}},
	}, {
		Name: "Astronomy", UnlockedBy: "Theology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 28000,
		}, {
			Name: "manuscript", Quantity: 65,
		}},
	}, {
		Name: "Navigation", UnlockedBy: "Astronomy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 35000,
		}, {
			Name: "manuscript", Quantity: 100,
		}},
	}, {
		Name: "Architecture", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 42000,
		}, {
			Name: "compendium", Quantity: 10,
		}},
	}, {
		Name: "Physics", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 50000,
		}, {
			Name: "compendium", Quantity: 35,
		}},
	}, {
		Name: "Metaphysics", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 55000,
		}, {
			Name: "unobtainium", Quantity: 5,
		}},
	}, {
		Name: "Chemistry", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 60000,
		}, {
			Name: "compendium", Quantity: 50,
		}},
	}, {
		Name: "Acoustics", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "science", Quantity: 60000,
		}, {
			Name: "compendium", Quantity: 60,
		}},
	}, {
		Name: "Geology", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 65000,
		}, {
			Name: "compendium", Quantity: 65,
		}},
	}, {
		Name: "Drama and Poetry", UnlockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 90000,
		}, {
			Name: "parchment", Quantity: 5000,
		}},
	}, {
		Name: "Electricity", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "compendium", Quantity: 85,
		}},
	}, {
		Name: "Biology", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 85000,
		}, {
			Name: "compendium", Quantity: 100,
		}},
	}, {
		Name: "Biochemistry", UnlockedBy: "Biology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 145000,
		}, {
			Name: "compendium", Quantity: 500,
		}},
	}, {
		Name: "Genetics", UnlockedBy: "Biochemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 190000,
		}, {
			Name: "compendium", Quantity: 1500,
		}},
	}, {
		Name: "Industrialization", UnlockedBy: "Electricity",
		Costs: []data.Resource{{
			Name: "science", Quantity: 10000,
		}, {
			Name: "blueprint", Quantity: 25,
		}},
	}, {
		Name: "Mechanization", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 115000,
		}, {
			Name: "blueprint", Quantity: 45,
		}},
	}, {
		Name: "Combustion", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 115000,
		}, {
			Name: "blueprint", Quantity: 45,
		}},
	}, {
		Name: "Metallurgy", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "blueprint", Quantity: 60,
		}},
	}, {
		Name: "Ecology", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "blueprint", Quantity: 55,
		}},
	}, {
		Name: "Electronics", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 135000,
		}, {
			Name: "blueprint", Quantity: 70,
		}},
	}, {
		Name: "Robotics", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 140000,
		}, {
			Name: "blueprint", Quantity: 80,
		}},
	}, {
		Name: "Artificial Intelligence", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "blueprint", Quantity: 150,
		}},
	}, {
		Name: "Quantum Cryptography", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1250000,
		}, {
			Name: "relic", Quantity: 1024,
		}},
	}, {
		Name: "Blackchain", UnlockedBy: "Quantum Cryptography",
		Costs: []data.Resource{{
			Name: "science", Quantity: 5000000,
		}, {
			Name: "relic", Quantity: 4096,
		}},
	}, {
		Name: "Nuclear Fission", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "blueprint", Quantity: 100,
		}},
	}, {
		Name: "Rocketry", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "blueprint", Quantity: 125,
		}},
	}, {
		Name: "Oil Processing", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 215000,
		}, {
			Name: "blueprint", Quantity: 150,
		}},
	}, {
		Name: "Satellites", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 190000,
		}, {
			Name: "blueprint", Quantity: 125,
		}},
	}, {
		Name: "Orbital Engineering", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "blueprint", Quantity: 250,
		}},
	}, {
		Name: "Thorium", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 375000,
		}, {
			Name: "blueprint", Quantity: 375,
		}},
	}, {
		Name: "Exogeology", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 275000,
		}, {
			Name: "blueprint", Quantity: 250,
		}},
	}, {
		Name: "Advanced Exogeology", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 325000,
		}, {
			Name: "blueprint", Quantity: 350,
		}},
	}, {
		Name: "Nanotechnology", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "blueprint", Quantity: 150,
		}},
	}, {
		Name: "Superconductors", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 225000,
		}, {
			Name: "blueprint", Quantity: 175,
		}},
	}, {
		Name: "Antimatter", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500000,
		}, {
			Name: "relic", Quantity: 1,
		}},
	}, {
		Name: "Terraformation", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Quantity: 750000,
		}, {
			Name: "relic", Quantity: 5,
		}},
	}, {
		Name: "Hydroponics", UnlockedBy: "Terraformation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1000000,
		}, {
			Name: "relic", Quantity: 25,
		}},
	}, {
		Name: "Exophysics", UnlockedBy: "Hydroponics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 25000000,
		}, {
			Name: "relic", Quantity: 500,
		}},
	}, {
		Name: "Particle Physics", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Quantity: 185000,
		}, {
			Name: "blueprint", Quantity: 135,
		}},
	}, {
		Name: "Dimensional Physics", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{Name: "science", Quantity: 235000}},
	}, {
		Name: "Chronophysics", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "time crystal", Quantity: 5,
		}},
	}, {
		Name: "Tachyon Theory", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 750000,
		}, {
			Name: "time crystal", Quantity: 25,
		}, {
			Name: "relic", Quantity: 1,
		}},
	}, {
		Name: "Cryptotheology", UnlockedBy: "Theology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 650000,
		}, {
			Name: "relic", Quantity: 5,
		}},
	}, {
		Name: "Void Space", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 800000,
		}, {
			Name: "time crystal", Quantity: 30,
		}, {
			Name: "void", Quantity: 100,
		}},
	}, {
		Name: "Paradox Theory", UnlockedBy: "Void Space",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1000000,
		}, {
			Name: "time crystal", Quantity: 40,
		}, {
			Name: "void", Quantity: 250,
		}},
	}})

	addWorkshops(g, []data.Action{{
		Name: "Mineral Hoes", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 275,
		}, {
			Name: "science", Quantity: 100,
		}},
	}, {
		Name: "Iron Hoes", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 25,
		}, {
			Name: "science", Quantity: 200,
		}},
	}, {
		Name: "Mineral Axe", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 500,
		}, {
			Name: "science", Quantity: 100,
		}},
	}, {
		Name: "Iron Axe", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 50,
		}, {
			Name: "science", Quantity: 200,
		}},
	}, {
		Name: "Expanded Barns", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 1000,
		}, {
			Name: "minerals", Quantity: 750,
		}, {
			Name: "iron", Quantity: 50,
		}, {
			Name: "science", Quantity: 500,
		}},
	}, {
		Name: "Reinforced Barns", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 100,
		}, {
			Name: "science", Quantity: 800,
		}, {
			Name: "beam", Quantity: 25,
		}, {
			Name: "slab", Quantity: 10,
		}},
	}, {
		Name: "Bolas", UnlockedBy: "Mining",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 50,
		}, {
			Name: "minerals", Quantity: 250,
		}, {
			Name: "science", Quantity: 1000,
		}},
	}, {
		Name: "Hunting Armor", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 750,
		}, {
			Name: "science", Quantity: 2000,
		}},
	}, {
		Name: "Reinforced Saw", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 1000,
		}, {
			Name: "science", Quantity: 2500,
		}},
	}, {
		Name: "Composite Bow", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 200,
		}, {
			Name: "iron", Quantity: 100,
		}, {
			Name: "science", Quantity: 500,
		}},
	}, {
		Name: "Catnip Enrichment", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 5000,
		}, {
			Name: "science", Quantity: 500,
		}},
	}})

	return g
}

func join[T any](slices ...[]T) []T {
	res := []T{}
	for _, slice := range slices {
		res = append(res, slice...)
	}
	return res
}

func resourceWithModulus(resource data.Resource, names []string) []data.Resource {
	res := []data.Resource{}
	resource.ProductionModulus = len(names)
	for i, name := range names {
		resource.Name = name
		resource.ProductionModulusEquals = i
		res = append(res, resource)
	}
	return res
}

func resourceWithName(resource data.Resource, names []string) []data.Resource {
	res := []data.Resource{}
	for _, name := range names {
		resource.Name = name
		res = append(res, resource)
	}
	return res
}

func addCrafts(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		name := action.Name
		action.Name = "@" + name
		action.Type = "Craft"
		action.Adds = []data.Resource{{
			Name: name, Quantity: 1,
			ProductionBonus: []data.Resource{{
				Name: "Workshop", ProductionFactor: 0.06,
			}, {
				Name: "Factory", ProductionFactor: 0.05,
			}},
		}}
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: name, Type: "Resource", Capacity: -1, ProducerAction: action.Name,
		})
	}
}

func addBuildings(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		name := action.Name
		isActive := false
		if strings.HasPrefix(action.Name, "Active ") {
			name = strings.TrimPrefix(action.Name, "Active ")
			isActive = true
		}
		action.Name = name
		action.Type = "Bonfire"
		action.Adds = append([]data.Resource{{
			Name: action.Name, Quantity: 1,
		}}, action.Adds...)
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Capacity: -1,
		})

		if !isActive {
			continue
		}

		action.Name = "Active " + name
		action.Costs = []data.Resource{{Name: "Idle " + name, Quantity: 1}}
		action.Adds = []data.Resource{{Name: action.Name, Quantity: 1}}
		action.UnlockedBy = name
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Capacity: -1,
		})

		g.AddResource(data.Resource{
			Name: "Idle " + name, Type: "Bonfire", Capacity: -1, StartQuantity: 1,
			Producers: []data.Resource{{
				Name: "", ProductionFactor: -1,
			}, {
				Name: name, ProductionFactor: 1,
			}, {
				Name: "Active " + name, ProductionFactor: -1,
			}},
		})
	}
}

func addJobs(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		action.Type = "Village"
		action.Costs = []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}}
		action.Adds = []data.Resource{{Name: action.Name, Quantity: 1}}
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Capacity: -1,
			OnGone: []data.Resource{{
				Name: "gone kitten", Quantity: 1,
			}, {
				Name: "kitten", Capacity: 1,
			}},
		})
	}
}

func addSciences(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		action.Type = "Science"
		action.Adds = []data.Resource{{Name: action.Name, Quantity: 1}}
		action.LockedBy = action.Name
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Capacity: 1,
		})
	}
}

func addWorkshops(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		action.Type = "Workshop"
		action.Adds = []data.Resource{{Name: action.Name, Quantity: 1}}
		action.LockedBy = action.Name
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Capacity: 1,
		})
	}
}

const (
	gather = iota
	refine
	field
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
		/*
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
			//*/
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
