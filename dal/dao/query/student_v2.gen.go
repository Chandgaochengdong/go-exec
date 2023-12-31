// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"com.shopline.richard/exec/dal/dao/model"
)

func newStudentV2(db *gorm.DB, opts ...gen.DOOption) studentV2 {
	_studentV2 := studentV2{}

	_studentV2.studentV2Do.UseDB(db, opts...)
	_studentV2.studentV2Do.UseModel(&model.StudentV2{})

	tableName := _studentV2.studentV2Do.TableName()
	_studentV2.ALL = field.NewAsterisk(tableName)
	_studentV2.ID = field.NewInt64(tableName, "id")
	_studentV2.Name = field.NewString(tableName, "name")
	_studentV2.Age = field.NewString(tableName, "age")
	_studentV2.Sex = field.NewString(tableName, "sex")
	_studentV2.UpdateTime = field.NewTime(tableName, "update_time")
	_studentV2.CreateTime = field.NewTime(tableName, "create_time")
	_studentV2.CreateBy = field.NewString(tableName, "create_by")
	_studentV2.UpdateBy = field.NewString(tableName, "update_by")

	_studentV2.fillFieldMap()

	return _studentV2
}

type studentV2 struct {
	studentV2Do studentV2Do

	ALL        field.Asterisk
	ID         field.Int64  // 主键id
	Name       field.String // 姓名
	Age        field.String // 年龄
	Sex        field.String // 性别
	UpdateTime field.Time   // 更新时间
	CreateTime field.Time   // 创建时间
	CreateBy   field.String // 创建人
	UpdateBy   field.String // 更新人

	fieldMap map[string]field.Expr
}

func (s studentV2) Table(newTableName string) *studentV2 {
	s.studentV2Do.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s studentV2) As(alias string) *studentV2 {
	s.studentV2Do.DO = *(s.studentV2Do.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *studentV2) updateTableName(table string) *studentV2 {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Age = field.NewString(table, "age")
	s.Sex = field.NewString(table, "sex")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.CreateTime = field.NewTime(table, "create_time")
	s.CreateBy = field.NewString(table, "create_by")
	s.UpdateBy = field.NewString(table, "update_by")

	s.fillFieldMap()

	return s
}

func (s *studentV2) WithContext(ctx context.Context) *studentV2Do {
	return s.studentV2Do.WithContext(ctx)
}

func (s studentV2) TableName() string { return s.studentV2Do.TableName() }

func (s studentV2) Alias() string { return s.studentV2Do.Alias() }

func (s studentV2) Columns(cols ...field.Expr) gen.Columns { return s.studentV2Do.Columns(cols...) }

func (s *studentV2) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *studentV2) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["age"] = s.Age
	s.fieldMap["sex"] = s.Sex
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["update_by"] = s.UpdateBy
}

func (s studentV2) clone(db *gorm.DB) studentV2 {
	s.studentV2Do.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s studentV2) replaceDB(db *gorm.DB) studentV2 {
	s.studentV2Do.ReplaceDB(db)
	return s
}

type studentV2Do struct{ gen.DO }

func (s studentV2Do) Debug() *studentV2Do {
	return s.withDO(s.DO.Debug())
}

func (s studentV2Do) WithContext(ctx context.Context) *studentV2Do {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studentV2Do) ReadDB() *studentV2Do {
	return s.Clauses(dbresolver.Read)
}

func (s studentV2Do) WriteDB() *studentV2Do {
	return s.Clauses(dbresolver.Write)
}

func (s studentV2Do) Session(config *gorm.Session) *studentV2Do {
	return s.withDO(s.DO.Session(config))
}

func (s studentV2Do) Clauses(conds ...clause.Expression) *studentV2Do {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studentV2Do) Returning(value interface{}, columns ...string) *studentV2Do {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studentV2Do) Not(conds ...gen.Condition) *studentV2Do {
	return s.withDO(s.DO.Not(conds...))
}

func (s studentV2Do) Or(conds ...gen.Condition) *studentV2Do {
	return s.withDO(s.DO.Or(conds...))
}

func (s studentV2Do) Select(conds ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.Select(conds...))
}

func (s studentV2Do) Where(conds ...gen.Condition) *studentV2Do {
	return s.withDO(s.DO.Where(conds...))
}

func (s studentV2Do) Order(conds ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.Order(conds...))
}

func (s studentV2Do) Distinct(cols ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studentV2Do) Omit(cols ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studentV2Do) Join(table schema.Tabler, on ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studentV2Do) LeftJoin(table schema.Tabler, on ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studentV2Do) RightJoin(table schema.Tabler, on ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studentV2Do) Group(cols ...field.Expr) *studentV2Do {
	return s.withDO(s.DO.Group(cols...))
}

func (s studentV2Do) Having(conds ...gen.Condition) *studentV2Do {
	return s.withDO(s.DO.Having(conds...))
}

func (s studentV2Do) Limit(limit int) *studentV2Do {
	return s.withDO(s.DO.Limit(limit))
}

func (s studentV2Do) Offset(offset int) *studentV2Do {
	return s.withDO(s.DO.Offset(offset))
}

func (s studentV2Do) Scopes(funcs ...func(gen.Dao) gen.Dao) *studentV2Do {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studentV2Do) Unscoped() *studentV2Do {
	return s.withDO(s.DO.Unscoped())
}

func (s studentV2Do) Create(values ...*model.StudentV2) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studentV2Do) CreateInBatches(values []*model.StudentV2, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studentV2Do) Save(values ...*model.StudentV2) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studentV2Do) First() (*model.StudentV2, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentV2), nil
	}
}

func (s studentV2Do) Take() (*model.StudentV2, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentV2), nil
	}
}

func (s studentV2Do) Last() (*model.StudentV2, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentV2), nil
	}
}

func (s studentV2Do) Find() ([]*model.StudentV2, error) {
	result, err := s.DO.Find()
	return result.([]*model.StudentV2), err
}

func (s studentV2Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StudentV2, err error) {
	buf := make([]*model.StudentV2, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studentV2Do) FindInBatches(result *[]*model.StudentV2, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studentV2Do) Attrs(attrs ...field.AssignExpr) *studentV2Do {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studentV2Do) Assign(attrs ...field.AssignExpr) *studentV2Do {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studentV2Do) Joins(fields ...field.RelationField) *studentV2Do {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studentV2Do) Preload(fields ...field.RelationField) *studentV2Do {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studentV2Do) FirstOrInit() (*model.StudentV2, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentV2), nil
	}
}

func (s studentV2Do) FirstOrCreate() (*model.StudentV2, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentV2), nil
	}
}

func (s studentV2Do) FindByPage(offset int, limit int) (result []*model.StudentV2, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s studentV2Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studentV2Do) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studentV2Do) Delete(models ...*model.StudentV2) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studentV2Do) withDO(do gen.Dao) *studentV2Do {
	s.DO = *do.(*gen.DO)
	return s
}
