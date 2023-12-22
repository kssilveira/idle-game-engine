package kittens

import (
	"strings"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

func NewGame(now game.Now) *game.Game {
	kittenNames := []string{
		"kitten", "woodcutter", "scholar", "farmer", "hunter", "miner", "priest", "geologist",
	}

	g := game.NewGame(now())

	g.AddResources(join([]data.Resource{{
		Name: "day", Type: "Calendar", IsHidden: true, Count: 0, Cap: -1,
		Producers: []data.Resource{{Factor: 0.5}},
	}, {
		Name: "year", Type: "Calendar", StartCount: 1, Cap: -1,
		Producers: []data.Resource{{Name: "day", Factor: 0.0025, ProductionFloor: true}},
	}}, resourceWithModulus(data.Resource{
		Type: "Calendar", StartCount: 1, Cap: -1,
		Producers: []data.Resource{{Name: "day", Factor: 0.01, ProductionFloor: true}},
	}, []string{
		"Spring", "Summer", "Autumn", "Winter"}), []data.Resource{{
		Name: "day_of_year", Type: "Calendar", StartCount: 1, Cap: -1,
		ProductionModulus: 400, ProductionModulusEquals: -1,
		Producers: []data.Resource{{Name: "day", ProductionFloor: true}},
	}}))

	addBonus(g, []data.Resource{{
		Name: "BarnBonus",
		Producers: []data.Resource{{
			Name: "Expanded Barns", Factor: 0.75,
		}, {
			Name: "Reinforced Barns", Factor: 0.80,
		}, {
			Name: "Titanium Barns",
		}, {
			Name: "Alloy Barns",
		}, {
			Name: "Concrete Barns", Factor: 0.75,
		}, {
			Name: "Concrete Pillars", Factor: 0.05,
		}},
	}, {
		Name: "WarehouseBonus", Type: "Resource", IsHidden: true, StartCountFromZero: true,
		Producers: []data.Resource{{
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
		}},
	}, {
		Name: "HarbourBonus", Type: "Resource", IsHidden: true, StartCountFromZero: true,
		Producers: []data.Resource{{
			Name: "ship", Factor: 0.01,
			Bonus: []data.Resource{{
				Name: "Expanded Cargo",
				Bonus: []data.Resource{{
					Name: "Reactor", Factor: 0.05,
					Bonus:               []data.Resource{{Name: "Reactor Vessel"}},
					BonusStartsFromZero: true,
				}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "BarnCatnipCapBonus", Type: "Resource", IsHidden: true, StartCountFromZero: true,
		Producers: []data.Resource{{
			Name: "Silos", Factor: 0.25,
			Bonus:               []data.Resource{{Name: "BarnBonus"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "HuntingBonus", Type: "Resource", IsHidden: true, StartCountFromZero: true,
		Producers: []data.Resource{{
			Name: "Bolas",
		}, {
			Name: "Hunting Armour", Factor: 2,
		}, {
			Name: "Steel Armour", Factor: 0.5,
		}, {
			Name: "Alloy Armour", Factor: 0.5,
		}, {
			Name: "Nanosuits", Factor: 0.5,
		}},
	}, {
		Name: "CraftRatio",
		Producers: []data.Resource{{
			Name: "Workshop", Factor: 0.06,
		}, {
			Name: "Factory", Factor: 0.05,
			Bonus: []data.Resource{{Name: "Factory Logistics", Factor: 0.20}},
		}},
	}, {
		Name: "SpaceElevatorOilBonus",
		Producers: []data.Resource{{
			Name: "Space Elevator", Factor: -0.05,
		}},
	}, {
		Name: "SpaceReactorScienceBonus",
		Producers: []data.Resource{{
			Name: "Antimatter Reactors", Factor: 0.95,
		}, {
			Name: "Advanced AM Reactors", Factor: 1.50,
		}, {
			Name: "Void Reactors", Factor: 4.00,
		}},
	}, {
		Name: "AcceleratorCapBonus",
		Producers: []data.Resource{{
			Name: "Stasis Chambers", Factor: 0.95,
		}, {
			Name: "Void Energy", Factor: 0.75,
		}, {
			Name: "Dark Energy", Factor: 2.50,
		}, {
			Name: "Tachyon Accelerators", Factor: 5.00,
		}},
	}, {
		Name: "MoonBaseCapBonus",
		Producers: []data.Resource{{
			Name: "AI Core", Factor: 0.10,
			Bonus:               []data.Resource{{Name: "AI Bases"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "CatnipCapBonus",
		Producers: []data.Resource{{
			Name: "Refrigeration", Factor: 0.75,
		}, {
			Name: "Hydroponics", Factor: 0.10,
		}},
	}, {
		Name: "ParagonCapBonus",
		Producers: []data.Resource{{
			Name: "paragon", Factor: 0.001,
		}, {
			Name: "burned paragon", Factor: 0.0005,
		}},
	}, {
		Name: "GlobalCapBonus",
		Producers: []data.Resource{{
			Name: "Void Rift", Factor: 0.02,
		}, {
			Name: "Event Horizon", Factor: 0.10,
		}},
	}, {
		Name:      "BaseMetalCapBonus",
		Producers: []data.Resource{{Name: "Sunforge", Factor: 0.01}},
	}, {
		Name: "ParagonProductionBonus",
		Producers: []data.Resource{{
			Name: "paragon", Factor: 0.01,
		}, {
			Name: "burned paragon", Factor: 0.01,
		}},
	}})

	g.AddResources([]data.Resource{{
		Name: "HutCostExponentBase", Type: "Resource", IsHidden: true, StartCount: 2.5,
		Bonus: []data.Resource{{
			Name: "Ironwood Huts", Factor: -0.50,
		}, {
			Name: "Concrete Huts", Factor: -0.30,
		}, {
			Name: "Unobtainium Huts", Factor: -0.25,
		}, {
			Name: "Eludium Huts", Factor: -0.10,
		}},
	}, {
		Name: "catnip cap", Type: "Resource", IsHidden: true, StartCount: 5000,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 5000,
			Bonus: []data.Resource{{Name: "BarnCatnipCapBonus"}},
		}, {
			Name: "Warehouse", Factor: 750,
			Bonus:               []data.Resource{{Name: "BarnCatnipCapBonus"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Harbour", Factor: 2500,
			Bonus: []data.Resource{{
				Name: "HarbourBonus",
			}, {
				Name: "BarnCatnipCapBonus",
			}},
		}, {
			Name: "Accelerator", Factor: 30000,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Moon Base", Factor: 45000, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "CatnipCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "catnip", Type: "Resource", CapResource: "catnip cap",
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
				Name: "happiness",
			}, {
				Name: "Mineral Hoes", Factor: 0.50,
			}, {
				Name: "Iron Hoes", Factor: 0.30,
			}},
		}, {
			Name: "Active Brewery", Factor: -1 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Bio Lab", Factor: -1 * 5, ProductionOnGone: true,
			Bonus:               []data.Resource{{Name: "Biofuel Processing"}},
			BonusStartsFromZero: true,
		}}),
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{
				Name: "Aqueduct", Factor: 0.03,
			}, {
				Name: "Hydroponics", Factor: 0.025,
			}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "wood cap", Type: "Resource", IsHidden: true, StartCount: 200,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 200,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Warehouse", Factor: 150,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Harbour", Factor: 700,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Moon Base", Factor: 25000, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 200000,
		}, {
			Name: "Accelerator", Factor: 20000,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "wood", Type: "Resource", CapResource: "wood cap",
		Producers: []data.Resource{{
			Name: "woodcutter", Factor: 0.018 * 5,
			Bonus: []data.Resource{{
				Name: "happiness",
			}, {
				Name: "Mineral Axe", Factor: 0.70,
			}, {
				Name: "Iron Axe", Factor: 0.50,
			}, {
				Name: "Steel Axe", Factor: 0.50,
			}, {
				Name: "Titanium Axe", Factor: 0.50,
			}, {
				Name: "Alloy Axe", Factor: 0.50,
			}},
		}, {
			Name: "Active Smelter", Factor: -0.05 * 5, ProductionOnGone: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{
				Name: "Lumber Mill", Factor: 0.10,
				Bonus: []data.Resource{{
					Name: "Reinforced Saw", Factor: 0.20,
				}, {
					Name: "Steel Saw", Factor: 0.20,
				}, {
					Name: "Titanium Saw", Factor: 0.15,
				}, {
					Name: "Alloy Saw", Factor: 0.15,
				}},
			}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "science cap", Type: "Resource", IsHidden: true, StartCount: 250,
		Producers: []data.Resource{{
			Name: "Library", Factor: 250,
			Bonus: []data.Resource{{
				Name: "Observatory", Factor: 0.02,
				Bonus: []data.Resource{{
					Name: "Titanium Reflectors",
				}, {
					Name: "Unobtainium Reflectors",
				}, {
					Name: "Eludium Reflectors",
				}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "compendium", Factor: 10,
		}, {
			Name: "Academy", Factor: 500,
		}, {
			Name: "Observatory", Factor: 1000,
			Bonus: []data.Resource{{
				Name: "Astrolabe", Factor: 0.50,
			}, {
				Name: "Satellite", Factor: 0.05,
			}},
		}, {
			Name: "Bio Lab", Factor: 1500,
			Bonus: []data.Resource{{
				Name: "Data Center", Factor: 0.01,
				Bonus: []data.Resource{{
					Name: "Uplink",
				}, {
					Name: "Starlink",
				}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Data Center", Factor: 750,
			Bonus: []data.Resource{{
				Name: "Bio Lab", Factor: 0.01,
				Bonus:               []data.Resource{{Name: "Uplink"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Observatory", Factor: 0.02,
				Bonus: []data.Resource{{
					Name: "Titanium Reflectors",
				}, {
					Name: "Unobtainium Reflectors",
				}, {
					Name: "Eludium Reflectors",
				}},
				BonusStartsFromZero: true,
			}, {
				Name: "AI Core", Factor: 0.10,
				Bonus:               []data.Resource{{Name: "Machine Learning"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Temple", Factor: 500,
			Bonus:               []data.Resource{{Name: "Scholasticism"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Accelerator", Factor: 2500,
			Bonus:               []data.Resource{{Name: "LHC"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Research Vessel", Factor: 10000, Bonus: []data.Resource{{Name: "SpaceReactorScienceBonus"}},
		}, {
			Name: "Space Beacon", Factor: 25000, Bonus: []data.Resource{{Name: "SpaceReactorScienceBonus"}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "science", Type: "Resource", CapResource: "science cap",
		Producers: []data.Resource{{
			Name: "scholar", Factor: 0.035 * 5,
			Bonus: []data.Resource{{Name: "happiness"}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{
				Name: "Library", Factor: 0.10,
			}, {
				Name: "Academy", Factor: 0.20,
			}, {
				Name: "Observatory", Factor: 0.25,
			}, {
				Name: "Bio Lab", Factor: 0.35,
			}, {
				Name: "Data Center", Factor: 0.10,
				Bonus: []data.Resource{{
					Name: "Bio Lab", Factor: 0.01,
					Bonus:               []data.Resource{{Name: "Uplink"}},
					BonusStartsFromZero: true,
				}, {
					Name: "AI Core", Factor: 0.10,
					Bonus:               []data.Resource{{Name: "Machine Learning"}},
					BonusStartsFromZero: true,
				}},
			}, {
				Name: "Space Station", Factor: 0.50,
			}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "catpower cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []data.Resource{{
			Name: "Hut", Factor: 75,
		}, {
			Name: "Log House", Factor: 50,
		}, {
			Name: "Mansion", Factor: 50,
		}, {
			Name: "Temple", Factor: 75,
			Bonus:               []data.Resource{{Name: "Templars"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "catpower", Type: "Resource", CapResource: "catpower cap",
		Producers: []data.Resource{{
			Name: "hunter", Factor: 0.06 * 5,
			Bonus: []data.Resource{{
				Name: "happiness",
			}, {
				Name: "Composite Bow", Factor: 0.50,
			}, {
				Name: "Crossbow", Factor: 0.25,
			}, {
				Name: "Railgun", Factor: 0.25,
			}},
		}, {
			Name: "Active Mint", Factor: -0.75 * 5, ProductionOnGone: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "mineral cap", Type: "Resource", IsHidden: true, StartCount: 250,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 250,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Warehouse", Factor: 200,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Harbour", Factor: 950,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Moon Base", Factor: 30000, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 200000,
		}, {
			Name: "Accelerator", Factor: 25000,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "mineral", Type: "Resource", CapResource: "mineral cap",
		Producers: []data.Resource{{
			Name: "miner", Factor: 0.05 * 5,
			Bonus: []data.Resource{{
				Name: "happiness",
			}, {
				Name: "Mine", Factor: 0.20,
			}, {
				Name: "Quarry", Factor: 0.35,
			}},
		}, {
			Name: "Active Smelter", Factor: -0.1 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Calciner", Factor: -1.5 * 5, ProductionOnGone: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "iron cap", Type: "Resource", IsHidden: true, StartCount: 50,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 50,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Warehouse", Factor: 25,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Harbour", Factor: 150,
			Bonus: []data.Resource{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Moon Base", Factor: 9000, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 50000,
		}, {
			Name: "Accelerator", Factor: 7500,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "iron", Type: "Resource", CapResource: "iron cap",
		Producers: []data.Resource{{
			Name: "Active Smelter", Factor: 0.02 * 5,
			Bonus: []data.Resource{{Name: "Electrolytic Smelting", Factor: 0.95}},
		}, {
			Name: "Active Calciner", Factor: 0.15 * 5,
			Bonus: []data.Resource{{
				Name: "Oxidation", Factor: 0.95,
			}, {
				Name: "Rotary Kiln", Factor: 0.75,
			}, {
				Name: "Fluoridized Reactors",
			}},
		}, {
			Name: "Active Calciner", Factor: -0.15 * 5 * 0.10, ProductionOnGone: true,
			Bonus:               []data.Resource{{Name: "Steel Plants"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "coal cap", Type: "Resource", IsHidden: true, StartCount: 60,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 60,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Warehouse", Factor: 30,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Harbour", Factor: 100,
			Bonus: []data.Resource{{
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}, {
				Name: "Barges", Factor: 0.50,
			}},
		}, {
			Name: "Moon Base", Factor: 3500, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 25000,
		}, {
			Name: "Accelerator", Factor: 2500,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "coal", Type: "Resource", CapResource: "coal cap",
		Producers: []data.Resource{{
			Name: "geologist", Factor: 0.015 * 5,
			Bonus: []data.Resource{{
				Name: "happiness",
			}, {
				Name: "Geodesy", Factor: 0.50,
			}, {
				Name: "Mining Drill", Factor: 0.66,
			}, {
				Name: "Unobtainium Drill",
			}},
		}, {
			Name: "Quarry", Factor: 0.015 * 5,
		}, {
			Name: "Active Smelter", Factor: 0.005 * 5,
			Bonus: []data.Resource{{
				Name:  "Coal Furnace",
				Bonus: []data.Resource{{Name: "Electrolytic Smelting", Factor: 0.95}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Mine", Factor: 0.003 * 5,
			Bonus:               []data.Resource{{Name: "Deep Mining"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Active Calciner", Factor: -0.15 * 5 * 0.10, ProductionOnGone: true,
			Bonus:               []data.Resource{{Name: "Steel Plants"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{
				Name: "Pyrolysis", Factor: 0.20,
			}, {
				Name: "Active Steamworks", Factor: -0.80, ProductionBoolean: true, ProductionOnGone: true,
				Bonus: []data.Resource{{
					Name: "High Pressure Engine", Factor: -0.25,
				}, {
					Name: "Fuel Injector", Factor: -0.25,
				}},
			}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "gold cap", Type: "Resource", IsHidden: true, StartCount: 10,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 10,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Warehouse", Factor: 5,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Harbour", Factor: 25,
			Bonus: []data.Resource{{
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Mint", Factor: 100,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Accelerator", Factor: 250,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "Sky Palace", Factor: 0.01}},
		}, {
			Bonus: []data.Resource{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "gold", Type: "Resource", CapResource: "gold cap",
		Producers: []data.Resource{{
			Name: "Active Mint", Factor: -0.005 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Smelter", Factor: 0.001 * 5,
			Bonus:               []data.Resource{{Name: "Gold Ore"}},
			BonusStartsFromZero: true,
		}, {
			Name: "geologist", Factor: 0.0008 * 5,
			Bonus: []data.Resource{{
				Name: "Geodesy",
				Bonus: []data.Resource{{
					Name: "Mining Drill", Factor: 0.625,
				}, {
					Name: "Unobtainium Drill", Factor: 0.625,
				}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "titanium cap", Type: "Resource", IsHidden: true, StartCount: 2,
		Producers: []data.Resource{{
			Name: "Barn", Factor: 2,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Warehouse", Factor: 10,
			Bonus: []data.Resource{{Name: "WarehouseBonus"}},
		}, {
			Name: "Harbour", Factor: 50,
			Bonus: []data.Resource{{
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Accelerator", Factor: 750,
			Bonus: []data.Resource{{
				Name: "Energy Rifts", Bonus: []data.Resource{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Moon Base", Factor: 1250, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 7500,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "titanium", Type: "Resource", CapResource: "titanium cap",
		Producers: []data.Resource{{
			Name: "Active Accelerator", Factor: -0.015 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Calciner", Factor: 0.0005 * 5,
			Bonus: []data.Resource{{
				Name: "Oxidation", Factor: 2.85,
			}, {
				Name: "Rotary Kiln", Factor: 2.25,
			}, {
				Name: "Fluoridized Reactors", Factor: 3.00,
			}},
		}, {
			Name: "Active Smelter", Factor: 0.0015 * 5,
			Bonus:               []data.Resource{{Name: "Nuclear Smelter"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "oil cap", Type: "Resource", IsHidden: true, StartCount: 1500,
		Producers: []data.Resource{{
			Name: "Oil Well", Factor: 1500,
		}, {
			Name: "tanker", Factor: 500,
		}, {
			Name: "Moon Base", Factor: 3500, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 7500,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "oil", Type: "Resource", CapResource: "oil cap",
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
			Name: "Active Magneto", Factor: -0.05 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Calciner", Factor: -0.024 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Bio Lab", Factor: 0.02 * 5,
			Bonus: []data.Resource{{
				Name:  "Biofuel Processing",
				Bonus: []data.Resource{{Name: "GM Catnip", Factor: 0.60}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Hydraulic Fracturer", Factor: 0.5 * 5,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "uranium cap", Type: "Resource", IsHidden: true, StartCount: 250,
		Producers: []data.Resource{{
			Name: "Reactor", Factor: 250,
		}, {
			Name: "Planet Cracker", Factor: 1750,
		}, {
			Name: "Cryostation", Factor: 5000,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "uranium", Type: "Resource", CapResource: "uranium cap",
		Producers: []data.Resource{{
			Name: "Active Accelerator", Factor: 0.0025 * 5,
		}, {
			Name: "Active Reactor", Factor: -0.001 * 5, ProductionOnGone: true,
			Bonus: []data.Resource{{Name: "Enriched Uranium", Factor: -0.25}},
		}, {
			Name: "Quarry", Factor: 0.0005 * 5,
			Bonus: []data.Resource{{
				Name:  "Orbital Geodesy",
				Bonus: []data.Resource{{Name: "Enriched Uranium", Factor: 0.25}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Active Lunar Outpost", Factor: -0.35 * 5, ProductionOnGone: true,
		}, {
			Name: "Planet Cracker", Factor: 0.3 * 5,
			Bonus: []data.Resource{{Name: "Planet Busters"}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "unobtainium cap", Type: "Resource", IsHidden: true, StartCount: 150,
		Producers: []data.Resource{{
			Name: "Moon Base", Factor: 150, Bonus: []data.Resource{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 750,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "unobtainium", Type: "Resource", CapResource: "unobtainium cap",
		Producers: []data.Resource{{Name: "Active Lunar Outpost", Factor: 0.035}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "Microwarp Reactors", Factor: 0.75}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "time crystal", Type: "Resource", Cap: -1,
	}, {
		Name: "antimatter cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []data.Resource{{
			Name: "Containment Chamber", Factor: 100,
			Bonus: []data.Resource{{Name: "Heatsink", Factor: 0.02}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonCapBonus"}},
		}, {
			Bonus: []data.Resource{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "antimatter", Type: "Resource", CapResource: "antimatter cap",
		Producers: []data.Resource{{Name: "Sunlifter", Factor: 1. / (2 * 100 * 4)}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "relic", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "Space Beacon", Factor: 0.01 / 2,
			Bonus: []data.Resource{{
				Name: "Relic Station",
				Bonus: []data.Resource{{
					Name: "Black Nexus", Factor: 0.10,
					Bonus:               []data.Resource{{Name: "Black Pyramid"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Hash Level", Factor: 0.25,
				}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "void", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "Chronosphere", Factor: 0.005 / (2 * 100),
			Bonus: []data.Resource{{Name: "Void Hoover"}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{
				Name:  "Chronocontrol",
				Bonus: []data.Resource{{Name: "Distortion", Factor: 2}},
			}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "temporal flux cap", Type: "Resource", IsHidden: true, StartCount: 3000,
		Bonus: []data.Resource{{
			Name: "Temporal Battery", Factor: 0.25,
		}},
	}, {
		Name: "temporal flux", Type: "Resource", IsHidden: true, CapResource: "temporal flux cap",
		Producers: []data.Resource{{
			Factor: 5. / (60 * 10),
			Bonus:  []data.Resource{{Name: "Temporal Accelerator", Factor: 0.05}},
		}, {
			Name: "Chronosphere", Factor: 1. / (2 * 100 * 4),
			Bonus:               []data.Resource{{Name: "Chronosurge"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "blackcoin", Type: "Resource", Cap: -1,
	}, {
		Name: "kitten", Type: "Resource", Cap: 0,
		Producers: []data.Resource{{Factor: 0.05}},
		Bonus: []data.Resource{{
			Name: "Venus of Willenfluff", Factor: 0.75,
		}, {
			Name: "Pawgan Rituals", Factor: 1.50,
		}},
		OnGone: []data.Resource{{Name: "gone kitten", Count: 1}},
	}, {
		Name: "all kittens", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true,
		Producers: resourceWithName(data.Resource{
			Name: "kitten", ProductionFloor: true,
		}, kittenNames),
	}, {
		Name: "fur", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.05,
			Bonus: []data.Resource{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Active Mint", Factor: 0.0000875,
			Bonus:               []data.Resource{{Name: "catpower cap"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "ivory", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.035,
			Bonus: []data.Resource{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Active Mint", Factor: 0.000021,
			Bonus:               []data.Resource{{Name: "catpower cap"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "spice", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.005,
			Bonus: []data.Resource{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Active Brewery", Factor: -0.1 * 5, ProductionOnGone: true,
		}, {
			Name: "Spice Refinery", Factor: 0.125,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "unicorn", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "Unic. Pasture", Factor: 0.001 * 5,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{
				Name: "Unicorn Selection", Factor: 0.25,
			}, {
				Name: "Unicorn Tomb", Factor: 0.05,
			}, {
				Name: "Ivory Tower", Factor: 0.10,
			}, {
				Name: "Ivory Citadel", Factor: 0.25,
			}, {
				Name: "Sky Palace", Factor: 0.50,
			}, {
				Name: "Unicorn Utopia", Factor: 2.50,
			}, {
				Name: "Sunspire", Factor: 5.00,
			}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "culture cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []data.Resource{{
			Name: "Library", Factor: 10,
		}, {
			Name: "Academy", Factor: 25,
		}, {
			Name: "Amphitheatre", Factor: 50,
		}, {
			Name: "Chapel", Factor: 200,
		}, {
			Name: "Data Center", Factor: 250,
			Bonus: []data.Resource{{
				Name: "Bio Lab", Factor: 0.01,
				Bonus:               []data.Resource{{Name: "Uplink"}},
				BonusStartsFromZero: true,
			}, {
				Name: "AI Core", Factor: 0.10,
				Bonus:               []data.Resource{{Name: "Machine Learning"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Temple", Factor: 125,
			Bonus:               []data.Resource{{Name: "Basilica"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []data.Resource{{
			Name: "Ziggurat", Factor: 0.08,
			Bonus: []data.Resource{{Name: "Unicorn Graveyard", Factor: 0.125}},
		}},
	}, {
		Name: "culture", Type: "Resource", CapResource: "culture cap",
		Producers: []data.Resource{{
			Name: "Amphitheatre", Factor: 0.005 * 5,
		}, {
			Name: "Chapel", Factor: 0.05 * 5,
		}, {
			Name: "Temple", Factor: 0.1 * 5,
			Bonus: []data.Resource{{
				Name: "Stained Glass", Factor: 0.50,
			}, {
				Name: "Basilica", Factor: 2.00,
			}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "faith cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []data.Resource{{
			Name: "Temple", Factor: 100,
			Bonus: []data.Resource{{
				Name: "Golden Spire", Factor: 0.50,
			}, {
				Name: "Sun Altar", Factor: 0.50,
			}},
		}},
	}, {
		Name: "faith", Type: "Resource", CapResource: "faith cap",
		Producers: []data.Resource{{
			Name: "priest", Factor: 0.0015 * 5,
			Bonus: []data.Resource{{Name: "happiness"}},
		}, {
			Name: "Temple", Factor: 0.0015 * 5,
			Bonus:               []data.Resource{{Name: "Theology"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Chapel", Factor: 0.005 * 5,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "Solar Chant", Factor: 0.10}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "starchart", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "scholar", Factor: 0.0005,
			Bonus:               []data.Resource{{Name: "Astrophysicists"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Satellite", Factor: 0.005,
		}, {
			Name: "Research Vessel", Factor: 0.05,
		}, {
			Name: "Space Beacon", Factor: 0.125,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "Hubble Space Telescope", Factor: 0.30}},
		}, {
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "gigaflop", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "AI Core", Factor: 0.02 * 5,
		}, {
			Name: "Entanglement Station", Factor: -0.1 * 5,
		}},
	}, {
		Name: "hash", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{Name: "Entanglement Station", Factor: 0.1 * 5}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "leviathan energy cap", Type: "Resource", IsHidden: true, StartCount: 1,
		Producers: []data.Resource{{Name: "Marker", Factor: 5}},
	}, {
		Name: "leviathan energy", Type: "Resource", CapResource: "leviathan energy cap",
	}, {
		Name: "tear", Type: "Resource", Cap: -1,
	}, {
		Name: "alicorn", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "Sky Palace", Factor: 0.00002 * 5,
			Bonus: []data.Resource{{
				Name: "Unicorn Utopia", Factor: 0.00015,
			}, {
				Name: "Sunspire", Factor: 0.0003,
			}, {
				Name: "Black Radiance", Factor: 0.03464,
				Bonus:               []data.Resource{{Name: "sorrow"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Unicorn Utopia", Factor: 0.000025 * 5,
		}, {
			Name: "Sunspire", Factor: 0.00005 * 5,
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "necrocorn", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "Marker", Factor: 0.000001 * 5,
			Bonus: []data.Resource{{Name: "Unicorn Graveyard", Factor: 0.10}},
		}},
		Bonus: []data.Resource{{
			Bonus: []data.Resource{{Name: "ParagonProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "sorrow cap", Type: "Resource", IsHidden: true, StartCount: 16,
		Producers: []data.Resource{{Name: "Black Core"}},
	}, {
		Name: "sorrow", Type: "Resource", CapResource: "sorrow cap",
	}, {
		Name: "chronoheat cap", Type: "Resource", IsHidden: true, StartCount: 1,
		Producers: []data.Resource{{
			Name: "Chrono Furnace", Factor: 100,
		}, {
			Name: "Time Boiler", Factor: 10,
		}},
	}, {
		Name: "chronoheat", Type: "Resource", CapResource: "chronoheat cap",
		Producers: []data.Resource{{Name: "Chrono Furnace", Factor: -0.02 * 5}},
	}, {
		Name: "chrono furnace fuel", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{
			Name: "Chrono Furnace", Factor: 0.02 * 5,
			Bonus:               []data.Resource{{Name: "chronoheat", ProductionBoolean: true}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "karma", Type: "Resource", Cap: -1,
	}, {
		Name: "paragon", Type: "Resource", Cap: -1,
		Producers: []data.Resource{{Factor: 1. / (2 * 100 * 4 * 1000)}},
	}, {
		Name: "burned paragon", Type: "Resource", Cap: -1,
	}, {
		Name: "gone kitten", Type: "Resource", Cap: -1,
	}, {
		Name: "happiness", Type: "Job", StartCount: 0.1, Cap: -1,
		Producers: []data.Resource{{
			Name: "all kittens", Factor: -0.02,
		}, {
			Name: "ivory", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "fur", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "spice", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "unicorn", Factor: 0.1, ProductionBoolean: true,
		}, {
			Name: "Amphitheatre", Factor: 0.048,
		}, {
			Name: "Broadcast Tower", Factor: 0.75,
		}, {
			Name: "Temple", Factor: 0.005,
			Bonus:               []data.Resource{{Name: "Sun Altar"}},
			BonusStartsFromZero: true,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Gather catnip", Type: "Building", LockedBy: "Catnip Field",
		Adds: []data.Resource{{Name: "catnip", Count: 10}},
	}, {
		Name: "Refine catnip", Type: "Building", UnlockedBy: "Catnip Field", LockedBy: "woodcutter",
		Costs: []data.Resource{{
			Name: "catnip", Count: 100,
			Bonus: []data.Resource{{Name: "Catnip Enrichment", Factor: -0.50}},
		}},
		Adds: []data.Resource{{
			Name: "wood", Count: 1,
			Bonus: []data.Resource{{Name: "Bio Lab", Factor: 0.10}},
		}},
	}})

	addBuildings(g, []data.Action{{
		Name: "Catnip Field", UnlockedBy: "catnip",
		Costs: []data.Resource{{Name: "catnip", Count: 10, CostExponentBase: 1.12}},
	}, {
		Name: "Hut", UnlockedBy: "Catnip Field",
		Costs: []data.Resource{{Name: "wood", Count: 5, CostExponentBaseResource: "HutCostExponentBase"}},
		Adds:  []data.Resource{{Name: "kitten", Cap: 2}},
	}, {
		Name: "Library", UnlockedBy: "Catnip Field",
		Costs: []data.Resource{{Name: "wood", Count: 25, CostExponentBase: 1.15}},
	}, {
		Name: "Barn", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "wood", Count: 50, CostExponentBase: 1.75}},
	}, {
		Name: "Mine", UnlockedBy: "Mining",
		Costs: []data.Resource{{Name: "wood", Count: 100, CostExponentBase: 1.15}},
	}, {
		Name: "Workshop", UnlockedBy: "Mining",
		Costs: []data.Resource{{
			Name: "wood", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "mineral", Count: 400, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Smelter", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{Name: "mineral", Count: 200, CostExponentBase: 1.15}},
	}, {
		Name: "Pasture", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{
			Name: "catnip", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "wood", Count: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Unic. Pasture", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "unicorn", Count: 2, CostExponentBase: 1.75}},
	}, {
		Name: "Academy", UnlockedBy: "Mathematics",
		Costs: []data.Resource{{
			Name: "wood", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "mineral", Count: 70, CostExponentBase: 1.15,
		}, {
			Name: "science", Count: 100, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Warehouse", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "beam", Count: 1.5, CostExponentBase: 1.15,
		}, {
			Name: "slab", Count: 2, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Log House", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Count: 200, CostExponentBase: 1.15,
		}, {
			Name: "mineral", Count: 250, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "kitten", Cap: 1}},
	}, {
		Name: "Aqueduct", UnlockedBy: "Engineering",
		Costs: []data.Resource{{Name: "mineral", Count: 75, CostExponentBase: 1.12}},
	}, {
		Name: "Mansion", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "slab", Count: 185, CostExponentBase: 1.15,
		}, {
			Name: "steel", Count: 75, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Count: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{Name: "kitten", Cap: 1}},
	}, {
		Name: "Observatory", UnlockedBy: "Astronomy",
		Costs: []data.Resource{{
			Name: "scaffold", Count: 50, CostExponentBase: 1.1,
		}, {
			Name: "slab", Count: 35, CostExponentBase: 1.1,
		}, {
			Name: "iron", Count: 750, CostExponentBase: 1.1,
		}, {
			Name: "science", Count: 1000, CostExponentBase: 1.1,
		}},
	}, {
		Name: "Active Bio Lab", UnlockedBy: "Biology",
		Costs: []data.Resource{{
			Name: "slab", Count: 100, CostExponentBase: 1.1,
		}, {
			Name: "alloy", Count: 25, CostExponentBase: 1.1,
		}, {
			Name: "science", Count: 1500, CostExponentBase: 1.1,
		}},
	}, {
		Name: "Harbour", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "scaffold", Count: 5, CostExponentBase: 1.15,
		}, {
			Name: "slab", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "plate", Count: 75, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Quarry", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "scaffold", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "steel", Count: 125, CostExponentBase: 1.15,
		}, {
			Name: "slab", Count: 1000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Lumber Mill", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "wood", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "iron", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "mineral", Count: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Oil Well", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "gear", Count: 25, CostExponentBase: 1.15,
		}, {
			Name: "scaffold", Count: 25, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Accelerator", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "titanium", Count: 7500, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Count: 125, CostExponentBase: 1.15,
		}, {
			Name: "uranium", Count: 25, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Steamworks", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "steel", Count: 65, CostExponentBase: 1.25,
		}, {
			Name: "gear", Count: 20, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Count: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Active Magneto", UnlockedBy: "Electricity",
		Costs: []data.Resource{{
			Name: "alloy", Count: 10, CostExponentBase: 1.25,
		}, {
			Name: "gear", Count: 5, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Count: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Active Calciner", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Count: 15, CostExponentBase: 1.15,
		}, {
			Name: "blueprint", Count: 1, CostExponentBase: 1.15,
		}, {
			Name: "oil", Count: 500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Factory", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "titanium", Count: 2000, CostExponentBase: 1.15,
		}, {
			Name: "plate", Count: 25000, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Count: 15, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Reactor", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "titanium", Count: 3500, CostExponentBase: 1.15,
		}, {
			Name: "plate", Count: 5000, CostExponentBase: 1.15,
		}, {
			Name: "concrete", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "blueprint", Count: 25, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Amphitheatre", UnlockedBy: "Writing",
		Costs: []data.Resource{{
			Name: "wood", Count: 200, CostExponentBase: 1.15,
		}, {
			Name: "mineral", Count: 1200, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Count: 3, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Chapel", UnlockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "mineral", Count: 2000, CostExponentBase: 1.15,
		}, {
			Name: "culture", Count: 250, CostExponentBase: 1.15,
		}, {
			Name: "parchment", Count: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Temple", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "slab", Count: 25, CostExponentBase: 1.15,
		}, {
			Name: "plate", Count: 15, CostExponentBase: 1.15,
		}, {
			Name: "gold", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "manuscript", Count: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Tradepost", UnlockedBy: "Currency",
		Costs: []data.Resource{{
			Name: "wood", Count: 500, CostExponentBase: 1.15,
		}, {
			Name: "mineral", Count: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Count: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Mint", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "mineral", Count: 5000, CostExponentBase: 1.15,
		}, {
			Name: "plate", Count: 200, CostExponentBase: 1.15,
		}, {
			Name: "gold", Count: 500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Active Brewery", UnlockedBy: "Drama and Poetry",
		Costs: []data.Resource{{
			Name: "wood", Count: 1000, CostExponentBase: 1.5,
		}, {
			Name: "culture", Count: 750, CostExponentBase: 1.5,
		}, {
			Name: "spice", Count: 5, CostExponentBase: 1.5,
		}, {
			Name: "parchment", Count: 375, CostExponentBase: 1.5,
		}},
	}, {
		Name: "Ziggurat", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "megalith", Count: 50, CostExponentBase: 1.25,
		}, {
			Name: "scaffold", Count: 50, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Count: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Chronosphere", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "unobtainium", Count: 2500, CostExponentBase: 1.25,
		}, {
			Name: "time crystal", Count: 1, CostExponentBase: 1.25,
		}, {
			Name: "blueprint", Count: 100, CostExponentBase: 1.25,
		}, {
			Name: "science", Count: 250000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "AI Core", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "antimatter", Count: 125, CostExponentBase: 1.15,
		}, {
			Name: "science", Count: 500000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Solar Farm", UnlockedBy: "Ecology",
		Costs: []data.Resource{{
			Name: "titanium", Count: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Hydro Plant", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "concrete", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Count: 2500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Data Center", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "concrete", Count: 10, CostExponentBase: 1.15,
		}, {
			Name: "steel", Count: 100, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Broadcast Tower", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "iron", Count: 1250, CostExponentBase: 1.18,
		}, {
			Name: "titanium", Count: 75, CostExponentBase: 1.18,
		}},
	}, {
		Name: "Unicorn Tomb", UnlockedBy: "Ziggurat",
		Costs: []data.Resource{{
			Name: "tear", Count: 5, CostExponentBase: 1.15,
		}, {
			Name: "ivory", Count: 500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Ivory Tower", UnlockedBy: "Unicorn Tomb",
		Costs: []data.Resource{{
			Name: "tear", Count: 25, CostExponentBase: 1.15,
		}, {
			Name: "ivory", Count: 25000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Ivory Citadel", UnlockedBy: "Ivory Tower",
		Costs: []data.Resource{{
			Name: "tear", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "ivory", Count: 50000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Sky Palace", UnlockedBy: "Ivory Citadel",
		Costs: []data.Resource{{
			Name: "tear", Count: 500, CostExponentBase: 1.15,
		}, {
			Name: "ivory", Count: 125000, CostExponentBase: 1.15,
		}, {
			Name: "megalith", Count: 5, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Unicorn Utopia", UnlockedBy: "Sky Palace",
		Costs: []data.Resource{{
			Name: "tear", Count: 5000, CostExponentBase: 1.15,
		}, {
			Name: "ivory", Count: 1000000, CostExponentBase: 1.15,
		}, {
			Name: "gold", Count: 500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Sunspire", UnlockedBy: "Unicorn Utopia",
		Costs: []data.Resource{{
			Name: "tear", Count: 25000, CostExponentBase: 1.15,
		}, {
			Name: "ivory", Count: 750000, CostExponentBase: 1.15,
		}, {
			Name: "gold", Count: 1000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Marker", UnlockedBy: "Megalomania",
		Costs: []data.Resource{{
			Name: "tear", Count: 5000, CostExponentBase: 1.15,
		}, {
			Name: "megalith", Count: 750, CostExponentBase: 1.15,
		}, {
			Name: "spice", Count: 50000, CostExponentBase: 1.15,
		}, {
			Name: "unobtainium", Count: 2500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Unicorn Graveyard", UnlockedBy: "Black Codex",
		Costs: []data.Resource{{
			Name: "necrocorn", Count: 5, CostExponentBase: 1.15,
		}, {
			Name: "megalith", Count: 1000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Unicorn Necropolis", UnlockedBy: "Unicorn Graveyard",
		Costs: []data.Resource{{
			Name: "necrocorn", Count: 15, CostExponentBase: 1.15,
		}, {
			Name: "megalith", Count: 2500, CostExponentBase: 1.15,
		}, {
			Name: "alicorn", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "void", Count: 5, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Black Pyramid", UnlockedBy: "Megalomania",
		Costs: []data.Resource{{
			Name: "sorrow", Count: 5, CostExponentBase: 1.15,
		}, {
			Name: "megalith", Count: 2500, CostExponentBase: 1.15,
		}, {
			Name: "spice", Count: 150000, CostExponentBase: 1.15,
		}, {
			Name: "unobtainium", Count: 5000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Solar Chant", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{Name: "faith", Count: 100, CostExponentBase: 2.5}},
	}, {
		Name: "Scholasticism", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{Name: "faith", Count: 250, CostExponentBase: 2.5}},
	}, {
		Name: "Golden Spire", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "faith", Count: 350, CostExponentBase: 2.5,
		}, {
			Name: "gold", Count: 150, CostExponentBase: 2.5,
		}},
	}, {
		Name: "Sun Altar", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "faith", Count: 500, CostExponentBase: 2.5,
		}, {
			Name: "gold", Count: 250, CostExponentBase: 2.5,
		}},
	}, {
		Name: "Stained Glass", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "faith", Count: 500, CostExponentBase: 2.5,
		}, {
			Name: "gold", Count: 250, CostExponentBase: 2.5,
		}},
	}, {
		Name: "Basilica", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "faith", Count: 1250, CostExponentBase: 2.5,
		}, {
			Name: "gold", Count: 750, CostExponentBase: 2.5,
		}},
	}, {
		Name: "Templars", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "faith", Count: 3500, CostExponentBase: 2.5,
		}, {
			Name: "gold", Count: 3000, CostExponentBase: 2.5,
		}},
	}, {
		Name: "Black Obelisk", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 100, CostExponentBase: 1.15}},
	}, {
		Name: "Black Nexus", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 5000, CostExponentBase: 1.15}},
	}, {
		Name: "Black Core", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 10000, CostExponentBase: 1.15}},
	}, {
		Name: "Event Horizon", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 25000, CostExponentBase: 1.15}},
	}, {
		Name: "Black Library", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 30000, CostExponentBase: 1.15}},
	}, {
		Name: "Black Radiance", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 37500, CostExponentBase: 1.15}},
	}, {
		Name: "Blazar", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{Name: "relic", Count: 50000, CostExponentBase: 1.15}},
	}, {
		Name: "Dark Nova", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{
			Name: "relic", Count: 75000, CostExponentBase: 1.15,
		}, {
			Name: "void", Count: 7500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Mausoleum", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{
			Name: "relic", Count: 50000, CostExponentBase: 1.15,
		}, {
			Name: "void", Count: 12500, CostExponentBase: 1.15,
		}, {
			Name: "necrocorn", Count: 10, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Holy Genocide", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{
			Name: "relic", Count: 100000, CostExponentBase: 1.15,
		}, {
			Name: "void", Count: 25000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Space Elevator", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "titanium", Count: 6000, CostExponentBase: 1.15,
		}, {
			Name: "science", Count: 75000, CostExponentBase: 1.15,
		}, {
			Name: "unobtainium", Count: 50, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Satellite", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "starchart", Count: 325, CostExponentBase: 1.08,
		}, {
			Name: "titanium", Count: 2500, CostExponentBase: 1.08,
		}, {
			Name: "science", Count: 100000, CostExponentBase: 1.08,
		}, {
			Name: "oil", Count: 15000, CostExponentBase: 1.05, Bonus: []data.Resource{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Space Station", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "starchart", Count: 425, CostExponentBase: 1.12,
		}, {
			Name: "alloy", Count: 750, CostExponentBase: 1.12,
		}, {
			Name: "science", Count: 150000, CostExponentBase: 1.12,
		}, {
			Name: "oil", Count: 35000, CostExponentBase: 1.05, Bonus: []data.Resource{{Name: "SpaceElevatorOilBonus"}},
		}},
		Adds: []data.Resource{{Name: "kitten", Cap: 2}},
	}, {
		Name: "Active Lunar Outpost", UnlockedBy: "Moon Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 650, CostExponentBase: 1.12,
		}, {
			Name: "uranium", Count: 500, CostExponentBase: 1.12,
		}, {
			Name: "alloy", Count: 750, CostExponentBase: 1.12,
		}, {
			Name: "concrete", Count: 150, CostExponentBase: 1.12,
		}, {
			Name: "science", Count: 100000, CostExponentBase: 1.12,
		}, {
			Name: "oil", Count: 55000, CostExponentBase: 1.05, Bonus: []data.Resource{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Moon Base", UnlockedBy: "Moon Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 700, CostExponentBase: 1.12,
		}, {
			Name: "titanium", Count: 9500, CostExponentBase: 1.12,
		}, {
			Name: "concrete", Count: 250, CostExponentBase: 1.12,
		}, {
			Name: "science", Count: 100000, CostExponentBase: 1.12,
		}, {
			Name: "unobtainium", Count: 50, CostExponentBase: 1.12,
		}, {
			Name: "oil", Count: 70000, CostExponentBase: 1.05, Bonus: []data.Resource{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Planet Cracker", UnlockedBy: "Dune Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 2500, CostExponentBase: 1.18,
		}, {
			Name: "alloy", Count: 1750, CostExponentBase: 1.18,
		}, {
			Name: "science", Count: 125000, CostExponentBase: 1.18,
		}, {
			Name: "kerosene", Count: 50, CostExponentBase: 1.18,
		}},
	}, {
		Name: "Hydraulic Fracturer", UnlockedBy: "Dune Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 750, CostExponentBase: 1.18,
		}, {
			Name: "alloy", Count: 1025, CostExponentBase: 1.18,
		}, {
			Name: "science", Count: 150000, CostExponentBase: 1.18,
		}, {
			Name: "kerosene", Count: 100, CostExponentBase: 1.18,
		}},
	}, {
		Name: "Spice Refinery", UnlockedBy: "Dune Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 500, CostExponentBase: 1.15,
		}, {
			Name: "alloy", Count: 500, CostExponentBase: 1.15,
		}, {
			Name: "science", Count: 75000, CostExponentBase: 1.15,
		}, {
			Name: "kerosene", Count: 125, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Research Vessel", UnlockedBy: "Piscine Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "alloy", Count: 2500, CostExponentBase: 1.15,
		}, {
			Name: "titanium", Count: 12500, CostExponentBase: 1.15,
		}, {
			Name: "kerosene", Count: 250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Orbital Array", UnlockedBy: "Piscine Mission",
		Costs: []data.Resource{{
			Name: "science", Count: 250000, CostExponentBase: 1.15,
		}, {
			Name: "eludium", Count: 100, CostExponentBase: 1.15,
		}, {
			Name: "kerosene", Count: 500, CostExponentBase: 1.15,
		}, {
			Name: "starchart", Count: 2000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Sunlifter", UnlockedBy: "Helios Mission",
		Costs: []data.Resource{{
			Name: "science", Count: 500000, CostExponentBase: 1.15,
		}, {
			Name: "eludium", Count: 225, CostExponentBase: 1.15,
		}, {
			Name: "kerosene", Count: 2500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Containment Chamber", UnlockedBy: "Helios Mission",
		Costs: []data.Resource{{
			Name: "science", Count: 500000, CostExponentBase: 1.125,
		}, {
			Name: "kerosene", Count: 2500, CostExponentBase: 1.125,
		}},
	}, {
		Name: "Heatsink", UnlockedBy: "Helios Mission",
		Costs: []data.Resource{{
			Name: "science", Count: 125000, CostExponentBase: 1.12,
		}, {
			Name: "thorium", Count: 12500, CostExponentBase: 1.12,
		}, {
			Name: "relic", Count: 1, CostExponentBase: 1.12,
		}, {
			Name: "kerosene", Count: 5000, CostExponentBase: 1.12,
		}},
	}, {
		Name: "Sunforge", UnlockedBy: "Helios Mission",
		Costs: []data.Resource{{
			Name: "science", Count: 100000, CostExponentBase: 1.12,
		}, {
			Name: "relic", Count: 1, CostExponentBase: 1.12,
		}, {
			Name: "kerosene", Count: 1250, CostExponentBase: 1.12,
		}, {
			Name: "antimatter", Count: 250, CostExponentBase: 1.12,
		}},
	}, {
		Name: "Cryostation", UnlockedBy: "T-Minus Mission",
		Costs: []data.Resource{{
			Name: "science", Count: 200000, CostExponentBase: 1.12,
		}, {
			Name: "eludium", Count: 25, CostExponentBase: 1.12,
		}, {
			Name: "concrete", Count: 1500, CostExponentBase: 1.12,
		}, {
			Name: "kerosene", Count: 500, CostExponentBase: 1.12,
		}},
	}, {
		Name: "Space Beacon", UnlockedBy: "Kairo Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 25000, CostExponentBase: 1.15,
		}, {
			Name: "antimatter", Count: 50, CostExponentBase: 1.15,
		}, {
			Name: "alloy", Count: 25000, CostExponentBase: 1.15,
		}, {
			Name: "kerosene", Count: 7500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Terraforming Station", UnlockedBy: "Terraformation",
		Costs: []data.Resource{{
			Name: "antimatter", Count: 25, CostExponentBase: 1.25,
		}, {
			Name: "uranium", Count: 5000, CostExponentBase: 1.25,
		}, {
			Name: "kerosene", Count: 5000, CostExponentBase: 1.25,
		}},
		Adds: []data.Resource{{
			Name: "kitten", Cap: 1,
			Bonus: []data.Resource{{Name: "Hydroponics", Factor: 0.01}},
		}},
	}, {
		Name: "Hydroponics", UnlockedBy: "Hydroponics Tech",
		Costs: []data.Resource{{
			Name: "unobtainium", Count: 1, CostExponentBase: 1.15,
		}, {
			Name: "kerosene", Count: 500, CostExponentBase: 1.15,
		}},
	}, {
		Name: "HR Harvester", UnlockedBy: "Umbra Mission",
		Costs: []data.Resource{{
			Name: "relic", Count: 25, CostExponentBase: 1.15,
		}, {
			Name: "antimatter", Count: 1250, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Entanglement Station", UnlockedBy: "Quantum Cryptography",
		Costs: []data.Resource{{
			Name: "relic", Count: 1250, CostExponentBase: 1.15,
		}, {
			Name: "antimatter", Count: 5250, CostExponentBase: 1.15,
		}, {
			Name: "eludium", Count: 5000, CostExponentBase: 1.15,
		}},
	}, {
		Name: "Tectonic", UnlockedBy: "Terraformation",
		Costs: []data.Resource{{
			Name: "science", Count: 600000, CostExponentBase: 1.25,
		}, {
			Name: "antimatter", Count: 500, CostExponentBase: 1.25,
		}, {
			Name: "thorium", Count: 75000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Molten Core", UnlockedBy: "Exophysics",
		Costs: []data.Resource{{
			Name: "science", Count: 25000000, CostExponentBase: 1.25,
		}, {
			Name: "uranium", Count: 5000000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Hash Level", UnlockedBy: "Entanglement Station",
		Costs: []data.Resource{{Name: "hash", Count: 1600, CostExponentBase: 1.6}},
	}, {
		Name: "Temporal Battery", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{Name: "time crystal", Count: 5, CostExponentBase: 1.25}},
	}, {
		Name: "Chrono Furnace", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 25, CostExponentBase: 1.25,
		}, {
			Name: "relic", Count: 5, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Time Boiler", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{Name: "time crystal", Count: 25000, CostExponentBase: 1.25}},
	}, {
		Name: "Temporal Accelerator", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 10, CostExponentBase: 1.25,
		}, {
			Name: "relic", Count: 1000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Time Impedance", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 100, CostExponentBase: 1.05,
		}, {
			Name: "relic", Count: 250, CostExponentBase: 1.05,
		}},
	}, {
		Name: "Resource Retrieval", UnlockedBy: "Paradox Theory",
		Costs: []data.Resource{{Name: "time crystal", Count: 1000, CostExponentBase: 1.3}},
	}, {
		Name: "Temporal Press", UnlockedBy: "Chronosurge",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 100, CostExponentBase: 1.1,
		}, {
			Name: "void", Count: 10, CostExponentBase: 1.1,
		}},
	}, {
		Name: "Cryochambers", UnlockedBy: "Void Space",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 2, CostExponentBase: 1.25,
		}, {
			Name: "void", Count: 100, CostExponentBase: 1.25,
		}, {
			Name: "karma", Count: 1, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Void Hoover", UnlockedBy: "Void Aspiration",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 10, CostExponentBase: 1.25,
		}, {
			Name: "void", Count: 250, CostExponentBase: 1.25,
		}, {
			Name: "antimatter", Count: 1000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Void Rift", UnlockedBy: "Void Aspiration",
		Costs: []data.Resource{{Name: "void", Count: 75, CostExponentBase: 1.3}},
	}, {
		Name: "Chronocontrol", UnlockedBy: "Paradox Theory",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 30, CostExponentBase: 1.25,
		}, {
			Name: "void", Count: 500, CostExponentBase: 1.25,
		}, {
			Name: "temporal flux", Count: 3000, CostExponentBase: 1.25,
		}},
	}, {
		Name: "Void Resonator", UnlockedBy: "Paradox Theory",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 1000, CostExponentBase: 1.25,
		}, {
			Name: "relic", Count: 10000, CostExponentBase: 1.25,
		}, {
			Name: "void", Count: 50, CostExponentBase: 1.25,
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
		Name: "Send hunters", Type: "Job", UnlockedBy: "Archery",
		Costs: []data.Resource{{Name: "catpower", Count: 100}},
		Adds: []data.Resource{{
			Name: "fur", Count: 39.5, Bonus: []data.Resource{{Name: "HuntingBonus"}},
		}, {
			Name: "ivory", Count: 10.78, Bonus: []data.Resource{{Name: "HuntingBonus"}},
		}, {
			Name: "unicorn", Count: 0.05,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Lizards", Type: "Trade", UnlockedBy: "Archery",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "mineral", Count: 1000,
		}},
		Adds: []data.Resource{{
			Name: "wood", Count: 500,
		}, {
			Name: "beam", Count: 10 * 0.15,
		}, {
			Name: "scaffold", Count: 1 * 0.10,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Sharks", Type: "Trade", UnlockedBy: "Archery",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "iron", Count: 100,
		}},
		Adds: []data.Resource{{
			Name: "catnip", Count: 35000,
		}, {
			Name: "parchment", Count: 5 * 0.25,
		}, {
			Name: "manuscript", Count: 3 * 0.15,
		}, {
			Name: "compendium", Count: 1 * 0.10,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Griffins", Type: "Trade", UnlockedBy: "Archery",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "wood", Count: 500,
		}},
		Adds: []data.Resource{{
			Name: "iron", Count: 250,
		}, {
			Name: "steel", Count: 25 * 0.25,
		}, {
			Name: "gear", Count: 5 * 0.10,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Nagas", Type: "Trade", UnlockedBy: "Writing",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "ivory", Count: 500,
		}},
		Adds: []data.Resource{{
			Name: "mineral", Count: 1000,
		}, {
			Name: "slab", Count: 5 * 0.75,
		}, {
			Name: "concrete", Count: 5 * 0.25,
		}, {
			Name: "megalith", Count: 1 * 0.10,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Zebras", Type: "Trade", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "slab", Count: 50,
		}},
		Adds: []data.Resource{{
			Name: "iron", Count: 300,
		}, {
			Name: "plate", Count: 2 * 0.65,
		}, {
			Name: "titanium", Count: 1.5 * 0.15,
			Bonus: []data.Resource{{Name: "ship", Factor: 0.03 * 0.0035}},
		}, {
			Name: "alloy", Count: 0.25 * 0.05,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Spiders", Type: "Trade", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "scaffold", Count: 50,
		}},
		Adds: []data.Resource{{
			Name: "coal", Count: 350,
		}, {
			Name: "oil", Count: 100 * 0.25,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Dragons", Type: "Trade", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "titanium", Count: 250,
		}},
		Adds: []data.Resource{{
			Name: "uranium", Count: 1 * 0.95,
		}, {
			Name: "thorium", Count: 1 * 0.50,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Leviathans", Type: "Trade", UnlockedBy: "Black Pyramid",
		Costs: []data.Resource{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "unobtainium", Count: 5000,
		}},
		Adds: []data.Resource{{
			Name: "starchart", Count: 250 * 0.50,
			Bonus: []data.Resource{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "time crystal", Count: 0.25 * 0.98,
			Bonus: []data.Resource{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "sorrow", Count: 1 * 0.15,
			Bonus: []data.Resource{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "relic", Count: 1 * 0.05,
			Bonus: []data.Resource{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Feed Leviathans", Type: "Craft", UnlockedBy: "Black Pyramid",
		Costs: []data.Resource{{Name: "necrocorn", Count: 1}},
		Adds:  []data.Resource{{Name: "leviathan energy", Count: 1}},
	}, {
		Name: "Sacrifice Unicorns", Type: "Craft", UnlockedBy: "Ziggurat",
		Costs: []data.Resource{{Name: "unicorn", Count: 2500}},
		Adds: []data.Resource{{
			Name: "tear", Count: 1,
			Bonus:               []data.Resource{{Name: "Ziggurat"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "Sacrifice Alicorns", Type: "Craft", UnlockedBy: "Ziggurat",
		Costs: []data.Resource{{Name: "alicorn", Count: 25}},
		Adds: []data.Resource{{
			Name: "time crystal", Count: 1,
			Bonus: []data.Resource{{
				Name: "Unicorn Utopia", Factor: 0.05,
			}, {
				Name: "Sunspire", Factor: 0.10,
			}},
		}},
	}, {
		Name: "Refine Tears", Type: "Craft", UnlockedBy: "Megalomania",
		Costs: []data.Resource{{Name: "tear", Count: 10000}},
		Adds:  []data.Resource{{Name: "sorrow", Count: 1}},
	}, {
		Name: "Refine Time Crystals", Type: "Craft", UnlockedBy: "Ziggurat",
		Costs: []data.Resource{{Name: "time crystal", Count: 25}},
		Adds: []data.Resource{{
			Name: "relic", Count: 1,
			Bonus: []data.Resource{{
				Name:                "Black Nexus",
				Bonus:               []data.Resource{{Name: "Black Pyramid"}},
				BonusStartsFromZero: true,
			}},
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Combust time crystal", Type: "Craft", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{Name: "time crystal", Count: 1}},
		Adds: []data.Resource{{
			Name: "day", Count: 400,
		}, {
			Name: "chronoheat", Count: 10,
		}},
	}, {
		Name: "Burn Chrono Furnace Fuel", Type: "Craft", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{Name: "chrono furnace fuel", Count: 100}},
		Adds:  []data.Resource{{Name: "day", Count: 400}},
	}, {
		Name: "Burn Paragon", Type: "Craft", UnlockedBy: "Chronoforge",
		Costs: []data.Resource{{Name: "paragon", Count: 1}},
		Adds:  []data.Resource{{Name: "burned paragon", Count: 1}},
	}})

	addTechs(g, []data.Action{{
		Name: "Calendar", UnlockedBy: "Library",
		Costs: []data.Resource{{Name: "science", Count: 30}},
	}, {
		Name: "Agriculture", UnlockedBy: "Calendar",
		Costs: []data.Resource{{Name: "science", Count: 100}},
	}, {
		Name: "Archery", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "science", Count: 300}},
	}, {
		Name: "Mining", UnlockedBy: "Agriculture",
		Costs: []data.Resource{{Name: "science", Count: 500}},
	}, {
		Name: "Animal Husbandry", UnlockedBy: "Archery",
		Costs: []data.Resource{{Name: "science", Count: 500}},
	}, {
		Name: "Metal Working", UnlockedBy: "Mining",
		Costs: []data.Resource{{Name: "science", Count: 900}},
	}, {
		Name: "Civil Service", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Count: 1500}},
	}, {
		Name: "Mathematics", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Count: 1000}},
	}, {
		Name: "Construction", UnlockedBy: "Animal Husbandry",
		Costs: []data.Resource{{Name: "science", Count: 1300}},
	}, {
		Name: "Currency", UnlockedBy: "Civil Service",
		Costs: []data.Resource{{Name: "science", Count: 2200}},
	}, {
		Name: "Celestial Mechanics", UnlockedBy: "Mathematics",
		Costs: []data.Resource{{Name: "science", Count: 250}},
	}, {
		Name: "Engineering", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "science", Count: 1500}},
	}, {
		Name: "Writing", UnlockedBy: "Engineering",
		Costs: []data.Resource{{Name: "science", Count: 3600}},
	}, {
		Name: "Philosophy", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Count: 9500}},
	}, {
		Name: "Steel", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Count: 12000}},
	}, {
		Name: "Machinery", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "science", Count: 15000}},
	}, {
		Name: "Theology", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "science", Count: 20000,
		}, {
			Name: "manuscript", Count: 35,
		}},
	}, {
		Name: "Astronomy", UnlockedBy: "Theology",
		Costs: []data.Resource{{
			Name: "science", Count: 28000,
		}, {
			Name: "manuscript", Count: 65,
		}},
	}, {
		Name: "Navigation", UnlockedBy: "Astronomy",
		Costs: []data.Resource{{
			Name: "science", Count: 35000,
		}, {
			Name: "manuscript", Count: 100,
		}},
	}, {
		Name: "Architecture", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 42000,
		}, {
			Name: "compendium", Count: 10,
		}},
	}, {
		Name: "Physics", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 50000,
		}, {
			Name: "compendium", Count: 35,
		}},
	}, {
		Name: "Metaphysics", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 55000,
		}, {
			Name: "unobtainium", Count: 5,
		}},
	}, {
		Name: "Chemistry", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 60000,
		}, {
			Name: "compendium", Count: 50,
		}},
	}, {
		Name: "Acoustics", UnlockedBy: "Architecture",
		Costs: []data.Resource{{
			Name: "science", Count: 60000,
		}, {
			Name: "compendium", Count: 60,
		}},
	}, {
		Name: "Geology", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 65000,
		}, {
			Name: "compendium", Count: 65,
		}},
	}, {
		Name: "Drama and Poetry", UnlockedBy: "Acoustics",
		Costs: []data.Resource{{
			Name: "science", Count: 90000,
		}, {
			Name: "parchment", Count: 5000,
		}},
	}, {
		Name: "Electricity", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "compendium", Count: 85,
		}},
	}, {
		Name: "Biology", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "science", Count: 85000,
		}, {
			Name: "compendium", Count: 100,
		}},
	}, {
		Name: "Biochemistry", UnlockedBy: "Biology",
		Costs: []data.Resource{{
			Name: "science", Count: 145000,
		}, {
			Name: "compendium", Count: 500,
		}},
	}, {
		Name: "Genetics", UnlockedBy: "Biochemistry",
		Costs: []data.Resource{{
			Name: "science", Count: 190000,
		}, {
			Name: "compendium", Count: 1500,
		}},
	}, {
		Name: "Industrialization", UnlockedBy: "Electricity",
		Costs: []data.Resource{{
			Name: "science", Count: 10000,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Mechanization", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Count: 115000,
		}, {
			Name: "blueprint", Count: 45,
		}},
	}, {
		Name: "Combustion", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Count: 115000,
		}, {
			Name: "blueprint", Count: 45,
		}},
	}, {
		Name: "Metallurgy", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "blueprint", Count: 60,
		}},
	}, {
		Name: "Ecology", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "blueprint", Count: 55,
		}},
	}, {
		Name: "Electronics", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Count: 135000,
		}, {
			Name: "blueprint", Count: 70,
		}},
	}, {
		Name: "Robotics", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 140000,
		}, {
			Name: "blueprint", Count: 80,
		}},
	}, {
		Name: "Artificial Intelligence", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "blueprint", Count: 150,
		}},
	}, {
		Name: "Quantum Cryptography", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Count: 1250000,
		}, {
			Name: "relic", Count: 1024,
		}},
	}, {
		Name: "Blackchain", UnlockedBy: "Quantum Cryptography",
		Costs: []data.Resource{{
			Name: "science", Count: 5000000,
		}, {
			Name: "relic", Count: 4096,
		}},
	}, {
		Name: "Nuclear Fission", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "blueprint", Count: 100,
		}},
	}, {
		Name: "Rocketry", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "blueprint", Count: 125,
		}},
	}, {
		Name: "Oil Processing", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Count: 215000,
		}, {
			Name: "blueprint", Count: 150,
		}},
	}, {
		Name: "Satellites", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Count: 190000,
		}, {
			Name: "blueprint", Count: 125,
		}},
	}, {
		Name: "Orbital Engineering", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "blueprint", Count: 250,
		}},
	}, {
		Name: "Thorium", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 375000,
		}, {
			Name: "blueprint", Count: 375,
		}},
	}, {
		Name: "Exogeology", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 275000,
		}, {
			Name: "blueprint", Count: 250,
		}},
	}, {
		Name: "Advanced Exogeology", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 325000,
		}, {
			Name: "blueprint", Count: 350,
		}},
	}, {
		Name: "Nanotechnology", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "blueprint", Count: 150,
		}},
	}, {
		Name: "Superconductors", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Count: 225000,
		}, {
			Name: "blueprint", Count: 175,
		}},
	}, {
		Name: "Antimatter", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Count: 500000,
		}, {
			Name: "relic", Count: 1,
		}},
	}, {
		Name: "Terraformation", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Count: 750000,
		}, {
			Name: "relic", Count: 5,
		}},
	}, {
		Name: "Hydroponics Tech", UnlockedBy: "Terraformation",
		Costs: []data.Resource{{
			Name: "science", Count: 1000000,
		}, {
			Name: "relic", Count: 25,
		}},
	}, {
		Name: "Exophysics", UnlockedBy: "Hydroponics Tech",
		Costs: []data.Resource{{
			Name: "science", Count: 25000000,
		}, {
			Name: "relic", Count: 500,
		}},
	}, {
		Name: "Particle Physics", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Count: 185000,
		}, {
			Name: "blueprint", Count: 135,
		}},
	}, {
		Name: "Dimensional Physics", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{Name: "science", Count: 235000}},
	}, {
		Name: "Chronophysics", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "time crystal", Count: 5,
		}},
	}, {
		Name: "Tachyon Theory", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "science", Count: 750000,
		}, {
			Name: "time crystal", Count: 25,
		}, {
			Name: "relic", Count: 1,
		}},
	}, {
		Name: "Cryptotheology", UnlockedBy: "Theology",
		Costs: []data.Resource{{
			Name: "science", Count: 650000,
		}, {
			Name: "relic", Count: 5,
		}},
	}, {
		Name: "Void Space", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Count: 800000,
		}, {
			Name: "time crystal", Count: 30,
		}, {
			Name: "void", Count: 100,
		}},
	}, {
		Name: "Paradox Theory", UnlockedBy: "Void Space",
		Costs: []data.Resource{{
			Name: "science", Count: 1000000,
		}, {
			Name: "time crystal", Count: 40,
		}, {
			Name: "void", Count: 250,
		}},
	}, {
		Name: "Mineral Hoes", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "mineral", Count: 275,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Iron Hoes", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Count: 25,
		}, {
			Name: "science", Count: 200,
		}},
	}, {
		Name: "Mineral Axe", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "mineral", Count: 500,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Iron Axe", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Count: 50,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Steel Axe", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "steel", Count: 75,
		}, {
			Name: "science", Count: 20000,
		}},
	}, {
		Name: "Reinforced Saw", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "iron", Count: 1000,
		}, {
			Name: "science", Count: 2500,
		}},
	}, {
		Name: "Steel Saw", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "steel", Count: 750,
		}, {
			Name: "science", Count: 52000,
		}},
	}, {
		Name: "Titanium Saw", UnlockedBy: "Steel Saw",
		Costs: []data.Resource{{
			Name: "titanium", Count: 500,
		}, {
			Name: "science", Count: 70000,
		}},
	}, {
		Name: "Alloy Saw", UnlockedBy: "Titanium Saw",
		Costs: []data.Resource{{
			Name: "alloy", Count: 75,
		}, {
			Name: "science", Count: 85000,
		}},
	}, {
		Name: "Titanium Axe", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 38000,
		}, {
			Name: "titanium", Count: 10,
		}},
	}, {
		Name: "Alloy Axe", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Count: 70000,
		}, {
			Name: "alloy", Count: 25,
		}},
	}, {
		Name: "Expanded Barns", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "wood", Count: 1000,
		}, {
			Name: "mineral", Count: 750,
		}, {
			Name: "iron", Count: 50,
		}, {
			Name: "science", Count: 500,
		}},
	}, {
		Name: "Reinforced Barns", UnlockedBy: "Workshop",
		Costs: []data.Resource{{
			Name: "iron", Count: 100,
		}, {
			Name: "science", Count: 800,
		}, {
			Name: "beam", Count: 25,
		}, {
			Name: "slab", Count: 10,
		}},
	}, {
		Name: "Reinforced Warehouses", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Count: 15000,
		}, {
			Name: "plate", Count: 50,
		}, {
			Name: "steel", Count: 50,
		}, {
			Name: "scaffold", Count: 25,
		}},
	}, {
		Name: "Titanium Barns", UnlockedBy: "Reinforced Barns",
		Costs: []data.Resource{{
			Name: "science", Count: 60000,
		}, {
			Name: "titanium", Count: 25,
		}, {
			Name: "steel", Count: 200,
		}, {
			Name: "scaffold", Count: 250,
		}},
	}, {
		Name: "Alloy Barns", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "alloy", Count: 20,
		}, {
			Name: "plate", Count: 750,
		}},
	}, {
		Name: "Concrete Barns", UnlockedBy: "Concrete Pillars",
		Costs: []data.Resource{{
			Name: "science", Count: 2000,
		}, {
			Name: "concrete", Count: 45,
		}, {
			Name: "titanium", Count: 2000,
		}},
	}, {
		Name: "Titanium Warehouses", UnlockedBy: "Silos",
		Costs: []data.Resource{{
			Name: "science", Count: 70000,
		}, {
			Name: "titanium", Count: 50,
		}, {
			Name: "steel", Count: 500,
		}, {
			Name: "scaffold", Count: 500,
		}},
	}, {
		Name: "Alloy Warehouses", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Count: 90000,
		}, {
			Name: "titanium", Count: 750,
		}, {
			Name: "alloy", Count: 50,
		}},
	}, {
		Name: "Concrete Warehouses", UnlockedBy: "Concrete Pillars",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 1250,
		}, {
			Name: "concrete", Count: 35,
		}},
	}, {
		Name: "Storage Bunkers", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 25000,
		}, {
			Name: "unobtainium", Count: 500,
		}, {
			Name: "concrete", Count: 1250,
		}},
	}, {
		Name: "Energy Rifts", UnlockedBy: "Dimensional Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "titanium", Count: 7500,
		}, {
			Name: "uranium", Count: 250,
		}},
	}, {
		Name: "Stasis Chambers", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "science", Count: 235000,
		}, {
			Name: "alloy", Count: 200,
		}, {
			Name: "uranium", Count: 2000,
		}, {
			Name: "time crystal", Count: 1,
		}},
	}, {
		Name: "Void Energy", UnlockedBy: "Stasis Chambers",
		Costs: []data.Resource{{
			Name: "science", Count: 275000,
		}, {
			Name: "alloy", Count: 250,
		}, {
			Name: "uranium", Count: 2500,
		}, {
			Name: "time crystal", Count: 2,
		}},
	}, {
		Name: "Dark Energy", UnlockedBy: "Void Energy",
		Costs: []data.Resource{{
			Name: "science", Count: 350000,
		}, {
			Name: "eludium", Count: 75,
		}, {
			Name: "time crystal", Count: 3,
		}},
	}, {
		Name: "Chronoforge", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Count: 500000,
		}, {
			Name: "relic", Count: 5,
		}, {
			Name: "time crystal", Count: 10,
		}},
	}, {
		Name: "Tachyon Accelerators", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Count: 500000,
		}, {
			Name: "eludium", Count: 125,
		}, {
			Name: "time crystal", Count: 10,
		}},
	}, {
		Name: "Flux Condensator", UnlockedBy: "Chronophysics",
		Costs: []data.Resource{{
			Name: "alloy", Count: 250,
		}, {
			Name: "unobtainium", Count: 5000,
		}, {
			Name: "time crystal", Count: 5,
		}},
	}, {
		Name: "LHC", UnlockedBy: "Dimensional Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 100,
		}, {
			Name: "alloy", Count: 150,
		}},
	}, {
		Name: "Photovoltaic Cells", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "titanium", Count: 5000,
		}},
	}, {
		Name: "Thin Film Cells", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "unobtainium", Count: 200,
		}, {
			Name: "uranium", Count: 1000,
		}},
	}, {
		Name: "Quantum Dot Cells", UnlockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "eludium", Count: 200,
		}, {
			Name: "thorium", Count: 1000,
		}},
	}, {
		Name: "Solar Satellites", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 225000,
		}, {
			Name: "alloy", Count: 750,
		}},
	}, {
		Name: "Expanded Cargo", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 55000,
		}, {
			Name: "blueprint", Count: 15,
		}},
	}, {
		Name: "Barges", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 1500,
		}, {
			Name: "blueprint", Count: 30,
		}},
	}, {
		Name: "Reactor Vessel", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Count: 135000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "uranium", Count: 125,
		}},
	}, {
		Name: "Ironwood Huts", UnlockedBy: "Reinforced Warehouses",
		Costs: []data.Resource{{
			Name: "science", Count: 30000,
		}, {
			Name: "wood", Count: 15000,
		}, {
			Name: "iron", Count: 3000,
		}},
	}, {
		Name: "Concrete Huts", UnlockedBy: "Concrete Pillars",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "concrete", Count: 45,
		}, {
			Name: "titanium", Count: 3000,
		}},
	}, {
		Name: "Unobtainium Huts", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "unobtainium", Count: 350,
		}, {
			Name: "titanium", Count: 15000,
		}},
	}, {
		Name: "Eludium Huts", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 275000,
		}, {
			Name: "eludium", Count: 125,
		}},
	}, {
		Name: "Silos", UnlockedBy: "Ironwood Huts",
		Costs: []data.Resource{{
			Name: "science", Count: 50000,
		}, {
			Name: "steel", Count: 125,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "Refrigeration", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 2500,
		}, {
			Name: "blueprint", Count: 15,
		}},
	}, {
		Name: "Composite Bow", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "science", Count: 500,
		}, {
			Name: "iron", Count: 100,
		}, {
			Name: "wood", Count: 200,
		}},
	}, {
		Name: "Crossbow", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "science", Count: 12000,
		}, {
			Name: "iron", Count: 1500,
		}},
	}, {
		Name: "Railgun", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Bolas", UnlockedBy: "Mining",
		Costs: []data.Resource{{
			Name: "science", Count: 1000,
		}, {
			Name: "mineral", Count: 250,
		}, {
			Name: "wood", Count: 50,
		}},
	}, {
		Name: "Hunting Armour", UnlockedBy: "Metal Working",
		Costs: []data.Resource{{
			Name: "science", Count: 2000,
		}, {
			Name: "iron", Count: 750,
		}},
	}, {
		Name: "Steel Armour", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Count: 10000,
		}, {
			Name: "steel", Count: 50,
		}},
	}, {
		Name: "Alloy Armour", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "science", Count: 50000,
		}, {
			Name: "alloy", Count: 25,
		}},
	}, {
		Name: "Nanosuits", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Count: 185000,
		}, {
			Name: "alloy", Count: 250,
		}},
	}, {
		Name: "Caravanserai", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 25000,
		}, {
			Name: "ivory", Count: 10000,
		}, {
			Name: "gold", Count: 250,
		}},
	}, {
		Name: "Catnip Enrichment", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "science", Count: 500,
		}, {
			Name: "catnip", Count: 5000,
		}},
	}, {
		Name: "Gold Ore", UnlockedBy: "Currency",
		Costs: []data.Resource{{
			Name: "science", Count: 1000,
		}, {
			Name: "mineral", Count: 800,
		}, {
			Name: "iron", Count: 100,
		}},
	}, {
		Name: "Geodesy", UnlockedBy: "Geology",
		Costs: []data.Resource{{
			Name: "science", Count: 90000,
		}, {
			Name: "titanium", Count: 250,
		}, {
			Name: "starchart", Count: 500,
		}},
	}, {
		Name: "Register", UnlockedBy: "Writing",
		Costs: []data.Resource{{
			Name: "science", Count: 500,
		}, {
			Name: "gold", Count: 10,
		}},
	}, {
		Name: "Concrete Pillars", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "concrete", Count: 50,
		}},
	}, {
		Name: "Mining Drill", UnlockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 1750,
		}, {
			Name: "steel", Count: 750,
		}},
	}, {
		Name: "Unobtainium Drill", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 250,
		}, {
			Name: "alloy", Count: 1250,
		}},
	}, {
		Name: "Coal Furnace", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Count: 5000,
		}, {
			Name: "mineral", Count: 5000,
		}, {
			Name: "iron", Count: 2000,
		}, {
			Name: "beam", Count: 35,
		}},
	}, {
		Name: "Deep Mining", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Count: 5000,
		}, {
			Name: "iron", Count: 1200,
		}, {
			Name: "beam", Count: 50,
		}},
	}, {
		Name: "Pyrolysis", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 35000,
		}, {
			Name: "compendium", Count: 5,
		}},
	}, {
		Name: "Electrolytic Smelting", UnlockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 2000,
		}},
	}, {
		Name: "Oxidation", UnlockedBy: "Metallurgy",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "steel", Count: 5000,
		}},
	}, {
		Name: "Steel Plants", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Count: 140000,
		}, {
			Name: "titanium", Count: 3500,
		}, {
			Name: "gear", Count: 750,
		}},
	}, {
		Name: "Automated Plants", UnlockedBy: "Steel Plants",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "alloy", Count: 750,
		}},
	}, {
		Name: "Nuclear Plants", UnlockedBy: "Automated Plants",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "uranium", Count: 10000,
		}},
	}, {
		Name: "Rotary Kiln", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Count: 145000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "gear", Count: 500,
		}},
	}, {
		Name: "Fluoridized Reactors", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "alloy", Count: 200,
		}},
	}, {
		Name: "Nuclear Smelter", UnlockedBy: "Nuclear Fission",
		Costs: []data.Resource{{
			Name: "science", Count: 165000,
		}, {
			Name: "uranium", Count: 250,
		}},
	}, {
		Name: "Orbital Geodesy", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "alloy", Count: 1000,
		}, {
			Name: "oil", Count: 35000,
		}},
	}, {
		Name: "Printing Press", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "science", Count: 7500,
		}, {
			Name: "gear", Count: 45,
		}},
	}, {
		Name: "Offset Press", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "oil", Count: 15000,
		}},
	}, {
		Name: "Photolithography", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "alloy", Count: 1250,
		}, {
			Name: "oil", Count: 50000,
		}, {
			Name: "uranium", Count: 250,
		}},
	}, {
		Name: "Uplink", UnlockedBy: "Satellites",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "alloy", Count: 1750,
		}},
	}, {
		Name: "Starlink", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "alloy", Count: 5000,
		}, {
			Name: "oil", Count: 25000,
		}},
	}, {
		Name: "Cryocomputing", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "eludium", Count: 15,
		}},
	}, {
		Name: "Machine Learning", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "eludium", Count: 25,
		}, {
			Name: "antimatter", Count: 125,
		}},
	}, {
		Name: "Workshop Automation", UnlockedBy: "Machinery",
		Costs: []data.Resource{{
			Name: "science", Count: 10000,
		}, {
			Name: "gear", Count: 25,
		}},
	}, {
		Name: "Advanced Automation", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 75,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Pneumatic Press", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 20000,
		}, {
			Name: "gear", Count: 30,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "High Pressure Engine", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "science", Count: 20000,
		}, {
			Name: "gear", Count: 25,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "Fuel Injector", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "oil", Count: 20000,
		}},
	}, {
		Name: "Factory Logistics", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "titanium", Count: 2000,
		}},
	}, {
		Name: "Carbon Sequestration", UnlockedBy: "Ecology",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "titanium", Count: 1250,
		}, {
			Name: "gear", Count: 125,
		}, {
			Name: "steel", Count: 4000,
		}, {
			Name: "alloy", Count: 1000,
		}},
	}, {
		Name: "Space Manufacturing", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "titanium", Count: 125000,
		}},
	}, {
		Name: "Astrolabe", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 25000,
		}, {
			Name: "titanium", Count: 5,
		}, {
			Name: "starchart", Count: 75,
		}},
	}, {
		Name: "Titanium Reflectors", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "science", Count: 20000,
		}, {
			Name: "titanium", Count: 15,
		}, {
			Name: "starchart", Count: 20,
		}},
	}, {
		Name: "Unobtainium Reflectors", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 75,
		}, {
			Name: "starchart", Count: 750,
		}},
	}, {
		Name: "Eludium Reflectors", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "eludium", Count: 15,
		}},
	}, {
		Name: "Hydro Plant Turbines", UnlockedBy: "Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 125,
		}},
	}, {
		Name: "Antimatter Bases", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "eludium", Count: 15,
		}, {
			Name: "antimatter", Count: 250,
		}},
	}, {
		Name: "AI Bases", UnlockedBy: "Antimatter Bases",
		Costs: []data.Resource{{
			Name: "science", Count: 750000,
		}, {
			Name: "antimatter", Count: 7500,
		}},
	}, {
		Name: "Antimatter Fission", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Count: 525000,
		}, {
			Name: "antimatter", Count: 175,
		}, {
			Name: "thorium", Count: 7500,
		}},
	}, {
		Name: "Antimatter Drive", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "science", Count: 450000,
		}, {
			Name: "antimatter", Count: 125,
		}},
	}, {
		Name: "Antimatter Reactors", UnlockedBy: "Antimatter",
		Costs: []data.Resource{{
			Name: "eludium", Count: 35,
		}, {
			Name: "antimatter", Count: 750,
		}},
	}, {
		Name: "Advanced AM Reactors", UnlockedBy: "Antimatter Reactors",
		Costs: []data.Resource{{
			Name: "eludium", Count: 70,
		}, {
			Name: "antimatter", Count: 1750,
		}},
	}, {
		Name: "Void Reactors", UnlockedBy: "Advanced AM Reactors",
		Costs: []data.Resource{{
			Name: "void", Count: 250,
		}, {
			Name: "antimatter", Count: 2500,
		}},
	}, {
		Name: "Relic Station", UnlockedBy: "Cryptotheology",
		Costs: []data.Resource{{
			Name: "eludium", Count: 100,
		}, {
			Name: "antimatter", Count: 5000,
		}},
	}, {
		Name: "Pumpjack", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 250,
		}, {
			Name: "gear", Count: 125,
		}},
	}, {
		Name: "Biofuel Processing", UnlockedBy: "Biochemistry",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 1250,
		}},
	}, {
		Name: "Unicorn Selection", UnlockedBy: "Genetics",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 1500,
		}},
	}, {
		Name: "GM Catnip", UnlockedBy: "Genetics",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 1500,
		}, {
			Name: "catnip", Count: 1000000,
		}},
	}, {
		Name: "CAD System", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 750,
		}},
	}, {
		Name: "SETI", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 250,
		}},
	}, {
		Name: "Logistics", UnlockedBy: "Industrialization",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 100,
		}, {
			Name: "scaffold", Count: 1000,
		}},
	}, {
		Name: "Augmentations", UnlockedBy: "Nanotechnology",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "uranium", Count: 50,
		}},
	}, {
		Name: "Cold Fusion", UnlockedBy: "Superconductors",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "eludium", Count: 25,
		}},
	}, {
		Name: "Thorium Reactors", UnlockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Count: 400000,
		}, {
			Name: "thorium", Count: 10000,
		}},
	}, {
		Name: "Enriched Uranium", UnlockedBy: "Particle Physics",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 7500,
		}, {
			Name: "uranium", Count: 150,
		}},
	}, {
		Name: "Enriched Thorium", UnlockedBy: "Thorium Reactors",
		Costs: []data.Resource{{
			Name: "science", Count: 12500,
		}, {
			Name: "thorium", Count: 500000,
		}},
	}, {
		Name: "Oil Refinery", UnlockedBy: "Combustion",
		Costs: []data.Resource{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 1250,
		}, {
			Name: "gear", Count: 500,
		}},
	}, {
		Name: "Hubble Space Telescope", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "alloy", Count: 1250,
		}, {
			Name: "oil", Count: 50000,
		}},
	}, {
		Name: "Satellite Navigation", UnlockedBy: "Hubble Space Telescope",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "alloy", Count: 750,
		}},
	}, {
		Name: "Satellite Radio", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 225000,
		}, {
			Name: "alloy", Count: 5000,
		}},
	}, {
		Name: "Astrophysicists", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 350,
		}},
	}, {
		Name: "Microwarp Reactors", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "eludium", Count: 50,
		}},
	}, {
		Name: "Planet Busters", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "science", Count: 275000,
		}, {
			Name: "eludium", Count: 250,
		}},
	}, {
		Name: "Thorium Drive", UnlockedBy: "Thorium",
		Costs: []data.Resource{{
			Name: "science", Count: 400000,
		}, {
			Name: "ship", Count: 10000,
		}, {
			Name: "gear", Count: 40000,
		}, {
			Name: "alloy", Count: 2000,
		}, {
			Name: "thorium", Count: 100000,
		}},
	}, {
		Name: "Oil Distillation", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 5000,
		}},
	}, {
		Name: "Factory Processing", UnlockedBy: "Oil Processing",
		Costs: []data.Resource{{
			Name: "science", Count: 195000,
		}, {
			Name: "titanium", Count: 7500,
		}, {
			Name: "concrete", Count: 125,
		}},
	}, {
		Name: "Factory Optimization", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "titanium", Count: 1250,
		}},
	}, {
		Name: "Space Engineers", UnlockedBy: "Orbital Engineering",
		Costs: []data.Resource{{
			Name: "science", Count: 225000,
		}, {
			Name: "alloy", Count: 200,
		}},
	}, {
		Name: "AI Engineers", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Count: 35000,
		}, {
			Name: "eludium", Count: 50,
		}, {
			Name: "antimatter", Count: 500,
		}},
	}, {
		Name: "Chronoengineers", UnlockedBy: "Tachyon Theory",
		Costs: []data.Resource{{
			Name: "science", Count: 500000,
		}, {
			Name: "time crystal", Count: 5,
		}, {
			Name: "eludium", Count: 100,
		}},
	}, {
		Name: "Telecommunication", UnlockedBy: "Electronics",
		Costs: []data.Resource{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "uranium", Count: 50,
		}},
	}, {
		Name: "Neural Network", UnlockedBy: "Artificial Intelligence",
		Costs: []data.Resource{{
			Name: "science", Count: 200000,
		}, {
			Name: "titanium", Count: 7500,
		}},
	}, {
		Name: "Robotic Assistance", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Count: 100000,
		}, {
			Name: "steel", Count: 10000,
		}, {
			Name: "gear", Count: 250,
		}},
	}, {
		Name: "Factory Robotics", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "science", Count: 75000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "titanium", Count: 1250,
		}},
	}, {
		Name: "Void Aspiration", UnlockedBy: "Void Energy",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 15,
		}, {
			Name: "antimatter", Count: 2000,
		}},
	}, {
		Name: "Distortion", UnlockedBy: "Paradox Theory",
		Costs: []data.Resource{{
			Name: "science", Count: 300000,
		}, {
			Name: "time crystal", Count: 25,
		}, {
			Name: "antimatter", Count: 2000,
		}, {
			Name: "void", Count: 1000,
		}},
	}, {
		Name: "Chronosurge", UnlockedBy: "Chronocontrol",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 25,
		}, {
			Name: "unobtainium", Count: 100000,
		}, {
			Name: "void", Count: 750,
		}, {
			Name: "temporal flux", Count: 6500,
		}},
	}, {
		Name: "Invisible Black Hand", UnlockedBy: "Blackchain",
		Costs: []data.Resource{{
			Name: "time crystal", Count: 128,
		}, {
			Name: "blackcoin", Count: 64,
		}, {
			Name: "void", Count: 32,
		}, {
			Name: "temporal flux", Count: 4096,
		}},
	}, {
		Name: "Orbital Launch", UnlockedBy: "Rocketry",
		Costs: []data.Resource{{
			Name: "starchart", Count: 250,
		}, {
			Name: "catpower", Count: 5000,
		}, {
			Name: "science", Count: 100000,
		}, {
			Name: "oil", Count: 15000, Bonus: []data.Resource{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Moon Mission", UnlockedBy: "Orbital Launch",
		Costs: []data.Resource{{
			Name: "starchart", Count: 500,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "science", Count: 125000,
		}, {
			Name: "oil", Count: 45000, Bonus: []data.Resource{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Dune Mission", UnlockedBy: "Moon Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 1000,
		}, {
			Name: "titanium", Count: 7000,
		}, {
			Name: "science", Count: 175000,
		}, {
			Name: "kerosene", Count: 75,
		}},
	}, {
		Name: "Piscine Mission", UnlockedBy: "Moon Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 1500,
		}, {
			Name: "titanium", Count: 9000,
		}, {
			Name: "science", Count: 200000,
		}, {
			Name: "kerosene", Count: 250,
		}},
	}, {
		Name: "Helios Mission", UnlockedBy: "Dune Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 3000,
		}, {
			Name: "titanium", Count: 15000,
		}, {
			Name: "science", Count: 250000,
		}, {
			Name: "kerosene", Count: 1250,
		}},
	}, {
		Name: "T-Minus Mission", UnlockedBy: "Piscine Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 2500,
		}, {
			Name: "titanium", Count: 12000,
		}, {
			Name: "science", Count: 225000,
		}, {
			Name: "kerosene", Count: 750,
		}},
	}, {
		Name: "Kairo Mission", UnlockedBy: "T-Minus Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 5000,
		}, {
			Name: "titanium", Count: 20000,
		}, {
			Name: "science", Count: 300000,
		}, {
			Name: "kerosene", Count: 7500,
		}},
	}, {
		Name: "Rorschach Mission", UnlockedBy: "Kairo Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 15000,
		}, {
			Name: "titanium", Count: 80000,
		}, {
			Name: "science", Count: 500000,
		}, {
			Name: "kerosene", Count: 25000,
		}},
	}, {
		Name: "Yarn Mission", UnlockedBy: "Helios Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 7500,
		}, {
			Name: "titanium", Count: 35000,
		}, {
			Name: "science", Count: 350000,
		}, {
			Name: "kerosene", Count: 12000,
		}},
	}, {
		Name: "Umbra Mission", UnlockedBy: "Yarn Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 25000,
		}, {
			Name: "science", Count: 500000,
		}, {
			Name: "kerosene", Count: 25000,
		}, {
			Name: "thorium", Count: 15000,
		}},
	}, {
		Name: "Charon Mission", UnlockedBy: "Umbra Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 75000,
		}, {
			Name: "science", Count: 750000,
		}, {
			Name: "kerosene", Count: 35000,
		}, {
			Name: "thorium", Count: 35000,
		}},
	}, {
		Name: "Centaurus Mission", UnlockedBy: "Rorschach Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 100000,
		}, {
			Name: "titanium", Count: 40000,
		}, {
			Name: "science", Count: 400000,
		}, {
			Name: "kerosene", Count: 50000,
		}, {
			Name: "thorium", Count: 50000,
		}},
	}, {
		Name: "Furthest Ring", UnlockedBy: "Centaurus Mission",
		Costs: []data.Resource{{
			Name: "starchart", Count: 500000,
		}, {
			Name: "science", Count: 1250000,
		}, {
			Name: "kerosene", Count: 75000,
		}, {
			Name: "thorium", Count: 75000,
		}},
	}, {
		Name: "Enlightenment", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 5}},
	}, {
		Name: "Codex Vox", UnlockedBy: "Enlightenment",
		Costs: []data.Resource{{Name: "paragon", Count: 25}},
	}, {
		Name: "Codex Logos", UnlockedBy: "Codex Vox",
		Costs: []data.Resource{{Name: "paragon", Count: 50}},
	}, {
		Name: "Codex Agrum", UnlockedBy: "Codex Logos",
		Costs: []data.Resource{{Name: "paragon", Count: 75}},
	}, {
		Name: "Megalomania", UnlockedBy: "Enlightenment",
		Costs: []data.Resource{{Name: "paragon", Count: 10}},
	}, {
		Name: "Black Codex", UnlockedBy: "Megalomania",
		Costs: []data.Resource{{Name: "paragon", Count: 25}},
	}, {
		Name: "Codex Leviathanus", UnlockedBy: "Codex Logos",
		Costs: []data.Resource{{Name: "paragon", Count: 75}},
	}, {
		Name: "Golden Ratio", UnlockedBy: "Enlightenment",
		Costs: []data.Resource{{Name: "paragon", Count: 50}},
	}, {
		Name: "Divine Proportion", UnlockedBy: "Golden Ratio",
		Costs: []data.Resource{{Name: "paragon", Count: 100}},
	}, {
		Name: "Vitruvian Feline", UnlockedBy: "Divine Proportion",
		Costs: []data.Resource{{Name: "paragon", Count: 250}},
	}, {
		Name: "Renaissance", UnlockedBy: "Vitruvian Feline",
		Costs: []data.Resource{{Name: "paragon", Count: 750}},
	}, {
		Name: "Diplomacy", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 5}},
	}, {
		Name: "Zebra Diplomacy", UnlockedBy: "Diplomacy",
		Costs: []data.Resource{{Name: "paragon", Count: 35}},
	}, {
		Name: "Zebra Covenant", UnlockedBy: "Zebra Diplomacy",
		Costs: []data.Resource{{Name: "paragon", Count: 75}},
	}, {
		Name: "Navigation Diplomacy", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 300}},
	}, {
		Name: "Chronomancy", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 25}},
	}, {
		Name: "Anachronomancy", UnlockedBy: "Chronomancy",
		Costs: []data.Resource{{Name: "paragon", Count: 125}},
	}, {
		Name: "Astromancy", UnlockedBy: "Chronomancy",
		Costs: []data.Resource{{Name: "paragon", Count: 50}},
	}, {
		Name: "Unicornmancy", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 125}},
	}, {
		Name: "Carnivals", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 25}},
	}, {
		Name: "Numerology", UnlockedBy: "Carnivals",
		Costs: []data.Resource{{Name: "paragon", Count: 50}},
	}, {
		Name: "Order of the Void", UnlockedBy: "Numerology",
		Costs: []data.Resource{{Name: "paragon", Count: 75}},
	}, {
		Name: "Venus of Willenfluff", UnlockedBy: "Numerology",
		Costs: []data.Resource{{Name: "paragon", Count: 150}},
	}, {
		Name: "Pawgan Rituals", UnlockedBy: "Venus of Willenfluff",
		Costs: []data.Resource{{Name: "paragon", Count: 400}},
	}, {
		Name: "Numeromancy", UnlockedBy: "Numerology",
		Costs: []data.Resource{{Name: "paragon", Count: 250}},
	}, {
		Name: "Malkuth", UnlockedBy: "Numeromancy",
		Costs: []data.Resource{{Name: "paragon", Count: 500}},
	}, {
		Name: "Yesod", UnlockedBy: "Malkuth",
		Costs: []data.Resource{{Name: "paragon", Count: 750}},
	}, {
		Name: "Hod", UnlockedBy: "Yesod",
		Costs: []data.Resource{{Name: "paragon", Count: 1250}},
	}, {
		Name: "Netzach", UnlockedBy: "Hod",
		Costs: []data.Resource{{Name: "paragon", Count: 1750}},
	}, {
		Name: "Tiferet", UnlockedBy: "Netzach",
		Costs: []data.Resource{{Name: "paragon", Count: 2500}},
	}, {
		Name: "Gevurah", UnlockedBy: "Tiferet",
		Costs: []data.Resource{{Name: "paragon", Count: 5000}},
	}, {
		Name: "Chesed", UnlockedBy: "Gevurah",
		Costs: []data.Resource{{Name: "paragon", Count: 7500}},
	}, {
		Name: "Binah", UnlockedBy: "Chesed",
		Costs: []data.Resource{{Name: "paragon", Count: 15000}},
	}, {
		Name: "Chokhmah", UnlockedBy: "Binah",
		Costs: []data.Resource{{Name: "paragon", Count: 30000}},
	}, {
		Name: "Keter", UnlockedBy: "Chokhmah",
		Costs: []data.Resource{{Name: "paragon", Count: 60000}},
	}, {
		Name: "Adjustment Bureau", UnlockedBy: "Metaphysics",
		Costs: []data.Resource{{Name: "paragon", Count: 5}},
	}, {
		Name: "ASCOH", UnlockedBy: "Adjustment Bureau",
		Costs: []data.Resource{{Name: "paragon", Count: 5}},
	}})

	addCrafts(g, []data.Action{{
		Name: "beam", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "wood", Count: 175}},
	}, {
		Name: "slab", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "mineral", Count: 250}},
	}, {
		Name: "concrete", UnlockedBy: "Mechanization",
		Costs: []data.Resource{{
			Name: "slab", Count: 2500,
		}, {
			Name: "steel", Count: 25,
		}},
	}, {
		Name: "plate", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "iron", Count: 125}},
	}, {
		Name: "steel", UnlockedBy: "Steel",
		Costs: []data.Resource{{
			Name: "iron", Count: 100,
		}, {
			Name: "coal", Count: 100,
		}},
		Producers: []data.Resource{{
			Name: "Active Calciner", Factor: 0.15 * 5 * 0.10 * 0.01,
			Bonus: []data.Resource{{
				Name: "Steel Plants",
				Bonus: []data.Resource{{
					Name: "Oxidation", Factor: 0.95,
				}, {
					Name: "CraftRatio", Factor: 0.25,
					Bonus:               []data.Resource{{Name: "Automated Plants"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Reactor", Factor: 0.02,
					Bonus:               []data.Resource{{Name: "Nuclear Plants"}},
					BonusStartsFromZero: true,
				}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "gear", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "steel", Count: 15}},
	}, {
		Name: "alloy", UnlockedBy: "Chemistry",
		Costs: []data.Resource{{
			Name: "steel", Count: 75,
		}, {
			Name: "titanium", Count: 10,
		}},
	}, {
		Name: "eludium", UnlockedBy: "Advanced Exogeology",
		Costs: []data.Resource{{
			Name: "alloy", Count: 2500,
		}, {
			Name: "unobtainium", Count: 1000,
		}},
	}, {
		Name: "scaffold", UnlockedBy: "Construction",
		Costs: []data.Resource{{Name: "beam", Count: 50}},
	}, {
		Name: "ship", UnlockedBy: "Navigation",
		Costs: []data.Resource{{
			Name: "scaffold", Count: 100,
		}, {
			Name: "plate", Count: 150,
		}, {
			Name: "starchart", Count: 25,
			Bonus: []data.Resource{{
				Name: "Satellite", Factor: 0.0125,
				Bonus:               []data.Resource{{Name: "Satellite Navigation"}},
				BonusStartsFromZero: true,
			}},
		}},
	}, {
		Name: "tanker", UnlockedBy: "Robotics",
		Costs: []data.Resource{{
			Name: "ship", Count: 200,
		}, {
			Name: "alloy", Count: 1250,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "kerosene", UnlockedBy: "Oil Processing",
		Costs: []data.Resource{{Name: "oil", Count: 7500}},
		Bonus: []data.Resource{{
			Name: "Factory", Factor: 0.05,
			Bonus:               []data.Resource{{Name: "Factory Processing"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "parchment", UnlockedBy: "Writing",
		Costs: []data.Resource{{Name: "fur", Count: 175}},
	}, {
		Name: "manuscript", UnlockedBy: "Construction",
		Producers: []data.Resource{{
			Name: "Steamworks", Factor: 0.0005 * 5,
			Bonus: []data.Resource{{
				Name: "Printing Press",
				Bonus: []data.Resource{{
					Name: "Offset Press", Factor: 4 - 1,
					Bonus: []data.Resource{{Name: "Photolithography", Factor: 4 - 1}},
				}},
			}},
			BonusStartsFromZero: true,
		}},
		Costs: []data.Resource{{
			Name: "culture", Count: 400,
		}, {
			Name: "parchment", Count: 25,
		}},
		Bonus: []data.Resource{{
			Name: "Codex Vox", Factor: 0.25,
		}, {
			Name: "Codex Logos", Factor: 0.25,
		}, {
			Name: "Codex Agrum", Factor: 0.25,
		}},
	}, {
		Name: "compendium", UnlockedBy: "Philosophy",
		Costs: []data.Resource{{
			Name: "manuscript", Count: 50,
		}, {
			Name: "science", Count: 10000,
		}},
		Bonus: []data.Resource{{
			Name: "Codex Logos", Factor: 0.25,
		}, {
			Name: "Codex Agrum", Factor: 0.25,
		}},
	}, {
		Name: "blueprint", UnlockedBy: "Physics",
		Costs: []data.Resource{{
			Name: "compendium", Count: 25,
		}, {
			Name: "science", Count: 25000,
		}},
		Bonus: []data.Resource{{
			Name: "CAD System", Factor: 0.01,
			Bonus: []data.Resource{{
				Name: "Library",
			}, {
				Name: "Data Center",
			}, {
				Name: "Academy",
			}, {
				Name: "Observatory",
			}, {
				Name: "Bio Lab",
			}},
		}, {
			Name: "Codex Agrum", Factor: 0.25,
		}},
	}, {
		Name: "thorium", UnlockedBy: "Thorium",
		Costs: []data.Resource{{Name: "uranium", Count: 250}},
		Producers: []data.Resource{{
			Name: "Active Reactor", Factor: -0.05 * 5,
			Bonus: []data.Resource{{
				Name:  "Thorium Reactors",
				Bonus: []data.Resource{{Name: "Enriched Thorium", Factor: -0.25}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "megalith", UnlockedBy: "Construction",
		Costs: []data.Resource{{
			Name: "beam", Count: 25,
		}, {
			Name: "slab", Count: 50,
		}, {
			Name: "plate", Count: 5,
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

func addBonus(g *game.Game, resources []data.Resource) {
	for _, resource := range resources {
		resource.Type = "Resource"
		resource.IsHidden = true
		resource.StartCountFromZero = true
		g.AddResource(resource)
	}
}

func addCrafts(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		name := action.Name
		action.Name = "@" + name
		action.Type = "Craft"
		action.Adds = []data.Resource{{
			Name: name, Count: 1,
			Bonus: join([]data.Resource{{Name: "CraftRatio"}}, action.Bonus),
		}}
		action.IsHidden = true
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: name, Type: "Resource", Cap: -1, ProducerAction: action.Name,
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
		action.Type = "Building"
		action.Adds = append([]data.Resource{{
			Name: action.Name, Count: 1,
		}}, action.Adds...)
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: -1,
		})

		if !isActive {
			continue
		}

		action.Name = "Active " + name
		action.Costs = []data.Resource{{Name: "Idle " + name, Count: 1}}
		action.Adds = []data.Resource{{Name: action.Name, Count: 1}}
		action.UnlockedBy = name
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: -1,
		})

		g.AddResource(data.Resource{
			Name: "Idle " + name, Type: "Building", IsHidden: true, Cap: -1, StartCountFromZero: true,
			Producers: []data.Resource{{
				Name: name,
			}, {
				Name: "Active " + name, Factor: -1,
			}},
		})
	}
}

func addJobs(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		action.Type = "Job"
		action.Costs = []data.Resource{{Name: "kitten", Count: 1, Cap: 1}}
		action.Adds = []data.Resource{{Name: action.Name, Count: 1}}
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: -1,
			OnGone: []data.Resource{{
				Name: "gone kitten", Count: 1,
			}, {
				Name: "kitten", Cap: 1,
			}},
		})
	}
}

func addTechs(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		action.Type = "Tech"
		action.Adds = []data.Resource{{Name: action.Name, Count: 1}}
		action.LockedBy = action.Name
		g.AddAction(action)
		g.AddResource(data.Resource{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: 1,
		})
	}
}
