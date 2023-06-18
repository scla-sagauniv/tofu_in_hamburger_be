package logic

import (
	"log"
	"tofu_in_hamburger_be/db"
)

func DeleteAllRecipe() error {
	log.Println("Start DeleteAllRecipe")

	query := "delete from recipes"
	log.Println("--- delete all rows query ---")
	log.Println(query)
	log.Println("-------------------------")

	if _, err := db.Db.Exec(query); err != nil {
		log.Fatal("Exec error: ", err)
		return err
	}

	return nil
}
