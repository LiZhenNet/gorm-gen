package dal

import (
	"time"

	"github.com/lizhennet/gorm-gen/example/internal/model"
	"gorm.io/gorm"
)

type ProjectModelCondition struct {
	where  []map[string]interface{}
	order  []string
	offset int
	limit  int
}

func NewProjectModelCondition() *ProjectModelCondition {
	return &ProjectModelCondition{}
}

func (c *ProjectModelCondition) clone() *ProjectModelCondition {
	clone := *c
	return &clone
}

func (c *ProjectModelCondition) Offset(offset int) *ProjectModelCondition {
	c.offset = offset
	return c
}

func (c *ProjectModelCondition) Limit(limit int) *ProjectModelCondition {
	c.limit = limit
	return c
}

func (c *ProjectModelCondition) Page(pageNo, pageSize int) *ProjectModelCondition {
	c.offset = (pageNo - 1) * pageSize
	c.limit = pageSize
	return c
}

func (c *ProjectModelCondition) Out(db *gorm.DB) *gorm.DB {
	for _, w := range c.where {
		for k, v := range w {
			db = db.Where(k, v)
		}
	}
	for _, o := range c.order {
		db = db.Order(o)
	}
	if c.offset != 0 {
		db = db.Offset(c.offset)
	}
	if c.limit != 0 {
		db = db.Limit(c.limit)
	}
	return db
}

func (c *ProjectModelCondition) IdEqual(v int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id = ?": v})
	return c
}

func (c *ProjectModelCondition) IdNotEqual(v int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id != ?": v})
	return c
}

func (c *ProjectModelCondition) IdLessThan(v int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id < ?": v})
	return c
}

func (c *ProjectModelCondition) IdLessThanOrEqualTo(v int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id <= ?": v})
	return c
}

func (c *ProjectModelCondition) IdGreaterThan(v int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id > ?": v})
	return c
}

func (c *ProjectModelCondition) IdGreaterThanOrEqualTo(v int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id >= ?": v})
	return c
}

func (c *ProjectModelCondition) IdIn(v []int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id IN (?)": v})
	return c
}

func (c *ProjectModelCondition) IdNotIn(v []int64) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"id NOT IN (?)": v})
	return c
}

func (c *ProjectModelCondition) OrderById() *ProjectModelCondition {
	c.order = append(c.order, "id")
	return c
}

func (c *ProjectModelCondition) OrderByIdDesc() *ProjectModelCondition {
	c.order = append(c.order, "id DESC")
	return c
}

func (c *ProjectModelCondition) ProjectNameEqual(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name = ?": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameNotEqual(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name != ?": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameLessThan(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name < ?": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameLessThanOrEqualTo(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name <= ?": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameGreaterThan(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name > ?": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameGreaterThanOrEqualTo(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name >= ?": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameIn(v []string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name IN (?)": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameNotIn(v []string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name NOT IN (?)": v})
	return c
}

func (c *ProjectModelCondition) ProjectNameLike(v string) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"name LIKE ? ": v})
	return c
}

func (c *ProjectModelCondition) OrderByProjectName() *ProjectModelCondition {
	c.order = append(c.order, "name")
	return c
}

func (c *ProjectModelCondition) OrderByProjectNameDesc() *ProjectModelCondition {
	c.order = append(c.order, "name DESC")
	return c
}

func (c *ProjectModelCondition) CreateTimeEqual(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time = ?": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeNotEqual(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time != ?": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeLessThan(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time < ?": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeLessThanOrEqualTo(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time <= ?": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeGreaterThan(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time > ?": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeGreaterThanOrEqualTo(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time >= ?": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeIn(v []time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time IN (?)": v})
	return c
}

func (c *ProjectModelCondition) CreateTimeNotIn(v []time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"create_time NOT IN (?)": v})
	return c
}

func (c *ProjectModelCondition) OrderByCreateTime() *ProjectModelCondition {
	c.order = append(c.order, "create_time")
	return c
}

func (c *ProjectModelCondition) OrderByCreateTimeDesc() *ProjectModelCondition {
	c.order = append(c.order, "create_time DESC")
	return c
}

func (c *ProjectModelCondition) UpdateTimeEqual(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time = ?": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeNotEqual(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time != ?": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeLessThan(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time < ?": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeLessThanOrEqualTo(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time <= ?": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeGreaterThan(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time > ?": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeGreaterThanOrEqualTo(v time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time >= ?": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeIn(v []time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time IN (?)": v})
	return c
}

func (c *ProjectModelCondition) UpdateTimeNotIn(v []time.Time) *ProjectModelCondition {
	c.where = append(c.where, map[string]interface{}{"update_time NOT IN (?)": v})
	return c
}

func (c *ProjectModelCondition) OrderByUpdateTime() *ProjectModelCondition {
	c.order = append(c.order, "update_time")
	return c
}

func (c *ProjectModelCondition) OrderByUpdateTimeDesc() *ProjectModelCondition {
	c.order = append(c.order, "update_time DESC")
	return c
}

type ProjectModelUpdater struct {
	values map[string]interface{}
}

func NewProjectModelUpdater() *ProjectModelUpdater {
	return &ProjectModelUpdater{values: map[string]interface{}{}}
}

func (u *ProjectModelUpdater) Values() map[string]interface{} {
	return u.values
}
func (u *ProjectModelUpdater) Id(v int64) *ProjectModelUpdater {
	u.values["id"] = v
	return u
}
func (u *ProjectModelUpdater) ProjectName(v string) *ProjectModelUpdater {
	u.values["name"] = v
	return u
}
func (u *ProjectModelUpdater) CreateTime(v time.Time) *ProjectModelUpdater {
	u.values["create_time"] = v
	return u
}
func (u *ProjectModelUpdater) UpdateTime(v time.Time) *ProjectModelUpdater {
	u.values["update_time"] = v
	return u
}

type ProjectModelCommonDal struct {
}

func (d ProjectModelCommonDal) Insert(db *gorm.DB, value *model.ProjectModel) (int64, error) {
	result := db.Create(value)
	return result.RowsAffected, result.Error
}

func (d ProjectModelCommonDal) BatchInsert(db *gorm.DB, value *[]model.ProjectModel) error {
	result := db.Create(value)
	return result.Error
}

func (d ProjectModelCommonDal) DeleteByCondition(db *gorm.DB, condition *ProjectModelCondition) (int64, error) {
	db = condition.Out(db)
	result := db.Delete(&model.ProjectModel{})
	return result.RowsAffected, result.Error
}

func (d ProjectModelCommonDal) UpdateByCondition(db *gorm.DB, condition *ProjectModelCondition, updater *ProjectModelUpdater) (int64, error) {
	db = condition.Out(db)
	result := db.Model(&model.ProjectModel{}).Updates(updater.Values())
	return result.RowsAffected, result.Error
}

func (d ProjectModelCommonDal) SelectByCondition(db *gorm.DB, condition *ProjectModelCondition) ([]*model.ProjectModel, error) {
	db = condition.Out(db)
	ret := make([]*model.ProjectModel, 0)
	result := db.Model(&model.ProjectModel{}).Find(&ret)
	return ret, result.Error
}

func (d ProjectModelCommonDal) CountByCondition(db *gorm.DB, condition *ProjectModelCondition) (int64, error) {
	db = condition.Out(db)
	var count int64
	result := db.Model(&model.ProjectModel{}).Count(&count)
	return count, result.Error
}
