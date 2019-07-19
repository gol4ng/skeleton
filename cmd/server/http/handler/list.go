package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gol4ng/skeleton/pkg/my_package/repository"
)

// View handler
func List(repository repository.DocumentRepository) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		//docs, err := repository.Find(request.Context())
		docs, err := repository.FindBy("title", "yes", request.Context())
		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(docs)
		if err != nil {
			panic(err)
		}
		response.WriteHeader(http.StatusOK)
		response.Write(data)
		request.Context()
	}
}
