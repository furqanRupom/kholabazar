package user

import (
	"encoding/json"
	"fmt"
	"kholabazar/repo"
	"kholabazar/utils"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser repo.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Please give me valid JSON!")
		return
	}
	createUser, err := h.userRepo.Create(newUser)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "User creation failed!")
	}
	utils.SendData(w, createUser, http.StatusCreated)

}
