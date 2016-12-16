package talk_training

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"fmt"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
    	w.Header().Set("Access-Control-Allow-Origin", "*")
	productIdStr :=  r.URL.Query().Get("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err!=nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	response := GetTalks(productId)
	json.NewEncoder(w).Encode(response)
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	userId, err := strconv.Atoi(r.FormValue("userId"))
	if err!=nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	shopId, err := strconv.Atoi(r.FormValue("shopId"))
	if err!=nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	productId, err := strconv.Atoi(r.FormValue("productId"))
	if err!=nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	message := r.FormValue("message")
	if err!=nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	createBy, err := strconv.Atoi(r.FormValue("createBy"))
	if err!=nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	response := AddTalks(userId, shopId, productId, message, createBy)
	json.NewEncoder(w).Encode(response)
}
