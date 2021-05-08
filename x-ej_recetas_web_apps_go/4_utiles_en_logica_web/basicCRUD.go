package crud

import (
	"encoding/json"
	"net/http"
)

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

	var products []*product.Product

	for results.Next() {
		product := &product.Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.Description)

		catch(err)
		products = append(products, product)
	}

	responseWithJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product product.Product
	json.NewDecoder(r.Body).Decode(&product)

	query, err := dbConn.Prepare(`INSERT products SET product_code=?, description=?`)
	catch(err)

	_, _err := query.Exec(product.ProductCode, product.Description)
	catch(_err)
	defer query.Close()

	responseWithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product product.Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := dbConn.Prepare(`UPDATE products SET product_code=?, description=? where id=?`)
	catch(err)

	_, _err := query.Exec(product.ProductCode, product.Description, id)
	catch(_err)
	defer query.Close()

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := dbConn.Prepare(`DELETE FROM products where id=?`)
	catch(err)

	_, _err := query.Exec(id)
	catch(_err)
	defer query.Close()

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}
