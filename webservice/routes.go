package webservice

import (
	"net/http"
)

// route struct with routes information
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// type Routes []Route

// Initialize our routes
var routes = []Route{
	Route{
		"GetCategories",
		"GET",
		"/categories",
		getCategories,
	},
	Route{
		"GetRecipes",
		"GET",
		"/recipes",
		getRecipes,
	},
	Route{
		"GetRecipe",
		"GET",
		"/recipes/{recipeId}",
		getRecipe,
	},
	Route{
		"GetRecipes",
		"POST",
		"/recipes",
		createRecipe,
	},
	Route{
		"GetRecipes",
		"PUT",
		"/recipes/{recipeId}",
		updateRecipe,
	},
	Route{
		"DeleteRecipe",
		"DELETE",
		"/recipes/{recipeId}",
		deleteRecipe,
	},
	Route{
		"AddIngredientToRecipe",
		"POST",
		"/recipes/{recipeId}/ingredients",
		addIngredientToRecipe,
	},
	Route{
		"RemoveIngredientToRecipe",
		"DELETE",
		"/recipes/{recipeId}/ingredients/{ingredientId}",
		removeIngredientToRecipe,
	},
}
