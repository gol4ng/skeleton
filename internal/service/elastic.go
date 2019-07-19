package service

import (
	"fmt"
	"net/http"

	"github.com/olivere/elastic"

	esRepo "github.com/gol4ng/skeleton/internal/repository/elastic"
	"github.com/gol4ng/skeleton/pkg/my_package/repository"
)

type loggerWrapper struct {
}

func (l loggerWrapper) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (container *Container) GetEsClient() *elastic.Client {
	if container.esClient == nil {
		var err error
		container.esClient, err = elastic.NewSimpleClient(
			elastic.SetURL(container.Cfg.ElasticURL),
			//elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewExponentialBackoff(10*time.Millisecond, 1*time.Second))),
			elastic.SetBasicAuth(container.Cfg.ElasticUsername, container.Cfg.ElasticPassword),
			elastic.SetHttpClient(&http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost: container.Cfg.ElasticMaxIdleConns,
				},
				Timeout: container.Cfg.ElasticTimeout,
			}),
			//elastic.SetErrorLog(loggerWrapper{}),
			//elastic.SetInfoLog(loggerWrapper{}),
			//elastic.SetTraceLog(loggerWrapper{}),
			//elastic.SetSniff(false),
		)
		if err != nil {
			panic(err)
		}
	}
	return container.esClient
}

func (container *Container) GetIndexManager() *esRepo.IndexManager {
	if container.indexManager == nil {
		container.indexManager = esRepo.NewIndexManager(
			container.GetEsClient(),
			container.Cfg.ElasticIndexPrefixURL,
		)
	}
	return container.indexManager
}

func (container *Container) GetDocumentRepository() repository.DocumentRepository {
	if container.documentRepository == nil {
		container.documentRepository = esRepo.NewDocumentRepository(
			container.Cfg.DocumentIndexName,
			container.GetEsClient(),
			container.GetIndexManager(),
		)
	}
	return container.documentRepository
}
