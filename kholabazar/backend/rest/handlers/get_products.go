package handlers

import (
	"kholabazar/database"
	"kholabazar/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), 200)
}
