package product

import (
	"encoding/json"
	"kholabazar/repo"
	"kholabazar/utils"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProductData repo.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateProductData)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Data is invalid!")
		return
	}

	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Please give me valid format id")
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to load product")
		return
	}
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found!")
		return
	}
	updateProductData.ID = pId
	h.productRepo.Update(updateProductData)

	utils.SendData(w, "Product updated successfully!", 200)

}
