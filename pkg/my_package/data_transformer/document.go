package data_transformer

import (
	"github.com/gol4ng/skeleton/pkg/my_package"
	protos "github.com/gol4ng/skeleton/pkg/my_package/protos"
)

func TransformDocument(document *my_package.Document) *protos.Document {
	return &protos.Document{
		Title:       document.Title,
		Type:        TransformType(document.Type),
		Description: document.Description,
	}
}

func ReverseTransformDocument(document *protos.Document) *my_package.Document {
	return &my_package.Document{
		Title:       document.Title,
		Type:        ReverseTransformType(document.Type),
		Description: document.Description,
	}
}
