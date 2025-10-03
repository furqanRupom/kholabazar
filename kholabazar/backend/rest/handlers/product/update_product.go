package product

import (
	"encoding/json"
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProductData database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateProductData)
	if err != nil {
		utils.SendError(w, 403, "Data is invalid!")
		return
	}

	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me valid format id", 400)
	}

	product := database.Get(pId)
	if product == nil {
		utils.SendError(w, 404, "Product not found!")
		return
	}
	updateProductData.ID = pId
	database.Update(updateProductData)
	utils.SendData(w, "Product updated successfully!", 200)

}
