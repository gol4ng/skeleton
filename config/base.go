package config

import (
	"time"
)

// global config to be used by the container
type Config struct {
	Debug
	Elastic
	Rabbit
	// App configurations
	Server
	NotifyDocumentCreatedConsumer
	// <otherAppConf>
}

// Add all debugging & profiling configurations here
type Debug struct {
	Debug   bool `config:"debug"`
	Profile bool `config:"profile"`
}

// Add all Elastic related configurations here
type Elastic struct {
	ElasticURL            string        `config:"elastic_url"`
	ElasticUsername       string        `config:"elastic_username"`
	ElasticPassword       string        `config:"elastic_password"`
	ElasticIndexPrefixURL string        `config:"elastic_index_prefix"`
	ElasticTimeout        time.Duration `config:"elastic_timeout"`
	ElasticMaxIdleConns   int           `config:"elastic_max_idle_connections"`
	ElasticMaxRetry       int           `config:"elastic_max_retry"`
	// elastic index names
	ElasticDocumentIndexName string `config:"elastic_document_index_name"`
	// Elastic<OtherDocument>IndexName string `config:"elastic_<other_document>_index_name"`
}

type Rabbit struct {
	RabbitUrl string `config:"rabbit_url"`
	// add exchange names here
	RabbitOtherAppExchangeName string `config:"rabbit_other_app_exchange_name"`
	// add queue names here
	RabbitNotifyDocumentCreatedQueueName string `config:"rabbit_notify_document_created_queue_name"`
	RabbitOtherFakeActionToDoQueueName   string `config:"rabbit_other_fake_action_to_do_queue_name"`
}
