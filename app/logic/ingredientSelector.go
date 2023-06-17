package logic

import (
	"log"
	"math/rand"
	"time"
	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"
)

var RecievedIngredients []*v1.Ingredient

func IngredientSelector() []*v1.Ingredient {
	/*
		RecievedIngredientsからランダムに 1 ~ 4 個選択して、返す
		個数よりRecievedIngredientsが少なかったら残り全てを返す
	*/

	var res []*v1.Ingredient
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(4) + 1

	if count >= len(RecievedIngredients) {
		tmp := RecievedIngredients
		RecievedIngredients = []*v1.Ingredient{}
		return tmp
	}

	c := 0
	for {
		log.Println(RecievedIngredients)
		log.Println(len(RecievedIngredients))
		f := rand.Intn(2)
		if f == 1 {
			res = append(res, RecievedIngredients[c])
			RecievedIngredients = append(RecievedIngredients[:c], RecievedIngredients[c+1:]...)
		}
		if len(res) >= count {
			break
		}
		c = (c + 1) % len(RecievedIngredients)
	}

	return res
}
