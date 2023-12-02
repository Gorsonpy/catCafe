package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Gorsonpy/catCafe/biz/model/cat"
	"github.com/cloudwego/kitex/pkg/klog"
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

// MarshalJSON 实现了 time.Time 的自定义 JSON 格式化
func (cat *Cat) MarshalJSON() ([]byte, error) {
	type Alias Cat // 别名，以防止无限递归调用
	return json.Marshal(&struct {
		CheckInDate string `json:"checkInDate"`
		*Alias
	}{
		CheckInDate: cat.CheckInDate.Format("2006-01-02 15:04:05"),
		Alias:       (*Alias)(cat),
	})
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
