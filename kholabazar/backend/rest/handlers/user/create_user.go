package user

import (
	"encoding/json"
	"fmt"
	"kholabazar/repo"
	"kholabazar/utils"
	"net/http"
)

type ReqUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Please give me valid JSON!")
		return
	}
	createUser, err := h.userRepo.Create(repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "User creation failed!")
	}
	utils.SendData(w, createUser, http.StatusCreated)

}
