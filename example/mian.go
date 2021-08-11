package main

import (
	"fmt"
	"github.com/lizhennet/gorm-gen/example/internal/dal"
	"github.com/lizhennet/gorm-gen/example/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:33066)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println(err)
	}
	db = db.Debug()
	commonDal := dal.ProjectModelCommonDal{}
	// INSERT INTO `project` (`name`,`create_time`,`update_time`) VALUES ('example','2021-08-10 19:55:36.255',NULL)
	row, err := commonDal.Insert(db, &model.ProjectModel{
		ProjectName: "example",
		CreateTime:  time.Now(),
		UpdateTime:  nil,
	})
	println(fmt.Sprintf("insert row:%v,err:%v", row, err))

	values := &[]model.ProjectModel{
		{
			ProjectName: "BatchInsertExample1",
			CreateTime:  time.Now(),
			UpdateTime:  nil,
		},
		{
			ProjectName: "BatchInsertExample2",
			CreateTime:  time.Now(),
			UpdateTime:  nil,
		}}
	// INSERT INTO `project` (`name`,`create_time`,`update_time`) VALUES ('BatchInsertExample1','2021-08-10 19:55:36.262',NULL),('BatchInsertExample2','2021-08-10 19:55:36.262',NULL)
	err = commonDal.BatchInsert(db, values)
	println(fmt.Sprintf("BatchInsert err:%v", err))

	c := dal.NewProjectModelCondition().IdGreaterThan(2).ProjectNameEqual("BatchInsertExample1")
	// DELETE FROM `project` WHERE id > 2 AND name = 'BatchInsertExample1'
	row, err = commonDal.DeleteByCondition(db, c)
	println(fmt.Sprintf("delete row:%v,err:%v", row, err))

	c2 := dal.NewProjectModelCondition().IdGreaterThan(2).ProjectNameEqual("BatchInsertExample2")
	// UPDATE `project` SET `name`='update' WHERE id > 2 AND name = 'BatchInsertExample2'
	row, err = commonDal.UpdateByCondition(db, c2, dal.NewProjectModelUpdater().ProjectName("update"))
	println(fmt.Sprintf("update row:%v,err:%v", row, err))

	c3 := dal.NewProjectModelCondition().IdGreaterThan(2).ProjectNameEqual("example").CreateTimeLessThanOrEqualTo(time.Now())
	// SELECT * FROM `project` WHERE id > 2 AND name = 'example' AND create_time <= '2021-08-10 19:59:50.197'
	data, err := commonDal.SelectByCondition(db, c3)
	println(fmt.Sprintf("select data len:%v,err:%v", len(data), err))

	c4 := dal.NewProjectModelCondition().IdGreaterThan(2).ProjectNameEqual("example").CreateTimeLessThanOrEqualTo(time.Now()).Page(2, 10)
	// SELECT * FROM `project` WHERE id > 2 AND name = 'example' AND create_time <= '2021-08-11 14:50:42.007' LIMIT 10 OFFSET 10
	data, err = commonDal.SelectByCondition(db, c4)
	println(fmt.Sprintf("select page data len:%v,err:%v", len(data), err))

	// SELECT count(*) FROM `project` WHERE id > 2 AND name = 'example' AND create_time <= '2021-08-10 20:02:14.061'
	count, err := commonDal.CountByCondition(db, c3)
	println(fmt.Sprintf("count  count:%v,err:%v", count, err))

}
