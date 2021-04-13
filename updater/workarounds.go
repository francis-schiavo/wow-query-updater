package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func InsertMissingItemClasses() {
	// Missing item class
	insertOnceUpdate(&datasets.ItemClass{
		ID: 10,
		Name: datasets.LocalizedField{
			EnUS: "Money(OBSOLETE)",
			EsMX: "Dinero (OBSOLETO)",
			PtBR: "Money(OBSOLETE)",
			DeDE: "Gold(ÜBERFLÜSSIG)",
			EnGB: "Money(OBSOLETE)",
			EsES: "Dinero (OBSOLETO)",
			FrFR: "Argent(OBSOLETE)",
			ItIT: "Denaro",
			RuRU: "Деньги (НЕ ИСП.)",
			KoKR: "돈",
			ZhTW: "金錢(廢棄)",
			ZhCN: "钱币（废弃）",
		},
	}, "name")

	// Missing item subclass
	insertOnceExpr(&datasets.ItemSubclass{
		ID:      0,
		ClassID: 10,
		DisplayName: datasets.LocalizedField{
			EnUS: "Money(OBSOLETE)",
			EsMX: "Dinero (OBSOLETO)",
			PtBR: "Money(OBSOLETE)",
			DeDE: "Gold(ÜBERFLÜSSIG)",
			EnGB: "Money(OBSOLETE)",
			EsES: "Dinero (OBSOLETO)",
			FrFR: "Argent(OBSOLETE)",
			ItIT: "Denaro",
			RuRU: "Деньги (НЕ ИСП.)",
			KoKR: "돈",
			ZhTW: "金錢(廢棄)",
			ZhCN: "钱币（废弃）",
		},
	}, "(id,class_id) DO UPDATE", "display_name")

	insertOnceExpr(&datasets.ItemSubclass{
		ID:      0,
		ClassID: 7,
		DisplayName: datasets.LocalizedField{
			EnUS: "Trade Goods (OBSOLETE)",
			EsMX: "Trade Goods (OBSOLETE)",
			PtBR: "Trade Goods (OBSOLETE)",
			DeDE: "Trade Goods (OBSOLETE)",
			EnGB: "Trade Goods (OBSOLETE)",
			EsES: "Trade Goods (OBSOLETE)",
			FrFR: "Trade Goods (OBSOLETE)",
			ItIT: "Trade Goods (OBSOLETE)",
			RuRU: "Trade Goods (OBSOLETE)",
			KoKR: "직업용품 (미사용)",
			ZhTW: "商品(廢棄)",
			ZhCN: "商品",
		},
	}, "(id,class_id) DO UPDATE", "display_name")

	insertOnceExpr(&datasets.ItemSubclass{
		ID:      12,
		ClassID: 15,
		DisplayName: datasets.LocalizedField{
			EnUS: "Miscellaneous",
			EsMX: "Miscelánea",
			PtBR: "Diversos",
			DeDE: "Verschiedenes",
			EnGB: "Miscellaneous",
			EsES: "Miscelánea",
			FrFR: "Divers",
			ItIT: "Varie",
			RuRU: "Разное",
			KoKR: "기타",
			ZhTW: "雜項",
			ZhCN: "杂项",
		},
	}, "(id,class_id) DO UPDATE", "display_name")
}

func InsertMissingReputationTiers() {
	// Missing reputation referenced by items
	insertOnceExpr(&datasets.ReputationTier{
		ReputationTierID: 9999999,
		Identifiable: datasets.Identifiable{
			ID: 9999999,
		},
		Name: datasets.LocalizedField{},
	}, "(reputation_tier_id,id) DO NOTHING")

	insertOnce(&datasets.ReputationFaction{
		Identifiable: datasets.Identifiable{
			ID: 2463,
		},
		ReputationTierID: 9999999,
		Name: datasets.LocalizedField{
			EnUS: "Marasmius",
			EsMX: "Marasmius",
			PtBR: "Marasmius",
			DeDE: "Marasmius",
			EnGB: "Marasmius",
			EsES: "Marasmius",
			FrFR: "Marasmius",
			ItIT: "Marasmio",
			RuRU: "Чесночник",
			KoKR: "마라스미우스",
			ZhTW: "瑪拉茲莫斯",
			ZhCN: "玛拉斯缪斯",
		},
		IsHeader: false,
	})

	insertOnce(&datasets.ReputationFaction{
		Identifiable: datasets.Identifiable{
			ID: 2464,
		},
		ReputationTierID: 9999999,
		Name: datasets.LocalizedField{
			EnUS: "Court of Night",
			EsMX: "Corte de la Noche",
			PtBR: "Corte da Noite",
			DeDE: "Hof der Nacht",
			EnGB: "Court of Night",
			EsES: "Corte de la Noche",
			FrFR: "Cour de la Nuit",
			ItIT: "Corte della Notte",
			RuRU: "Двор Ночи",
			KoKR: "밤의 궁정",
			ZhTW: "暗夜之廷",
			ZhCN: "魅夜王庭",
		},
		IsHeader: false,
	})
}

func InsertMissingSpells() {
	// Missing spell referenced by items
	insertOnce(&datasets.Spell{
		Identifiable: datasets.Identifiable{
			ID: 33388,
		},
		Name: datasets.LocalizedField{
			EnUS: "Apprentice Riding",
			EsMX: "Aprendiz jinete",
			PtBR: "Aprendiz de Montaria",
			DeDE: "Unerfahrenes Reiten",
			EnGB: "Apprentice Riding",
			EsES: "Aprendiz jinete",
			FrFR: "Apprenti cavalier",
			ItIT: "Apprendista in Equitazione",
			RuRU: "Верховая езда (ученик)",
			KoKR: "초급 타기",
			ZhTW: "初級騎術",
			ZhCN: "初级骑术",
		},
	})

	insertOnce(&datasets.Spell{
		Identifiable: datasets.Identifiable{
			ID: 119467,
		},
		Name: datasets.LocalizedField{
			EnUS: "Battle Pet Training",
			EsMX: "Entrenamiento de mascotas de duelo",
			PtBR: "Treinamento de Mascote de Batalha",
			DeDE: "Kampfhaustiertraining",
			EnGB: "Battle Pet Training",
			EsES: "Entrenamiento de mascotas de duelo",
			FrFR: "Entraînement de mascotte de combat",
			ItIT: "Addestramento Mascotte",
			RuRU: "Обучение боевых питомцев",
			KoKR: "전투 애완동물 조련",
			ZhTW: "戰寵訓練師",
			ZhCN: "战斗宠物训练",
		},
	})

	insertOnce(&datasets.Spell{
		Identifiable: datasets.Identifiable{
			ID: 156593,
		},
		Name: datasets.LocalizedField{
			EnUS: "Alchemical Catalyst - Lotus",
			EsMX: "Catalizador alquímico - Loto",
			PtBR: "Catalisador Alquímico - Lótus",
			DeDE: "Alchemischer Katalysator - Lotus",
			EnGB: "Alchemical Catalyst - Lotus",
			EsES: "Catalizador alquímico - Loto",
			FrFR: "Catalyseur alchimique (lotus)",
			ItIT: "Catalizzatore Alchemico - Loto",
			RuRU: "Алхимический катализатор – лотос",
			KoKR: "연금술 촉매 - 카멜레온 연꽃",
			ZhTW: "鍊金催化劑 - 變色龍蓮花",
			ZhCN: "炼金催化剂 - 拟态莲",
		},
	})

	insertOnce(&datasets.Spell{
		Identifiable: datasets.Identifiable{
			ID: 163024,
		},
		Name: datasets.LocalizedField{
			EnUS: "Warforged Nightmare",
			EsMX: "Tormento de guerra",
			PtBR: "Pesadelo Forjado para a Guerra",
			DeDE: "Kriegsgeschmiedeter Nachtmahr",
			EnGB: "Warforged Nightmare",
			EsES: "Tormento de guerra",
			FrFR: "Méca-cavale de cauchemar",
			ItIT: "Destriero dell'Incubo Guerraforgiato",
			RuRU: "Боевой конь кошмаров",
			KoKR: "전쟁벼림 악몽마",
			ZhTW: "戰鑄夢魘戰馬",
			ZhCN: "战火梦魇兽",
		},
	})
}

func InsertMissingStats()  {
	// Missing Item Stats
	insertOnce(&datasets.Stat{
		ID:   "ALL_RESISTANCE",
	})
}

func UpdateTechTalentUsingTree(data *blizzard_api.ApiResponse) {
	var techTalentTree datasets.TechTalentTree
	data.Parse(&techTalentTree)

	for _, talent := range techTalentTree.Talents {
		var techTalent datasets.TechTalent
		talentData := connections.WowClient.TechTalent(talent.ID, &blizzard_api.RequestOptions{})
		talentData.Parse(&techTalent)

		techTalent.TechTalentTreeID = techTalentTree.ID
		if techTalent.PrerequisiteTalent != nil {
			techTalent.PrerequisiteTalentID = techTalent.PrerequisiteTalent.ID
		}

		if techTalent.SpellTooltip != nil {
			updateSpellTooltip(techTalent.SpellTooltip)
			techTalent.SpellTooltipID = techTalent.SpellTooltip.ID
		}

		insertOnceUpdate(&techTalent, "name", "description", "spell_tooltip_id", "tier", "display_order", "prerequisite_talent_id", "tech_talent_tree_id")
	}
}