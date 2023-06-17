package logic

import (
	"log"
	"time"
	"tofu_in_hamburger_be/db"
	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func SelectAllIngredient() ([]*v1.IngredientOnDb, error) {
	log.Println("Start SelectAllIngredient")

	query := "select * from materials"
	log.Println("--- select all rows query ---")
	log.Println(query)
	log.Println("-------------------------")

	rows, err := db.Db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var ingredients []*v1.IngredientOnDb
	for rows.Next() {
		var ingredient v1.IngredientOnDb
		var createdAt time.Time
		var updatedAt time.Time
		err := rows.Scan(&ingredient.Id, &ingredient.Titile, &ingredient.Description, &ingredient.ImageUrl, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		ingredient.CreatedAt = timestamppb.New(createdAt)
		ingredient.UpdatedAt = timestamppb.New(updatedAt)
		ingredients = append(ingredients, &ingredient)
	}
	return ingredients, nil
}
