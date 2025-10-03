package product

import (
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), 200)
}
