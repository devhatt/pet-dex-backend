package petcontroller

import (
	"net/http"
	"pet-dex-backend/v2/infra/db"
)

func CreatePet(w http.ResponseWriter, r *http.Request) {
	// Exemplo de Criacao Pet Use Case
	usecase := NewPetUseCase(db.NewPetRepository())

	err := usecase.Do()

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
}
