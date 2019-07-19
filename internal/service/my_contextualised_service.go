package service

import (
	"context"
)

var myContextualizedServiceKey struct{}

// you need to define this kinf of service only when the service only live in context
func (container *Container) GetMyContextualizedService(ctx context.Context /* you can add some service injection here*/) *MyContextualizeService {
	if myContextualizedService, ok := ctx.Value(myContextualizedServiceKey).(*MyContextualizeService); ok {
		return myContextualizedService
	}
	myContextualizedService := &MyContextualizeService{}
	context.WithValue(ctx, myContextualizedServiceKey, myContextualizedService)
	return myContextualizedService
}

type MyContextualizeService struct {
}

func (m *MyContextualizeService) Find() {
	// your service method
}
