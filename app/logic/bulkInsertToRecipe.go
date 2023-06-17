package logic

import (
	"fmt"
	"log"
	"tofu_in_hamburger_be/db"
	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"

	"github.com/golang/protobuf/ptypes"
)

func BulkInsertToRecipe(recipes []*v1.RecipeOnDb) error {
	log.Println("Start BulkInsertToRecipe")
	insert := "INSERT INTO recipes(id, title, recipe_url, image_url, pickup, nickname, materials, material_ids, publishday, ranking, recipe_indication_id, recipe_cost_id) VALUES "

	var vals []any
	for _, recipe := range recipes {
		insert += fmt.Sprintf(`(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?), `)
		tm, err := ptypes.Timestamp(recipe.GetPublishday())
		if err != nil {
			log.Fatal(err)
			return err
		}
		mysqlDateTime := tm.Format("2006-01-02 15:04:05")
		vals = append(vals, recipe.GetId(), recipe.GetTitle(), recipe.GetRecipeUrl(), recipe.GetImageUrl(), recipe.GetPickup(), recipe.GetNickname(), recipe.GetMaterials(), recipe.GetMaterialIds(), mysqlDateTime, recipe.GetRanking(), recipe.GetRecipeIndicationId(), recipe.GetRecipeCostId())
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
