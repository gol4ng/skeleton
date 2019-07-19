package config

import (
	"time"
)

type Base struct {
	Debug             bool   `config:"debug"`
	DocumentIndexName string `config:"document_index_name"`

	ElasticURL            string        `config:"elastic_url"`
	ElasticUsername       string        `config:"elastic_username"`
	ElasticPassword       string        `config:"elastic_password"`
	ElasticIndexPrefixURL string        `config:"elastic_index_prefix"`
	ElasticTimeout        time.Duration `config:"elastic_timeout"`
	ElasticMaxIdleConns   int           `config:"elastic_max_idle_connections"`
	ElasticMaxRetry       int           `config:"elastic_max_retry"`

	Server
}
