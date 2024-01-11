package petcontroller

import (
	"net/http"
	"pet-dex-backend/v2/usecase"
)

type ExampleController struct {
	Usecase *usecase.ExampleUsecase
}

func NewExampleController(usecase *usecase.ExampleUsecase) *ExampleController {
	return &ExampleController{
		Usecase: usecase,
	}
}

func (e *ExampleController) ExampleHandler(w http.ResponseWriter, r *http.Request) {
	output := e.Usecase.Do()

	w.WriteHeader(200)
	w.Write([]byte(output))
}
