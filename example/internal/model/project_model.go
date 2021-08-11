package model

import "time"

type ProjectModel struct {
	Id          int64      ` gorm:"column:id"  json:"id" `                   //id
	ProjectName string     ` gorm:"column:name"  json:"name" `               //project name
	CreateTime  time.Time  ` gorm:"column:create_time"  json:"create_time" ` //project create time
	UpdateTime  *time.Time ` gorm:"column:update_time"  json:"update_time" ` //
}

func (ProjectModel) TableName() string {
	return "project"
}
