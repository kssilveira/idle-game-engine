package kittens

import (
	"strings"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

var (
	CraftRatio = []data.Resource{{
		Name: "Workshop", Factor: 0.06,
	}, {
		Name: "Factory", Factor: 0.05,
		Bonus: []data.Resource{{
			Name: "Factory Logistics", Factor: 0.2,
		}},
	}}
)

func NewGame(now game.Now) *game.Game {
	BarnBonus := []data.Resource{{
		Name: "Expanded Barns", Factor: 0.75,
	}, {
		Name: "Reinforced Barns", Factor: 0.80,
	}, {
		Name: "Titanium Barns", Factor: 1.00,
	}, {
		Name: "Alloy Barns", Factor: 1.00,
	}, {
		Name: "Concrete Barns", Factor: 0.75,
	}, {
		Name: "Concrete Pillars", Factor: 0.05,
	}}
	WarehouseBonus := []data.Resource{{
		Name: "Reinforced Warehouses", Factor: 0.25,
	}, {
		Name: "Titanium Warehouses", Factor: 0.50,
	}, {
		Name: "Alloy Warehouses", Factor: 0.45,
	}, {
		Name: "Concrete Warehouses", Factor: 0.35,
	}, {
		Name: "Storage Bunkers", Factor: 0.20,
	}, {
		Name: "Concrete Pillars", Factor: 0.05,
	}}
	HarbourBonus := []data.Resource{{
		Name: "trade ship", Factor: 0.01,
		Bonus: []data.Resource{{
			Name: "Expanded Cargo", Factor: 1,
		}},
		BonusStartsFromZero: true,
	}}
	HuntingBonus := []data.Resource{{
		Name: "Bolas", Factor: 1,
	}, {
		Name: "Hunting Armour", Factor: 2,
	}, {
		Name: "Steel Armour", Factor: 0.5,
	}, {
		Name: "Alloy Armour", Factor: 0.5,
	}, {
		Name: "Nanosuits", Factor: 0.5,
	}}
	CatnipCapacityBonus := []data.Resource{{
		Name: "Refrigeration", Factor: 0.75,
	}}
	CultureCapacityBonus := []data.Resource{{Name: "Ziggurat", Factor: 0.08}}
	kittenNames := []string{
		"kitten", "woodcutter", "scholar", "farmer", "hunter", "miner", "priest", "geologist",
	}

	g := game.NewGame(now())

	g.AddResources(join([]data.Resource{{
		Name: "day", Type: "Calendar", IsHidden: true, Quantity: 0, Capacity: -1,
		Producers: []data.Resource{{Factor: 0.5}},
	}, {
		Name: "year", Type: "Calendar", StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{Name: "day", Factor: 0.0025, ProductionFloor: true}},
	}}, resourceWithModulus(data.Resource{
		Type: "Calendar", StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{Name: "day", Factor: 0.01, ProductionFloor: true}},
	}, []string{
		"Spring", "Summer", "Autumn", "Winter"}), []data.Resource{{
		Name: "day_of_year", Type: "Calendar", StartQuantity: 1, Capacity: -1,
		ProductionModulus: 400, ProductionModulusEquals: -1,
		Producers: []data.Resource{{Name: "day", Factor: 1, ProductionFloor: true}},
	}, {
		Name: "catnip", Type: "Resource", StartCapacity: 5000,
		Producers: join([]data.Resource{{
			Name: "Catnip Field", Factor: 0.125 * 5,
			Bonus: []data.Resource{{
				Name: "Spring", Factor: 0.50,
			}, {
				Name: "Winter", Factor: -0.75,
			}},
		}}, resourceWithName(data.Resource{
			Factor: -4.25, ProductionFloor: true, ProductionOnGone: true,
			Bonus: []data.Resource{{
				Name: "Pasture", Factor: -0.005,
			}, {
				Name: "Unic. Pasture", Factor: -0.0015,
			}, {
				Name: "Robotic Assistance", Factor: -0.25,
			}},
		}, kittenNames), []data.Resource{{
			Name: "farmer", Factor: 1 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}, {
				Name: "Mineral Hoes", Factor: 0.5,
			}, {
				Name: "Iron Hoes", Factor: 0.3,
			}},
		}, {
			Name: "Brewery", Factor: -1 * 5,
		}, {
			Name: "Bio Lab", Factor: -1 * 5,
			Bonus: []data.Resource{{
				Name: "Biofuel Processing", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}}),
		Bonus: []data.Resource{{
			Name: "Aqueduct", Factor: 0.03,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 5000, Bonus: CatnipCapacityBonus,
		}, {
			Name: "Harbour", Factor: 2500, Bonus: join(HarbourBonus, CatnipCapacityBonus),
		}},
	}, {
		Name: "wood", Type: "Resource", StartCapacity: 200,
		Producers: []data.Resource{{
			Name: "woodcutter", Factor: 0.018 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}, {
				Name: "Mineral Axe", Factor: 0.7,
			}, {
				Name: "Iron Axe", Factor: 0.5,
			}, {
				Name: "Steel Axe", Factor: 0.5,
			}, {
				Name: "Titanium Axe", Factor: 0.5,
			}, {
				Name: "Alloy Axe", Factor: 0.5,
			}},
		}, {
			Name: "Active Smelter", Factor: -0.05 * 5,
		}},
		Bonus: []data.Resource{{
			Name: "Lumber Mill", Factor: 0.10,
			Bonus: []data.Resource{{
				Name: "Reinforced Saw", Factor: 0.2,
			}, {
				Name: "Steel Saw", Factor: 0.2,
			}, {
				Name: "Titanium Saw", Factor: 0.15,
			}, {
				Name: "Alloy Saw", Factor: 0.15,
			}},
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 200, Bonus: join(BarnBonus, WarehouseBonus),
		}, {
			Name: "Warehouse", Factor: 150, Bonus: join(BarnBonus, WarehouseBonus),
		}, {
			Name: "Harbour", Factor: 700, Bonus: join(BarnBonus, WarehouseBonus, HarbourBonus),
		}},
	}, {
		Name: "science", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "scholar", Factor: 0.035 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}},
		}},
		Bonus: []data.Resource{{
			Name: "Library", Factor: 0.1,
		}, {
			Name: "Academy", Factor: 0.2,
		}, {
			Name: "Observatory", Factor: 0.25,
		}, {
			Name: "Bio Lab", Factor: 0.35,
		}, {
			Name: "Data Center", Factor: 0.10,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Library", Factor: 250,
			Bonus: []data.Resource{{
				Name: "Observatory", Factor: 0.02,
				Bonus: []data.Resource{{
					Name: "Titanium Reflectors", Factor: 1,
				}, {
					Name: "Unobtainium Reflectors", Factor: 1,
				}, {
					Name: "Eludium Reflectors", Factor: 1,
				}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Academy", Factor: 500,
		}, {
			Name: "Observatory", Factor: 1000,
			Bonus: []data.Resource{{
				Name: "Astrolabe", Factor: 0.5,
			}},
		}, {
			Name: "Bio Lab", Factor: 1500,
			Bonus: []data.Resource{{
				Name: "Data Center", Factor: 0.01,
				Bonus: []data.Resource{{
					Name: "Uplink", Factor: 1,
				}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Data Center", Factor: 750,
			Bonus: []data.Resource{{
				Name: "Bio Lab", Factor: 0.01,
				Bonus: []data.Resource{{
					Name: "Uplink", Factor: 1,
				}},
				BonusStartsFromZero: true,
			}},
		}},
	}, {
		Name: "catpower", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "hunter", Factor: 0.06 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}, {
				Name: "Composite Bow", Factor: 0.5,
			}, {
				Name: "Crossbow", Factor: 0.25,
			}, {
				Name: "Railgun", Factor: 0.25,
			}},
		}, {
			Name: "Mint", Factor: -0.75 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Hut", Factor: 75,
		}, {
			Name: "Log House", Factor: 50,
		}, {
			Name: "Mansion", Factor: 50,
		}},
	}, {
		Name: "mineral", Type: "Resource", StartCapacity: 250,
		Producers: []data.Resource{{
			Name: "miner", Factor: 0.05 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}, {
				Name: "Mine", Factor: 0.2,
			}, {
				Name: "Quarry", Factor: 0.35,
			}},
		}, {
			Name: "Active Smelter", Factor: -0.1 * 5,
		}, {
			Name: "Calciner", Factor: -1.5 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 250, Bonus: join(BarnBonus, WarehouseBonus),
		}, {
			Name: "Warehouse", Factor: 200, Bonus: join(BarnBonus, WarehouseBonus),
		}, {
			Name: "Harbour", Factor: 950, Bonus: join(BarnBonus, WarehouseBonus, HarbourBonus),
		}},
	}, {
		Name: "iron", Type: "Resource", StartCapacity: 50,
		Producers: []data.Resource{{
			Name: "Active Smelter", Factor: 0.02 * 5,
			Bonus: []data.Resource{{Name: "Electrolytic Smelting", Factor: 0.95}},
		}, {
			Name: "Calciner", Factor: 0.15 * 5,
			Bonus: []data.Resource{{
				Name: "Oxidation", Factor: 0.95,
			}, {
				Name: "Rotary Kiln", Factor: 0.75,
			}, {
				Name: "Fluoridized Reactors", Factor: 1,
			}},
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 50, Bonus: join(BarnBonus, WarehouseBonus),
		}, {
			Name: "Warehouse", Factor: 25, Bonus: join(BarnBonus, WarehouseBonus),
		}, {
			Name: "Harbour", Factor: 150, Bonus: join(BarnBonus, WarehouseBonus, HarbourBonus),
		}},
	}, {
		Name: "coal", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "geologist", Factor: 0.015 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}, {
				Name: "Geodesy", Factor: 0.5,
			}, {
				Name: "Mining Drill", Factor: 10,
			}, {
				Name: "Unobtainium Drill", Factor: 2,
			}},
		}, {
			Name: "Quarry", Factor: 0.015 * 5,
		}, {
			Name: "Active Smelter", Factor: 0.005 * 5,
			Bonus: []data.Resource{{
				Name: "Coal Furnace", Factor: 1,
				Bonus: []data.Resource{{Name: "Electrolytic Smelting", Factor: 0.95}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Mine", Factor: 0.003 * 5,
			Bonus: []data.Resource{{
				Name: "Deep Mining", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Name: "Pyrolysis", Factor: 0.2,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 60, Bonus: WarehouseBonus,
		}, {
			Name: "Warehouse", Factor: 30, Bonus: WarehouseBonus,
		}, {
			Name: "Harbour", Factor: 100, Bonus: join(WarehouseBonus, HarbourBonus, []data.Resource{{
				Name: "Barges", Factor: 0.5,
			}}),
		}},
	}, {
		Name: "gold", Type: "Resource", StartCapacity: 20,
		Producers: []data.Resource{{
			Name: "Mint", Factor: -0.005 * 5,
		}, {
			Name: "Active Smelter", Factor: 0.001 * 5,
			Bonus: []data.Resource{{
				Name: "Gold Ore", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "geologist", Factor: 0.0008 * 5,
			Bonus: []data.Resource{{
				Name: "Geodesy", Factor: 1,
				Bonus: []data.Resource{{
					Name: "Mining Drill", Factor: 0.6,
				}, {
					Name: "Unobtainium Drill", Factor: 0.6,
				}},
			}},
			BonusStartsFromZero: true,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 10, Bonus: WarehouseBonus,
		}, {
			Name: "Warehouse", Factor: 5, Bonus: WarehouseBonus,
		}, {
			Name: "Harbour", Factor: 25, Bonus: join(WarehouseBonus, HarbourBonus),
		}, {
			Name: "Mint", Factor: 100, Bonus: WarehouseBonus,
		}},
	}, {
		Name: "titanium", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Accelerator", Factor: -0.015 * 5,
		}, {
			Name: "Calciner", Factor: 0.0005 * 5,
			Bonus: []data.Resource{{
				Name: "Oxidation", Factor: 2.85,
			}, {
				Name: "Rotary Kiln", Factor: 2.25,
			}, {
				Name: "Fluoridized Reactors", Factor: 3,
			}},
		}, {
			Name: "Active Smelter", Factor: 0.0015 * 5,
			Bonus: []data.Resource{{
				Name: "Nuclear Smelter", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Barn", Factor: 2, Bonus: WarehouseBonus,
		}, {
			Name: "Warehouse", Factor: 10, Bonus: WarehouseBonus,
		}, {
			Name: "Harbour", Factor: 50, Bonus: join(WarehouseBonus, HarbourBonus),
		}},
	}, {
		Name: "oil", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Oil Well", Factor: 0.02 * 5,
			Bonus: []data.Resource{{
				Name: "Pumpjack", Factor: 0.45,
			}, {
				Name: "Oil Refinery", Factor: 0.35,
			}, {
				Name: "Oil Distillation", Factor: 0.75,
			}},
		}, {
			Name: "Magneto", Factor: -0.05 * 5,
		}, {
			Name: "Calciner", Factor: -0.024 * 5,
		}, {
			Name: "Bio Lab", Factor: 0.10,
			Bonus: []data.Resource{{
				Name: "Biofuel Processing", Factor: 1,
				Bonus: []data.Resource{{
					Name: "GM Catnip", Factor: 0.60,
				}},
			}},
			BonusStartsFromZero: true,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Oil Well", Factor: 1500,
		}},
	}, {
		Name: "uranium", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "Accelerator", Factor: 0.0025 * 5,
		}, {
			Name: "Reactor", Factor: -0.001 * 5,
			Bonus: []data.Resource{{
				Name: "Enriched Uranium", Factor: -0.25,
			}},
		}, {
			Name: "Quarry", Factor: 0.0005 * 5,
			Bonus: []data.Resource{{
				Name: "Orbital Geodesy", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Reactor", Factor: 250,
		}},
	}, {
		Name: "unobtainium", Type: "Resource", Capacity: 1,
		Bonus: []data.Resource{{
			Name: "Microwarp Reactors", Factor: 0.75,
		}},
	}, {
		Name: "time crystal", Type: "Resource", Capacity: 1,
	}, {
		Name: "antimatter", Type: "Resource", Capacity: 1,
	}, {
		Name: "relic", Type: "Resource", Capacity: 1,
	}, {
		Name: "void", Type: "Resource", Capacity: 1,
	}, {
		Name: "temporal flux", Type: "Resource", Capacity: 1,
		Producers: []data.Resource{{
			Name: "Chronosphere", Factor: 1. / (2 * 400 * 4),
			Bonus: []data.Resource{{
				Name: "Chronosurge", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "blackcoin", Type: "Resource", Capacity: 1,
	}, {
		Name: "kitten", Type: "Resource", Capacity: 0,
		Producers: []data.Resource{{
			Name: "", Factor: 0.05,
		}},
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}},
	}, {
		Name: "all kittens", Type: "Resource", IsHidden: true, Capacity: -1, StartQuantity: 1,
		Producers: join([]data.Resource{{
			Name: "", Factor: -1,
		}}, resourceWithName(data.Resource{
			Name: "kitten", Factor: 1, ProductionFloor: true,
		}, kittenNames)),
	}, {
		Name: "furs", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.05,
			Bonus: []data.Resource{{Name: "Tradepost", Factor: -0.04}},
		}},
	}, {
		Name: "ivory", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.035,
			Bonus: []data.Resource{{Name: "Tradepost", Factor: -0.04}},
		}},
	}, {
		Name: "spice", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.005,
			Bonus: []data.Resource{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Brewery", Factor: -0.1 * 5,
		}},
	}, {
		Name: "unicorns", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "Unic. Pasture", Factor: 0.001 * 5,
			Bonus: []data.Resource{{
				Name: "Unicorn Selection", Factor: 0.25,
			}},
		}},
	}, {
		Name: "culture", Type: "Resource", StartCapacity: 575,
		Producers: []data.Resource{{
			Name: "Amphitheatre", Factor: 0.005 * 5,
		}, {
			Name: "Chapel", Factor: 0.05 * 5,
		}, {
			Name: "Temple", Factor: 0.1 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Library", Factor: 10, Bonus: CultureCapacityBonus,
		}, {
			Name: "Academy", Factor: 25, Bonus: CultureCapacityBonus,
		}, {
			Name: "Amphitheatre", Factor: 50, Bonus: CultureCapacityBonus,
		}, {
			Name: "Chapel", Factor: 200, Bonus: CultureCapacityBonus,
		}, {
			Name: "Data Center", Factor: 250,
			Bonus: join(CultureCapacityBonus, []data.Resource{{
				Name: "Bio Lab", Factor: 0.01,
				Bonus: []data.Resource{{
					Name: "Uplink", Factor: 1,
				}},
				BonusStartsFromZero: true,
			}}),
		}},
	}, {
		Name: "faith", Type: "Resource", StartCapacity: 1,
		Producers: []data.Resource{{
			Name: "priest", Factor: 0.0015 * 5,
			Bonus: []data.Resource{{
				Name: "happiness", Factor: 1,
			}},
		}, {
			Name: "Chapel", Factor: 0.005 * 5,
		}},
		CapacityProducers: []data.Resource{{
			Name: "Temple", Factor: 100,
		}},
	}, {
		Name: "starchart", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "scholar", Factor: 0.0005,
			Bonus: []data.Resource{{
				Name: "Astrophysicists", Factor: 1,
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Name: "Hubble Space Telescope", Factor: 0.30,
		}},
	}, {
		Name: "steel", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "Calciner", Factor: 0.15 * 5 * 0.1,
			Bonus: []data.Resource{{
				Name: "Steel Plants", Factor: 1,
				Bonus: []data.Resource{{
					Name: "Oxidation", Factor: 0.95,
				}, {
					Name: "Automated Plants", Factor: 0.25,
					Bonus:               CraftRatio,
					BonusStartsFromZero: true,
				}, {
					Name: "Reactor", Factor: 0.02,
					Bonus: []data.Resource{{
						Name: "Nuclear Plants", Factor: 1,
					}},
					BonusStartsFromZero: true,
				}},
			}},
			BonusStartsFromZero: true,
		}},
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
		Name: "trade ship", Type: "Resource", Capacity: -1,
	}, {
		Name: "eludium", Type: "Resource", Capacity: 1,
	}, {
		Name: "thorium", Type: "Resource", Capacity: 1,
		Producers: []data.Resource{{
			Name: "Reactor", Factor: -0.25 * 5,
			Bonus: []data.Resource{{
				Name: "Thorium Reactors", Factor: 1,
				Bonus: []data.Resource{{
					Name: "Enriched Thorium", Factor: 0.25,
				}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "gigaflops", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "AI Core", Factor: 0.02 * 5,
		}},
	}, {
		Name: "gone kitten", Type: "Resource", Capacity: -1,
	}, {
		Name: "happiness", Type: "Village", StartQuantity: 0.1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.02,
		}, {
			Name: "ivory", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "furs", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "spice", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "unicorns", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "Amphitheatre", Factor: 0.048,
		}, {
			Name: "Broadcast Tower", Factor: 0.75,
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
			Name: "mineral", Quantity: 400, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Smelter", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{Name: "mineral", Quantity: 200, CostExponentBase: 1.15}},
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
			Name: "mineral", Quantity: 70, CostExponentBase: 1.15,
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
			Name: "mineral", Quantity: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "kitten", Capacity: 1}},
	}, {
		Name: "Aqueduct", UnlockedBy: "Engineering",
		Costs: []data.Resource{{Name: "mineral", Quantity: 75, CostExponentBase: 1.12}},
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
			Name: "mineral", Quantity: 250, CostExponentBase: 1.15,
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
			Name: "mineral", Quantity: 1200, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Quantity: 3, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Chapel", UnlockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "mineral", Quantity: 2000, CostExponentBase: 1.15,
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
			Name: "mineral", Quantity: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Quantity: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Mint", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "mineral", Quantity: 5000, CostExponentBase: 1.15,
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
	}, {
		Name: "Chronocontrol", UnlockedBy: "Paradox Theory",
		Costs: []data.Resource{{}},
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
			Name: "furs", Quantity: 39.5, Bonus: HuntingBonus,
		}, {
			Name: "ivory", Quantity: 10.78, Bonus: HuntingBonus,
		}, {
			Name: "unicorns", Quantity: 0.05,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Lizards", Type: "Trade", UnlockedBy: "Archery",
		Costs: []data.Resource{{
			Name: "catpower", Quantity: 50,
		}, {
			Name: "gold", Quantity: 15,
		}, {
			Name: "mineral", Quantity: 1000,
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
			Name: "mineral", Quantity: 275,
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
			Name: "mineral", Quantity: 500,
		}, {
			Name: "science", Quantity: 100,
		}},
	}, {
		Name: "Iron Axe", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 50,
		}, {
			Name: "science", Quantity: 100,
		}},
	}, {
		Name: "Steel Axe", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 75,
		}, {
			Name: "science", Quantity: 20000,
		}},
	}, {
		Name: "Reinforced Saw", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "iron", Quantity: 1000,
		}, {
			Name: "science", Quantity: 2500,
		}},
	}, {
		Name: "Steel Saw", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "steel", Quantity: 750,
		}, {
			Name: "science", Quantity: 52000,
		}},
	}, {
		Name: "Titanium Saw", UnlockedBy: "Steel Saw",
		Costs: []data.Resource{{
			Name: "titanium", Quantity: 500,
		}, {
			Name: "science", Quantity: 70000,
		}},
	}, {
		Name: "Alloy Saw", UnlockedBy: "Titanium Saw",
		Costs: []data.Resource{{
			Name: "alloy", Quantity: 75,
		}, {
			Name: "science", Quantity: 85000,
		}},
	}, {
		Name: "Titanium Axe", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 38000,
		}, {
			Name: "titanium", Quantity: 10,
		}},
	}, {
		Name: "Alloy Axe", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 70000,
		}, {
			Name: "alloy", Quantity: 25,
		}},
	}, {
		Name: "Expanded Barns", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "wood", Quantity: 1000,
		}, {
			Name: "mineral", Quantity: 750,
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
		Name: "Reinforced Warehouses", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Quantity: 15000,
		}, {
			Name: "plate", Quantity: 50,
		}, {
			Name: "steel", Quantity: 50,
		}, {
			Name: "scaffold", Quantity: 25,
		}},
	}, {
		Name: "Titanium Barns", UnlockedBy: "Reinforced Barns",
		Costs: []data.Resource{{
			Name: "science", Quantity: 60000,
		}, {
			Name: "titanium", Quantity: 25,
		}, {
			Name: "steel", Quantity: 200,
		}, {
			Name: "scaffold", Quantity: 250,
		}},
	}, {
		Name: "Alloy Barns", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "alloy", Quantity: 20,
		}, {
			Name: "plate", Quantity: 750,
		}},
	}, {
		Name: "Concrete Barns", UnlockedBy: "Concrete Pillars",
		Costs: []data.Resource{{
			Name: "science", Quantity: 2000,
		}, {
			Name: "concrete", Quantity: 45,
		}, {
			Name: "titanium", Quantity: 2000,
		}},
	}, {
		Name: "Titanium Warehouses", UnlockedBy: "Silos",
		Costs: []data.Resource{{
			Name: "science", Quantity: 70000,
		}, {
			Name: "titanium", Quantity: 50,
		}, {
			Name: "steel", Quantity: 500,
		}, {
			Name: "scaffold", Quantity: 500,
		}},
	}, {
		Name: "Alloy Warehouses", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 90000,
		}, {
			Name: "titanium", Quantity: 750,
		}, {
			Name: "alloy", Quantity: 50,
		}},
	}, {
		Name: "Concrete Warehouses", UnlockedBy: "Concrete Pillars",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "titanium", Quantity: 1250,
		}, {
			Name: "concrete", Quantity: 35,
		}},
	}, {
		Name: "Storage Bunkers", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 25000,
		}, {
			Name: "unobtainium", Quantity: 500,
		}, {
			Name: "concrete", Quantity: 1250,
		}},
	}, {
		Name: "Energy Rifts", UnlockedBy: "Dimensional Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "titanium", Quantity: 7500,
		}, {
			Name: "uranium", Quantity: 250,
		}},
	}, {
		Name: "Stasis Chambers", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 235000,
		}, {
			Name: "alloy", Quantity: 200,
		}, {
			Name: "uranium", Quantity: 2000,
		}, {
			Name: "time crystal", Quantity: 1,
		}},
	}, {
		Name: "Void Energy", UnlockedBy: "Stasis Chambers",
		Costs: []data.Resource{{
			Name: "science", Quantity: 275000,
		}, {
			Name: "alloy", Quantity: 250,
		}, {
			Name: "uranium", Quantity: 2500,
		}, {
			Name: "time crystal", Quantity: 2,
		}},
	}, {
		Name: "Dark Energy", UnlockedBy: "Void Energy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 350000,
		}, {
			Name: "eludium", Quantity: 75,
		}, {
			Name: "time crystal", Quantity: 3,
		}},
	}, {
		Name: "Chronoforge", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500000,
		}, {
			Name: "relic", Quantity: 5,
		}, {
			Name: "time crystal", Quantity: 10,
		}},
	}, {
		Name: "Tachyon Accelerators", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500000,
		}, {
			Name: "eludium", Quantity: 125,
		}, {
			Name: "time crystal", Quantity: 10,
		}},
	}, {
		Name: "Flux Condensator", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "alloy", Quantity: 250,
		}, {
			Name: "unobtainium", Quantity: 5000,
		}, {
			Name: "time crystal", Quantity: 5,
		}},
	}, {
		Name: "LHC", UnlockedBy: "Dimensional Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "unobtainium", Quantity: 100,
		}, {
			Name: "alloy", Quantity: 150,
		}},
	}, {
		Name: "Photovoltaic Cells", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "titanium", Quantity: 5000,
		}},
	}, {
		Name: "Thin Film Cells", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "unobtainium", Quantity: 200,
		}, {
			Name: "uranium", Quantity: 1000,
		}},
	}, {
		Name: "Quantum Dot Cells", UnlockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "eludium", Quantity: 200,
		}, {
			Name: "thorium", Quantity: 1000,
		}},
	}, {
		Name: "Solar Satellites", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 225000,
		}, {
			Name: "alloy", Quantity: 750,
		}},
	}, {
		Name: "Expanded Cargo", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 55000,
		}, {
			Name: "blueprint", Quantity: 15,
		}},
	}, {
		Name: "Barges", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "titanium", Quantity: 1500,
		}, {
			Name: "blueprint", Quantity: 30,
		}},
	}, {
		Name: "Reactor Vessel", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Quantity: 135000,
		}, {
			Name: "titanium", Quantity: 5000,
		}, {
			Name: "uranium", Quantity: 125,
		}},
	}, {
		Name: "Ironwood Huts", UnlockedBy: "Reinforced Warehouses",
		Costs: []data.Resource{{
			Name: "science", Quantity: 30000,
		}, {
			Name: "wood", Quantity: 15000,
		}, {
			Name: "iron", Quantity: 3000,
		}},
	}, {
		Name: "Concrete Huts", UnlockedBy: "Concrete Pillars",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "concrete", Quantity: 45,
		}, {
			Name: "titanium", Quantity: 3000,
		}},
	}, {
		Name: "Unobtainium Huts", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "unobtainium", Quantity: 350,
		}, {
			Name: "titanium", Quantity: 15000,
		}},
	}, {
		Name: "Eludium Huts", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 275000,
		}, {
			Name: "eludium", Quantity: 125,
		}},
	}, {
		Name: "Silos", UnlockedBy: "Ironwood Huts",
		Costs: []data.Resource{{
			Name: "science", Quantity: 50000,
		}, {
			Name: "steel", Quantity: 125,
		}, {
			Name: "blueprint", Quantity: 5,
		}},
	}, {
		Name: "Refrigeration", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "titanium", Quantity: 2500,
		}, {
			Name: "blueprint", Quantity: 15,
		}},
	}, {
		Name: "Composite Bow", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500,
		}, {
			Name: "iron", Quantity: 100,
		}, {
			Name: "wood", Quantity: 200,
		}},
	}, {
		Name: "Crossbow", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "science", Quantity: 12000,
		}, {
			Name: "iron", Quantity: 1500,
		}},
	}, {
		Name: "Railgun", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "titanium", Quantity: 5000,
		}, {
			Name: "blueprint", Quantity: 25,
		}},
	}, {
		Name: "Bolas", UnlockedBy: "Mining",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1000,
		}, {
			Name: "mineral", Quantity: 250,
		}, {
			Name: "wood", Quantity: 50,
		}},
	}, {
		Name: "Hunting Armour", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{
			Name: "science", Quantity: 2000,
		}, {
			Name: "iron", Quantity: 750,
		}},
	}, {
		Name: "Steel Armour", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Quantity: 10000,
		}, {
			Name: "steel", Quantity: 50,
		}},
	}, {
		Name: "Alloy Armour", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 50000,
		}, {
			Name: "alloy", Quantity: 25,
		}},
	}, {
		Name: "Nanosuits", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 185000,
		}, {
			Name: "alloy", Quantity: 250,
		}},
	}, {
		Name: "Caravanserai", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 25000,
		}, {
			Name: "ivory", Quantity: 10000,
		}, {
			Name: "gold", Quantity: 250,
		}},
	}, {
		Name: "Catnip Enrichment", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500,
		}, {
			Name: "catnip", Quantity: 500,
		}},
	}, {
		Name: "Gold Ore", UnlockedBy: "Currency",
		Costs: []data.Resource{{
			Name: "science", Quantity: 1000,
		}, {
			Name: "mineral", Quantity: 800,
		}, {
			Name: "iron", Quantity: 100,
		}},
	}, {
		Name: "Geodesy", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 90000,
		}, {
			Name: "titanium", Quantity: 250,
		}, {
			Name: "starchart", Quantity: 500,
		}},
	}, {
		Name: "Register", UnlockedBy: "Writing",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500,
		}, {
			Name: "gold", Quantity: 10,
		}},
	}, {
		Name: "Concrete Pillars", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "concrete", Quantity: 50,
		}},
	}, {
		Name: "Mining Drill", UnlockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "titanium", Quantity: 1750,
		}, {
			Name: "steel", Quantity: 750,
		}},
	}, {
		Name: "Unobtainium Drill", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "unobtainium", Quantity: 250,
		}, {
			Name: "alloy", Quantity: 1250,
		}},
	}, {
		Name: "Coal Furnace", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Quantity: 5000,
		}, {
			Name: "mineral", Quantity: 5000,
		}, {
			Name: "iron", Quantity: 2000,
		}, {
			Name: "beam", Quantity: 35,
		}},
	}, {
		Name: "Deep Mining", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Quantity: 5000,
		}, {
			Name: "iron", Quantity: 1200,
		}, {
			Name: "beam", Quantity: 50,
		}},
	}, {
		Name: "Pyrolysis", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 35000,
		}, {
			Name: "compendium", Quantity: 5,
		}},
	}, {
		Name: "Electrolytic Smelting", UnlockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "titanium", Quantity: 2000,
		}},
	}, {
		Name: "Oxidation", UnlockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "steel", Quantity: 5000,
		}},
	}, {
		Name: "Steel Plants", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 140000,
		}, {
			Name: "titanium", Quantity: 3500,
		}, {
			Name: "gear", Quantity: 750,
		}},
	}, {
		Name: "Automated Plants", UnlockedBy: "Steel Plants",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "alloy", Quantity: 750,
		}},
	}, {
		Name: "Nuclear Plants", UnlockedBy: "Automated Plants",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "uranium", Quantity: 10000,
		}},
	}, {
		Name: "Rotary Kiln", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 145000,
		}, {
			Name: "titanium", Quantity: 5000,
		}, {
			Name: "gear", Quantity: 500,
		}},
	}, {
		Name: "Fluoridized Reactors", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "alloy", Quantity: 200,
		}},
	}, {
		Name: "Nuclear Smelter", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Quantity: 165000,
		}, {
			Name: "uranium", Quantity: 250,
		}},
	}, {
		Name: "Orbital Geodesy", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "alloy", Quantity: 1000,
		}, {
			Name: "oil", Quantity: 35000,
		}},
	}, {
		Name: "Printing Press", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "science", Quantity: 7500,
		}, {
			Name: "gear", Quantity: 45,
		}},
	}, {
		Name: "Offset Press", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "gear", Quantity: 250,
		}, {
			Name: "oil", Quantity: 15000,
		}},
	}, {
		Name: "Photolithography", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "alloy", Quantity: 1250,
		}, {
			Name: "oil", Quantity: 50000,
		}, {
			Name: "uranium", Quantity: 250,
		}},
	}, {
		Name: "Uplink", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "alloy", Quantity: 1750,
		}},
	}, {
		Name: "Cryocomputing", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "eludium", Quantity: 15,
		}},
	}, {
		Name: "Machine Learning", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "eludium", Quantity: 25,
		}, {
			Name: "antimatter", Quantity: 125,
		}},
	}, {
		Name: "Workshop Automation", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "science", Quantity: 10000,
		}, {
			Name: "gear", Quantity: 25,
		}},
	}, {
		Name: "Advanced Automation", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "gear", Quantity: 75,
		}, {
			Name: "blueprint", Quantity: 25,
		}},
	}, {
		Name: "Pneumatic Press", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 20000,
		}, {
			Name: "gear", Quantity: 30,
		}, {
			Name: "blueprint", Quantity: 5,
		}},
	}, {
		Name: "High Pressure Engine", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Quantity: 20000,
		}, {
			Name: "gear", Quantity: 25,
		}, {
			Name: "blueprint", Quantity: 5,
		}},
	}, {
		Name: "Fuel Injector", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "gear", Quantity: 250,
		}, {
			Name: "oil", Quantity: 20000,
		}},
	}, {
		Name: "Factory Logistics", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "gear", Quantity: 250,
		}, {
			Name: "titanium", Quantity: 2000,
		}},
	}, {
		Name: "Carbon Sequestration", UnlockedBy: "Ecology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "titanium", Quantity: 1250,
		}, {
			Name: "gear", Quantity: 125,
		}, {
			Name: "steel", Quantity: 4000,
		}, {
			Name: "alloy", Quantity: 1000,
		}},
	}, {
		Name: "Space Manufacturing", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "titanium", Quantity: 125000,
		}},
	}, {
		Name: "Astrolabe", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 25000,
		}, {
			Name: "titanium", Quantity: 5,
		}, {
			Name: "starchart", Quantity: 75,
		}},
	}, {
		Name: "Titanium Reflectors", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Quantity: 20000,
		}, {
			Name: "titanium", Quantity: 15,
		}, {
			Name: "starchart", Quantity: 20,
		}},
	}, {
		Name: "Unobtainium Reflectors", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "unobtainium", Quantity: 75,
		}, {
			Name: "starchart", Quantity: 750,
		}},
	}, {
		Name: "Eludium Reflectors", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "eludium", Quantity: 15,
		}},
	}, {
		Name: "Hydro Plant Turbines", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "unobtainium", Quantity: 125,
		}},
	}, {
		Name: "Antimatter Bases", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "eludium", Quantity: 15,
		}, {
			Name: "antimatter", Quantity: 250,
		}},
	}, {
		Name: "AI Bases", UnlockedBy: "Antimatter Bases",
		Costs: []data.Resource{{
			Name: "science", Quantity: 750000,
		}, {
			Name: "antimatter", Quantity: 7500,
		}},
	}, {
		Name: "Antimatter Fission", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Quantity: 525000,
		}, {
			Name: "antimatter", Quantity: 175,
		}, {
			Name: "thorium", Quantity: 7500,
		}},
	}, {
		Name: "Antimatter Drive", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Quantity: 525000,
		}, {
			Name: "antimatter", Quantity: 125,
		}},
	}, {
		Name: "Antimatter Reactors", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "eludium", Quantity: 35,
		}, {
			Name: "antimatter", Quantity: 750,
		}},
	}, {
		Name: "Advanced AM Reactors", UnlockedBy: "Antimatter Reactors",
		Costs: []data.Resource{{
			Name: "eludium", Quantity: 70,
		}, {
			Name: "antimatter", Quantity: 1750,
		}},
	}, {
		Name: "Void Reactors", UnlockedBy: "Advanced AM Reactors",
		Costs: []data.Resource{{
			Name: "void", Quantity: 250,
		}, {
			Name: "antimatter", Quantity: 2500,
		}},
	}, {
		Name: "Relic Station", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{
			Name: "eludium", Quantity: 100,
		}, {
			Name: "antimatter", Quantity: 5000,
		}},
	}, {
		Name: "Pumpjack", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "titanium", Quantity: 250,
		}, {
			Name: "gear", Quantity: 125,
		}},
	}, {
		Name: "Biofuel Processing", UnlockedBy: "Biochemistry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "titanium", Quantity: 1250,
		}},
	}, {
		Name: "Unicorn Selection", UnlockedBy: "Genetics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "titanium", Quantity: 1500,
		}},
	}, {
		Name: "GM Catnip", UnlockedBy: "Genetics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "titanium", Quantity: 1500,
		}, {
			Name: "catnip", Quantity: 1000000,
		}},
	}, {
		Name: "CAD System", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "titanium", Quantity: 750,
		}},
	}, {
		Name: "SETI", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "titanium", Quantity: 250,
		}},
	}, {
		Name: "Logistics", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "gear", Quantity: 100,
		}, {
			Name: "scaffold", Quantity: 1000,
		}},
	}, {
		Name: "Augmentations", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "titanium", Quantity: 5000,
		}, {
			Name: "uranium", Quantity: 50,
		}},
	}, {
		Name: "Cold Fusion", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "eludium", Quantity: 25,
		}},
	}, {
		Name: "Thorium Reactors", UnlockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Quantity: 400000,
		}, {
			Name: "thorium", Quantity: 10000,
		}},
	}, {
		Name: "Enriched Uranium", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "titanium", Quantity: 7500,
		}, {
			Name: "uranium", Quantity: 150,
		}},
	}, {
		Name: "Enriched Thorium", UnlockedBy: "Thorium Reactors",
		Costs: []data.Resource{{
			Name: "science", Quantity: 12500,
		}, {
			Name: "thorium", Quantity: 500000,
		}},
	}, {
		Name: "Oil Refinery", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Quantity: 125000,
		}, {
			Name: "titanium", Quantity: 1250,
		}, {
			Name: "gear", Quantity: 500,
		}},
	}, {
		Name: "Hubble Space Telescope", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "alloy", Quantity: 1250,
		}, {
			Name: "oil", Quantity: 50000,
		}},
	}, {
		Name: "Satellite Navigation", UnlockedBy: "Hubble Space Telescope",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "alloy", Quantity: 750,
		}},
	}, {
		Name: "Satellite Radio", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 225000,
		}, {
			Name: "alloy", Quantity: 5000,
		}},
	}, {
		Name: "Astrophysicists", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 250000,
		}, {
			Name: "unobtainium", Quantity: 350,
		}},
	}, {
		Name: "Microwarp Reactors", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "eludium", Quantity: 50,
		}},
	}, {
		Name: "Planet Busters", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Quantity: 275000,
		}, {
			Name: "eludium", Quantity: 250,
		}},
	}, {
		Name: "Thorium Drive", UnlockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Quantity: 400000,
		}, {
			Name: "trade ship", Quantity: 10000,
		}, {
			Name: "gear", Quantity: 40000,
		}, {
			Name: "alloy", Quantity: 2000,
		}, {
			Name: "thorium", Quantity: 100000,
		}},
	}, {
		Name: "Oil Distillation", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Quantity: 175000,
		}, {
			Name: "titanium", Quantity: 5000,
		}},
	}, {
		Name: "Factory Processing", UnlockedBy: "Oil Processing",
		Costs: []data.Resource{{
			Name: "science", Quantity: 195000,
		}, {
			Name: "titanium", Quantity: 7500,
		}, {
			Name: "concrete", Quantity: 125,
		}},
	}, {
		Name: "Factory Optimization", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "gear", Quantity: 250,
		}, {
			Name: "titanium", Quantity: 1250,
		}},
	}, {
		Name: "Space Engineers", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Quantity: 225000,
		}, {
			Name: "alloy", Quantity: 200,
		}},
	}, {
		Name: "AI Engineers", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Quantity: 35000,
		}, {
			Name: "eludium", Quantity: 50,
		}, {
			Name: "antimatter", Quantity: 500,
		}},
	}, {
		Name: "Chronoengineers", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 500000,
		}, {
			Name: "time crystal", Quantity: 5,
		}, {
			Name: "eludium", Quantity: 100,
		}},
	}, {
		Name: "Telecommunication", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 150000,
		}, {
			Name: "titanium", Quantity: 5000,
		}, {
			Name: "uranium", Quantity: 50,
		}},
	}, {
		Name: "Neural Network", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Quantity: 200000,
		}, {
			Name: "titanium", Quantity: 7500,
		}},
	}, {
		Name: "Robotic Assistance", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 100000,
		}, {
			Name: "steel", Quantity: 10000,
		}, {
			Name: "gear", Quantity: 250,
		}},
	}, {
		Name: "Factory Robotics", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Quantity: 75000,
		}, {
			Name: "gear", Quantity: 250,
		}, {
			Name: "titanium", Quantity: 1250,
		}},
	}, {
		Name: "Void Aspiration", UnlockedBy: "Void Energy",
		Costs: []data.Resource{{
			Name: "time crystal", Quantity: 15,
		}, {
			Name: "antimatter", Quantity: 2000,
		}},
	}, {
		Name: "Distortion", UnlockedBy: "Paradox Theory",
		Costs: []data.Resource{{
			Name: "science", Quantity: 300000,
		}, {
			Name: "time crystal", Quantity: 25,
		}, {
			Name: "antimatter", Quantity: 2000,
		}, {
			Name: "void", Quantity: 1000,
		}},
	}, {
		Name: "Chronosurge", UnlockedBy: "Chronocontrol",
		Costs: []data.Resource{{
			Name: "time crystal", Quantity: 25,
		}, {
			Name: "unobtainium", Quantity: 100000,
		}, {
			Name: "void", Quantity: 750,
		}, {
			Name: "temporal flux", Quantity: 6500,
		}},
	}, {
		Name: "Invisible Black Hand", UnlockedBy: "Blackchain",
		Costs: []data.Resource{{
			Name: "time crystal", Quantity: 128,
		}, {
			Name: "blackcoin", Quantity: 64,
		}, {
			Name: "void", Quantity: 32,
		}, {
			Name: "temporal flux", Quantity: 4096,
		}},
	}})

	addCrafts(g, []data.Action{{
		Name: "beam", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "wood", Quantity: 175}},
	}, {
		Name: "slab", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "mineral", Quantity: 250}},
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
		Producers: []data.Resource{{
			Name: "Steamworks", Factor: 0.0025,
			Bonus: []data.Resource{{
				Name: "Printing Press", Factor: 1,
				Bonus: []data.Resource{{
					Name: "Offset Press", Factor: 4 - 1,
					Bonus: []data.Resource{{
						Name: "Photolithography", Factor: 4 - 1,
					}},
				}},
			}},
			BonusStartsFromZero: true,
		}},
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
			Bonus: CraftRatio,
		}}
		action.IsHidden = true
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: name, Type: "Resource", Capacity: -1, ProducerAction: action.Name,
			Producers: action.Producers,
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
			Name: "Idle " + name, Type: "Bonfire", IsHidden: true, Capacity: -1, StartQuantity: 1,
			Producers: []data.Resource{{
				Name: "", Factor: -1,
			}, {
				Name: name, Factor: 1,
			}, {
				Name: "Active " + name, Factor: -1,
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
