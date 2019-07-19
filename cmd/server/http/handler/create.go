package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gol4ng/skeleton/pkg/my_package"
	"github.com/gol4ng/skeleton/pkg/my_package/repository"
)

// View handler
func Create(repository repository.DocumentRepository) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		doc := &my_package.Document{
			Title:       "yes",
			Type:        my_package.IMAGE,
			Description: "youpi",
		}
		if err := repository.Store(doc, request.Context()); err != nil {
			panic(err)
		}
		data, err := json.Marshal(doc)
		if err != nil {
			panic(err)
		}
		response.WriteHeader(http.StatusCreated)
		response.Write(data)
		request.Context()
	}
}
