package MyRest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json: "name"`
	Ingredients string `json:"ingredients"`
}

var Rolls []Roll

func CreateSushiRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var newroll Roll

	json.NewDecoder(r.Body).Decode(&newroll)
	newroll.ID = strconv.Itoa(len(Rolls) + 1)

	Rolls = append(Rolls, newroll)

	json.NewEncoder(w).Encode(newroll)

}
func GetSushiRollsDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(Rolls)
}

func GetSushiRollDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// func postSushiDetails() {

// }

// func DeleteSushiDetails() {

// }
