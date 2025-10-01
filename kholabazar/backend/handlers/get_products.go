package handlers 
import (
	"net/http"
	"kholabazar/database"
	"kholabazar/utils"
)
func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}
	utils.SendData(w, database.ProductList, 200)
}