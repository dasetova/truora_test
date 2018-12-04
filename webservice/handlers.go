package webservice

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/dasetova/truora_test/models"
)

func getCategories(w http.ResponseWriter, req *http.Request) {
	categories, err := models.GetCategories()

	if err != nil {
		convertError(w, http.StatusInternalServerError, err.Error())
	}

	convertResponse(w, http.StatusOK, categories)
}

func getRecipes(w http.ResponseWriter, req *http.Request) {
	filters := req.URL.Query()
	recipe_name := filters.Get("name")

	recipes, err := models.GetRecipes(recipe_name)

	if err != nil {
		convertError(w, http.StatusInternalServerError, err.Error())
	}

	convertResponse(w, http.StatusOK, recipes)
}

func getRecipe(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	recipe_id, err := strconv.Atoi(params["recipeId"])

	if err != nil {
		convertError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	recipe := models.Recipe{ID: recipe_id}

	err = recipe.GetRecipe()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			convertError(w, http.StatusNotFound, "Recipe not found")
		default:
			convertError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	convertResponse(w, http.StatusOK, recipe)
}

func createRecipe(w http.ResponseWriter, req *http.Request) {
	var recipe models.Recipe
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&recipe); err != nil {
		convertError(w, http.StatusBadRequest, "Invalid params")
		return
	}
	defer req.Body.Close()
	if err := recipe.CreateRecipe(); err != nil {

		convertError(w, http.StatusInternalServerError, err.Error())
		return
	}
	convertResponse(w, http.StatusCreated, recipe)
}

func updateRecipe(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["recipeId"])
	if err != nil {
		convertError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	var recipe models.Recipe
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&recipe); err != nil {
		convertError(w, http.StatusBadRequest, "Invalid params")
		return
	}
	defer req.Body.Close()
	recipe.ID = id

	result, err := recipe.UpdateRecipe()

	if err != nil {
		convertError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		convertError(w, http.StatusNotFound, "Recipe not found")
		return
	}
	convertResponse(w, http.StatusOK, recipe)
}

func deleteRecipe(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	recipe_id, err := strconv.Atoi(params["recipeId"])
	if err != nil {
		convertError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	recipe := models.Recipe{ID: recipe_id}

	defer req.Body.Close()

	result, err := recipe.DeleteRecipe()

	if err != nil {
		convertError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		convertError(w, http.StatusNotFound, "Recipe not found")
		return
	}

	convertResponse(w, http.StatusOK, map[string]string{"result": "recipe delete succesfully"})
}

func addIngredientToRecipe(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	recipe_id, err := strconv.Atoi(params["recipeId"])
	if err != nil {
		convertError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	var ingredient models.Ingredient
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		convertError(w, http.StatusBadRequest, "Invalid params")
		return
	}

	ingredient.RecipeID = recipe_id

	defer req.Body.Close()
	if err := ingredient.AddIngredientToRecipe(); err != nil {
		convertError(w, http.StatusInternalServerError, err.Error())
		return
	}
	convertResponse(w, http.StatusCreated, ingredient)
}

func removeIngredientToRecipe(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	recipe_id, err := strconv.Atoi(params["recipeId"])
	if err != nil {
		convertError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}
	ingredient_id, err := strconv.Atoi(params["ingredientId"])

	ingredient := models.Ingredient{ID: ingredient_id, RecipeID: recipe_id}

	defer req.Body.Close()

	result, err := ingredient.RemoveIngredientToRecipe()
	if err != nil {
		convertError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		convertError(w, http.StatusNotFound, "Ingredient not found for given receip")
		return
	}
	convertResponse(w, http.StatusOK, map[string]string{"result": "ingredient removed from receipt"})
}

func convertError(w http.ResponseWriter, code int, message string) {
	convertResponse(w, code, map[string]string{"error": message})
}

func convertResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}
