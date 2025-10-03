package user

import (
	"encoding/json"
	"fmt"
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
		return
	}
	createUser := newUser.Store()
	utils.SendData(w, createUser, 201)

}