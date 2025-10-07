package product

import (
	"kholabazar/domain"
	"kholabazar/utils"
	"net/http"
	"strconv"
)

type PaginatedData struct {
	Meta Pagination        `json:"meta"`
	Data []*domain.Product `json:"data"`
}

type Pagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalPages int64 `json:"totalPages"`
	TotalItems int64 `json:"totalItems"`
}

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

	paginatedData := PaginatedData{
		Meta: Pagination{
			Page:       page,
			Limit:      limit,
			TotalPages: count / limit,
			TotalItems: count,
		},
		Data: productList,
	}
	utils.SendData(w, http.StatusOK, paginatedData)
}
