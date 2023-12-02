package es

import (
	"context"
	"fmt"
	"time"
)

type Cat struct {
	CatId        int64     `json:"catId" gorm:"primary_key"`
	Name         string    `json:"name"`
	Breed        string    `json:"breed"`
	Gender       string    `json:"gender"`
	Age          int64     `json:"age"`
	HealthStatus string    `json:"healthStatus"`
	PhotoUrl     string    `json:"photoUrl"`
	CheckInDate  time.Time `json:"checkInDate"`
	AppointmentNum int64   `json:"appointmentNum"`
}

func AddCat(cat *Cat) error {
	var err error

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
