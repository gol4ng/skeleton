package elastic

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"

	"github.com/gol4ng/skeleton/pkg/my_package"
)

type DocumentRepository struct {
	es *elastic.Client

	idxManager *IndexManager
	IdxName    string
}

func (d *DocumentRepository) getIndexName() string {
	return d.idxManager.GenerateName(d.IdxName)
}

func (d *DocumentRepository) ensureIndexExist(ctx context.Context) error {
	if err := d.idxManager.EnsureIndexExist(d.getIndexName(), DocumentMapping, ctx); err != nil {
		return err
	}
	return nil
}

func (d *DocumentRepository) indexService() *elastic.IndexService {
	return d.es.Index().Index(d.getIndexName()).Type(DocumentType)
}

func (d *DocumentRepository) searchService() *elastic.SearchService {
	return d.es.Search(d.getIndexName()).Type(DocumentType)
}

func (d *DocumentRepository) Store(document *my_package.Document, ctx context.Context) error {
	_, err := d.indexService().BodyJson(document).Do(ctx)
	return err
}

func (d *DocumentRepository) Find(ctx context.Context) ([]*my_package.Document, error) {
	var documents []*my_package.Document

	result, err := d.searchService().Query(elastic.NewMatchAllQuery()).Do(ctx)

	if err != nil {
		return documents, err
	}
	if result == nil || result.TotalHits() == 0 {
		return documents, nil
	}
	for _, hit := range result.Hits.Hits {
		doc := &my_package.Document{}
		err := json.Unmarshal(*hit.Source, doc)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func (d *DocumentRepository) FindBy(field string, value string, ctx context.Context) ([]*my_package.Document, error) {
	var documents []*my_package.Document

	result, err := d.searchService().Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery(field, value))).Do(ctx)

	if err != nil {
		return documents, err
	}
	if result == nil || result.TotalHits() == 0 {
		return documents, nil
	}
	for _, hit := range result.Hits.Hits {
		doc := &my_package.Document{}
		err := json.Unmarshal(*hit.Source, doc)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func NewDocumentRepository(idxname string, es *elastic.Client, idxManager *IndexManager) *DocumentRepository {
	return &DocumentRepository{
		es:         es,
		idxManager: idxManager,
		IdxName:    idxname,
	}
}
