package product

import (
	"kholabazar/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList,error := h.productRepo.List()
	if error != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to load product list")
		return
	}
	utils.SendData(w, productList, 200)
}
