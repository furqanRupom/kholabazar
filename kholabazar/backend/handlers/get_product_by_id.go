package handlers

import (
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me valid format id", 400)
		return
	}
	for _, product := range database.ProductList {
		if product.ID == pId {
			utils.SendData(w, product, 200)
			return
		}
	}
		utils.SendData(w, "data not found", 404)
}
