package handlers

import (
	"encoding/json"
	"fmt"
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
		return
	}
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	utils.SendData(w, newProduct, 201)

}
