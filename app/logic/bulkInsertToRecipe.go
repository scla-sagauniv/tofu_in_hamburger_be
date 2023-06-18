package logic

import (
	"fmt"
	"log"
	"tofu_in_hamburger_be/db"
	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func timestampToString(pb *timestamppb.Timestamp) (string, error) {
	tm, err := ptypes.Timestamp(pb)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tm.Format("2006-01-02 15:04:05"), nil
}

func boolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

func BulkInsertToRecipe(recipes []*v1.RecipeOnDb) error {
	log.Println("Start BulkInsertToRecipe")
	insert := "INSERT INTO recipes (id, title, recipe_url, image_url, pickup, nickname, materials, material_ids, publishday, ranking, recipe_indication_id, recipe_cost_id) VALUES "
	// insert := "INSERT INTO recipes (id, title, recipe_url, image_url, pickup, nickname, materials, material_ids, publishday, ranking, recipe_indication_id, recipe_cost_id, created_at, updated_at) VALUES "

	var vals []any
	for _, recipe := range recipes {
		insert += fmt.Sprintf(`(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?), `)
		publishday, err := timestampToString(recipe.GetPublishday())
		if err != nil {
			log.Fatal(err)
			return err
		}
		// createdAt, err := timestampToString(recipe.GetCreatedAt())
		// if err != nil {
		// 	log.Fatal(err)
		// 	return err
		// }
		// updatedAt, err := timestampToString(recipe.GetUpdatedAt())
		// if err != nil {
		// 	log.Fatal(err)
		// 	return err
		// }
		vals = append(vals, recipe.GetId(), recipe.GetTitle(), recipe.GetRecipeUrl(), recipe.GetImageUrl(), recipe.GetPickup(), recipe.GetNickname(), recipe.GetMaterials(), recipe.GetMaterialIds(), publishday, recipe.GetRanking(), recipe.GetRecipeIndicationId(), recipe.GetRecipeCostId())
		// insert += fmt.Sprintf(`(%d, '%s', '%s', '%s', %d, '%s', '%s', '%s', '%s', %d, %d, %d, '%s', '%s'), `, recipe.GetId(), recipe.GetTitle(), recipe.GetRecipeUrl(), recipe.GetImageUrl(), boolToInt(recipe.GetPickup()), recipe.GetNickname(), recipe.GetMaterials(), recipe.GetMaterialIds(), publishday, recipe.GetRanking(), recipe.GetRecipeIndicationId(), recipe.GetRecipeCostId(), createdAt, updatedAt)
	}
	insert = insert[:len(insert)-2]

	fmt.Println(insert)

	log.Println("--- bulk insert query ---")
	log.Println(insert)
	log.Println("-------------------------")
	stmt, err := db.Db.Prepare(insert)
	if err != nil {
		log.Fatal("Prepare error: ", err)
		return err
	}

	if _, err := stmt.Exec(vals...); err != nil {
		log.Fatal("Exec error: ", err)
		return err
	}

	return nil
}
