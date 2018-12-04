package utils

import (
	"log"
	"os"
)

var (
	CategoryId   = 0
	RecipeId     = 0
	IngredientId = 0
)

// Validate EnvVar to Test Database is seted
func ValidateTestEnvVars() error {
	if len(os.Getenv("DATABASE_URL_TEST")) == 0 {
		return errEnvDBUrlTest
	}
	return nil
}

// Clean Test DB to tests executions
func CleanDB() {
	db, err := ConnectToDB()

	if err != nil {
		log.Fatal(ErrDBConnection)
	}

	db.Exec("DROP TABLE IF EXISTS ingredients;")
	db.Exec("DROP TABLE IF EXISTS recipes;")
	db.Exec("DROP TABLE IF EXISTS categories;")
	db.Exec("DROP TABLE IF EXISTS gomigrate;")
	db.Close()
}

// Set DATABASE_URL pointing to DATABASE_URL_TEST for tests executions
func SetTestEnviroment() {
	ValidateTestEnvVars()
	os.Setenv("DATABASE_URL", os.Getenv("DATABASE_URL_TEST"))
}

// Creates seed data to tests
func SeedDB() {
	db, err := ConnectToDB()

	if err != nil {
		log.Fatal(ErrDBConnection)
	}

	err = db.QueryRow(
		"INSERT INTO categories(name) VALUES($1) RETURNING id",
		"Test Category").Scan(&CategoryId)

	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow(
		"INSERT INTO recipes(name, description, category_id, indications) VALUES($1, $2, $3, $4) RETURNING id",
		"Test Recipe", "No description", CategoryId, "No indications").Scan(&RecipeId)

	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow(
		"INSERT INTO ingredients(description, quantity, measure_unit, recipe_id) VALUES($1, $2, $3, $4) RETURNING id",
		"Test Ingredient", "1", "Units", RecipeId).Scan(&IngredientId)

	if err != nil {
		log.Fatal(err)
	}

}
