package repository

import (
	"context"

	"github.com/gol4ng/skeleton/pkg/my_package"
)

type DocumentRepository interface {
	Store(document *my_package.Document, ctx context.Context) error
	Find(ctx context.Context) ([]*my_package.Document, error)
	FindBy(field string, value string, ctx context.Context) ([]*my_package.Document, error)
}
