package product

import (
	"kholabazar/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me valid format id", 400)
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Product deletion failed!")
		return
	}
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found!")
		return
	}

	h.productRepo.Delete(product.ID)

	utils.SendData(w, "Product deleted successfully!", http.StatusOK)

}
