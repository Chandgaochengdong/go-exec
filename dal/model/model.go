package model

import (
	"time"
)

type Student struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Age        int32     `json:"age"`
	Sex        string    `json:"sex"`
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	CreateBy   string    `json:"create_by"`
	UpdateBy   string    `json:"update_by"`
}
