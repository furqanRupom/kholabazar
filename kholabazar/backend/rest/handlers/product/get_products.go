package product

import (
	"kholabazar/utils"
	"net/http"
	"strconv"
)


func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageStr := reqQuery.Get("page")
	limitStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to load product list")
		return
	}
	count, _ := h.svc.Count()

	utils.SendPage(w, productList,page,limit,count)
}
