// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"fmt"
	"gorm.io/gorm"
)

const TableNameStudent = "student"

// Student mapped from table <student>
type Student struct {
	ID   int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name" json:"name"` // 学生姓名
	Age  int8   `gorm:"column:age" json:"age"`   // 年龄
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`//enable the gorm soft delete
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli" json:"updated_at"`

}

// TableName Student's table name
func (*Student) TableName() string {
	return TableNameStudent
}

func (student *Student) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Printf("准备更新 %v\n", student)
	return
}

func (student *Student) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Printf("更新完成 %v\n", student)
	return
}

func (student *Student) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Printf("成功删除 %v\n", student)
	return
}

func (student *Student) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("开始删除 %v\n", student)
	return
}




