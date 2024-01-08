package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pet-dex-backend/v2/api/routes"
)

const (
	port = 8081
)

func main() {
	r := chi.NewRouter()
	routes.InitRouter(r)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
		if err != nil {
			log.Fatal("Erro ao iniciar o servidor:", err)
		}
	}()

	<-stop

	fmt.Println("\nServidor encerrado.")
}
