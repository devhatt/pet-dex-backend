package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"pet-dex-backend/v2/entity"
)

var Pets map[string]entity.Pet

func init() {
	Pets = make(map[string]entity.Pet)
}

func AtualizarAniversario(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pet, ok := Pets[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Erro ao decodificar a solicitação", http.StatusBadRequest)
		return
	}
	if aniversario, ok := requestBody["aniversario"]; ok {
		log.Printf("ID do Pet: %s", id)
		log.Printf("Aniversário antes da atualização: %s", pet.Aniversario)
		log.Printf("Novo Aniversário recebido: %s", aniversario)

		// Aplicar a atualização
		pet.Aniversario = aniversario
		Pets[id] = pet // Atualize o mapa global com o pet modificado

		log.Printf("Aniversário após a atualização: %s", pet.Aniversario)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Aniversário atualizado com sucesso"))
	} else {
		log.Println("Campo 'aniversario' não encontrado no corpo da requisição")
		http.Error(w, "Campo 'aniversario' não encontrado no corpo da requisição", http.StatusBadRequest)
	}
}

func AtualizarDiadeAdocao(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pet, ok := Pets[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Erro ao decodificar a solicitação", http.StatusBadRequest)
		return
	}
	if diaDoacao, ok := requestBody["dia_doacao"]; ok {
		log.Printf("ID do Pet: %s", id)
		log.Printf("Dia de Doação antes da atualização: %s", pet.DiaDoacao)
		log.Printf("Novo Dia de Doação recebido: %s", diaDoacao)

		// Aplicar a atualização
		pet.DiaDoacao = diaDoacao
		Pets[id] = pet // Atualize o mapa global com o pet modificado

		log.Printf("Dia de Doação após a atualização: %s", pet.DiaDoacao)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Dia de Doação atualizado com sucesso"))
	} else {
		log.Println("Campo 'dia_doacao' não encontrado no corpo da requisição")
		http.Error(w, "Campo 'dia_doacao' não encontrado no corpo da requisição", http.StatusBadRequest)
	}
}
