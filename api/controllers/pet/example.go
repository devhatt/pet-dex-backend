package petcontroller

import (
	"net/http"
	"pet-dex-backend/v2/usecase/pet"
)

type ExampleController struct {
	Usecase *pet.ExampleUsecase
}

func NewExampleController(usecase *pet.ExampleUsecase) *ExampleController {
	return &ExampleController{
		Usecase: usecase,
	}
}

func (e *ExampleController) ExampleHandler(w http.ResponseWriter, r *http.Request) {
	output := e.Usecase.Do()

	w.WriteHeader(200)
	w.Write([]byte(output))
}
