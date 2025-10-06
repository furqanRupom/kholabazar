package product

import (
	"kholabazar/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList,err := h.svc.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to load product list")
		return
	}
	utils.SendData(w, productList, 200)
}
