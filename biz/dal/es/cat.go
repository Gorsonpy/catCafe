package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Gorsonpy/catCafe/biz/model/cat"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/olivere/elastic/v7"
)

type Cat struct {
	CatId          int64     `json:"catId" gorm:"primary_key"`
	Name           string    `json:"name"`
	Breed          string    `json:"breed"`
	Gender         string    `json:"gender"`
	Age            int64     `json:"age"`
	HealthStatus   string    `json:"healthStatus"`
	PhotoUrl       string    `json:"photoUrl"`
	CheckInDate    time.Time `json:"checkInDate"`
	AppointmentNum int64     `json:"appointmentNum"`
}

func QueryCats(req *cat.QueryCatsReq) ([]*Cat, error) {
	// 构建查询条件
	query := elastic.NewBoolQuery().
		Should(
			elastic.NewMatchQuery("name", req.SearchContent).Fuzziness("AUTO"),
			elastic.NewMatchQuery("name.pinyin", req.SearchContent).Fuzziness("AUTO"),
			elastic.NewMatchQuery("breed", req.SearchContent).Fuzziness("AUTO"),
			elastic.NewMatchQuery("breed.pinyin", req.SearchContent).Fuzziness("AUTO"),
			elastic.NewMatchQuery("gender", req.SearchContent).Fuzziness("AUTO"),
			elastic.NewMatchQuery("healthStatus", req.SearchContent).Fuzziness("AUTO"),
			elastic.NewMatchQuery("healthStatus.pinyin", req.SearchContent).Fuzziness("AUTO"),
		)

	if req.LAge != 0 {
		// 添加范围过滤条件
		query = query.Filter(elastic.NewRangeQuery("age").Gte(req.LAge))
	}

	if req.RAge != 0 {
		query = query.Filter(elastic.NewRangeQuery("age").Lte(req.RAge))
	}

	// 添加性别过滤条件
	if req.Gender != "" {
		query = query.Filter(elastic.NewTermQuery("gender.keyword", req.Gender))
	}

	// 添加品种过滤条件
	if req.Breed != "" {
		query = query.Filter(elastic.NewTermQuery("breed.keyword", req.Breed))
	}

	// 构建搜索请求
	searchResult, err := EsClient.Search().
		Index("cat_index"). // 替换为实际的索引名称
		Query(query).
		Size(int(req.Limit)). // 设置返回的文档数量限制
		Do(context.Background())
	if err != nil {
		// 处理错误
		return nil, err
	}

	// 处理搜索结果
	var cats []*Cat
	for _, hit := range searchResult.Hits.Hits {
		// 将 Source 转为字符串
		sourceStr := string(hit.Source)
		// 打印文档内容
		fmt.Println("Document Source:", sourceStr)

		var cat Cat
		err := json.Unmarshal(hit.Source, &cat)
		if err != nil {
			// 处理解析错误
			return nil, err
		}
		cats = append(cats, &cat)
	}

	// 返回查询结果
	return cats, nil
}

func CatToModel(c []*Cat) []*cat.CatModel {
	catList := make([]*cat.CatModel, 0)
	for _, val := range c {
		tmp := cat.NewCatModel()
		tmp.CatId = val.CatId
		tmp.Name = val.Name
		tmp.Breed = val.Breed
		tmp.Gender = val.Gender
		tmp.Age = val.Age
		tmp.HealthStatus = val.HealthStatus
		tmp.PhotoUrl = val.PhotoUrl
		tmp.CheckInDate = val.CheckInDate.Format(time.DateTime)
		tmp.AppointmentNum = val.AppointmentNum
		catList = append(catList, tmp)
	}
	return catList
}

func AddCat(cat *Cat) error {
	var err error

	a, _ := json.Marshal(cat)
	klog.Info(string(a))
	// 添加文档到索引
	_, err = EsClient.Index().
		Index("cat_index").
		Type("_doc").
		Id(fmt.Sprintf("%d", cat.CatId)).
		BodyJson(cat).
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
