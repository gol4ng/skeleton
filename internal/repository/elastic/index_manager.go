package elastic

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

const IndexSuffixTime = "-2006-01-02-15"

type IndexManager struct {
	es          *elastic.Client
	indexPrefix string
}

func (m *IndexManager) GenerateName(name string) string {
	return m.indexPrefix + "_" + name
}

func (m *IndexManager) GenerateTimeName(name string, t time.Time) string {
	return m.GenerateName(name) + t.Format(IndexSuffixTime)
}

func (m *IndexManager) GenerateNowName(name string) string {
	return m.GenerateTimeName(name, time.Now())
}

func (m *IndexManager) IndexExist(name string, ctx context.Context) (bool, error) {
	return m.es.IndexExists(name).Do(ctx)
}

func (m *IndexManager) EnsureIndexExist(rawName string, mapping string, ctx context.Context) error {
	exist, err := m.es.IndexExists(rawName).Do(ctx)

	if err == nil && !exist {
		fmt.Printf("create elastic searchService indexService \"%s\"(mapping: %s)\n", rawName, mapping)
		_, err = m.es.CreateIndex(rawName).Body(mapping).Do(ctx)
	}

	return err
}

func (m *IndexManager) MoveAlias(aliasName string, indexName string, ctx context.Context) error {
	res, err := m.es.Aliases().Index("_all").Do(ctx)
	if err != nil {
		return err
	}

	aliasService := m.es.Alias()
	for _, i := range res.IndicesByAlias(aliasName) {
		aliasService.Remove(i, aliasName)
	}
	_, err = aliasService.Add(indexName, aliasName).Do(ctx)
	return err
}

func NewIndexManager(es *elastic.Client, indexPrefix string) *IndexManager {
	return &IndexManager{
		es:          es,
		indexPrefix: indexPrefix,
	}
}
