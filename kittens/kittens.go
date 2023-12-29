package kittens

import (
	"strings"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

type R = data.Resource

func NewGame(now game.Now) *game.Game {
	kittenNames := []string{
		"kitten", "woodcutter", "scholar", "farmer", "hunter", "miner", "priest", "geologist",
	}

	g := game.NewGame(now())

	g.AddResources(join([]R{{
		Name: "day", Type: "Calendar", IsHidden: true, Count: 0, Cap: -1,
		Producers: []R{{Factor: 0.5}},
	}}, resourceWithModulus(R{
		Type: "Calendar", StartCount: 1, Cap: -1,
		Producers: []R{{Name: "day", Factor: 1. / (100 * 4 * 5), ProductionFloor: true}},
	}, []string{
		"Charon", "Umbra", "Yarn", "Helios", "Cath", "Redmoon", "Dune", "Piscine", "Termogus",
		"Kairo"}), []R{{
		Name: "year", Type: "Calendar", StartCount: 1, Cap: -1,
		Producers: []R{{Name: "day", Factor: 0.0025, ProductionFloor: true}},
	}}, resourceWithModulus(R{
		Type: "Calendar", StartCount: 1, Cap: -1,
		Producers: []R{{Name: "day", Factor: 0.01, ProductionFloor: true}},
	}, []string{
		"Spring", "Summer", "Autumn", "Winter"}), []R{{
		Name: "day_of_year", Type: "Calendar", StartCount: 1, Cap: -1,
		ProductionModulus: 400, ProductionModulusEquals: -1,
		Producers: []R{{Name: "day", ProductionFloor: true}},
	}}))

	addBonus(g, []R{{
		Name: "BarnBonus",
		Producers: []R{{
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
		Producers: []R{{
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
		Producers: []R{{
			Name: "ship", Factor: 0.01,
			Bonus: []R{{
				Name: "Expanded Cargo",
				Bonus: []R{{
					Name: "Reactor", Factor: 0.05,
					Bonus:               []R{{Name: "Reactor Vessel"}},
					BonusStartsFromZero: true,
				}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "BarnCatnipCapBonus", Type: "Resource", IsHidden: true, StartCountFromZero: true,
		Producers: []R{{
			Name: "Silos", Factor: 0.25,
			Bonus:               []R{{Name: "BarnBonus"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "HuntingBonus", Type: "Resource", IsHidden: true, StartCountFromZero: true,
		Producers: []R{{
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
		Producers: []R{{
			Name: "Workshop", Factor: 0.06,
		}, {
			Name: "Factory", Factor: 0.05,
			Bonus: []R{{Name: "Factory Logistics", Factor: 0.20}},
		}},
	}, {
		Name: "SpaceElevatorOilBonus",
		Producers: []R{{
			Name: "Space Elevator", Factor: -0.05,
			Bonus: []R{{
				Name:                "Cath",
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Kairo", Factor: -0.50,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}},
	}, {
		Name: "SpaceReactorScienceBonus",
		Producers: []R{{
			Name: "Antimatter Reactors", Factor: 0.95,
		}, {
			Name: "Advanced AM Reactors", Factor: 1.50,
		}, {
			Name: "Void Reactors", Factor: 4.00,
		}},
	}, {
		Name: "AcceleratorCapBonus",
		Producers: []R{{
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
		Producers: []R{{
			Name: "AI Core", Factor: 0.10,
			Bonus:               []R{{Name: "AI Bases"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "CatnipCapBonus",
		Producers: []R{{
			Name: "Refrigeration", Factor: 0.75,
		}, {
			Name: "Hydroponics", Factor: 0.10,
		}},
	}, {
		Name: "ParagonEffectBonus",
		Producers: []R{{
			Name: "Malkuth", Factor: 0.05,
		}, {
			Name: "Yesod", Factor: 0.05,
		}, {
			Name: "Hod", Factor: 0.05,
		}, {
			Name: "Netzach", Factor: 0.05,
		}, {
			Name: "Tiferet", Factor: 0.05,
		}, {
			Name: "Gevurah", Factor: 0.05,
		}, {
			Name: "Chesed", Factor: 0.05,
		}, {
			Name: "Binah", Factor: 0.05,
		}, {
			Name: "Chokhmah", Factor: 0.05,
		}, {
			Name: "Keter", Factor: 0.05,
		}},
	}, {
		Name: "ParagonCapBonus",
		Producers: []R{{
			Name: "paragon", Factor: 0.001,
		}, {
			Name: "burned paragon", Factor: 0.0005,
		}},
		Bonus: []R{{Name: "ParagonEffectBonus"}},
	}, {
		Name: "OtherCapBonus",
		Producers: []R{{
			Name: "Void Rift", Factor: 0.02,
		}, {
			Name: "Event Horizon", Factor: 0.10,
		}},
	}, {
		Name:      "GlobalCapBonus",
		Producers: []R{{}},
		Bonus: []R{{
			Bonus: []R{{
				Bonus: []R{{Name: "ParagonCapBonus"}},
			}, {
				Bonus: []R{{Name: "OtherCapBonus"}},
			}},
			BonusIsMultiplicative: true,
		}, {
			Factor: -1,
		}},
		BonusStartsFromZero: true,
	}, {
		Name:      "BaseMetalCapBonus",
		Producers: []R{{Name: "Sunforge", Factor: 0.01}},
	}, {
		Name: "SolarRevolutionProductionBonus",
		Producers: []R{{
			Name: "worship", Factor: 0.01 / 1000,
			Bonus: []R{{
				Name: "Solar Revolution",
				Bonus: []R{{
					Name: "Black Obelisk", Factor: 0.05,
					Bonus:               []R{{Name: "Transcendence Level"}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
	}, {
		Name: "ParagonProductionBonus",
		Producers: []R{{
			Name: "paragon", Factor: 0.01,
		}, {
			Name: "burned paragon", Factor: 0.01,
		}},
		Bonus: []R{{Name: "ParagonEffectBonus"}},
	}, {
		Name: "OtherProductionBonus",
		Producers: []R{{
			Name: "Active Magneto", Factor: 0.02,
			Bonus: []R{{Name: "Steamworks", Factor: 0.15}},
		}, {
			Name: "Active Reactor", Factor: 0.05,
		}},
	}, {
		Name:      "GlobalProductionBonus",
		Producers: []R{{}},
		Bonus: []R{{
			Bonus: []R{{
				Bonus: []R{{Name: "SolarRevolutionProductionBonus"}},
			}, {
				Bonus: []R{{Name: "ParagonProductionBonus"}},
			}, {
				Bonus: []R{{Name: "OtherProductionBonus"}},
			}},
			BonusIsMultiplicative: true,
		}, {
			Factor: -1,
		}},
		BonusStartsFromZero: true,
	}, {
		Name: "SpaceProductionBonus",
		Producers: []R{{
			Name: "Space Elevator", Factor: 0.01,
		}, {
			Name: "Orbital Array", Factor: 0.02,
		}},
		Bonus: []R{{
			Name: "Factory", Factor: 0.0375,
			Bonus: []R{{
				Name:  "Space Manufacturing",
				Bonus: []R{{Name: "Factory Logistics", Factor: 4.5/3.75 - 1}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "CryostationStorageBonus",
		Producers: []R{{
			Name: "Helios", Factor: -0.10,
			Bonus:               []R{{Name: "Numerology"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Termogus", Factor: 0.20,
			Bonus:               []R{{Name: "Numerology"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "AstronomicalEventBonus",
		Producers: []R{{
			Name: "Observatory", Factor: 0.002,
		}, {
			Name: "Celestial Mechanics", Factor: 0.20,
		}, {
			Name: "Chronomancy", Factor: 0.10,
		}, {
			Name: "Astromancy",
		}, {
			Name: "Temporal Accelerator", Factor: 0.0125,
		}, {
			Name: "Blazar", Factor: 0.025,
		}},
	}, {
		Name: "UnicornEventBonus",
		Producers: []R{{
			Name: "Unicornmancy", Factor: 0.10,
		}, {
			Name: "Temporal Accelerator", Factor: 0.0125,
		}, {
			Name: "Blazar", Factor: 0.025,
		}},
	}, {
		Name: "PriceRatioBonus",
		Producers: []R{{
			Name: "Enlightenment", Factor: -0.01,
		}, {
			Name: "Golden Ratio", Factor: -0.01618,
		}, {
			Name: "Divine Proportion", Factor: -0.017777,
		}, {
			Name: "Vitruvian Feline", Factor: -0.02,
		}, {
			Name: "Renaissance", Factor: -0.0225,
		}},
	}})

	addResets(g, []string{"catnip", "wood", "mineral", "coal", "iron", "titanium", "gold", "oil", "uranium", "unobtainium", "antimatter", "catpower", "science", "culture", "faith", "starchart", "relic", "void", "blackcoin"})

	g.AddResources([]R{{
		Name: "catnip cap", Type: "Resource", IsHidden: true, StartCount: 5000,
		Producers: []R{{
			Name: "Barn", Factor: 5000,
			Bonus: []R{{Name: "BarnCatnipCapBonus"}},
		}, {
			Name: "Warehouse", Factor: 750,
			Bonus:               []R{{Name: "BarnCatnipCapBonus"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Harbour", Factor: 2500,
			Bonus: []R{{
				Name: "HarbourBonus",
			}, {
				Name: "BarnCatnipCapBonus",
			}},
		}, {
			Name: "Accelerator", Factor: 30000,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Moon Base", Factor: 45000, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "CatnipCapBonus"}},
		}, {
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "catnip", Type: "Resource", CapResource: "catnip cap", ResetResource: "reset catnip",
		Producers: join([]R{{
			Name: "Catnip Field", Factor: 0.125 * 5,
			Bonus: []R{{
				Name: "Spring", Factor: 0.50,
			}, {
				Name: "Winter", Factor: -0.75,
			}},
		}}, resourceWithName(R{
			Factor: -4.25, ProductionFloor: true, ProductionOnGone: true,
			Bonus: []R{{
				Name: "Pasture", Factor: -0.005,
			}, {
				Name: "Unic. Pasture", Factor: -0.0015,
			}, {
				Name: "Robotic Assistance", Factor: -0.25,
			}},
		}, kittenNames), []R{{
			Name: "farmer", Factor: 1 * 5,
			Bonus: []R{{
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
			Bonus:               []R{{Name: "Biofuel Processing"}},
			BonusStartsFromZero: true,
		}}),
		Bonus: []R{{
			Bonus: []R{{
				Name: "Aqueduct", Factor: 0.03,
			}, {
				Name: "Hydroponics", Factor: 0.025,
				Bonus: []R{{
					Name:                "Yarn",
					Bonus:               []R{{Name: "Numerology"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Piscine", Factor: -0.50,
					Bonus:               []R{{Name: "Numerology"}},
					BonusStartsFromZero: true,
				}},
			}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Charon", Factor: 0.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "wood cap", Type: "Resource", IsHidden: true, StartCount: 200,
		Producers: []R{{
			Name: "Barn", Factor: 200,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Warehouse", Factor: 150,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Harbour", Factor: 700,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Moon Base", Factor: 25000, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 200000, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}, {
			Name: "Accelerator", Factor: 20000,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "wood", Type: "Resource", CapResource: "wood cap", ResetResource: "reset wood",
		Producers: []R{{
			Name: "woodcutter", Factor: 0.018 * 5,
			Bonus: []R{{
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
		}, {
			Name: "Active Steamworks", Factor: -0.02 / (2 * 100 * 4), ProductionOnGone: true,
			Bonus: []R{{
				Name: "Workshop Automation",
				Bonus: []R{{
					Name: "wood cap",
					Bonus: []R{{
						Name: "Advanced Automation",
					}},
				}},
				BonusStartsFromZero: true,
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{
				Name: "Lumber Mill", Factor: 0.10,
				Bonus: []R{{
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
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Charon", Factor: 0.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "science cap", Type: "Resource", IsHidden: true, StartCount: 250,
		Producers: []R{{
			Name: "Library", Factor: 250,
			Bonus: []R{{
				Name: "Observatory", Factor: 0.02,
				Bonus: []R{{
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
			Bonus: []R{{
				Name: "Astrolabe", Factor: 0.50,
			}, {
				Name: "Satellite", Factor: 0.05,
				Bonus: []R{{
					Name:                "Charon",
					Bonus:               []R{{Name: "Numerology"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Kairo", Factor: -0.25,
					Bonus:               []R{{Name: "Numerology"}},
					BonusStartsFromZero: true,
				}},
			}},
		}, {
			Name: "Bio Lab", Factor: 1500,
			Bonus: []R{{
				Name: "Data Center", Factor: 0.01,
				Bonus: []R{{
					Name: "Uplink",
				}, {
					Name: "Starlink",
				}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Data Center", Factor: 750,
			Bonus: []R{{
				Name: "Bio Lab", Factor: 0.01,
				Bonus:               []R{{Name: "Uplink"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Observatory", Factor: 0.02,
				Bonus: []R{{
					Name: "Titanium Reflectors",
				}, {
					Name: "Unobtainium Reflectors",
				}, {
					Name: "Eludium Reflectors",
				}},
				BonusStartsFromZero: true,
			}, {
				Name: "AI Core", Factor: 0.10,
				Bonus:               []R{{Name: "Machine Learning"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Temple", Factor: 500,
			Bonus:               []R{{Name: "Scholasticism"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Accelerator", Factor: 2500,
			Bonus:               []R{{Name: "LHC"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Research Vessel", Factor: 10000, Bonus: []R{{Name: "SpaceReactorScienceBonus"}},
		}, {
			Name: "Space Beacon", Factor: 25000, Bonus: []R{{Name: "SpaceReactorScienceBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "science", Type: "Resource", CapResource: "science cap", ResetResource: "reset science",
		Producers: []R{{
			Name: "scholar", Factor: 0.035 * 5,
			Bonus: []R{{Name: "happiness"}},
		}, {
			Factor: 0.0025 * 25. / 2.,
			Bonus:  []R{{Name: "AstronomicalEventBonus"}},
		}, {
			Name: "Celestial Mechanics", Factor: 0.001 * 15. / 2.,
			Bonus: []R{{Name: "AstronomicalEventBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{
				Name: "Library", Factor: 0.10,
			}, {
				Name: "Academy", Factor: 0.20,
			}, {
				Name: "Observatory", Factor: 0.25,
			}, {
				Name: "Bio Lab", Factor: 0.35,
			}, {
				Name: "Data Center", Factor: 0.10,
				Bonus: []R{{
					Name: "Bio Lab", Factor: 0.01,
					Bonus:               []R{{Name: "Uplink"}},
					BonusStartsFromZero: true,
				}, {
					Name: "AI Core", Factor: 0.10,
					Bonus:               []R{{Name: "Machine Learning"}},
					BonusStartsFromZero: true,
				}},
			}, {
				Name: "Space Station", Factor: 0.50,
				Bonus: []R{{
					Name: "Cath", Factor: 0.50,
					Bonus:               []R{{Name: "Numerology"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Kairo", Factor: -0.25,
					Bonus:               []R{{Name: "Numerology"}},
					BonusStartsFromZero: true,
				}},
			}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Piscine",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "catpower cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []R{{
			Name: "Hut", Factor: 75,
		}, {
			Name: "Log House", Factor: 50,
		}, {
			Name: "Mansion", Factor: 50,
		}, {
			Name: "Temple", Factor: 75,
			Bonus:               []R{{Name: "Templars"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "catpower", Type: "Resource", CapResource: "catpower cap", ResetResource: "reset catpower",
		Producers: []R{{
			Name: "hunter", Factor: 0.06 * 5,
			Bonus: []R{{
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
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Cath",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "mineral cap", Type: "Resource", IsHidden: true, StartCount: 250,
		Producers: []R{{
			Name: "Barn", Factor: 250,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Warehouse", Factor: 200,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Harbour", Factor: 950,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Moon Base", Factor: 30000, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 200000, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}, {
			Name: "Accelerator", Factor: 25000,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "mineral", Type: "Resource", CapResource: "mineral cap", ResetResource: "reset mineral",
		Producers: []R{{
			Name: "miner", Factor: 0.05 * 5,
			Bonus: []R{{
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
		}, {
			Factor: 0.001 * 25. / 2.,
			Bonus:  []R{{Name: "AstronomicalEventBonus"}},
		}, {
			Name: "Active Steamworks", Factor: -0.02 / (2 * 100 * 4), ProductionOnGone: true,
			Bonus: []R{{
				Name: "Workshop Automation",
				Bonus: []R{{
					Name: "mineral cap",
					Bonus: []R{{
						Name: "Advanced Automation",
					}},
				}},
				BonusStartsFromZero: true,
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Charon", Factor: 0.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "iron cap", Type: "Resource", IsHidden: true, StartCount: 50,
		Producers: []R{{
			Name: "Barn", Factor: 50,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Warehouse", Factor: 25,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}},
		}, {
			Name: "Harbour", Factor: 150,
			Bonus: []R{{
				Name: "BarnBonus",
			}, {
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Moon Base", Factor: 9000, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 50000, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}, {
			Name: "Accelerator", Factor: 7500,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "iron", Type: "Resource", CapResource: "iron cap", ResetResource: "reset iron",
		Producers: []R{{
			Name: "Active Smelter", Factor: 0.02 * 5,
			Bonus: []R{{Name: "Electrolytic Smelting", Factor: 0.95}},
		}, {
			Name: "Active Calciner", Factor: 0.15 * 5,
			Bonus: []R{{
				Name: "Oxidation", Factor: 0.95,
			}, {
				Name: "Rotary Kiln", Factor: 0.75,
			}, {
				Name: "Fluoridized Reactors",
			}},
		}, {
			Name: "Active Calciner", Factor: -0.15 * 5 * 0.10, ProductionOnGone: true,
			Bonus:               []R{{Name: "Steel Plants"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Active Steamworks", Factor: -0.02 / (2 * 100 * 4), ProductionOnGone: true,
			Bonus: []R{{
				Name: "Workshop Automation",
				Bonus: []R{{
					Name: "Pneumatic Press",
					Bonus: []R{{
						Name: "iron cap",
						Bonus: []R{{
							Name: "Advanced Automation",
						}},
					}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Umbra", Factor: 0.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "coal cap", Type: "Resource", IsHidden: true, StartCount: 60,
		Producers: []R{{
			Name: "Barn", Factor: 60,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Warehouse", Factor: 30,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Harbour", Factor: 100,
			Bonus: []R{{
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}, {
				Name: "Barges", Factor: 0.50,
			}},
		}, {
			Name: "Moon Base", Factor: 3500, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 25000, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}, {
			Name: "Accelerator", Factor: 2500,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "coal", Type: "Resource", CapResource: "coal cap", ResetResource: "reset coal",
		Producers: []R{{
			Name: "geologist", Factor: 0.015 * 5,
			Bonus: []R{{
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
			Bonus: []R{{
				Name:  "Coal Furnace",
				Bonus: []R{{Name: "Electrolytic Smelting", Factor: 0.95}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Mine", Factor: 0.003 * 5,
			Bonus:               []R{{Name: "Deep Mining"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Active Calciner", Factor: -0.15 * 5 * 0.10, ProductionOnGone: true,
			Bonus:               []R{{Name: "Steel Plants"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{
				Name: "Pyrolysis", Factor: 0.20,
			}, {
				Name: "Active Steamworks", Factor: -0.80, ProductionBoolean: true, ProductionOnGone: true,
				Bonus: []R{{
					Name: "High Pressure Engine", Factor: -0.25,
				}, {
					Name: "Fuel Injector", Factor: -0.25,
				}},
			}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Umbra", Factor: 0.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "gold cap", Type: "Resource", IsHidden: true, StartCount: 10,
		Producers: []R{{
			Name: "Barn", Factor: 10,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Warehouse", Factor: 5,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Harbour", Factor: 25,
			Bonus: []R{{
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Mint", Factor: 100,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Accelerator", Factor: 250,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "Sky Palace", Factor: 0.01}},
		}, {
			Bonus: []R{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "gold", Type: "Resource", CapResource: "gold cap", ResetResource: "reset gold",
		Producers: []R{{
			Name: "Active Mint", Factor: -0.005 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Smelter", Factor: 0.001 * 5,
			Bonus:               []R{{Name: "Gold Ore"}},
			BonusStartsFromZero: true,
		}, {
			Name: "geologist", Factor: 0.0008 * 5,
			Bonus: []R{{
				Name: "Geodesy",
				Bonus: []R{{
					Name: "Mining Drill", Factor: 0.625,
				}, {
					Name: "Unobtainium Drill", Factor: 0.625,
				}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Umbra", Factor: 0.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "titanium cap", Type: "Resource", IsHidden: true, StartCount: 2,
		Producers: []R{{
			Name: "Barn", Factor: 2,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Warehouse", Factor: 10,
			Bonus: []R{{Name: "WarehouseBonus"}},
		}, {
			Name: "Harbour", Factor: 50,
			Bonus: []R{{
				Name: "WarehouseBonus",
			}, {
				Name: "HarbourBonus",
			}},
		}, {
			Name: "Accelerator", Factor: 750,
			Bonus: []R{{
				Name: "Energy Rifts", Bonus: []R{{Name: "AcceleratorCapBonus"}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Moon Base", Factor: 1250, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 7500, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "titanium", Type: "Resource", CapResource: "titanium cap", ResetResource: "reset titanium",
		Producers: []R{{
			Name: "Active Accelerator", Factor: -0.015 * 5, ProductionOnGone: true,
		}, {
			Name: "Active Calciner", Factor: 0.0005 * 5,
			Bonus: []R{{
				Name: "Oxidation", Factor: 2.85,
			}, {
				Name: "Rotary Kiln", Factor: 2.25,
			}, {
				Name: "Fluoridized Reactors", Factor: 3.00,
			}},
		}, {
			Name: "Active Smelter", Factor: 0.0015 * 5,
			Bonus:               []R{{Name: "Nuclear Smelter"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Charon", Factor: 1.50,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "oil cap", Type: "Resource", IsHidden: true, StartCount: 1500,
		Producers: []R{{
			Name: "Oil Well", Factor: 1500,
		}, {
			Name: "tanker", Factor: 500,
		}, {
			Name: "Moon Base", Factor: 3500, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 7500, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "oil", Type: "Resource", CapResource: "oil cap", ResetResource: "reset oil",
		Producers: []R{{
			Name: "Oil Well", Factor: 0.02 * 5,
			Bonus: []R{{
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
			Bonus: []R{{
				Name:  "Biofuel Processing",
				Bonus: []R{{Name: "GM Catnip", Factor: 0.60}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Hydraulic Fracturer", Factor: 0.5 * 5,
			Bonus: []R{{
				Name: "SpaceProductionBonus",
			}, {
				Name: "Umbra", Factor: -0.25,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Dune", Factor: 0.50,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Termogus",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "uranium cap", Type: "Resource", IsHidden: true, StartCount: 250,
		Producers: []R{{
			Name: "Reactor", Factor: 250,
		}, {
			Name: "Planet Cracker", Factor: 1750,
		}, {
			Name: "Cryostation", Factor: 5000, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "uranium", Type: "Resource", CapResource: "uranium cap", ResetResource: "reset uranium",
		Producers: []R{{
			Name: "Active Accelerator", Factor: 0.0025 * 5,
		}, {
			Name: "Active Reactor", Factor: -0.001 * 5, ProductionOnGone: true,
			Bonus: []R{{Name: "Enriched Uranium", Factor: -0.25}},
		}, {
			Name: "Quarry", Factor: 0.0005 * 5,
			Bonus: []R{{
				Name:  "Orbital Geodesy",
				Bonus: []R{{Name: "Enriched Uranium", Factor: 0.25}},
			}},
			BonusStartsFromZero: true,
		}, {
			Name: "Active Lunar Outpost", Factor: -0.35 * 5, ProductionOnGone: true,
			Bonus: []R{{Name: "SpaceProductionBonus"}},
		}, {
			Name: "Planet Cracker", Factor: 0.3 * 5,
			Bonus: []R{{
				Name: "Planet Busters",
			}, {
				Name: "SpaceProductionBonus",
			}, {
				Name: "Umbra", Factor: -0.10,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Dune", Factor: 0.10,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Dune",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "unobtainium cap", Type: "Resource", IsHidden: true, StartCount: 150,
		Producers: []R{{
			Name: "Moon Base", Factor: 150, Bonus: []R{{Name: "MoonBaseCapBonus"}},
		}, {
			Name: "Cryostation", Factor: 750, Bonus: []R{{Name: "CryostationStorageBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "BaseMetalCapBonus"}},
		}, {
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "unobtainium", Type: "Resource", CapResource: "unobtainium cap", ResetResource: "reset unobtainium",
		Producers: []R{{
			Name: "Active Lunar Outpost", Factor: 0.035,
			Bonus: []R{{
				Name: "SpaceProductionBonus",
			}, {
				Name: "Charon", Factor: -0.10,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Redmoon", Factor: 0.20,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "Microwarp Reactors", Factor: 0.75}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Redmoon",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "reset time crystal", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true, ResetResource: "reset time crystal",
		Producers: []R{{
			Name:                "time crystal",
			Bonus:               []R{{Name: "Anachronomancy"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "time crystal", Type: "Resource", Cap: -1, ResetResource: "reset time crystal",
	}, {
		Name: "antimatter cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []R{{
			Name: "Containment Chamber", Factor: 100,
			Bonus: []R{{Name: "Heatsink", Factor: 0.02}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalCapBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "antimatter", Type: "Resource", CapResource: "antimatter cap", ResetResource: "reset antimatter",
		Producers: []R{{Name: "Sunlifter", Factor: 1. / (2 * 100 * 4)}},
	}, {
		Name: "relic", Type: "Resource", Cap: -1, ResetResource: "reset relic",
		Producers: []R{{
			Name: "Space Beacon", Factor: 0.01 / 2,
			Bonus: []R{{
				Name: "Relic Station",
				Bonus: []R{{
					Name: "Black Nexus", Factor: 0.10,
					Bonus:               []R{{Name: "Black Pyramid"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Hash Level", Factor: 0.25,
				}, {
					Name: "SpaceProductionBonus",
				}},
			}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "void", Type: "Resource", Cap: -1, ResetResource: "reset void",
		Producers: []R{{
			Name: "Chronosphere", Factor: 0.005 / (2 * 100),
			Bonus: []R{{Name: "Void Hoover"}},
		}},
		Bonus: []R{{
			Bonus: []R{{
				Name:  "Chronocontrol",
				Bonus: []R{{Name: "Distortion", Factor: 2}},
			}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "temporal flux cap", Type: "Resource", IsHidden: true, StartCount: 3000,
		Bonus: []R{{
			Name: "Temporal Battery", Factor: 0.25,
		}},
	}, {
		Name: "temporal flux", Type: "Resource", IsHidden: true, CapResource: "temporal flux cap", ResetResource: "temporal flux",
		Producers: []R{{
			Factor: 5. / (60 * 10),
			Bonus:  []R{{Name: "Temporal Accelerator", Factor: 0.05}},
		}, {
			Name: "Chronosphere", Factor: 1. / (2 * 100 * 4),
			Bonus:               []R{{Name: "Chronosurge"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "blackcoin", Type: "Resource", Cap: -1, ResetResource: "reset blackcoin",
	}, {
		Name: "kitten", Type: "Resource", Cap: 0,
		Producers: []R{{Factor: 0.05}},
		Bonus: []R{{
			Name: "Venus of Willenfluff", Factor: 0.75,
		}, {
			Name: "Pawgan Rituals", Factor: 1.50,
		}, {
			Name: "Active Brewery", Factor: 0.01,
			Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
			BonusStartsFromZero: true,
		}},
		OnGone: []R{{Name: "gone kitten", Count: 1}},
	}, {
		Name: "all kittens", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true,
		Producers: resourceWithName(R{
			Name: "kitten", ProductionFloor: true,
		}, kittenNames),
	}, {
		Name: "fur", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "all kittens", Factor: -0.05,
			Bonus: []R{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Active Mint", Factor: 0.0000875,
			Bonus:               []R{{Name: "catpower cap"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "ivory", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "all kittens", Factor: -0.035,
			Bonus: []R{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Active Mint", Factor: 0.000021,
			Bonus:               []R{{Name: "catpower cap"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Ivory Citadel", Factor: 0.0005 * 250. / 2.,
			Bonus: []R{{Name: "UnicornEventBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "spice", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "all kittens", Factor: -0.005,
			Bonus: []R{{Name: "Tradepost", Factor: -0.04}},
		}, {
			Name: "Active Brewery", Factor: -0.1 * 5, ProductionOnGone: true,
		}, {
			Name: "Spice Refinery", Factor: 0.125,
			Bonus: []R{{Name: "SpaceProductionBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "unicorn", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "Unic. Pasture", Factor: 0.001 * 5,
		}, {
			Name: "Ivory Tower", Factor: 0.0005 * 500. / 2.,
			Bonus: []R{{Name: "UnicornEventBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{
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
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Helios", Factor: 0.25,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "culture cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []R{{
			Name: "Library", Factor: 10,
		}, {
			Name: "Academy", Factor: 25,
		}, {
			Name: "Amphitheatre", Factor: 50,
		}, {
			Name: "Chapel", Factor: 200,
		}, {
			Name: "Data Center", Factor: 250,
			Bonus: []R{{
				Name: "Bio Lab", Factor: 0.01,
				Bonus:               []R{{Name: "Uplink"}},
				BonusStartsFromZero: true,
			}, {
				Name: "AI Core", Factor: 0.10,
				Bonus:               []R{{Name: "Machine Learning"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Temple", Factor: 125,
			Bonus:               []R{{Name: "Basilica"}},
			BonusStartsFromZero: true,
		}},
		Bonus: []R{{
			Name: "Ziggurat", Factor: 0.08,
			Bonus: []R{{Name: "Unicorn Graveyard", Factor: 0.125}},
		}},
	}, {
		Name: "culture", Type: "Resource", CapResource: "culture cap", ResetResource: "reset culture",
		Producers: []R{{
			Name: "Amphitheatre", Factor: 0.005 * 5,
		}, {
			Name: "Chapel", Factor: 0.05 * 5,
		}, {
			Name: "Temple", Factor: 0.1 * 5,
			Bonus: []R{{
				Name: "Stained Glass", Factor: 0.50,
			}, {
				Name: "Basilica", Factor: 2.00,
			}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Yarn",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "faith cap", Type: "Resource", IsHidden: true, StartCount: 100,
		Producers: []R{{
			Name: "Temple", Factor: 100,
			Bonus: []R{{
				Name: "Golden Spire", Factor: 0.50,
			}, {
				Name: "Sun Altar", Factor: 0.50,
			}},
		}},
	}, {
		Name: "faith", Type: "Resource", CapResource: "faith cap", ResetResource: "reset faith",
		Producers: []R{{
			Name: "priest", Factor: 0.0015 * 5,
			Bonus: []R{{Name: "happiness"}},
		}, {
			Name: "Temple", Factor: 0.0015 * 5,
			Bonus:               []R{{Name: "Theology"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Chapel", Factor: 0.005 * 5,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "Solar Chant", Factor: 0.10}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Helios",
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "worship", Type: "Resource", Cap: -1,
	}, {
		Name: "epiphany", Type: "Resource", Cap: -1, ResetResource: "epiphany",
	}, {
		Name: "starchart", Type: "Resource", Cap: -1, ResetResource: "reset starchart",
		Producers: []R{{
			Name: "scholar", Factor: 0.0005,
			Bonus:               []R{{Name: "Astrophysicists"}},
			BonusStartsFromZero: true,
		}, {
			Name: "Satellite", Factor: 0.005,
			Bonus: []R{{
				Name:                "Cath",
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Kairo", Factor: -0.25,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Research Vessel", Factor: 0.05,
			Bonus: []R{{
				Name: "SpaceProductionBonus",
			}, {
				Name: "Yarn", Factor: -0.50,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Piscine", Factor: 0.50,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Space Beacon", Factor: 0.125,
			Bonus: []R{{
				Name: "SpaceProductionBonus",
			}, {
				Name: "Cath", Factor: -0.50,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Kairo", Factor: 4.00,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Astronomy", Factor: 0.0025 * 1 / 2.,
			Bonus: []R{{Name: "AstronomicalEventBonus"}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "Hubble Space Telescope", Factor: 0.30}},
		}, {
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}, {
			Bonus: []R{{
				Name: "Kairo", Factor: 4.00,
				Bonus: []R{{
					Name:                "Numeromancy",
					Bonus:               []R{{Name: "festival day", ProductionBoolean: true}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "festival day", Type: "Resource", Cap: -1,
		Producers: []R{{Factor: -0.5}},
	}, {
		Name: "gigaflop", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "AI Core", Factor: 0.02 * 5,
		}, {
			Name: "Entanglement Station", Factor: -0.1 * 5,
			Bonus: []R{{
				Name:                "Charon",
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}, {
				Name: "Redmoon", Factor: -0.50,
				Bonus:               []R{{Name: "Numerology"}},
				BonusStartsFromZero: true,
			}},
		}},
	}, {
		Name: "hash", Type: "Resource", Cap: -1,
		Producers: []R{{Name: "Entanglement Station", Factor: 0.1 * 5}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "leviathan energy cap", Type: "Resource", IsHidden: true, StartCount: 1,
		Producers: []R{{Name: "Marker", Factor: 5}},
	}, {
		Name: "leviathan energy", Type: "Resource", CapResource: "leviathan energy cap",
	}, {
		Name: "tear", Type: "Resource", Cap: -1,
	}, {
		Name: "alicorn", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "Sky Palace", Factor: 0.00002 * 5,
			Bonus: []R{{
				Name: "Unicorn Utopia", Factor: 0.00015,
			}, {
				Name: "Sunspire", Factor: 0.0003,
			}, {
				Name: "Black Radiance", Factor: 0.03464,
				Bonus:               []R{{Name: "sorrow"}},
				BonusStartsFromZero: true,
			}},
		}, {
			Name: "Unicorn Utopia", Factor: 0.000025 * 5,
		}, {
			Name: "Sunspire", Factor: 0.00005 * 5,
		}, {
			Name: "Sky Palace", Factor: 0.0001 / 2,
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "necrocorn", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "Marker", Factor: 0.000001 * 5,
			Bonus: []R{{Name: "Unicorn Graveyard", Factor: 0.10}},
		}},
		Bonus: []R{{
			Bonus: []R{{Name: "GlobalProductionBonus"}},
		}},
		BonusIsMultiplicative: true,
	}, {
		Name: "sorrow cap", Type: "Resource", IsHidden: true, StartCount: 16,
		Producers: []R{{Name: "Black Core"}},
	}, {
		Name: "sorrow", Type: "Resource", CapResource: "sorrow cap", ResetResource: "sorrow",
	}, {
		Name: "chronoheat cap", Type: "Resource", IsHidden: true, StartCount: 1,
		Producers: []R{{
			Name: "Chrono Furnace", Factor: 100,
		}, {
			Name: "Time Boiler", Factor: 10,
		}},
	}, {
		Name: "chronoheat", Type: "Resource", CapResource: "chronoheat cap",
		Producers: []R{{Name: "Chrono Furnace", Factor: -0.02 * 5}},
	}, {
		Name: "chrono furnace fuel", Type: "Resource", Cap: -1,
		Producers: []R{{
			Name: "Chrono Furnace", Factor: 0.02 * 5,
			Bonus:               []R{{Name: "chronoheat", ProductionBoolean: true}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "kitten minus 35", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true,
		Producers: []R{{
			Factor: -35,
		}, {
			Name: "all kittens",
		}},
	}, {
		Name: "reset karma kitten", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true,
		Producers: []R{{
			Name:                "kitten minus 35",
			Bonus:               []R{{Name: "kitten minus 35", ProductionBoolean: true}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "karma kitten", Type: "Resource", Cap: -1, ResetResource: "reset karma kitten",
	}, {
		Name: "kitten minus 70", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true,
		Producers: []R{{
			Factor: -70,
		}, {
			Name: "all kittens",
		}},
	}, {
		Name: "reset paragon", Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true,
		Producers: []R{{
			Name:                "kitten minus 70",
			Bonus:               []R{{Name: "kitten minus 70", ProductionBoolean: true}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "paragon", Type: "Resource", Cap: -1,
		Producers: []R{{Factor: 1. / (2 * 100 * 4 * 1000)}},
	}, {
		Name: "burned paragon", Type: "Resource", Cap: -1, ResetResource: "burned paragon",
	}, {
		Name: "gone kitten", Type: "Resource", Cap: -1,
	}, {
		Name: "happiness", Type: "Job", StartCount: 0.1, Cap: -1,
		Producers: []R{{
			Name: "all kittens", Factor: -0.02,
		}, {
			Name: "ivory", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "fur", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "spice", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "unicorn", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "alicorn", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "necrocorn", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "karma", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "karma", Factor: 0.01,
		}, {
			Name: "tear", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "relic", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "blackcoin", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "void", Factor: 0.10, ProductionBoolean: true,
		}, {
			Name: "festival day", Factor: 0.10, ProductionBoolean: true,
			Bonus: []R{{Name: "Active Brewery", Factor: 0.01}},
		}, {
			Name: "Amphitheatre", Factor: 0.048,
		}, {
			Name: "Broadcast Tower", Factor: 0.75,
		}, {
			Name: "Temple", Factor: 0.005,
			Bonus:               []R{{Name: "Sun Altar"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "reset Chronophysics", Type: "Resource", IsHidden: true, Cap: 1, StartCountFromZero: true, ResetResource: "reset Chronophysics",
		Producers: []R{{
			Name:                "Chronophysics",
			Bonus:               []R{{Name: "Anachronomancy"}},
			BonusStartsFromZero: true,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Gather catnip", Type: "Building", LockedBy: "Catnip Field",
		Adds: []R{{Name: "catnip", Count: 10}},
	}, {
		Name: "Refine catnip", Type: "Building", UnlockedBy: "Catnip Field", LockedBy: "woodcutter",
		Costs: []R{{
			Name: "catnip", Count: 100,
			Bonus: []R{{Name: "Catnip Enrichment", Factor: -0.50}},
		}},
		Adds: []R{{
			Name: "wood", Count: 5,
			Bonus: []R{{Name: "Bio Lab", Factor: 0.10}},
		}},
	}})

	addBuildings(g, []data.Action{{
		Name: "Catnip Field", UnlockedBy: "catnip",
		CostExponentBaseResource: getCEBR(1.12),
		Costs:                    []R{{Name: "catnip", Count: 10}},
	}, {
		Name: "Hut", UnlockedBy: "Catnip Field",
		CostExponentBaseResource: R{
			Factor: 2.5,
			Bonus: []R{{
				Name: "Ironwood Huts", Factor: -0.50,
			}, {
				Name: "Concrete Huts", Factor: -0.30,
			}, {
				Name: "Unobtainium Huts", Factor: -0.25,
			}, {
				Name: "Eludium Huts", Factor: -0.10,
			}, {
				Name: "PriceRatioBonus",
			}},
		},
		Costs: []R{{Name: "wood", Count: 5}},
		Adds:  []R{{Name: "kitten", Cap: 2}},
	}, {
		Name: "Library", UnlockedBy: "Catnip Field",
		CostExponentBaseResource: getCEBR(1.15),
		Costs:                    []R{{Name: "wood", Count: 25}},
	}, {
		Name: "Barn", UnlockedBy: "Agriculture",
		CostExponentBaseResource: getCEBR(1.75),
		Costs:                    []R{{Name: "wood", Count: 50}},
	}, {
		Name: "Mine", UnlockedBy: "Mining",
		CostExponentBaseResource: getCEBR(1.15),
		Costs:                    []R{{Name: "wood", Count: 100}},
	}, {
		Name: "Workshop", UnlockedBy: "Mining",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "wood", Count: 100,
		}, {
			Name: "mineral", Count: 400,
		}},
	}, {
		Name: "Active Smelter", UnlockedBy: "Metal Working",
		CostExponentBaseResource: getCEBR(1.15),
		Costs:                    []R{{Name: "mineral", Count: 200}},
	}, {
		Name: "Pasture", UnlockedBy: "Animal Husbandry",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "catnip", Count: 100,
		}, {
			Name: "wood", Count: 10,
		}},
	}, {
		Name: "Unic. Pasture", UnlockedBy: "Animal Husbandry",
		CostExponentBaseResource: getCEBR(1.75),
		Costs:                    []R{{Name: "unicorn", Count: 2}},
	}, {
		Name: "Academy", UnlockedBy: "Mathematics",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "wood", Count: 50,
		}, {
			Name: "mineral", Count: 70,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Warehouse", UnlockedBy: "Construction",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "beam", Count: 1.5,
		}, {
			Name: "slab", Count: 2,
		}},
	}, {
		Name: "Log House", UnlockedBy: "Construction",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "wood", Count: 200,
		}, {
			Name: "mineral", Count: 250,
		}},
		Adds: []R{{Name: "kitten", Cap: 1}},
	}, {
		Name: "Aqueduct", UnlockedBy: "Engineering",
		CostExponentBaseResource: getCEBR(1.12),
		Costs:                    []R{{Name: "mineral", Count: 75}},
	}, {
		Name: "Mansion", UnlockedBy: "Architecture",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "slab", Count: 185,
		}, {
			Name: "steel", Count: 75,
		}, {
			Name: "titanium", Count: 25,
		}},
		Adds: []R{{Name: "kitten", Cap: 1}},
	}, {
		Name: "Observatory", UnlockedBy: "Astronomy",
		CostExponentBaseResource: getCEBR(1.1),
		Costs: []R{{
			Name: "scaffold", Count: 50,
		}, {
			Name: "slab", Count: 35,
		}, {
			Name: "iron", Count: 750,
		}, {
			Name: "science", Count: 1000,
		}},
	}, {
		Name: "Active Bio Lab", UnlockedBy: "Biology",
		CostExponentBaseResource: getCEBR(1.1),
		Costs: []R{{
			Name: "slab", Count: 100,
		}, {
			Name: "alloy", Count: 25,
		}, {
			Name: "science", Count: 1500,
		}},
	}, {
		Name: "Harbour", UnlockedBy: "Navigation",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "scaffold", Count: 5,
		}, {
			Name: "slab", Count: 50,
		}, {
			Name: "plate", Count: 75,
		}},
	}, {
		Name: "Quarry", UnlockedBy: "Geology",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "scaffold", Count: 50,
		}, {
			Name: "steel", Count: 125,
		}, {
			Name: "slab", Count: 1000,
		}},
	}, {
		Name: "Lumber Mill", UnlockedBy: "Construction",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "wood", Count: 100,
		}, {
			Name: "iron", Count: 50,
		}, {
			Name: "mineral", Count: 250,
		}},
	}, {
		Name: "Oil Well", UnlockedBy: "Chemistry",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "steel", Count: 50,
		}, {
			Name: "gear", Count: 25,
		}, {
			Name: "scaffold", Count: 25,
		}},
	}, {
		Name: "Active Accelerator", UnlockedBy: "Particle Physics",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "titanium", Count: 7500,
		}, {
			Name: "concrete", Count: 125,
		}, {
			Name: "uranium", Count: 25,
		}},
	}, {
		Name: "Active Steamworks", UnlockedBy: "Machinery",
		CostExponentBaseResource: getCEBR(1.25),
		Costs: []R{{
			Name: "steel", Count: 65,
		}, {
			Name: "gear", Count: 20,
		}, {
			Name: "blueprint", Count: 1,
		}},
	}, {
		Name: "Active Magneto", UnlockedBy: "Electricity",
		CostExponentBaseResource: getCEBR(1.25),
		Costs: []R{{
			Name: "alloy", Count: 10,
		}, {
			Name: "gear", Count: 5,
		}, {
			Name: "blueprint", Count: 1,
		}},
	}, {
		Name: "Active Calciner", UnlockedBy: "Chemistry",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "steel", Count: 100,
		}, {
			Name: "titanium", Count: 15,
		}, {
			Name: "blueprint", Count: 1,
		}, {
			Name: "oil", Count: 500,
		}},
	}, {
		Name: "Factory", UnlockedBy: "Mechanization",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "titanium", Count: 2000,
		}, {
			Name: "plate", Count: 25000,
		}, {
			Name: "concrete", Count: 15,
		}},
	}, {
		Name: "Active Reactor", UnlockedBy: "Nuclear Fission",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "titanium", Count: 3500,
		}, {
			Name: "plate", Count: 5000,
		}, {
			Name: "concrete", Count: 50,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Amphitheatre", UnlockedBy: "Writing",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "wood", Count: 200,
		}, {
			Name: "mineral", Count: 1200,
		}, {
			Name: "parchment", Count: 3,
		}},
	}, {
		Name: "Chapel", UnlockedBy: "Acoustics",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "mineral", Count: 2000,
		}, {
			Name: "culture", Count: 250,
		}, {
			Name: "parchment", Count: 250,
		}},
	}, {
		Name: "Temple", UnlockedBy: "Philosophy",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "slab", Count: 25,
		}, {
			Name: "plate", Count: 15,
		}, {
			Name: "gold", Count: 50,
		}, {
			Name: "manuscript", Count: 10,
		}},
	}, {
		Name: "Tradepost", UnlockedBy: "Currency",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "wood", Count: 500,
		}, {
			Name: "mineral", Count: 200,
		}, {
			Name: "gold", Count: 10,
		}},
	}, {
		Name: "Active Mint", UnlockedBy: "Architecture",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "mineral", Count: 5000,
		}, {
			Name: "plate", Count: 200,
		}, {
			Name: "gold", Count: 500,
		}},
	}, {
		Name: "Active Brewery", UnlockedBy: "Drama and Poetry",
		CostExponentBaseResource: getCEBR(1.5),
		Costs: []R{{
			Name: "wood", Count: 1000,
		}, {
			Name: "culture", Count: 750,
		}, {
			Name: "spice", Count: 5,
		}, {
			Name: "parchment", Count: 375,
		}},
	}, {
		Name: "Ziggurat", UnlockedBy: "Construction",
		CostExponentBaseResource: getCEBR(1.25),
		Costs: []R{{
			Name: "megalith", Count: 50,
		}, {
			Name: "scaffold", Count: 50,
		}, {
			Name: "blueprint", Count: 1,
		}},
	}, {
		Name: "Chronosphere", UnlockedBy: "Chronophysics",
		CostExponentBaseResource: getCEBR(1.25),
		Costs: []R{{
			Name: "unobtainium", Count: 2500,
		}, {
			Name: "time crystal", Count: 1,
		}, {
			Name: "blueprint", Count: 100,
		}, {
			Name: "science", Count: 250000,
		}},
	}, {
		Name: "AI Core", UnlockedBy: "Artificial Intelligence",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "antimatter", Count: 125,
		}, {
			Name: "science", Count: 500000,
		}},
	}, {
		Name: "Solar Farm", UnlockedBy: "Ecology",
		CostExponentBaseResource: getCEBR(1.15),
		Costs:                    []R{{Name: "titanium", Count: 250}},
	}, {
		Name: "Hydro Plant", UnlockedBy: "Robotics",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "concrete", Count: 100,
		}, {
			Name: "titanium", Count: 2500,
		}},
	}, {
		Name: "Data Center", UnlockedBy: "Electronics",
		CostExponentBaseResource: getCEBR(1.15),
		Costs: []R{{
			Name: "concrete", Count: 10,
		}, {
			Name: "steel", Count: 100,
		}},
	}, {
		Name: "Broadcast Tower", UnlockedBy: "Electronics",
		CostExponentBaseResource: getCEBR(1.18),
		Costs: []R{{
			Name: "iron", Count: 1250,
		}, {
			Name: "titanium", Count: 75,
		}},
	}, {
		Name: "Unicorn Tomb", UnlockedBy: "Ziggurat",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 5,
		}, {
			Name: "ivory", Count: 500,
		}},
	}, {
		Name: "Ivory Tower", UnlockedBy: "Unicorn Tomb",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 25,
		}, {
			Name: "ivory", Count: 25000,
		}},
	}, {
		Name: "Ivory Citadel", UnlockedBy: "Ivory Tower",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 50,
		}, {
			Name: "ivory", Count: 50000,
		}},
	}, {
		Name: "Sky Palace", UnlockedBy: "Ivory Citadel",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 500,
		}, {
			Name: "ivory", Count: 125000,
		}, {
			Name: "megalith", Count: 5,
		}},
	}, {
		Name: "Unicorn Utopia", UnlockedBy: "Sky Palace",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 5000,
		}, {
			Name: "ivory", Count: 1000000,
		}, {
			Name: "gold", Count: 500,
		}},
	}, {
		Name: "Sunspire", UnlockedBy: "Unicorn Utopia",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 25000,
		}, {
			Name: "ivory", Count: 750000,
		}, {
			Name: "gold", Count: 1000,
		}},
	}, {
		Name: "Marker", UnlockedBy: "Megalomania",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "tear", Count: 5000,
		}, {
			Name: "megalith", Count: 750,
		}, {
			Name: "spice", Count: 50000,
		}, {
			Name: "unobtainium", Count: 2500,
		}},
	}, {
		Name: "Unicorn Graveyard", UnlockedBy: "Black Codex",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "necrocorn", Count: 5,
		}, {
			Name: "megalith", Count: 1000,
		}},
	}, {
		Name: "Unicorn Necropolis", UnlockedBy: "Unicorn Graveyard",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "necrocorn", Count: 15,
		}, {
			Name: "megalith", Count: 2500,
		}, {
			Name: "alicorn", Count: 100,
		}, {
			Name: "void", Count: 5,
		}},
	}, {
		Name: "Black Pyramid", UnlockedBy: "Megalomania",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "sorrow", Count: 5,
		}, {
			Name: "megalith", Count: 2500,
		}, {
			Name: "spice", Count: 150000,
		}, {
			Name: "unobtainium", Count: 5000,
		}},
	}, {
		Name: "Solar Chant", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs:            []R{{Name: "faith", Count: 100}},
	}, {
		Name: "Scholasticism", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs:            []R{{Name: "faith", Count: 250}},
	}, {
		Name: "Golden Spire", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs: []R{{
			Name: "faith", Count: 350,
		}, {
			Name: "gold", Count: 150,
		}},
	}, {
		Name: "Sun Altar", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs: []R{{
			Name: "faith", Count: 500,
		}, {
			Name: "gold", Count: 250,
		}},
	}, {
		Name: "Stained Glass", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs: []R{{
			Name: "faith", Count: 500,
		}, {
			Name: "gold", Count: 250,
		}},
	}, {
		Name: "Basilica", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs: []R{{
			Name: "faith", Count: 1250,
		}, {
			Name: "gold", Count: 750,
		}},
	}, {
		Name: "Templars", UnlockedBy: "Philosophy",
		CostExponentBase: 2.5,
		Costs: []R{{
			Name: "faith", Count: 3500,
		}, {
			Name: "gold", Count: 3000,
		}},
	}, {
		Name: "Transcendence Level", UnlockedBy: "Transcendence", ResetResource: "Transcendence Level",
		CostExponentBase: 3,
		Costs:            []R{{Name: "epiphany", Count: 1}},
	}, {
		Name: "Black Obelisk", UnlockedBy: "Cryptotheology", ResetResource: "Black Obelisk",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 100}},
	}, {
		Name: "Black Nexus", UnlockedBy: "Cryptotheology", ResetResource: "Black Nexus",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 5000}},
	}, {
		Name: "Black Core", UnlockedBy: "Cryptotheology", ResetResource: "Black Core",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 10000}},
	}, {
		Name: "Event Horizon", UnlockedBy: "Cryptotheology", ResetResource: "Event Horizon",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 25000}},
	}, {
		Name: "Black Library", UnlockedBy: "Cryptotheology", ResetResource: "Black Library",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 30000}},
	}, {
		Name: "Black Radiance", UnlockedBy: "Cryptotheology", ResetResource: "Black Radiance",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 37500}},
	}, {
		Name: "Blazar", UnlockedBy: "Cryptotheology", ResetResource: "Blazar",
		CostExponentBase: 1.15,
		Costs:            []R{{Name: "relic", Count: 50000}},
	}, {
		Name: "Dark Nova", UnlockedBy: "Cryptotheology", ResetResource: "Dark Nova",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "relic", Count: 75000,
		}, {
			Name: "void", Count: 7500,
		}},
	}, {
		Name: "Mausoleum", UnlockedBy: "Cryptotheology", ResetResource: "Mausoleum",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "relic", Count: 50000,
		}, {
			Name: "void", Count: 12500,
		}, {
			Name: "necrocorn", Count: 10,
		}},
	}, {
		Name: "Holy Genocide", UnlockedBy: "Cryptotheology", ResetResource: "Holy Genocide",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "relic", Count: 100000,
		}, {
			Name: "void", Count: 25000,
		}},
	}, {
		Name: "Space Elevator", UnlockedBy: "Orbital Engineering",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "titanium", Count: 6000,
		}, {
			Name: "science", Count: 75000,
		}, {
			Name: "unobtainium", Count: 50,
		}},
	}, {
		Name: "Satellite", UnlockedBy: "Satellites",
		CostExponentBase: 1.08,
		Costs: []R{{
			Name: "starchart", Count: 325,
		}, {
			Name: "titanium", Count: 2500,
		}, {
			Name: "science", Count: 100000,
		}, {
			Name: "oil", Count: 15000, Bonus: []R{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Space Station", UnlockedBy: "Orbital Engineering",
		CostExponentBase: 1.12,
		Costs: []R{{
			Name: "starchart", Count: 425,
		}, {
			Name: "alloy", Count: 750,
		}, {
			Name: "science", Count: 150000,
		}, {
			Name: "oil", Count: 35000, Bonus: []R{{Name: "SpaceElevatorOilBonus"}},
		}},
		Adds: []R{{Name: "kitten", Cap: 2}},
	}, {
		Name: "Active Lunar Outpost", UnlockedBy: "Moon Mission",
		CostExponentBase: 1.12,
		Costs: []R{{
			Name: "starchart", Count: 650,
		}, {
			Name: "uranium", Count: 500,
		}, {
			Name: "alloy", Count: 750,
		}, {
			Name: "concrete", Count: 150,
		}, {
			Name: "science", Count: 100000,
		}, {
			Name: "oil", Count: 55000, Bonus: []R{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Moon Base", UnlockedBy: "Moon Mission",
		CostExponentBase: 1.12,
		Costs: []R{{
			Name: "starchart", Count: 700,
		}, {
			Name: "titanium", Count: 9500,
		}, {
			Name: "concrete", Count: 250,
		}, {
			Name: "science", Count: 100000,
		}, {
			Name: "unobtainium", Count: 50,
		}, {
			Name: "oil", Count: 70000, Bonus: []R{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Planet Cracker", UnlockedBy: "Dune Mission",
		CostExponentBase: 1.18,
		Costs: []R{{
			Name: "starchart", Count: 2500,
		}, {
			Name: "alloy", Count: 1750,
		}, {
			Name: "science", Count: 125000,
		}, {
			Name: "kerosene", Count: 50,
		}},
	}, {
		Name: "Hydraulic Fracturer", UnlockedBy: "Dune Mission",
		CostExponentBase: 1.18,
		Costs: []R{{
			Name: "starchart", Count: 750,
		}, {
			Name: "alloy", Count: 1025,
		}, {
			Name: "science", Count: 150000,
		}, {
			Name: "kerosene", Count: 100,
		}},
	}, {
		Name: "Spice Refinery", UnlockedBy: "Dune Mission",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "starchart", Count: 500,
		}, {
			Name: "alloy", Count: 500,
		}, {
			Name: "science", Count: 75000,
		}, {
			Name: "kerosene", Count: 125,
		}},
	}, {
		Name: "Research Vessel", UnlockedBy: "Piscine Mission",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "starchart", Count: 100,
		}, {
			Name: "alloy", Count: 2500,
		}, {
			Name: "titanium", Count: 12500,
		}, {
			Name: "kerosene", Count: 250,
		}},
	}, {
		Name: "Orbital Array", UnlockedBy: "Piscine Mission",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "eludium", Count: 100,
		}, {
			Name: "kerosene", Count: 500,
		}, {
			Name: "starchart", Count: 2000,
		}},
	}, {
		Name: "Sunlifter", UnlockedBy: "Helios Mission",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "science", Count: 500000,
		}, {
			Name: "eludium", Count: 225,
		}, {
			Name: "kerosene", Count: 2500,
		}},
	}, {
		Name: "Containment Chamber", UnlockedBy: "Helios Mission",
		CostExponentBase: 1.125,
		Costs: []R{{
			Name: "science", Count: 500000,
		}, {
			Name: "kerosene", Count: 2500,
		}},
	}, {
		Name: "Heatsink", UnlockedBy: "Helios Mission",
		CostExponentBase: 1.12,
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "thorium", Count: 12500,
		}, {
			Name: "relic", Count: 1,
		}, {
			Name: "kerosene", Count: 5000,
		}},
	}, {
		Name: "Sunforge", UnlockedBy: "Helios Mission",
		CostExponentBase: 1.12,
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "relic", Count: 1,
		}, {
			Name: "kerosene", Count: 1250,
		}, {
			Name: "antimatter", Count: 250,
		}},
	}, {
		Name: "Cryostation", UnlockedBy: "T-Minus Mission",
		CostExponentBase: 1.12,
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "eludium", Count: 25,
		}, {
			Name: "concrete", Count: 1500,
		}, {
			Name: "kerosene", Count: 500,
		}},
	}, {
		Name: "Space Beacon", UnlockedBy: "Kairo Mission",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "starchart", Count: 25000,
		}, {
			Name: "antimatter", Count: 50,
		}, {
			Name: "alloy", Count: 25000,
		}, {
			Name: "kerosene", Count: 7500,
		}},
	}, {
		Name: "Terraforming Station", UnlockedBy: "Terraformation",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "antimatter", Count: 25,
		}, {
			Name: "uranium", Count: 5000,
		}, {
			Name: "kerosene", Count: 5000,
		}},
		Adds: []R{{
			Name: "kitten", Cap: 1,
			Bonus: []R{{Name: "Hydroponics", Factor: 0.01}},
		}},
	}, {
		Name: "Hydroponics", UnlockedBy: "Hydroponics Tech",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "unobtainium", Count: 1,
		}, {
			Name: "kerosene", Count: 500,
		}},
	}, {
		Name: "HR Harvester", UnlockedBy: "Umbra Mission",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "relic", Count: 25,
		}, {
			Name: "antimatter", Count: 1250,
		}},
	}, {
		Name: "Entanglement Station", UnlockedBy: "Quantum Cryptography",
		CostExponentBase: 1.15,
		Costs: []R{{
			Name: "relic", Count: 1250,
		}, {
			Name: "antimatter", Count: 5250,
		}, {
			Name: "eludium", Count: 5000,
		}},
	}, {
		Name: "Tectonic", UnlockedBy: "Terraformation",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "science", Count: 600000,
		}, {
			Name: "antimatter", Count: 500,
		}, {
			Name: "thorium", Count: 75000,
		}},
	}, {
		Name: "Molten Core", UnlockedBy: "Exophysics",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "science", Count: 25000000,
		}, {
			Name: "uranium", Count: 5000000,
		}},
	}, {
		Name: "Hash Level", UnlockedBy: "Entanglement Station",
		CostExponentBase: 1.6,
		Costs:            []R{{Name: "hash", Count: 1600}},
	}, {
		Name: "Temporal Battery", UnlockedBy: "Chronoforge",
		CostExponentBase: 1.25,
		Costs:            []R{{Name: "time crystal", Count: 5}},
	}, {
		Name: "Chrono Furnace", UnlockedBy: "Chronoforge",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "time crystal", Count: 25,
		}, {
			Name: "relic", Count: 5,
		}},
	}, {
		Name: "Time Boiler", UnlockedBy: "Chronoforge",
		CostExponentBase: 1.25,
		Costs:            []R{{Name: "time crystal", Count: 25000}},
	}, {
		Name: "Temporal Accelerator", UnlockedBy: "Chronoforge",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "time crystal", Count: 10,
		}, {
			Name: "relic", Count: 1000,
		}},
	}, {
		Name: "Time Impedance", UnlockedBy: "Chronoforge",
		CostExponentBase: 1.05,
		Costs: []R{{
			Name: "time crystal", Count: 100,
		}, {
			Name: "relic", Count: 250,
		}},
	}, {
		Name: "Resource Retrieval", UnlockedBy: "Paradox Theory",
		CostExponentBase: 1.3,
		Costs:            []R{{Name: "time crystal", Count: 1000}},
	}, {
		Name: "Temporal Press", UnlockedBy: "Chronosurge",
		CostExponentBase: 1.1,
		Costs: []R{{
			Name: "time crystal", Count: 100,
		}, {
			Name: "void", Count: 10,
		}},
	}, {
		Name: "Cryochambers", UnlockedBy: "Void Space",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "time crystal", Count: 2,
		}, {
			Name: "void", Count: 100,
		}, {
			Name: "karma", Count: 1,
		}},
	}, {
		Name: "Void Hoover", UnlockedBy: "Void Aspiration",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "time crystal", Count: 10,
		}, {
			Name: "void", Count: 250,
		}, {
			Name: "antimatter", Count: 1000,
		}},
	}, {
		Name: "Void Rift", UnlockedBy: "Void Aspiration",
		CostExponentBase: 1.3,
		Costs:            []R{{Name: "void", Count: 75}},
	}, {
		Name: "Chronocontrol", UnlockedBy: "Paradox Theory",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "time crystal", Count: 30,
		}, {
			Name: "void", Count: 500,
		}, {
			Name: "temporal flux", Count: 3000,
		}},
	}, {
		Name: "Void Resonator", UnlockedBy: "Paradox Theory",
		CostExponentBase: 1.25,
		Costs: []R{{
			Name: "time crystal", Count: 1000,
		}, {
			Name: "relic", Count: 10000,
		}, {
			Name: "void", Count: 50,
		}},
	}, {
		Name: "karma", UnlockedBy: "Hut", ResetResource: "karma",
		CostExponentBase: 1.13,
		Costs:            []R{{Name: "karma kitten", Count: 1}},
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
		Costs: []R{{Name: "catpower", Count: 100}},
		Adds: []R{{
			Name: "fur", Count: 39.5, Bonus: []R{{Name: "HuntingBonus"}},
		}, {
			Name: "ivory", Count: 10.78, Bonus: []R{{Name: "HuntingBonus"}},
		}, {
			Name: "unicorn", Count: 0.05,
		}},
	}, {
		Name: "Festival", Type: "Job", UnlockedBy: "Drama and Poetry",
		Costs: []R{{
			Name: "catpower", Count: 1500,
		}, {
			Name: "culture", Count: 5000,
		}, {
			Name: "parchment", Count: 2500,
		}},
		Adds: []R{{
			Name: "festival day", Count: 400,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Lizards", Type: "Trade", UnlockedBy: "Archery",
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "mineral", Count: 1000,
		}},
		Adds: []R{{
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
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "iron", Count: 100,
		}},
		Adds: []R{{
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
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "wood", Count: 500,
		}},
		Adds: []R{{
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
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "ivory", Count: 500,
		}},
		Adds: []R{{
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
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "slab", Count: 50,
		}},
		Adds: []R{{
			Name: "iron", Count: 300,
		}, {
			Name: "plate", Count: 2 * 0.65,
		}, {
			Name: "titanium", Count: 1.5 * 0.15,
			Bonus: []R{{Name: "ship", Factor: 0.03 * 0.0035}},
		}, {
			Name: "alloy", Count: 0.25 * 0.05,
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}, {
		Name: "Spiders", Type: "Trade", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "scaffold", Count: 50,
		}},
		Adds: []R{{
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
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "titanium", Count: 250,
		}},
		Adds: []R{{
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
		Costs: []R{{
			Name: "catpower", Count: 50,
		}, {
			Name: "gold", Count: 15,
		}, {
			Name: "unobtainium", Count: 5000,
		}},
		Adds: []R{{
			Name: "starchart", Count: 250 * 0.50,
			Bonus: []R{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "time crystal", Count: 0.25 * 0.98,
			Bonus: []R{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "sorrow", Count: 1 * 0.15,
			Bonus: []R{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "relic", Count: 1 * 0.05,
			Bonus: []R{{Name: "leviathan energy", Factor: 0.02}},
		}, {
			Name: "blueprint", Count: 0.10,
		}, {
			Name: "spice", Count: 8.75,
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Feed Leviathans", Type: "Craft", UnlockedBy: "Black Pyramid",
		Costs: []R{{Name: "necrocorn", Count: 1}},
		Adds:  []R{{Name: "leviathan energy", Count: 1}},
	}, {
		Name: "Sacrifice Unicorns", Type: "Craft", UnlockedBy: "Ziggurat",
		Costs: []R{{Name: "unicorn", Count: 2500}},
		Adds: []R{{
			Name: "tear", Count: 1,
			Bonus:               []R{{Name: "Ziggurat"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "Praise the sun!", Type: "Craft", UnlockedBy: "Philosophy",
		Costs: []R{{Name: "faith", Count: 1}},
		Adds: []R{{
			Name: "worship", Count: 1,
			Bonus: []R{{Name: "epiphany"}},
		}},
	}, {
		Name: "Adore the galaxy", Type: "Craft", UnlockedBy: "Apocrypha",
		Costs: []R{{Name: "worship", Count: 1}},
		Adds: []R{{
			Name: "epiphany", Count: 1. / 1000000,
			Bonus: []R{{
				Name:  "Transcendence Level",
				Bonus: []R{{Name: "Transcendence Level"}},
			}},
		}},
	}, {
		Name: "Sacrifice Alicorns", Type: "Craft", UnlockedBy: "Ziggurat",
		Costs: []R{{Name: "alicorn", Count: 25}},
		Adds: []R{{
			Name: "time crystal", Count: 1,
			Bonus: []R{{
				Name: "Unicorn Utopia", Factor: 0.05,
			}, {
				Name: "Sunspire", Factor: 0.10,
			}},
		}},
	}, {
		Name: "Refine Tears", Type: "Craft", UnlockedBy: "Megalomania",
		Costs: []R{{Name: "tear", Count: 10000}},
		Adds:  []R{{Name: "sorrow", Count: 1}},
	}, {
		Name: "Refine Time Crystals", Type: "Craft", UnlockedBy: "Ziggurat",
		Costs: []R{{Name: "time crystal", Count: 25}},
		Adds: []R{{
			Name: "relic", Count: 1,
			Bonus: []R{{
				Name:                "Black Nexus",
				Bonus:               []R{{Name: "Black Pyramid"}},
				BonusStartsFromZero: true,
			}},
		}},
	}})

	g.AddActions([]data.Action{{
		Name: "Combust time crystal", Type: "Craft", UnlockedBy: "Chronoforge",
		Costs: []R{{Name: "time crystal", Count: 1}},
		Adds: []R{{
			Name: "day", Count: 400,
		}, {
			Name: "chronoheat", Count: 10,
		}},
	}, {
		Name: "Burn Chrono Furnace Fuel", Type: "Craft", UnlockedBy: "Chronoforge",
		Costs: []R{{Name: "chrono furnace fuel", Count: 100}},
		Adds:  []R{{Name: "day", Count: 400}},
	}, {
		Name: "Burn Paragon", Type: "Craft", UnlockedBy: "Chronoforge",
		Costs: []R{{Name: "paragon", Count: 1}},
		Adds:  []R{{Name: "burned paragon", Count: 1}},
	}})

	addTechs(g, []data.Action{{
		Name: "Calendar", UnlockedBy: "Library",
		Costs: []R{{Name: "science", Count: 30}},
	}, {
		Name: "Agriculture", UnlockedBy: "Calendar",
		Costs: []R{{Name: "science", Count: 100}},
	}, {
		Name: "Archery", UnlockedBy: "Agriculture",
		Costs: []R{{Name: "science", Count: 300}},
	}, {
		Name: "Mining", UnlockedBy: "Agriculture",
		Costs: []R{{Name: "science", Count: 500}},
	}, {
		Name: "Animal Husbandry", UnlockedBy: "Archery",
		Costs: []R{{Name: "science", Count: 500}},
	}, {
		Name: "Metal Working", UnlockedBy: "Mining",
		Costs: []R{{Name: "science", Count: 900}},
	}, {
		Name: "Civil Service", UnlockedBy: "Animal Husbandry",
		Costs: []R{{Name: "science", Count: 1500}},
	}, {
		Name: "Mathematics", UnlockedBy: "Animal Husbandry",
		Costs: []R{{Name: "science", Count: 1000}},
	}, {
		Name: "Construction", UnlockedBy: "Animal Husbandry",
		Costs: []R{{Name: "science", Count: 1300}},
	}, {
		Name: "Currency", UnlockedBy: "Civil Service",
		Costs: []R{{Name: "science", Count: 2200}},
	}, {
		Name: "Celestial Mechanics", UnlockedBy: "Mathematics",
		Costs: []R{{Name: "science", Count: 250}},
	}, {
		Name: "Engineering", UnlockedBy: "Construction",
		Costs: []R{{Name: "science", Count: 1500}},
	}, {
		Name: "Writing", UnlockedBy: "Engineering",
		Costs: []R{{Name: "science", Count: 3600}},
	}, {
		Name: "Philosophy", UnlockedBy: "Writing",
		Costs: []R{{Name: "science", Count: 9500}},
	}, {
		Name: "Steel", UnlockedBy: "Writing",
		Costs: []R{{Name: "science", Count: 12000}},
	}, {
		Name: "Machinery", UnlockedBy: "Writing",
		Costs: []R{{Name: "science", Count: 15000}},
	}, {
		Name: "Theology", UnlockedBy: "Philosophy",
		Costs: []R{{
			Name: "science", Count: 20000,
		}, {
			Name: "manuscript", Count: 35,
		}},
	}, {
		Name: "Astronomy", UnlockedBy: "Theology",
		Costs: []R{{
			Name: "science", Count: 28000,
		}, {
			Name: "manuscript", Count: 65,
		}},
	}, {
		Name: "Navigation", UnlockedBy: "Astronomy",
		Costs: []R{{
			Name: "science", Count: 35000,
		}, {
			Name: "manuscript", Count: 100,
		}},
	}, {
		Name: "Architecture", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 42000,
		}, {
			Name: "compendium", Count: 10,
		}},
	}, {
		Name: "Physics", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 50000,
		}, {
			Name: "compendium", Count: 35,
		}},
	}, {
		Name: "Metaphysics", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "science", Count: 55000,
		}, {
			Name: "unobtainium", Count: 5,
		}},
	}, {
		Name: "Chemistry", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "science", Count: 60000,
		}, {
			Name: "compendium", Count: 50,
		}},
	}, {
		Name: "Acoustics", UnlockedBy: "Architecture",
		Costs: []R{{
			Name: "science", Count: 60000,
		}, {
			Name: "compendium", Count: 60,
		}},
	}, {
		Name: "Geology", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 65000,
		}, {
			Name: "compendium", Count: 65,
		}},
	}, {
		Name: "Drama and Poetry", UnlockedBy: "Acoustics",
		Costs: []R{{
			Name: "science", Count: 90000,
		}, {
			Name: "parchment", Count: 5000,
		}},
	}, {
		Name: "Electricity", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "science", Count: 75000,
		}, {
			Name: "compendium", Count: 85,
		}},
	}, {
		Name: "Biology", UnlockedBy: "Geology",
		Costs: []R{{
			Name: "science", Count: 85000,
		}, {
			Name: "compendium", Count: 100,
		}},
	}, {
		Name: "Biochemistry", UnlockedBy: "Biology",
		Costs: []R{{
			Name: "science", Count: 145000,
		}, {
			Name: "compendium", Count: 500,
		}},
	}, {
		Name: "Genetics", UnlockedBy: "Biochemistry",
		Costs: []R{{
			Name: "science", Count: 190000,
		}, {
			Name: "compendium", Count: 1500,
		}},
	}, {
		Name: "Industrialization", UnlockedBy: "Electricity",
		Costs: []R{{
			Name: "science", Count: 10000,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Mechanization", UnlockedBy: "Industrialization",
		Costs: []R{{
			Name: "science", Count: 115000,
		}, {
			Name: "blueprint", Count: 45,
		}},
	}, {
		Name: "Combustion", UnlockedBy: "Industrialization",
		Costs: []R{{
			Name: "science", Count: 115000,
		}, {
			Name: "blueprint", Count: 45,
		}},
	}, {
		Name: "Metallurgy", UnlockedBy: "Industrialization",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "blueprint", Count: 60,
		}},
	}, {
		Name: "Ecology", UnlockedBy: "Combustion",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "blueprint", Count: 55,
		}},
	}, {
		Name: "Electronics", UnlockedBy: "Mechanization",
		Costs: []R{{
			Name: "science", Count: 135000,
		}, {
			Name: "blueprint", Count: 70,
		}},
	}, {
		Name: "Robotics", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 140000,
		}, {
			Name: "blueprint", Count: 80,
		}},
	}, {
		Name: "Artificial Intelligence", UnlockedBy: "Robotics",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "blueprint", Count: 150,
		}},
	}, {
		Name: "Quantum Cryptography", UnlockedBy: "Artificial Intelligence",
		Costs: []R{{
			Name: "science", Count: 1250000,
		}, {
			Name: "relic", Count: 1024,
		}},
	}, {
		Name: "Blackchain", UnlockedBy: "Quantum Cryptography",
		Costs: []R{{
			Name: "science", Count: 5000000,
		}, {
			Name: "relic", Count: 4096,
		}},
	}, {
		Name: "Nuclear Fission", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "blueprint", Count: 100,
		}},
	}, {
		Name: "Rocketry", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "blueprint", Count: 125,
		}},
	}, {
		Name: "Oil Processing", UnlockedBy: "Rocketry",
		Costs: []R{{
			Name: "science", Count: 215000,
		}, {
			Name: "blueprint", Count: 150,
		}},
	}, {
		Name: "Satellites", UnlockedBy: "Rocketry",
		Costs: []R{{
			Name: "science", Count: 190000,
		}, {
			Name: "blueprint", Count: 125,
		}},
	}, {
		Name: "Orbital Engineering", UnlockedBy: "Satellites",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "blueprint", Count: 250,
		}},
	}, {
		Name: "Thorium", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 375000,
		}, {
			Name: "blueprint", Count: 375,
		}},
	}, {
		Name: "Exogeology", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 275000,
		}, {
			Name: "blueprint", Count: 250,
		}},
	}, {
		Name: "Advanced Exogeology", UnlockedBy: "Exogeology",
		Costs: []R{{
			Name: "science", Count: 325000,
		}, {
			Name: "blueprint", Count: 350,
		}},
	}, {
		Name: "Nanotechnology", UnlockedBy: "Nuclear Fission",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "blueprint", Count: 150,
		}},
	}, {
		Name: "Superconductors", UnlockedBy: "Nanotechnology",
		Costs: []R{{
			Name: "science", Count: 225000,
		}, {
			Name: "blueprint", Count: 175,
		}},
	}, {
		Name: "Antimatter", UnlockedBy: "Superconductors",
		Costs: []R{{
			Name: "science", Count: 500000,
		}, {
			Name: "relic", Count: 1,
		}},
	}, {
		Name: "Terraformation", UnlockedBy: "Antimatter",
		Costs: []R{{
			Name: "science", Count: 750000,
		}, {
			Name: "relic", Count: 5,
		}},
	}, {
		Name: "Hydroponics Tech", UnlockedBy: "Terraformation",
		Costs: []R{{
			Name: "science", Count: 1000000,
		}, {
			Name: "relic", Count: 25,
		}},
	}, {
		Name: "Exophysics", UnlockedBy: "Hydroponics Tech",
		Costs: []R{{
			Name: "science", Count: 25000000,
		}, {
			Name: "relic", Count: 500,
		}},
	}, {
		Name: "Particle Physics", UnlockedBy: "Nuclear Fission",
		Costs: []R{{
			Name: "science", Count: 185000,
		}, {
			Name: "blueprint", Count: 135,
		}},
	}, {
		Name: "Dimensional Physics", UnlockedBy: "Particle Physics",
		Costs: []R{{Name: "science", Count: 235000}},
	}, {
		Name: "Chronophysics", UnlockedBy: "Particle Physics", ResetResource: "reset Chronophysics",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "time crystal", Count: 5,
		}},
	}, {
		Name: "Tachyon Theory", UnlockedBy: "Chronophysics",
		Costs: []R{{
			Name: "science", Count: 750000,
		}, {
			Name: "time crystal", Count: 25,
		}, {
			Name: "relic", Count: 1,
		}},
	}, {
		Name: "Cryptotheology", UnlockedBy: "Theology",
		Costs: []R{{
			Name: "science", Count: 650000,
		}, {
			Name: "relic", Count: 5,
		}},
	}, {
		Name: "Void Space", UnlockedBy: "Tachyon Theory",
		Costs: []R{{
			Name: "science", Count: 800000,
		}, {
			Name: "time crystal", Count: 30,
		}, {
			Name: "void", Count: 100,
		}},
	}, {
		Name: "Paradox Theory", UnlockedBy: "Void Space",
		Costs: []R{{
			Name: "science", Count: 1000000,
		}, {
			Name: "time crystal", Count: 40,
		}, {
			Name: "void", Count: 250,
		}},
	}, {
		Name: "Mineral Hoes", UnlockedBy: "Workshop",
		Costs: []R{{
			Name: "mineral", Count: 275,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Iron Hoes", UnlockedBy: "Workshop",
		Costs: []R{{
			Name: "iron", Count: 25,
		}, {
			Name: "science", Count: 200,
		}},
	}, {
		Name: "Mineral Axe", UnlockedBy: "Workshop",
		Costs: []R{{
			Name: "mineral", Count: 500,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Iron Axe", UnlockedBy: "Workshop",
		Costs: []R{{
			Name: "iron", Count: 50,
		}, {
			Name: "science", Count: 100,
		}},
	}, {
		Name: "Steel Axe", UnlockedBy: "Steel",
		Costs: []R{{
			Name: "steel", Count: 75,
		}, {
			Name: "science", Count: 20000,
		}},
	}, {
		Name: "Reinforced Saw", UnlockedBy: "Construction",
		Costs: []R{{
			Name: "iron", Count: 1000,
		}, {
			Name: "science", Count: 2500,
		}},
	}, {
		Name: "Steel Saw", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "steel", Count: 750,
		}, {
			Name: "science", Count: 52000,
		}},
	}, {
		Name: "Titanium Saw", UnlockedBy: "Steel Saw",
		Costs: []R{{
			Name: "titanium", Count: 500,
		}, {
			Name: "science", Count: 70000,
		}},
	}, {
		Name: "Alloy Saw", UnlockedBy: "Titanium Saw",
		Costs: []R{{
			Name: "alloy", Count: 75,
		}, {
			Name: "science", Count: 85000,
		}},
	}, {
		Name: "Titanium Axe", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 38000,
		}, {
			Name: "titanium", Count: 10,
		}},
	}, {
		Name: "Alloy Axe", UnlockedBy: "Chemistry",
		Costs: []R{{
			Name: "science", Count: 70000,
		}, {
			Name: "alloy", Count: 25,
		}},
	}, {
		Name: "Expanded Barns", UnlockedBy: "Workshop",
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 75000,
		}, {
			Name: "alloy", Count: 20,
		}, {
			Name: "plate", Count: 750,
		}},
	}, {
		Name: "Concrete Barns", UnlockedBy: "Concrete Pillars",
		Costs: []R{{
			Name: "science", Count: 2000,
		}, {
			Name: "concrete", Count: 45,
		}, {
			Name: "titanium", Count: 2000,
		}},
	}, {
		Name: "Titanium Warehouses", UnlockedBy: "Silos",
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 90000,
		}, {
			Name: "titanium", Count: 750,
		}, {
			Name: "alloy", Count: 50,
		}},
	}, {
		Name: "Concrete Warehouses", UnlockedBy: "Concrete Pillars",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 1250,
		}, {
			Name: "concrete", Count: 35,
		}},
	}, {
		Name: "Storage Bunkers", UnlockedBy: "Exogeology",
		Costs: []R{{
			Name: "science", Count: 25000,
		}, {
			Name: "unobtainium", Count: 500,
		}, {
			Name: "concrete", Count: 1250,
		}},
	}, {
		Name: "Energy Rifts", UnlockedBy: "Dimensional Physics",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "titanium", Count: 7500,
		}, {
			Name: "uranium", Count: 250,
		}},
	}, {
		Name: "Stasis Chambers", UnlockedBy: "Chronophysics",
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 350000,
		}, {
			Name: "eludium", Count: 75,
		}, {
			Name: "time crystal", Count: 3,
		}},
	}, {
		Name: "Chronoforge", UnlockedBy: "Tachyon Theory",
		Costs: []R{{
			Name: "science", Count: 500000,
		}, {
			Name: "relic", Count: 5,
		}, {
			Name: "time crystal", Count: 10,
		}},
	}, {
		Name: "Tachyon Accelerators", UnlockedBy: "Tachyon Theory",
		Costs: []R{{
			Name: "science", Count: 500000,
		}, {
			Name: "eludium", Count: 125,
		}, {
			Name: "time crystal", Count: 10,
		}},
	}, {
		Name: "Flux Condensator", UnlockedBy: "Chronophysics",
		Costs: []R{{
			Name: "alloy", Count: 250,
		}, {
			Name: "unobtainium", Count: 5000,
		}, {
			Name: "time crystal", Count: 5,
		}},
	}, {
		Name: "LHC", UnlockedBy: "Dimensional Physics",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 100,
		}, {
			Name: "alloy", Count: 150,
		}},
	}, {
		Name: "Photovoltaic Cells", UnlockedBy: "Nanotechnology",
		Costs: []R{{
			Name: "science", Count: 75000,
		}, {
			Name: "titanium", Count: 5000,
		}},
	}, {
		Name: "Thin Film Cells", UnlockedBy: "Satellites",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "unobtainium", Count: 200,
		}, {
			Name: "uranium", Count: 1000,
		}},
	}, {
		Name: "Quantum Dot Cells", UnlockedBy: "Thorium",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "eludium", Count: 200,
		}, {
			Name: "thorium", Count: 1000,
		}},
	}, {
		Name: "Solar Satellites", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 225000,
		}, {
			Name: "alloy", Count: 750,
		}},
	}, {
		Name: "Expanded Cargo", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 55000,
		}, {
			Name: "blueprint", Count: 15,
		}},
	}, {
		Name: "Barges", UnlockedBy: "Industrialization",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 1500,
		}, {
			Name: "blueprint", Count: 30,
		}},
	}, {
		Name: "Reactor Vessel", UnlockedBy: "Nuclear Fission",
		Costs: []R{{
			Name: "science", Count: 135000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "uranium", Count: 125,
		}},
	}, {
		Name: "Ironwood Huts", UnlockedBy: "Reinforced Warehouses",
		Costs: []R{{
			Name: "science", Count: 30000,
		}, {
			Name: "wood", Count: 15000,
		}, {
			Name: "iron", Count: 3000,
		}},
	}, {
		Name: "Concrete Huts", UnlockedBy: "Concrete Pillars",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "concrete", Count: 45,
		}, {
			Name: "titanium", Count: 3000,
		}},
	}, {
		Name: "Unobtainium Huts", UnlockedBy: "Exogeology",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "unobtainium", Count: 350,
		}, {
			Name: "titanium", Count: 15000,
		}},
	}, {
		Name: "Eludium Huts", UnlockedBy: "Advanced Exogeology",
		Costs: []R{{
			Name: "science", Count: 275000,
		}, {
			Name: "eludium", Count: 125,
		}},
	}, {
		Name: "Silos", UnlockedBy: "Ironwood Huts",
		Costs: []R{{
			Name: "science", Count: 50000,
		}, {
			Name: "steel", Count: 125,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "Refrigeration", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 2500,
		}, {
			Name: "blueprint", Count: 15,
		}},
	}, {
		Name: "Composite Bow", UnlockedBy: "Construction",
		Costs: []R{{
			Name: "science", Count: 500,
		}, {
			Name: "iron", Count: 100,
		}, {
			Name: "wood", Count: 200,
		}},
	}, {
		Name: "Crossbow", UnlockedBy: "Machinery",
		Costs: []R{{
			Name: "science", Count: 12000,
		}, {
			Name: "iron", Count: 1500,
		}},
	}, {
		Name: "Railgun", UnlockedBy: "Particle Physics",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Bolas", UnlockedBy: "Mining",
		Costs: []R{{
			Name: "science", Count: 1000,
		}, {
			Name: "mineral", Count: 250,
		}, {
			Name: "wood", Count: 50,
		}},
	}, {
		Name: "Hunting Armour", UnlockedBy: "Metal Working",
		Costs: []R{{
			Name: "science", Count: 2000,
		}, {
			Name: "iron", Count: 750,
		}},
	}, {
		Name: "Steel Armour", UnlockedBy: "Steel",
		Costs: []R{{
			Name: "science", Count: 10000,
		}, {
			Name: "steel", Count: 50,
		}},
	}, {
		Name: "Alloy Armour", UnlockedBy: "Chemistry",
		Costs: []R{{
			Name: "science", Count: 50000,
		}, {
			Name: "alloy", Count: 25,
		}},
	}, {
		Name: "Nanosuits", UnlockedBy: "Nanotechnology",
		Costs: []R{{
			Name: "science", Count: 185000,
		}, {
			Name: "alloy", Count: 250,
		}},
	}, {
		Name: "Caravanserai", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 25000,
		}, {
			Name: "ivory", Count: 10000,
		}, {
			Name: "gold", Count: 250,
		}},
	}, {
		Name: "Catnip Enrichment", UnlockedBy: "Construction",
		Costs: []R{{
			Name: "science", Count: 500,
		}, {
			Name: "catnip", Count: 5000,
		}},
	}, {
		Name: "Gold Ore", UnlockedBy: "Currency",
		Costs: []R{{
			Name: "science", Count: 1000,
		}, {
			Name: "mineral", Count: 800,
		}, {
			Name: "iron", Count: 100,
		}},
	}, {
		Name: "Geodesy", UnlockedBy: "Geology",
		Costs: []R{{
			Name: "science", Count: 90000,
		}, {
			Name: "titanium", Count: 250,
		}, {
			Name: "starchart", Count: 500,
		}},
	}, {
		Name: "Register", UnlockedBy: "Writing",
		Costs: []R{{
			Name: "science", Count: 500,
		}, {
			Name: "gold", Count: 10,
		}},
	}, {
		Name: "Concrete Pillars", UnlockedBy: "Mechanization",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "concrete", Count: 50,
		}},
	}, {
		Name: "Mining Drill", UnlockedBy: "Metallurgy",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 1750,
		}, {
			Name: "steel", Count: 750,
		}},
	}, {
		Name: "Unobtainium Drill", UnlockedBy: "Exogeology",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 250,
		}, {
			Name: "alloy", Count: 1250,
		}},
	}, {
		Name: "Coal Furnace", UnlockedBy: "Steel",
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 5000,
		}, {
			Name: "iron", Count: 1200,
		}, {
			Name: "beam", Count: 50,
		}},
	}, {
		Name: "Pyrolysis", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "science", Count: 35000,
		}, {
			Name: "compendium", Count: 5,
		}},
	}, {
		Name: "Electrolytic Smelting", UnlockedBy: "Metallurgy",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 2000,
		}},
	}, {
		Name: "Oxidation", UnlockedBy: "Metallurgy",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "steel", Count: 5000,
		}},
	}, {
		Name: "Steel Plants", UnlockedBy: "Robotics",
		Costs: []R{{
			Name: "science", Count: 140000,
		}, {
			Name: "titanium", Count: 3500,
		}, {
			Name: "gear", Count: 750,
		}},
	}, {
		Name: "Automated Plants", UnlockedBy: "Steel Plants",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "alloy", Count: 750,
		}},
	}, {
		Name: "Nuclear Plants", UnlockedBy: "Automated Plants",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "uranium", Count: 10000,
		}},
	}, {
		Name: "Rotary Kiln", UnlockedBy: "Robotics",
		Costs: []R{{
			Name: "science", Count: 145000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "gear", Count: 500,
		}},
	}, {
		Name: "Fluoridized Reactors", UnlockedBy: "Nanotechnology",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "alloy", Count: 200,
		}},
	}, {
		Name: "Nuclear Smelter", UnlockedBy: "Nuclear Fission",
		Costs: []R{{
			Name: "science", Count: 165000,
		}, {
			Name: "uranium", Count: 250,
		}},
	}, {
		Name: "Orbital Geodesy", UnlockedBy: "Satellites",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "alloy", Count: 1000,
		}, {
			Name: "oil", Count: 35000,
		}},
	}, {
		Name: "Printing Press", UnlockedBy: "Machinery",
		Costs: []R{{
			Name: "science", Count: 7500,
		}, {
			Name: "gear", Count: 45,
		}},
	}, {
		Name: "Offset Press", UnlockedBy: "Combustion",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "oil", Count: 15000,
		}},
	}, {
		Name: "Photolithography", UnlockedBy: "Satellites",
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 75000,
		}, {
			Name: "alloy", Count: 1750,
		}},
	}, {
		Name: "Starlink", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "alloy", Count: 5000,
		}, {
			Name: "oil", Count: 25000,
		}},
	}, {
		Name: "Cryocomputing", UnlockedBy: "Superconductors",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "eludium", Count: 15,
		}},
	}, {
		Name: "Machine Learning", UnlockedBy: "Artificial Intelligence",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "eludium", Count: 25,
		}, {
			Name: "antimatter", Count: 125,
		}},
	}, {
		Name: "Workshop Automation", UnlockedBy: "Machinery",
		Costs: []R{{
			Name: "science", Count: 10000,
		}, {
			Name: "gear", Count: 25,
		}},
	}, {
		Name: "Advanced Automation", UnlockedBy: "Industrialization",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 75,
		}, {
			Name: "blueprint", Count: 25,
		}},
	}, {
		Name: "Pneumatic Press", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "science", Count: 20000,
		}, {
			Name: "gear", Count: 30,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "High Pressure Engine", UnlockedBy: "Steel",
		Costs: []R{{
			Name: "science", Count: 20000,
		}, {
			Name: "gear", Count: 25,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "Fuel Injector", UnlockedBy: "Combustion",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "oil", Count: 20000,
		}},
	}, {
		Name: "Factory Logistics", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "titanium", Count: 2000,
		}},
	}, {
		Name: "Carbon Sequestration", UnlockedBy: "Ecology",
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "titanium", Count: 125000,
		}},
	}, {
		Name: "Astrolabe", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 25000,
		}, {
			Name: "titanium", Count: 5,
		}, {
			Name: "starchart", Count: 75,
		}},
	}, {
		Name: "Titanium Reflectors", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "science", Count: 20000,
		}, {
			Name: "titanium", Count: 15,
		}, {
			Name: "starchart", Count: 20,
		}},
	}, {
		Name: "Unobtainium Reflectors", UnlockedBy: "Exogeology",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 75,
		}, {
			Name: "starchart", Count: 750,
		}},
	}, {
		Name: "Eludium Reflectors", UnlockedBy: "Advanced Exogeology",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "eludium", Count: 15,
		}},
	}, {
		Name: "Hydro Plant Turbines", UnlockedBy: "Exogeology",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 125,
		}},
	}, {
		Name: "Antimatter Bases", UnlockedBy: "Antimatter",
		Costs: []R{{
			Name: "eludium", Count: 15,
		}, {
			Name: "antimatter", Count: 250,
		}},
	}, {
		Name: "AI Bases", UnlockedBy: "Antimatter Bases",
		Costs: []R{{
			Name: "science", Count: 750000,
		}, {
			Name: "antimatter", Count: 7500,
		}},
	}, {
		Name: "Antimatter Fission", UnlockedBy: "Antimatter",
		Costs: []R{{
			Name: "science", Count: 525000,
		}, {
			Name: "antimatter", Count: 175,
		}, {
			Name: "thorium", Count: 7500,
		}},
	}, {
		Name: "Antimatter Drive", UnlockedBy: "Antimatter",
		Costs: []R{{
			Name: "science", Count: 450000,
		}, {
			Name: "antimatter", Count: 125,
		}},
	}, {
		Name: "Antimatter Reactors", UnlockedBy: "Antimatter",
		Costs: []R{{
			Name: "eludium", Count: 35,
		}, {
			Name: "antimatter", Count: 750,
		}},
	}, {
		Name: "Advanced AM Reactors", UnlockedBy: "Antimatter Reactors",
		Costs: []R{{
			Name: "eludium", Count: 70,
		}, {
			Name: "antimatter", Count: 1750,
		}},
	}, {
		Name: "Void Reactors", UnlockedBy: "Advanced AM Reactors",
		Costs: []R{{
			Name: "void", Count: 250,
		}, {
			Name: "antimatter", Count: 2500,
		}},
	}, {
		Name: "Relic Station", UnlockedBy: "Cryptotheology",
		Costs: []R{{
			Name: "eludium", Count: 100,
		}, {
			Name: "antimatter", Count: 5000,
		}},
	}, {
		Name: "Pumpjack", UnlockedBy: "Mechanization",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "titanium", Count: 250,
		}, {
			Name: "gear", Count: 125,
		}},
	}, {
		Name: "Biofuel Processing", UnlockedBy: "Biochemistry",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 1250,
		}},
	}, {
		Name: "Unicorn Selection", UnlockedBy: "Genetics",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 1500,
		}},
	}, {
		Name: "GM Catnip", UnlockedBy: "Genetics",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 1500,
		}, {
			Name: "catnip", Count: 1000000,
		}},
	}, {
		Name: "CAD System", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 750,
		}},
	}, {
		Name: "SETI", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 250,
		}},
	}, {
		Name: "Logistics", UnlockedBy: "Industrialization",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "gear", Count: 100,
		}, {
			Name: "scaffold", Count: 1000,
		}},
	}, {
		Name: "Augmentations", UnlockedBy: "Nanotechnology",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "uranium", Count: 50,
		}},
	}, {
		Name: "Cold Fusion", UnlockedBy: "Superconductors",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "eludium", Count: 25,
		}},
	}, {
		Name: "Thorium Reactors", UnlockedBy: "Thorium",
		Costs: []R{{
			Name: "science", Count: 400000,
		}, {
			Name: "thorium", Count: 10000,
		}},
	}, {
		Name: "Enriched Uranium", UnlockedBy: "Particle Physics",
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 7500,
		}, {
			Name: "uranium", Count: 150,
		}},
	}, {
		Name: "Enriched Thorium", UnlockedBy: "Thorium Reactors",
		Costs: []R{{
			Name: "science", Count: 12500,
		}, {
			Name: "thorium", Count: 500000,
		}},
	}, {
		Name: "Oil Refinery", UnlockedBy: "Combustion",
		Costs: []R{{
			Name: "science", Count: 125000,
		}, {
			Name: "titanium", Count: 1250,
		}, {
			Name: "gear", Count: 500,
		}},
	}, {
		Name: "Hubble Space Telescope", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "alloy", Count: 1250,
		}, {
			Name: "oil", Count: 50000,
		}},
	}, {
		Name: "Satellite Navigation", UnlockedBy: "Hubble Space Telescope",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "alloy", Count: 750,
		}},
	}, {
		Name: "Satellite Radio", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 225000,
		}, {
			Name: "alloy", Count: 5000,
		}},
	}, {
		Name: "Astrophysicists", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 250000,
		}, {
			Name: "unobtainium", Count: 350,
		}},
	}, {
		Name: "Microwarp Reactors", UnlockedBy: "Advanced Exogeology",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "eludium", Count: 50,
		}},
	}, {
		Name: "Planet Busters", UnlockedBy: "Advanced Exogeology",
		Costs: []R{{
			Name: "science", Count: 275000,
		}, {
			Name: "eludium", Count: 250,
		}},
	}, {
		Name: "Thorium Drive", UnlockedBy: "Thorium",
		Costs: []R{{
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
		Costs: []R{{
			Name: "science", Count: 175000,
		}, {
			Name: "titanium", Count: 5000,
		}},
	}, {
		Name: "Factory Processing", UnlockedBy: "Oil Processing",
		Costs: []R{{
			Name: "science", Count: 195000,
		}, {
			Name: "titanium", Count: 7500,
		}, {
			Name: "concrete", Count: 125,
		}},
	}, {
		Name: "Factory Optimization", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 75000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "titanium", Count: 1250,
		}},
	}, {
		Name: "Space Engineers", UnlockedBy: "Orbital Engineering",
		Costs: []R{{
			Name: "science", Count: 225000,
		}, {
			Name: "alloy", Count: 200,
		}},
	}, {
		Name: "AI Engineers", UnlockedBy: "Artificial Intelligence",
		Costs: []R{{
			Name: "science", Count: 35000,
		}, {
			Name: "eludium", Count: 50,
		}, {
			Name: "antimatter", Count: 500,
		}},
	}, {
		Name: "Chronoengineers", UnlockedBy: "Tachyon Theory",
		Costs: []R{{
			Name: "science", Count: 500000,
		}, {
			Name: "time crystal", Count: 5,
		}, {
			Name: "eludium", Count: 100,
		}},
	}, {
		Name: "Telecommunication", UnlockedBy: "Electronics",
		Costs: []R{{
			Name: "science", Count: 150000,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "uranium", Count: 50,
		}},
	}, {
		Name: "Neural Network", UnlockedBy: "Artificial Intelligence",
		Costs: []R{{
			Name: "science", Count: 200000,
		}, {
			Name: "titanium", Count: 7500,
		}},
	}, {
		Name: "Robotic Assistance", UnlockedBy: "Robotics",
		Costs: []R{{
			Name: "science", Count: 100000,
		}, {
			Name: "steel", Count: 10000,
		}, {
			Name: "gear", Count: 250,
		}},
	}, {
		Name: "Factory Robotics", UnlockedBy: "Robotics",
		Costs: []R{{
			Name: "science", Count: 75000,
		}, {
			Name: "gear", Count: 250,
		}, {
			Name: "titanium", Count: 1250,
		}},
	}, {
		Name: "Void Aspiration", UnlockedBy: "Void Energy",
		Costs: []R{{
			Name: "time crystal", Count: 15,
		}, {
			Name: "antimatter", Count: 2000,
		}},
	}, {
		Name: "Distortion", UnlockedBy: "Paradox Theory",
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
			Name: "time crystal", Count: 128,
		}, {
			Name: "blackcoin", Count: 64,
		}, {
			Name: "void", Count: 32,
		}, {
			Name: "temporal flux", Count: 4096,
		}},
	}, {
		Name: "Solar Revolution", UnlockedBy: "Philosophy",
		Costs: []R{{
			Name: "faith", Count: 750,
		}, {
			Name: "gold", Count: 500,
		}},
	}, {
		Name: "Apocrypha", UnlockedBy: "Philosophy",
		Costs: []R{{
			Name: "faith", Count: 5000,
		}, {
			Name: "gold", Count: 5000,
		}},
	}, {
		Name: "Transcendence", UnlockedBy: "Philosophy",
		Costs: []R{{
			Name: "faith", Count: 7500,
		}, {
			Name: "gold", Count: 7500,
		}},
	}, {
		Name: "Orbital Launch", UnlockedBy: "Rocketry",
		Costs: []R{{
			Name: "starchart", Count: 250,
		}, {
			Name: "catpower", Count: 5000,
		}, {
			Name: "science", Count: 100000,
		}, {
			Name: "oil", Count: 15000, Bonus: []R{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Moon Mission", UnlockedBy: "Orbital Launch",
		Costs: []R{{
			Name: "starchart", Count: 500,
		}, {
			Name: "titanium", Count: 5000,
		}, {
			Name: "science", Count: 125000,
		}, {
			Name: "oil", Count: 45000, Bonus: []R{{Name: "SpaceElevatorOilBonus"}},
		}},
	}, {
		Name: "Dune Mission", UnlockedBy: "Moon Mission",
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
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
		Costs: []R{{
			Name: "starchart", Count: 500000,
		}, {
			Name: "science", Count: 1250000,
		}, {
			Name: "kerosene", Count: 75000,
		}, {
			Name: "thorium", Count: 75000,
		}},
	}, {
		Name: "Enlightenment", UnlockedBy: "Metaphysics", ResetResource: "Enlightenment",
		Costs: []R{{Name: "paragon", Count: 5}},
	}, {
		Name: "Codex Vox", UnlockedBy: "Enlightenment", ResetResource: "Codex Vox",
		Costs: []R{{Name: "paragon", Count: 25}},
	}, {
		Name: "Codex Logos", UnlockedBy: "Codex Vox", ResetResource: "Codex Logos",
		Costs: []R{{Name: "paragon", Count: 50}},
	}, {
		Name: "Codex Agrum", UnlockedBy: "Codex Logos", ResetResource: "Codex Agrum",
		Costs: []R{{Name: "paragon", Count: 75}},
	}, {
		Name: "Megalomania", UnlockedBy: "Enlightenment", ResetResource: "Megalomania",
		Costs: []R{{Name: "paragon", Count: 10}},
	}, {
		Name: "Black Codex", UnlockedBy: "Megalomania", ResetResource: "Black Codex",
		Costs: []R{{Name: "paragon", Count: 25}},
	}, {
		Name: "Codex Leviathanus", UnlockedBy: "Codex Logos", ResetResource: "Codex Leviathanus",
		Costs: []R{{Name: "paragon", Count: 75}},
	}, {
		Name: "Golden Ratio", UnlockedBy: "Enlightenment", ResetResource: "Golden Ratio",
		Costs: []R{{Name: "paragon", Count: 50}},
	}, {
		Name: "Divine Proportion", UnlockedBy: "Golden Ratio", ResetResource: "Divine Proportion",
		Costs: []R{{Name: "paragon", Count: 100}},
	}, {
		Name: "Vitruvian Feline", UnlockedBy: "Divine Proportion", ResetResource: "Vitruvian Feline",
		Costs: []R{{Name: "paragon", Count: 250}},
	}, {
		Name: "Renaissance", UnlockedBy: "Vitruvian Feline", ResetResource: "Renaissance",
		Costs: []R{{Name: "paragon", Count: 750}},
	}, {
		Name: "Diplomacy", UnlockedBy: "Metaphysics", ResetResource: "Diplomacy",
		Costs: []R{{Name: "paragon", Count: 5}},
	}, {
		Name: "Zebra Diplomacy", UnlockedBy: "Diplomacy", ResetResource: "Zebra Diplomacy",
		Costs: []R{{Name: "paragon", Count: 35}},
	}, {
		Name: "Zebra Covenant", UnlockedBy: "Zebra Diplomacy", ResetResource: "Zebra Covenant",
		Costs: []R{{Name: "paragon", Count: 75}},
	}, {
		Name: "Navigation Diplomacy", UnlockedBy: "Metaphysics", ResetResource: "Navigation Diplomacy",
		Costs: []R{{Name: "paragon", Count: 300}},
	}, {
		Name: "Chronomancy", UnlockedBy: "Metaphysics", ResetResource: "Chronomancy",
		Costs: []R{{Name: "paragon", Count: 25}},
	}, {
		Name: "Anachronomancy", UnlockedBy: "Chronomancy", ResetResource: "Anachronomancy",
		Costs: []R{{Name: "paragon", Count: 125}},
	}, {
		Name: "Astromancy", UnlockedBy: "Chronomancy", ResetResource: "Astromancy",
		Costs: []R{{Name: "paragon", Count: 50}},
	}, {
		Name: "Unicornmancy", UnlockedBy: "Metaphysics", ResetResource: "Unicornmancy",
		Costs: []R{{Name: "paragon", Count: 125}},
	}, {
		Name: "Carnivals", UnlockedBy: "Metaphysics", ResetResource: "Carnivals",
		Costs: []R{{Name: "paragon", Count: 25}},
	}, {
		Name: "Numerology", UnlockedBy: "Carnivals", ResetResource: "Numerology",
		Costs: []R{{Name: "paragon", Count: 50}},
	}, {
		Name: "Order of the Void", UnlockedBy: "Numerology", ResetResource: "Order of the Void",
		Costs: []R{{Name: "paragon", Count: 75}},
	}, {
		Name: "Venus of Willenfluff", UnlockedBy: "Numerology", ResetResource: "Venus of Willenfluff",
		Costs: []R{{Name: "paragon", Count: 150}},
	}, {
		Name: "Pawgan Rituals", UnlockedBy: "Venus of Willenfluff", ResetResource: "Pawgan Rituals",
		Costs: []R{{Name: "paragon", Count: 400}},
	}, {
		Name: "Numeromancy", UnlockedBy: "Numerology", ResetResource: "Numeromancy",
		Costs: []R{{Name: "paragon", Count: 250}},
	}, {
		Name: "Malkuth", UnlockedBy: "Numeromancy", ResetResource: "Malkuth",
		Costs: []R{{Name: "paragon", Count: 500}},
	}, {
		Name: "Yesod", UnlockedBy: "Malkuth", ResetResource: "Yesod",
		Costs: []R{{Name: "paragon", Count: 750}},
	}, {
		Name: "Hod", UnlockedBy: "Yesod", ResetResource: "Hod",
		Costs: []R{{Name: "paragon", Count: 1250}},
	}, {
		Name: "Netzach", UnlockedBy: "Hod", ResetResource: "Netzach",
		Costs: []R{{Name: "paragon", Count: 1750}},
	}, {
		Name: "Tiferet", UnlockedBy: "Netzach", ResetResource: "Tiferet",
		Costs: []R{{Name: "paragon", Count: 2500}},
	}, {
		Name: "Gevurah", UnlockedBy: "Tiferet", ResetResource: "Gevurah",
		Costs: []R{{Name: "paragon", Count: 5000}},
	}, {
		Name: "Chesed", UnlockedBy: "Gevurah", ResetResource: "Chesed",
		Costs: []R{{Name: "paragon", Count: 7500}},
	}, {
		Name: "Binah", UnlockedBy: "Chesed", ResetResource: "Binah",
		Costs: []R{{Name: "paragon", Count: 15000}},
	}, {
		Name: "Chokhmah", UnlockedBy: "Binah", ResetResource: "Chokhmah",
		Costs: []R{{Name: "paragon", Count: 30000}},
	}, {
		Name: "Keter", UnlockedBy: "Chokhmah", ResetResource: "Keter",
		Costs: []R{{Name: "paragon", Count: 60000}},
	}, {
		Name: "Adjustment Bureau", UnlockedBy: "Metaphysics", ResetResource: "Adjustment Bureau",
		Costs: []R{{Name: "paragon", Count: 5}},
	}, {
		Name: "ASCOH", UnlockedBy: "Adjustment Bureau", ResetResource: "ASCOH",
		Costs: []R{{Name: "paragon", Count: 5}},
	}})

	addCrafts(g, []data.Action{{
		Name: "beam", UnlockedBy: "Construction",
		Costs: []R{{Name: "wood", Count: 175}},
		Producers: []R{{
			Name: "Active Steamworks", Factor: 1. / (2 * 100 * 4 * 175), ProductionOnGone: true,
			Bonus: []R{{
				Name: "Workshop Automation",
				Bonus: []R{{
					Name: "wood cap",
					Bonus: []R{{
						Name: "Advanced Automation",
					}, {
						Name: "CraftRatio",
					}},
				}},
				BonusStartsFromZero: true,
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "slab", UnlockedBy: "Construction",
		Costs: []R{{Name: "mineral", Count: 250}},
		Producers: []R{{
			Name: "Active Steamworks", Factor: 1. / (2 * 100 * 4 * 250), ProductionOnGone: true,
			Bonus: []R{{
				Name: "Workshop Automation",
				Bonus: []R{{
					Name: "mineral cap",
					Bonus: []R{{
						Name: "Advanced Automation",
					}, {
						Name: "CraftRatio",
					}},
				}},
				BonusStartsFromZero: true,
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "concrete", UnlockedBy: "Mechanization",
		Costs: []R{{
			Name: "slab", Count: 2500,
		}, {
			Name: "steel", Count: 25,
		}},
	}, {
		Name: "plate", UnlockedBy: "Construction",
		Costs: []R{{Name: "iron", Count: 125}},
		Producers: []R{{
			Name: "Active Steamworks", Factor: 1. / (2 * 100 * 4 * 125), ProductionOnGone: true,
			Bonus: []R{{
				Name: "Workshop Automation",
				Bonus: []R{{
					Name: "Pneumatic Press",
					Bonus: []R{{
						Name: "iron cap",
						Bonus: []R{{
							Name: "Advanced Automation",
						}, {
							Name: "CraftRatio",
						}},
					}},
					BonusStartsFromZero: true,
				}},
				BonusStartsFromZero: true,
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "steel", UnlockedBy: "Steel",
		Costs: []R{{
			Name: "iron", Count: 100,
		}, {
			Name: "coal", Count: 100,
		}},
		Producers: []R{{
			Name: "Active Calciner", Factor: 0.15 * 5 * 0.10 * 0.01,
			Bonus: []R{{
				Name: "Steel Plants",
				Bonus: []R{{
					Name: "Oxidation", Factor: 0.95,
				}, {
					Name: "CraftRatio", Factor: 0.25,
					Bonus:               []R{{Name: "Automated Plants"}},
					BonusStartsFromZero: true,
				}, {
					Name: "Reactor", Factor: 0.02,
					Bonus:               []R{{Name: "Nuclear Plants"}},
					BonusStartsFromZero: true,
				}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "gear", UnlockedBy: "Construction",
		Costs: []R{{Name: "steel", Count: 15}},
	}, {
		Name: "alloy", UnlockedBy: "Chemistry",
		Costs: []R{{
			Name: "steel", Count: 75,
		}, {
			Name: "titanium", Count: 10,
		}},
	}, {
		Name: "eludium", UnlockedBy: "Advanced Exogeology",
		Costs: []R{{
			Name: "alloy", Count: 2500,
		}, {
			Name: "unobtainium", Count: 1000,
		}},
	}, {
		Name: "scaffold", UnlockedBy: "Construction",
		Costs: []R{{Name: "beam", Count: 50}},
	}, {
		Name: "ship", UnlockedBy: "Navigation",
		Costs: []R{{
			Name: "scaffold", Count: 100,
		}, {
			Name: "plate", Count: 150,
		}, {
			Name: "starchart", Count: 25,
			Bonus: []R{{
				Name: "Satellite", Factor: 0.0125,
				Bonus:               []R{{Name: "Satellite Navigation"}},
				BonusStartsFromZero: true,
			}},
		}},
	}, {
		Name: "tanker", UnlockedBy: "Robotics",
		Costs: []R{{
			Name: "ship", Count: 200,
		}, {
			Name: "alloy", Count: 1250,
		}, {
			Name: "blueprint", Count: 5,
		}},
	}, {
		Name: "kerosene", UnlockedBy: "Oil Processing",
		Costs: []R{{Name: "oil", Count: 7500}},
		Bonus: []R{{
			Name: "Factory", Factor: 0.05,
			Bonus:               []R{{Name: "Factory Processing"}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "parchment", UnlockedBy: "Writing",
		Costs: []R{{Name: "fur", Count: 175}},
	}, {
		Name: "manuscript", UnlockedBy: "Construction",
		Producers: []R{{
			Name: "Steamworks", Factor: 0.0005 * 5,
			Bonus: []R{{
				Name: "Printing Press",
				Bonus: []R{{
					Name: "Offset Press", Factor: 4 - 1,
					Bonus: []R{{Name: "Photolithography", Factor: 4 - 1}},
				}},
			}},
			BonusStartsFromZero: true,
		}},
		Costs: []R{{
			Name: "culture", Count: 400,
		}, {
			Name: "parchment", Count: 25,
		}},
		Bonus: []R{{
			Name: "Codex Vox", Factor: 0.25,
		}, {
			Name: "Codex Logos", Factor: 0.25,
		}, {
			Name: "Codex Agrum", Factor: 0.25,
		}},
	}, {
		Name: "compendium", UnlockedBy: "Philosophy",
		Costs: []R{{
			Name: "manuscript", Count: 50,
		}, {
			Name: "science", Count: 10000,
		}},
		Bonus: []R{{
			Name: "Codex Logos", Factor: 0.25,
		}, {
			Name: "Codex Agrum", Factor: 0.25,
		}},
	}, {
		Name: "blueprint", UnlockedBy: "Physics",
		Costs: []R{{
			Name: "compendium", Count: 25,
		}, {
			Name: "science", Count: 25000,
		}},
		Bonus: []R{{
			Name: "CAD System", Factor: 0.01,
			Bonus: []R{{
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
		Costs: []R{{Name: "uranium", Count: 250}},
		Producers: []R{{
			Name: "Active Reactor", Factor: -0.05 * 5,
			Bonus: []R{{
				Name:  "Thorium Reactors",
				Bonus: []R{{Name: "Enriched Thorium", Factor: -0.25}},
			}},
			BonusStartsFromZero: true,
		}},
	}, {
		Name: "megalith", UnlockedBy: "Construction",
		Costs: []R{{
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

func resourceWithModulus(resource R, names []string) []R {
	res := []R{}
	resource.ProductionModulus = len(names)
	for i, name := range names {
		resource.Name = name
		resource.ProductionModulusEquals = i
		res = append(res, resource)
	}
	return res
}

func resourceWithName(resource R, names []string) []R {
	res := []R{}
	for _, name := range names {
		resource.Name = name
		res = append(res, resource)
	}
	return res
}

func addBonus(g *game.Game, resources []R) {
	for _, resource := range resources {
		resource.Type = "Resource"
		resource.IsHidden = true
		resource.StartCountFromZero = true
		g.AddResource(resource)
	}
}

func addResets(g *game.Game, names []string) {
	for _, name := range names {
		g.AddResource(R{
			Name: "reset " + name, Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true, ResetResource: "reset " + name,
			Producers: []R{{
				Name:                name,
				Bonus:               []R{{Name: "Chronosphere", Factor: 0.015}},
				BonusStartsFromZero: true,
			}},
		})
	}
}

func addCrafts(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		name := action.Name
		action.Name = "@" + name
		action.Type = "Craft"
		action.Adds = []R{{
			Name: name, Count: 1,
			Bonus: join([]R{{Name: "CraftRatio"}}, action.Bonus),
		}}
		action.IsHidden = true
		g.AddAction(action)
		g.AddResource(R{
			Name: name, Type: "Resource", Cap: -1, ProducerAction: action.Name,
			Producers: action.Producers,
		})
		g.AddResource(R{
			Name: "reset " + name, Type: "Resource", IsHidden: true, Cap: -1, StartCountFromZero: true, ResetResource: "reset " + name,
			Producers: []R{{
				Name:                name,
				Bonus:               []R{{Name: "Chronosphere", Factor: 0.015}},
				BonusStartsFromZero: true,
			}},
		})
	}
}

func addBuildings(g *game.Game, actions []data.Action) {
	for _, action := range actions {
		name := action.Name
		activeName := ""
		if strings.HasPrefix(action.Name, "Active ") {
			activeName = action.Name
			name = strings.TrimPrefix(action.Name, "Active ")
		}
		action.Name = name
		action.Type = "Building"
		action.Adds = append([]R{{
			Name: action.Name, Count: 1,
		}}, action.Adds...)
		g.AddAction(action)
		g.AddResource(R{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: -1,
		})

		if activeName == "" {
			continue
		}

		g.AddAction(data.Action{
			Name:       activeName,
			Type:       action.Type,
			Costs:      []R{{Name: "Idle " + name, Count: 1}},
			Adds:       []R{{Name: activeName, Count: 1}},
			UnlockedBy: name,
		})
		g.AddResource(R{
			Name: activeName, Type: action.Type, IsHidden: true, Cap: -1,
		})

		g.AddResource(R{
			Name: "Idle " + name, Type: "Building", IsHidden: true, Cap: -1, StartCountFromZero: true,
			Producers: []R{{
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
		action.Costs = []R{{Name: "kitten", Count: 1, Cap: 1}}
		action.Adds = []R{{Name: action.Name, Count: 1}}
		g.AddAction(action)
		g.AddResource(R{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: -1,
			OnGone: []R{{
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
		action.Adds = []R{{Name: action.Name, Count: 1}}
		action.LockedBy = action.Name
		g.AddAction(action)
		g.AddResource(R{
			Name: action.Name, Type: action.Type, IsHidden: true, Cap: 1,
		})
	}
}

func getCEBR(base float64) R {
	return R{
		Factor: base,
		Bonus:  []R{{Name: "PriceRatioBonus"}},
	}
}
