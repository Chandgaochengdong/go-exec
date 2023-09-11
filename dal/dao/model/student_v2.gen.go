// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStudentV2 = "student_v2"

// StudentV2 mapped from table <student_v2>
type StudentV2 struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键id" json:"id"`                           // 主键id
	Name       string    `gorm:"column:name;not null;comment:姓名" json:"name"`                                              // 姓名
	Age        string    `gorm:"column:age;not null;comment:年龄" json:"age"`                                                // 年龄
	Sex        string    `gorm:"column:sex;not null;comment:性别" json:"sex"`                                                // 性别
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP(3);comment:更新时间" json:"update_time"` // 更新时间
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP(3);comment:创建时间" json:"create_time"` // 创建时间
	CreateBy   string    `gorm:"column:create_by;comment:创建人" json:"create_by"`                                            // 创建人
	UpdateBy   string    `gorm:"column:update_by;default:1;comment:更新人" json:"update_by"`                                  // 更新人
}

// TableName StudentV2's table name
func (*StudentV2) TableName() string {
	return TableNameStudentV2
}