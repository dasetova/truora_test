package webservice

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dasetova/truora_test/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func Setup() {
	utils.SetTestEnviroment()
	utils.CleanDB()
	utils.MigrateDB()
	utils.SeedDB()
}

func TestGetRecipe(t *testing.T) {
	Setup()
	url := fmt.Sprintf("/recipes/%v", utils.RecipeId)
	Convey("Given a HTTP request for /recipes/123", t, func() {
		req := httptest.NewRequest("GET", url, nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
			})
		})
	})
}

func TestGetCategories(t *testing.T) {
	Setup()

	Convey("Given a HTTP request for /categories", t, func() {
		req := httptest.NewRequest("GET", "/categories", nil)

		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
			})
		})
	})
}

func TestCreateRecipe(t *testing.T) {
	Setup()

	Convey("Given a HTTP request for /recipes", t, func() {
		body := fmt.Sprintf(`{"name":"Recipe 2", "description":"Recipe description", "category_id":%v, "indications":"test", "ingredients":[{"description":"onion","quantity":3,"measure_unit":"units"}]}`, utils.CategoryId)
		bodyReader := strings.NewReader(body)

		req := httptest.NewRequest("POST", "/recipes", bodyReader)

		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 201", func() {
				So(resp.Code, ShouldEqual, 201)
			})
		})
	})
}

func TestAddIngredientToRecipe(t *testing.T) {
	Setup()

	url := fmt.Sprintf("/recipes/%v/ingredients", utils.RecipeId)

	Convey("Given a HTTP request for /recipes/id/ingredients", t, func() {
		body := fmt.Sprintf(`{
			"description": "onion",
			"quantity": 1,
			"measure_unit": "unit"
		}`)
		bodyReader := strings.NewReader(body)

		req := httptest.NewRequest("POST", url, bodyReader)

		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 201", func() {
				So(resp.Code, ShouldEqual, 201)
			})
		})
	})
}

func TestRemoveIngredientToRecipe(t *testing.T) {
	Setup()

	url := fmt.Sprintf("/recipes/%v/ingredients/%v", utils.RecipeId, utils.IngredientId)

	Convey("Given a HTTP request for /recipes/id/ingredients", t, func() {

		req := httptest.NewRequest("DELETE", url, nil)

		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
			})
		})
	})
}
