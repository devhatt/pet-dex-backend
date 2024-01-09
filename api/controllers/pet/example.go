package petcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/infra/config"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Message string `json:"message"`
}

func ExampleController(w http.ResponseWriter, r *http.Request) {

	db := config.GetDB()

	_, err := db.Exec("INSERT INTO user (id, name) VALUES (?,?)", 1, "ola")
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	defer db.Close()

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func PatchController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(400)
		return
	}

	var response Response

	response.Message = "Oi"

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&response)
}
