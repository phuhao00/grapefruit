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

	"grapefruit/internal/domain/po"
)

func newJob(db *gorm.DB, opts ...gen.DOOption) job {
	_job := job{}

	_job.jobDo.UseDB(db, opts...)
	_job.jobDo.UseModel(&po.Job{})

	tableName := _job.jobDo.TableName()
	_job.ALL = field.NewAsterisk(tableName)
	_job.ID = field.NewInt32(tableName, "id")
	_job.Name = field.NewString(tableName, "name")
	_job.Desc = field.NewString(tableName, "desc")
	_job.MinSalary = field.NewFloat64(tableName, "min_salary")
	_job.MaxSalary = field.NewFloat64(tableName, "max_salary")
	_job.CompanyID = field.NewInt32(tableName, "company_id")
	_job.Require = field.NewString(tableName, "require")
	_job.Publiser = field.NewInt32(tableName, "publiser")

	_job.fillFieldMap()

	return _job
}

type job struct {
	jobDo jobDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String
	Desc      field.String  // 职位描述
	MinSalary field.Float64 // 最低薪资
	MaxSalary field.Float64 // 最高薪资
	CompanyID field.Int32
	Require   field.String // 职位要求
	Publiser  field.Int32  // 发布职位的招聘者ID

	fieldMap map[string]field.Expr
}

func (j job) Table(newTableName string) *job {
	j.jobDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j job) As(alias string) *job {
	j.jobDo.DO = *(j.jobDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *job) updateTableName(table string) *job {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewInt32(table, "id")
	j.Name = field.NewString(table, "name")
	j.Desc = field.NewString(table, "desc")
	j.MinSalary = field.NewFloat64(table, "min_salary")
	j.MaxSalary = field.NewFloat64(table, "max_salary")
	j.CompanyID = field.NewInt32(table, "company_id")
	j.Require = field.NewString(table, "require")
	j.Publiser = field.NewInt32(table, "publiser")

	j.fillFieldMap()

	return j
}

func (j *job) WithContext(ctx context.Context) IJobDo { return j.jobDo.WithContext(ctx) }

func (j job) TableName() string { return j.jobDo.TableName() }

func (j job) Alias() string { return j.jobDo.Alias() }

func (j job) Columns(cols ...field.Expr) gen.Columns { return j.jobDo.Columns(cols...) }

func (j *job) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *job) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 8)
	j.fieldMap["id"] = j.ID
	j.fieldMap["name"] = j.Name
	j.fieldMap["desc"] = j.Desc
	j.fieldMap["min_salary"] = j.MinSalary
	j.fieldMap["max_salary"] = j.MaxSalary
	j.fieldMap["company_id"] = j.CompanyID
	j.fieldMap["require"] = j.Require
	j.fieldMap["publiser"] = j.Publiser
}

func (j job) clone(db *gorm.DB) job {
	j.jobDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j job) replaceDB(db *gorm.DB) job {
	j.jobDo.ReplaceDB(db)
	return j
}

type jobDo struct{ gen.DO }

type IJobDo interface {
	gen.SubQuery
	Debug() IJobDo
	WithContext(ctx context.Context) IJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJobDo
	WriteDB() IJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJobDo
	Not(conds ...gen.Condition) IJobDo
	Or(conds ...gen.Condition) IJobDo
	Select(conds ...field.Expr) IJobDo
	Where(conds ...gen.Condition) IJobDo
	Order(conds ...field.Expr) IJobDo
	Distinct(cols ...field.Expr) IJobDo
	Omit(cols ...field.Expr) IJobDo
	Join(table schema.Tabler, on ...field.Expr) IJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJobDo
	Group(cols ...field.Expr) IJobDo
	Having(conds ...gen.Condition) IJobDo
	Limit(limit int) IJobDo
	Offset(offset int) IJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJobDo
	Unscoped() IJobDo
	Create(values ...*po.Job) error
	CreateInBatches(values []*po.Job, batchSize int) error
	Save(values ...*po.Job) error
	First() (*po.Job, error)
	Take() (*po.Job, error)
	Last() (*po.Job, error)
	Find() ([]*po.Job, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.Job, err error)
	FindInBatches(result *[]*po.Job, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*po.Job) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJobDo
	Assign(attrs ...field.AssignExpr) IJobDo
	Joins(fields ...field.RelationField) IJobDo
	Preload(fields ...field.RelationField) IJobDo
	FirstOrInit() (*po.Job, error)
	FirstOrCreate() (*po.Job, error)
	FindByPage(offset int, limit int) (result []*po.Job, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jobDo) Debug() IJobDo {
	return j.withDO(j.DO.Debug())
}

func (j jobDo) WithContext(ctx context.Context) IJobDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobDo) ReadDB() IJobDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobDo) WriteDB() IJobDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobDo) Session(config *gorm.Session) IJobDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobDo) Clauses(conds ...clause.Expression) IJobDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobDo) Returning(value interface{}, columns ...string) IJobDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobDo) Not(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobDo) Or(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobDo) Select(conds ...field.Expr) IJobDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobDo) Where(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobDo) Order(conds ...field.Expr) IJobDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobDo) Distinct(cols ...field.Expr) IJobDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobDo) Omit(cols ...field.Expr) IJobDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobDo) Join(table schema.Tabler, on ...field.Expr) IJobDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJobDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobDo) RightJoin(table schema.Tabler, on ...field.Expr) IJobDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobDo) Group(cols ...field.Expr) IJobDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobDo) Having(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobDo) Limit(limit int) IJobDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobDo) Offset(offset int) IJobDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJobDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobDo) Unscoped() IJobDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobDo) Create(values ...*po.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobDo) CreateInBatches(values []*po.Job, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobDo) Save(values ...*po.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobDo) First() (*po.Job, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.Job), nil
	}
}

func (j jobDo) Take() (*po.Job, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.Job), nil
	}
}

func (j jobDo) Last() (*po.Job, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.Job), nil
	}
}

func (j jobDo) Find() ([]*po.Job, error) {
	result, err := j.DO.Find()
	return result.([]*po.Job), err
}

func (j jobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.Job, err error) {
	buf := make([]*po.Job, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobDo) FindInBatches(result *[]*po.Job, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobDo) Attrs(attrs ...field.AssignExpr) IJobDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobDo) Assign(attrs ...field.AssignExpr) IJobDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobDo) Joins(fields ...field.RelationField) IJobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobDo) Preload(fields ...field.RelationField) IJobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobDo) FirstOrInit() (*po.Job, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.Job), nil
	}
}

func (j jobDo) FirstOrCreate() (*po.Job, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.Job), nil
	}
}

func (j jobDo) FindByPage(offset int, limit int) (result []*po.Job, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobDo) Delete(models ...*po.Job) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobDo) withDO(do gen.Dao) *jobDo {
	j.DO = *do.(*gen.DO)
	return j
}
