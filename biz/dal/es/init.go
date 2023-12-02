package es

import (
	"github.com/olivere/elastic/v7"
)

var (
	EsClient *elastic.Client
)

func Init() {
	var err error

	// 创建 Elasticsearch 客户端
	EsClient, err = elastic.NewClient(elastic.SetURL("http://175.178.96.246:9200"))

	if err != nil {
		panic(err)
	}
}
