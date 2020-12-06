package config

import (
	"context"

	"github.com/gol4ng/skeleton/pkg/config"
)

// this config is related to cmd/job/notify_document_created_consumer
type NotifyDocumentCreatedConsumer struct {
	// add specific totoConsumerConfig
}

func NewNotifyDocumentCreatedConsumer() *Config {
	cfg := &Config{}

	config.LoadOrFatal(context.Background(), cfg)
	// @TODO validate config here
	return cfg
}

// Validate NotifyDocumentCreatedConsumer config in internal ?
func ValidateNotifyDocumentCreatedConsumer() {
}
