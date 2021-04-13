package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/datasets"
)

func updateRenownReward(reward *datasets.RenownReward) {
	insertOnceUpdate(reward, "name")
}

func updateCovenantAbility(covenant *datasets.Covenant, ability *datasets.CovenantAbility) {
	ability.SpellTooltip.ID = ability.SpellTooltip.Spell.ID
	updateSpellTooltip(ability.SpellTooltip)

	if ability.PlayableClass != nil {
		ability.PlayableClassID = ability.PlayableClass.ID
	}

	ability.CovenantID = covenant.ID
	ability.SpellTooltipID = ability.SpellTooltip.ID
	insertOnceUpdate(ability, "covenant_id", "playable_class_id", "spell_tooltip_id")
}

func UpdateCovenant(data *blizzard_api.ApiResponse) {
	var covenant datasets.Covenant
	data.Parse(&covenant)

	covenant.SignatureAbilityID = covenant.SignatureAbility.ID
	insertOnceUpdate(&covenant, "name", "description", "signature_ability_id")

	covenant.Media.CovenantID = covenant.ID
	insertOnce(covenant.Media)

	for _, classAbility := range covenant.ClassAbilities {
		updateCovenantAbility(&covenant, classAbility)
	}

	for _, covenantReward := range covenant.RenownRewards {
		updateRenownReward(covenantReward.Reward)

		covenantReward.CovenantID = covenant.ID
		covenantReward.RewardID = covenantReward.Reward.ID
		insertOnceExpr(covenantReward, "(covenant_id,reward_id) DO UPDATE", "level")
	}
}

func UpdateSoulbind(data *blizzard_api.ApiResponse) {
	var soulbind datasets.Soulbind
	data.Parse(&soulbind)

	if soulbind.Follower != nil {
		insertOnce(soulbind.Follower)
		soulbind.FollowerID = soulbind.Follower.ID
	}

	soulbind.CovenantID = soulbind.Covenant.ID
	soulbind.CreatureID = soulbind.Creature.ID

	soulbind.TechTalentTreeID = soulbind.TechTalentTree.ID

	insertOnceUpdate(&soulbind, "name", "covenant_id", "creature_id", "follower_id", "tech_talent_tree_id")
}

func UpdateConduit(data *blizzard_api.ApiResponse) {
	var conduit datasets.Conduit
	data.Parse(&conduit)

	insertOnceUpdate(conduit.SocketType, "name")
	conduit.ItemID = conduit.Item.ID
	conduit.SocketTypeID = conduit.SocketType.ID
	insertOnceUpdate(&conduit, "name", "item_id", "socket_type_id")

	for _, rank := range conduit.Ranks {
		if rank.SpellTooltip != nil {
			updateSpellTooltip(rank.SpellTooltip)
			rank.SpellTooltipID = rank.SpellTooltip.ID
		}
		rank.ConduitID = conduit.ID
		insertOnceUpdate(rank, "tier", "spell_tooltip_id", "conduit_id")
	}
}

func UpdateCovenantMedia(data *blizzard_api.ApiResponse, id int) {
	var media datasets.CovenantMedia
	data.Parse(&media)

	for _, asset := range media.Assets {
		asset.CovenantMediaID = id
		insertOnceExpr(&asset, "(covenant_media_id,key) DO UPDATE", "value")
	}
}