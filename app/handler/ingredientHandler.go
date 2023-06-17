package handler

import "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1/ingredientRainv1connect"

type IngredientHandler struct {
	ingredientRainv1connect.UnimplementedIngredientServiceHandler
}
