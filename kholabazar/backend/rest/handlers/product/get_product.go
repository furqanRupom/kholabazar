package product

import (
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me valid format id", 400)
		return
	}

	product := database.Get(pId)
	if product == nil {
		utils.SendError(w, 404,"Product not found!")
		return
	}
	utils.SendData(w, product, 200)
}
