package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func updateRecipe(data *blizzard_api.ApiResponse) {
	var recipe datasets.Recipe
	data.Parse(&recipe)

	insertOnceUpdate(&recipe, "name", "crafted_item_id")

	for _, reagent := range recipe.Reagents {
		reagent.RecipeID = recipe.ID
		reagent.ReagentID = reagent.Reagent.ID
		insertOnceExpr(&reagent, "(recipe_id,reagent_id) DO UPDATE", "quantity")
	}

	if recipe.Media != nil {
		recipe.Media.RecipeID = recipe.ID
		insertOnceUpdate(recipe.Media, "recipe_id")
	}
}

func updateCategory(category datasets.ProfessionCategory) {
	category.TierID = category.Tier.ID
	insertOnceUpdate(&category, "name", "tier_id")

	for _, recipe := range category.Recipes {
		recipeResponse :=  connections.WowClient.Recipe(recipe.ID, nil)
		updateRecipe(recipeResponse)

		insertOnceExpr(&datasets.ProfessionTierRecipes{
			ProfessionTierID: category.TierID,
			CategoryID:       category.ID,
			RecipeID:         recipe.ID,
		}, "(profession_tier_id,recipe_id) DO UPDATE", "category_id")
	}
}

func updateSkillTier(data *blizzard_api.ApiResponse, categories bool) {
	var skillTier datasets.ProfessionTier
	data.Parse(&skillTier)

	insertOnceUpdate(&skillTier, "name", "minimum_skill_level", "maximum_skill_level")

	if categories {
		for _, category := range skillTier.Categories {
			category.Tier = &skillTier
			updateCategory(category)
		}
	}
}

func UpdateProfession(data *blizzard_api.ApiResponse) {
	var profession datasets.Profession
	data.Parse(&profession)
	if profession.Type != nil {
		profession.TypeID = profession.Type.ID
		insertOnce(profession.Type)
	}
	insertOnceUpdate(&profession, "name", "description", "type_id")

	if profession.SkillTiers != nil {
		for _, tier := range profession.SkillTiers {
			tierResponse := connections.WowClient.ProfessionSkillTier(profession.ID, tier.ID, nil)
			updateSkillTier(tierResponse, false)
		}
	}

	if profession.Media != nil {
		profession.Media.ProfessionID = profession.ID
		insertOnceUpdate(profession.Media, "profession_id")
	}
}

func UpdateProfessionTiers(data *blizzard_api.ApiResponse) {
	var profession datasets.Profession
	data.Parse(&profession)

	if profession.SkillTiers != nil {
		for _, tier := range profession.SkillTiers {
			tierResponse := connections.WowClient.ProfessionSkillTier(profession.ID, tier.ID, nil)
			updateSkillTier(tierResponse, true)
		}
	}
}

func UpdateProfessionMedia(data *blizzard_api.ApiResponse, id int) {
	var professionMedia datasets.ProfessionMedia
	data.Parse(&professionMedia)

	for _, asset := range professionMedia.Assets {
		asset.ProfessionMediaID = professionMedia.ID
		insertOnceExpr(&asset, "(profession_media_id,key) DO UPDATE", "value")
	}
}

func UpdateRecipeMedia(data *blizzard_api.ApiResponse, id int) {
	var recipeMedia datasets.RecipeMedia
	data.Parse(&recipeMedia)

	for _, asset := range recipeMedia.Assets {
		asset.RecipeMediaID = recipeMedia.ID
		insertOnceExpr(&asset, "(recipe_media_id,key) DO UPDATE", "value")
	}
}