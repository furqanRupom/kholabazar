package product

import (
	"encoding/json"
	"kholabazar/repo"
	"kholabazar/utils"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
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
	req.ID = pId
	_, err = h.productRepo.Update(repo.Product(req))
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Product Update failed")
		return
	}

	utils.SendData(w, "Product updated successfully!", 200)

}
