package elastic

const (
	DocumentType    = "document"
	DocumentMapping = `{
		"mappings" : {
			"document" : {
			  	"properties" : {
					"title":{
						"type" : "keyword",
						"store" : true,
						"indexService": true
					},
					"type":{
						"type" : "keyword",
						"store" : true,
						"indexService": true
					},
					"description":{
						"type" : "keyword",
						"store" : true,
						"indexService": true
					}
			  	}
			}
		}
	}`
)
