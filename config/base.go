package config

import (
	"time"

	"github.com/gol4ng/skeleton/pkg/config"
)

type Base struct {
	Debug             bool   `config:"debug"`
	DocumentIndexName string `config:"document_index_name"`

	ElasticURL            string        `config:"elastic_url"`
	ElasticUsername       string        `config:"elastic_username" print:"-"`
	ElasticPassword       string        `config:"elastic_password" print:"-"`
	ElasticIndexPrefixURL string        `config:"elastic_index_prefix"`
	ElasticTimeout        time.Duration `config:"elastic_timeout"`
	ElasticMaxIdleConns   int           `config:"elastic_max_idle_connections"`
	ElasticMaxRetry       int           `config:"elastic_max_retry"`

	Server
}

func (c *Base) String() string {
	return config.ToString(c)
}
