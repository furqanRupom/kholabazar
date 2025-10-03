package user

import (
	"encoding/json"
	"kholabazar/config"
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	cnf := config.GetConfig()
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
		return
	}
	accessToken, err := utils.CreateJWT(utils.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	}, cnf.JWTSecret)
	if err != nil {
		http.Error(w, "Internal Server", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, accessToken, http.StatusOK)

}
