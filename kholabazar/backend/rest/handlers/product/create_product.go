package product

import (
	"encoding/json"
	"fmt"
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
		return
	}
	createProduct := database.Store(newProduct)
	utils.SendData(w, createProduct, 201)

}
