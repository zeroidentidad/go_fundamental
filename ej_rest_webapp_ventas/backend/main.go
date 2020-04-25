package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zeroidentidad/backend/database"
	"github.com/zeroidentidad/backend/models"
)

var dbConn *sql.DB

func main() {

	dbConn = database.InitDB()
	defer dbConn.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", AllProducts)
	r.Post("/products", CreateProduct)
	http.ListenAndServe(":3000", r)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'') FROM products`
	results, err := dbConn.Query(sql)
	catch(err)

	var products []*models.Product

	for results.Next() {
		product := &models.Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}

	responseWithJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	query, err := dbConn.Prepare(`INSERT products SET product_code=?, description=?`)
	catch(err)

	_, _err := query.Exec(product.Product_Code, product.Description)
	catch(_err)
	defer query.Close()

	responseWithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}
