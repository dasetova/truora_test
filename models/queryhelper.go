package models

import (
	"database/sql"
	"log"

	"github.com/dasetova/truora_test/utils"
	_ "github.com/lib/pq"
)

func GetCategories() ([]Category, error) {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

// Creates recipe and ingredients (if they are related)
func (recipe *Recipe) CreateRecipe() error {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	err = tx.QueryRow(
		"INSERT INTO recipes(name, description, indications, category_id) VALUES($1, $2, $3, $4) RETURNING id",
		recipe.Name, recipe.Description, recipe.Indications, recipe.CategoryID).Scan(&recipe.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(recipe.Ingredients) > 0 {
		for i := 0; i < len(recipe.Ingredients); i++ {
			err = recipe.Ingredients[i].addIngredientToRecipeTx(tx, recipe.ID)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()

	return nil
}

// Public function used to add ingredient to an existing recipe
func (ingredient *Ingredient) AddIngredientToRecipe() error {
	db, err := utils.ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(
		"INSERT INTO ingredients(description, quantity, measure_unit, recipe_id) VALUES($1, $2, $3, $4) RETURNING id",
		ingredient.Description, ingredient.Quantity, ingredient.MeasureUnit, ingredient.RecipeID).Scan(&ingredient.ID)

	if err != nil {
		return err
	}

	return nil
}

//Removes the given ingredient
func (ingredient *Ingredient) RemoveIngredientToRecipe() (sql.Result, error) {
	db, err := utils.ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db.Exec("DELETE FROM ingredients WHERE id = $1 and recipe_id = $2", ingredient.ID, ingredient.RecipeID)
}

//Private function used in Recipe's deletion
func (ingredient *Ingredient) removeIngredientToRecipeTx(tx *sql.Tx, recipe_id int) (sql.Result, error) {
	return tx.Exec("DELETE FROM ingredients WHERE id = $1 and recipe_id = $2", ingredient.ID, recipe_id)
}

// Private function used in Recipe's creation
func (ingredient *Ingredient) addIngredientToRecipeTx(tx *sql.Tx, recipe_id int) error {

	err := tx.QueryRow(
		"INSERT INTO ingredients(description, quantity, measure_unit, recipe_id) VALUES($1, $2, $3, $4) RETURNING id",
		ingredient.Description, ingredient.Quantity, ingredient.MeasureUnit, recipe_id).Scan(&ingredient.ID)

	if err != nil {
		return err
	}

	return nil
}

// Get recipes with filter by name
func GetRecipes(name string) ([]Recipe, error) {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query(`SELECT * FROM recipes r inner join categories c on r.category_id = c.id WHERE r.name LIKE '%' || $1 || '%'`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []Recipe{}
	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.CategoryID, &recipe.Indications, &recipe.Category.ID, &recipe.Category.Name)
		if err != nil {
			return nil, err
		}
		err = recipe.preloadRecipeIngredients()
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}

// Get recipe by ID
func (recipe *Recipe) GetRecipe() error {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.QueryRow("SELECT * FROM recipes r inner join categories c on r.category_id = c.id WHERE r.id = $1",
		recipe.ID).Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.CategoryID, &recipe.Indications, &recipe.Category.ID, &recipe.Category.Name)

	if err != nil {
		return err
	}
	return recipe.preloadRecipeIngredients()
}

// Deletes a recipe with its ingredients
func (recipe *Recipe) DeleteRecipe() (result sql.Result, err error) {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	result, err = db.Exec("DELETE FROM ingredients WHERE recipe_id=$1", recipe.ID)

	if err != nil {
		tx.Rollback()
		return result, err
	}

	result, err = db.Exec("DELETE FROM recipes WHERE id=$1", recipe.ID)

	if err != nil {
		tx.Rollback()
		return result, err
	}

	tx.Commit()

	return result, nil
}

// preload ingredients related with a recipe
func (recipe *Recipe) preloadRecipeIngredients() error {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM ingredients WHERE recipe_id = $1`, recipe.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ingredient Ingredient
		err := rows.Scan(&ingredient.ID, &ingredient.Description, &ingredient.Quantity, &ingredient.MeasureUnit, &ingredient.RecipeID)
		if err != nil {
			return err
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (recipe *Recipe) UpdateRecipe() (sql.Result, error) {
	db, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db.Exec("UPDATE recipes SET name = $1, description = $2, indications = $3, category_id = $4 WHERE id=$5", recipe.Name, recipe.Description, recipe.Indications, recipe.CategoryID, recipe.ID)
}
