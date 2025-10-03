package handlers

import (
	"encoding/json"
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid data !", 400)
		return
	}
	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "User not found!", 404)
	}
	utils.SendData(w, usr, http.StatusCreated)

}
