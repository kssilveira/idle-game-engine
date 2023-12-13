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
	}, []string{"Spring", "Summer", "Autumn", "Winter"}), []data.Resource{{
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
		}, []string{
			"kitten", "woodcutter", "scholar", "farmer", "hunter", "miner", "priest", "geologist",
		}), []data.Resource{{
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
		}, {
			Name: "priest", ProductionFactor: 1,
		}, {
			Name: "geologist", ProductionFactor: 1,
		}},
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
		Name: "priest", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "geologist", Type: "Village", IsHidden: true, Capacity: -1,
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
		Name: "Steel", Type: "Science", IsHidden: true, Capacity: 1,
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
		Name: "Theology", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Physics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Metaphysics", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Chemistry", Type: "Science", IsHidden: true, Capacity: 1,
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
	g.AddActions([]data.Action{{
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
		Name: "priest", Type: "Village", UnlockedBy: "Theology",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "priest", Quantity: 1}},
	}, {
		Name: "geologist", Type: "Village", UnlockedBy: "Geology",
		Costs: []data.Resource{{Name: "kitten", Quantity: 1, Capacity: 1}},
		Adds:  []data.Resource{{Name: "geologist", Quantity: 1}},
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
	}, {
		Name: "Calendar", Type: "Science", UnlockedBy: "Library", LockedBy: "Calendar",
		Costs: []data.Resource{{Name: "science", Quantity: 30}},
		Adds:  []data.Resource{{Name: "Calendar", Quantity: 1}},
	}, {
		Name: "Agriculture", Type: "Science", UnlockedBy: "Calendar", LockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "science", Quantity: 100}},
		Adds:  []data.Resource{{Name: "Agriculture", Quantity: 1}},
	}, {
		Name: "Archery", Type: "Science", UnlockedBy: "Agriculture", LockedBy: "Archery",
		Costs: []data.Resource{{Name: "science", Quantity: 300}},
		Adds:  []data.Resource{{Name: "Archery", Quantity: 1}},
	}, {
		Name: "Mining", Type: "Science", UnlockedBy: "Agriculture", LockedBy: "Mining",
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
		Name: "Currency", Type: "Science", UnlockedBy: "Civil Service", LockedBy: "Currency",
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
		Name: "Writing", Type: "Science", UnlockedBy: "Engineering", LockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Quantity: 3600}},
		Adds:  []data.Resource{{Name: "Writing", Quantity: 1}},
	}, {
		Name: "Philosophy", Type: "Science", UnlockedBy: "Writing", LockedBy: "Philosophy",
		Costs: []data.Resource{{Name: "science", Quantity: 9500}},
		Adds:  []data.Resource{{Name: "Philosophy", Quantity: 1}},
	}, {
		Name: "Steel", Type: "Science", UnlockedBy: "Writing", LockedBy: "Steel",
		Costs: []data.Resource{{Name: "science", Quantity: 12000}},
		Adds:  []data.Resource{{Name: "Steel", Quantity: 1}},
	}, {
		Name: "Machinery", Type: "Science", UnlockedBy: "Writing", LockedBy: "Machinery",
		Costs: []data.Resource{{Name: "science", Quantity: 15000}},
		Adds:  []data.Resource{{Name: "Machinery", Quantity: 1}},
	}, {
		Name: "Theology", Type: "Science", UnlockedBy: "Philosophy", LockedBy: "Theology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 20000,
		}, {
			Name: "manuscript", Quantity: 35,
		}},
		Adds: []data.Resource{{Name: "Theology", Quantity: 1}},
	}, {
		Name: "Astronomy", Type: "Science", UnlockedBy: "Theology", LockedBy: "Astronomy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 28000,
		}, {
			Name: "manuscript", Quantity: 65,
		}},
		Adds: []data.Resource{{Name: "Astronomy", Quantity: 1}},
	}, {
		Name: "Navigation", Type: "Science", UnlockedBy: "Astronomy", LockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 35000,
		}, {
			Name: "manuscript", Quantity: 100,
		}},
		Adds: []data.Resource{{Name: "Navigation", Quantity: 1}},
	}, {
		Name: "Architecture", Type: "Science", UnlockedBy: "Navigation", LockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "science", Quantity: 42000,
		}, {
			Name: "compendium", Quantity: 10,
		}},
		Adds: []data.Resource{{Name: "Architecture", Quantity: 1}},
	}, {
		Name: "Physics", Type: "Science", UnlockedBy: "Navigation", LockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 50000,
		}, {
			Name: "compendium", Quantity: 35,
		}},
		Adds: []data.Resource{{Name: "Physics", Quantity: 1}},
	}, {
		Name: "Metaphysics", Type: "Science", UnlockedBy: "Physics", LockedBy: "Metaphysics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 55000,
		}, {
			Name: "unobtainium", Quantity: 5,
		}},
		Adds: []data.Resource{{Name: "Metaphysics", Quantity: 1}},
	}, {
		Name: "Chemistry", Type: "Science", UnlockedBy: "Physics", LockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 60000,
		}, {
			Name: "compendium", Quantity: 50,
		}},
		Adds: []data.Resource{{Name: "Chemistry", Quantity: 1}},
	}, {
		Name: "Acoustics", Type: "Science", UnlockedBy: "Architecture", LockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 60000,
		}, {
			Name: "compendium", Quantity: 60,
		}},
		Adds: []data.Resource{{Name: "Acoustics", Quantity: 1}},
	}, {
		Name: "Geology", Type: "Science", UnlockedBy: "Navigation", LockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 65000,
		}, {
			Name: "compendium", Quantity: 65,
		}},
		Adds: []data.Resource{{Name: "Geology", Quantity: 1}},
	}, {
		Name: "Drama and Poetry", Type: "Science", UnlockedBy: "Acoustics", LockedBy: "Drama and Poetry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 90000,
		}, {
			Name: "parchment", Quantity: 5000,
		}},
		Adds: []data.Resource{{Name: "Drama and Poetry", Quantity: 1}},
	}, {
		Name: "Electricity", Type: "Science", UnlockedBy: "Physics", LockedBy: "Electricity",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "compendium", Quantity: 85,
		}},
		Adds: []data.Resource{{Name: "Electricity", Quantity: 1}},
	}, {
		Name: "Biology", Type: "Science", UnlockedBy: "Geology", LockedBy: "Biology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 85000,
		}, {
			Name: "compendium", Quantity: 100,
		}},
		Adds: []data.Resource{{Name: "Biology", Quantity: 1}},
	}, {
		Name: "Biochemistry", Type: "Science", UnlockedBy: "Biology", LockedBy: "Biochemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 145000,
		}, {
			Name: "compendium", Quantity: 500,
		}},
		Adds: []data.Resource{{Name: "Biochemistry", Quantity: 1}},
	}, {
		Name: "Genetics", Type: "Science", UnlockedBy: "Biochemistry", LockedBy: "Genetics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 190000,
		}, {
			Name: "compendium", Quantity: 1500,
		}},
		Adds: []data.Resource{{Name: "Genetics", Quantity: 1}},
	}, {
		Name: "Industrialization", Type: "Science", UnlockedBy: "Electricity", LockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 10000,
		}, {
			Name: "blueprint", Quantity: 25,
		}},
		Adds: []data.Resource{{Name: "Industrialization", Quantity: 1}},
	}, {
		Name: "Mechanization", Type: "Science", UnlockedBy: "Industrialization", LockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 115000,
		}, {
			Name: "blueprint", Quantity: 45,
		}},
		Adds: []data.Resource{{Name: "Mechanization", Quantity: 1}},
	}, {
		Name: "Combustion", Type: "Science", UnlockedBy: "Industrialization", LockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Quantity: 115000,
		}, {
			Name: "blueprint", Quantity: 45,
		}},
		Adds: []data.Resource{{Name: "Combustion", Quantity: 1}},
	}, {
		Name: "Metallurgy", Type: "Science", UnlockedBy: "Industrialization", LockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "blueprint", Quantity: 60,
		}},
		Adds: []data.Resource{{Name: "Metallurgy", Quantity: 1}},
	}, {
		Name: "Ecology", Type: "Science", UnlockedBy: "Combustion", LockedBy: "Ecology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "blueprint", Quantity: 55,
		}},
		Adds: []data.Resource{{Name: "Ecology", Quantity: 1}},
	}, {
		Name: "Electronics", Type: "Science", UnlockedBy: "Mechanization", LockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 135000,
		}, {
			Name: "blueprint", Quantity: 70,
		}},
		Adds: []data.Resource{{Name: "Electronics", Quantity: 1}},
	}, {
		Name: "Robotics", Type: "Science", UnlockedBy: "Electronics", LockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 140000,
		}, {
			Name: "blueprint", Quantity: 80,
		}},
		Adds: []data.Resource{{Name: "Robotics", Quantity: 1}},
	}, {
		Name: "Artificial Intelligence", Type: "Science", UnlockedBy: "Robotics", LockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "blueprint", Quantity: 150,
		}},
		Adds: []data.Resource{{Name: "Artificial Intelligence", Quantity: 1}},
	}, {
		Name: "Quantum Cryptography", Type: "Science", UnlockedBy: "Artificial Intelligence", LockedBy: "Quantum Cryptography",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1250000,
		}, {
			Name: "relic", Quantity: 1024,
		}},
		Adds: []data.Resource{{Name: "Quantum Cryptography", Quantity: 1}},
	}, {
		Name: "Blackchain", Type: "Science", UnlockedBy: "Quantum Cryptography", LockedBy: "Blackchain",
		Costs: []data.Resource{{
			Name: "science", Quantity: 5000000,
		}, {
			Name: "relic", Quantity: 4096,
		}},
		Adds: []data.Resource{{Name: "Blackchain", Quantity: 1}},
	}, {
		Name: "Nuclear Fission", Type: "Science", UnlockedBy: "Electronics", LockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "blueprint", Quantity: 100,
		}},
		Adds: []data.Resource{{Name: "Nuclear Fission", Quantity: 1}},
	}, {
		Name: "Rocketry", Type: "Science", UnlockedBy: "Electronics", LockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "blueprint", Quantity: 125,
		}},
		Adds: []data.Resource{{Name: "Rocketry", Quantity: 1}},
	}, {
		Name: "Oil Processing", Type: "Science", UnlockedBy: "Rocketry", LockedBy: "Oil Processing",
		Costs: []data.Resource{{
			Name: "science", Quantity: 215000,
		}, {
			Name: "blueprint", Quantity: 150,
		}},
		Adds: []data.Resource{{Name: "Oil Processing", Quantity: 1}},
	}, {
		Name: "Satellites", Type: "Science", UnlockedBy: "Rocketry", LockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Quantity: 190000,
		}, {
			Name: "blueprint", Quantity: 125,
		}},
		Adds: []data.Resource{{Name: "Satellites", Quantity: 1}},
	}, {
		Name: "Orbital Engineering", Type: "Science", UnlockedBy: "Satellites", LockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "blueprint", Quantity: 250,
		}},
		Adds: []data.Resource{{Name: "Orbital Engineering", Quantity: 1}},
	}, {
		Name: "Thorium", Type: "Science", UnlockedBy: "Orbital Engineering", LockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Quantity: 375000,
		}, {
			Name: "blueprint", Quantity: 375,
		}},
		Adds: []data.Resource{{Name: "Thorium", Quantity: 1}},
	}, {
		Name: "Exogeology", Type: "Science", UnlockedBy: "Orbital Engineering", LockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 275000,
		}, {
			Name: "blueprint", Quantity: 250,
		}},
		Adds: []data.Resource{{Name: "Exogeology", Quantity: 1}},
	}, {
		Name: "Advanced Exogeology", Type: "Science", UnlockedBy: "Exogeology", LockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 325000,
		}, {
			Name: "blueprint", Quantity: 350,
		}},
		Adds: []data.Resource{{Name: "Advanced Exogeology", Quantity: 1}},
	}, {
		Name: "Nanotechnology", Type: "Science", UnlockedBy: "Nuclear Fission", LockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "blueprint", Quantity: 150,
		}},
		Adds: []data.Resource{{Name: "Nanotechnology", Quantity: 1}},
	}, {
		Name: "Superconductors", Type: "Science", UnlockedBy: "Nanotechnology", LockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Quantity: 225000,
		}, {
			Name: "blueprint", Quantity: 175,
		}},
		Adds: []data.Resource{{Name: "Superconductors", Quantity: 1}},
	}, {
		Name: "Antimatter", Type: "Science", UnlockedBy: "Superconductors", LockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500000,
		}, {
			Name: "relic", Quantity: 1,
		}},
		Adds: []data.Resource{{Name: "Antimatter", Quantity: 1}},
	}, {
		Name: "Terraformation", Type: "Science", UnlockedBy: "Antimatter", LockedBy: "Terraformation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 750000,
		}, {
			Name: "relic", Quantity: 5,
		}},
		Adds: []data.Resource{{Name: "Terraformation", Quantity: 1}},
	}, {
		Name: "Hydroponics", Type: "Science", UnlockedBy: "Terraformation", LockedBy: "Hydroponics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1000000,
		}, {
			Name: "relic", Quantity: 25,
		}},
		Adds: []data.Resource{{Name: "Hydroponics", Quantity: 1}},
	}, {
		Name: "Exophysics", Type: "Science", UnlockedBy: "Hydroponics", LockedBy: "Exophysics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 25000000,
		}, {
			Name: "relic", Quantity: 500,
		}},
		Adds: []data.Resource{{Name: "Exophysics", Quantity: 1}},
	}, {
		Name: "Particle Physics", Type: "Science", UnlockedBy: "Nuclear Fission", LockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 185000,
		}, {
			Name: "blueprint", Quantity: 135,
		}},
		Adds: []data.Resource{{Name: "Particle Physics", Quantity: 1}},
	}, {
		Name: "Dimensional Physics", Type: "Science", UnlockedBy: "Particle Physics", LockedBy: "Dimensional Physics",
		Costs: []data.Resource{{Name: "science", Quantity: 235000}},
		Adds:  []data.Resource{{Name: "Dimensional Physics", Quantity: 1}},
	}, {
		Name: "Chronophysics", Type: "Science", UnlockedBy: "Particle Physics", LockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "time crystal", Quantity: 5,
		}},
		Adds: []data.Resource{{Name: "Chronophysics", Quantity: 1}},
	}, {
		Name: "Tachyon Theory", Type: "Science", UnlockedBy: "Chronophysics", LockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 750000,
		}, {
			Name: "time crystal", Quantity: 25,
		}, {
			Name: "relic", Quantity: 1,
		}},
		Adds: []data.Resource{{Name: "Tachyon Theory", Quantity: 1}},
	}, {
		Name: "Cryptotheology", Type: "Science", UnlockedBy: "Theology", LockedBy: "Cryptotheology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 650000,
		}, {
			Name: "relic", Quantity: 5,
		}},
		Adds: []data.Resource{{Name: "Cryptotheology", Quantity: 1}},
	}, {
		Name: "Void Space", Type: "Science", UnlockedBy: "Tachyon Theory", LockedBy: "Void Space",
		Costs: []data.Resource{{
			Name: "science", Quantity: 800000,
		}, {
			Name: "time crystal", Quantity: 30,
		}, {
			Name: "void", Quantity: 100,
		}},
		Adds: []data.Resource{{Name: "Void Space", Quantity: 1}},
	}, {
		Name: "Paradox Theory", Type: "Science", UnlockedBy: "Void Space", LockedBy: "Paradox Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1000000,
		}, {
			Name: "time crystal", Quantity: 40,
		}, {
			Name: "void", Quantity: 250,
		}},
		Adds: []data.Resource{{Name: "Paradox Theory", Quantity: 1}},
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
		Name: "Bolas", Type: "Workshop", UnlockedBy: "Mining", LockedBy: "Bolas",
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
		action.Adds = append(action.Adds, data.Resource{
			Name: action.Name, Quantity: 1,
		})
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Capacity: -1,
		})

		if !isActive {
			continue
		}

		action.Name = "Active " + name
		action.Costs = []data.Resource{{
			Name: "Idle " + name, Quantity: 1,
		}}
		action.Adds = []data.Resource{{
			Name: action.Name, Quantity: 1,
		}}
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
