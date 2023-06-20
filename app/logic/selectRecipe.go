package logic

import (
	"log"
	"strings"
	"time"
	"tofu_in_hamburger_be/db"
	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertToRecipeIndicationId(recipeIndicationId int64) string {
	switch recipeIndicationId {
	case 0:
		return "指定なし"
	case 1:
		return "5分以内"
	case 2:
		return "約10分"
	case 3:
		return "約15分"
	case 4:
		return "約30分"
	case 5:
		return "約1時間"
	case 6:
		return "1時間以上"
	}
	return "指定なし"
}

func convertToRecipeCostId(recipeCostId int64) string {
	switch recipeCostId {
	case 0:
		return "指定なし"
	case 1:
		return "100円以下"
	case 2:
		return "300円前後"
	case 3:
		return "500円前後"
	case 4:
		return "1,000円前後"
	case 5:
		return "2,000円前後"
	case 6:
		return "3,000円前後"
	case 7:
		return "5,000円前後"
	case 8:
		return "10,000円以上"
	}
	return "指定なし"
}

func SelectRecipe(ingredients []*v1.Ingredient) ([]*v1.Recipe, error) {
	log.Println("Start SelectRecipe")

	query := "select * from recipes order by ranking"
	log.Println("--- select all recipe query ---")
	log.Println(query)
	log.Println("-------------------------")

	rows, err := db.Db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var recipes []*v1.RecipeOnDb
	for rows.Next() {
		var recipe v1.RecipeOnDb
		var publishday time.Time
		var createdAt time.Time
		var updatedAt time.Time
		err := rows.Scan(
			&recipe.Id,
			&recipe.Title,
			&recipe.RecipeUrl,
			&recipe.ImageUrl,
			&recipe.Pickup,
			&recipe.Nickname,
			&recipe.Materials,
			&recipe.MaterialIds,
			&publishday,
			&recipe.Ranking,
			&recipe.RecipeIndicationId,
			&recipe.RecipeCostId,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		recipe.Publishday = timestamppb.New(publishday)
		recipe.CreatedAt = timestamppb.New(createdAt)
		recipe.UpdatedAt = timestamppb.New(updatedAt)
		recipes = append(recipes, &recipe)
	}
	var resultRecipes []*v1.Recipe
	for _, recipe := range recipes {
		f := false
		for _, ingredient := range ingredients {
			if !strings.Contains(recipe.GetMaterials(), ingredient.GetTitle()) {
				f = true
				break
			}
		}
		if !f {
			r := v1.Recipe{
				Id:               recipe.GetId(),
				Title:            recipe.GetTitle(),
				RecipeUrl:        recipe.GetRecipeUrl(),
				ImageUrl:         recipe.GetImageUrl(),
				Pickup:           recipe.GetPickup(),
				Nickname:         recipe.GetNickname(),
				Materials:        recipe.GetMaterials(),
				MaterialIds:      []int64{0},
				Publishday:       recipe.GetPublishday(),
				Rank:             recipe.GetRanking(),
				RecipeIndication: convertToRecipeIndicationId(recipe.GetRecipeIndicationId()),
				RecipeCost:       convertToRecipeCostId(recipe.GetRecipeCostId()),
				CreatedAt:        recipe.GetCreatedAt(),
				UpdatedAt:        recipe.GetUpdatedAt(),
			}
			resultRecipes = append(resultRecipes, &r)
		}
		f = false
		if len(resultRecipes) > 5 {
			break
		}
	}
	return resultRecipes, nil
}
