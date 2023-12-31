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

func newPhoto(db *gorm.DB, opts ...gen.DOOption) photo {
	_photo := photo{}

	_photo.photoDo.UseDB(db, opts...)
	_photo.photoDo.UseModel(&po.Photo{})

	tableName := _photo.photoDo.TableName()
	_photo.ALL = field.NewAsterisk(tableName)
	_photo.ID = field.NewInt64(tableName, "Id")
	_photo.Desc = field.NewString(tableName, "desc")
	_photo.URL = field.NewString(tableName, "url")
	_photo.UpdatedTime = field.NewString(tableName, "updated_time")
	_photo.CreatedTime = field.NewString(tableName, "created_time")
	_photo.LikeTimes = field.NewInt64(tableName, "like_times")

	_photo.fillFieldMap()

	return _photo
}

type photo struct {
	photoDo photoDo

	ALL         field.Asterisk
	ID          field.Int64
	Desc        field.String
	URL         field.String
	UpdatedTime field.String
	CreatedTime field.String
	LikeTimes   field.Int64

	fieldMap map[string]field.Expr
}

func (p photo) Table(newTableName string) *photo {
	p.photoDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p photo) As(alias string) *photo {
	p.photoDo.DO = *(p.photoDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *photo) updateTableName(table string) *photo {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "Id")
	p.Desc = field.NewString(table, "desc")
	p.URL = field.NewString(table, "url")
	p.UpdatedTime = field.NewString(table, "updated_time")
	p.CreatedTime = field.NewString(table, "created_time")
	p.LikeTimes = field.NewInt64(table, "like_times")

	p.fillFieldMap()

	return p
}

func (p *photo) WithContext(ctx context.Context) IPhotoDo { return p.photoDo.WithContext(ctx) }

func (p photo) TableName() string { return p.photoDo.TableName() }

func (p photo) Alias() string { return p.photoDo.Alias() }

func (p photo) Columns(cols ...field.Expr) gen.Columns { return p.photoDo.Columns(cols...) }

func (p *photo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *photo) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["Id"] = p.ID
	p.fieldMap["desc"] = p.Desc
	p.fieldMap["url"] = p.URL
	p.fieldMap["updated_time"] = p.UpdatedTime
	p.fieldMap["created_time"] = p.CreatedTime
	p.fieldMap["like_times"] = p.LikeTimes
}

func (p photo) clone(db *gorm.DB) photo {
	p.photoDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p photo) replaceDB(db *gorm.DB) photo {
	p.photoDo.ReplaceDB(db)
	return p
}

type photoDo struct{ gen.DO }

type IPhotoDo interface {
	gen.SubQuery
	Debug() IPhotoDo
	WithContext(ctx context.Context) IPhotoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPhotoDo
	WriteDB() IPhotoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPhotoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPhotoDo
	Not(conds ...gen.Condition) IPhotoDo
	Or(conds ...gen.Condition) IPhotoDo
	Select(conds ...field.Expr) IPhotoDo
	Where(conds ...gen.Condition) IPhotoDo
	Order(conds ...field.Expr) IPhotoDo
	Distinct(cols ...field.Expr) IPhotoDo
	Omit(cols ...field.Expr) IPhotoDo
	Join(table schema.Tabler, on ...field.Expr) IPhotoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPhotoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPhotoDo
	Group(cols ...field.Expr) IPhotoDo
	Having(conds ...gen.Condition) IPhotoDo
	Limit(limit int) IPhotoDo
	Offset(offset int) IPhotoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPhotoDo
	Unscoped() IPhotoDo
	Create(values ...*po.Photo) error
	CreateInBatches(values []*po.Photo, batchSize int) error
	Save(values ...*po.Photo) error
	First() (*po.Photo, error)
	Take() (*po.Photo, error)
	Last() (*po.Photo, error)
	Find() ([]*po.Photo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.Photo, err error)
	FindInBatches(result *[]*po.Photo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*po.Photo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPhotoDo
	Assign(attrs ...field.AssignExpr) IPhotoDo
	Joins(fields ...field.RelationField) IPhotoDo
	Preload(fields ...field.RelationField) IPhotoDo
	FirstOrInit() (*po.Photo, error)
	FirstOrCreate() (*po.Photo, error)
	FindByPage(offset int, limit int) (result []*po.Photo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPhotoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p photoDo) Debug() IPhotoDo {
	return p.withDO(p.DO.Debug())
}

func (p photoDo) WithContext(ctx context.Context) IPhotoDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p photoDo) ReadDB() IPhotoDo {
	return p.Clauses(dbresolver.Read)
}

func (p photoDo) WriteDB() IPhotoDo {
	return p.Clauses(dbresolver.Write)
}

func (p photoDo) Session(config *gorm.Session) IPhotoDo {
	return p.withDO(p.DO.Session(config))
}

func (p photoDo) Clauses(conds ...clause.Expression) IPhotoDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p photoDo) Returning(value interface{}, columns ...string) IPhotoDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p photoDo) Not(conds ...gen.Condition) IPhotoDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p photoDo) Or(conds ...gen.Condition) IPhotoDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p photoDo) Select(conds ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p photoDo) Where(conds ...gen.Condition) IPhotoDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p photoDo) Order(conds ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p photoDo) Distinct(cols ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p photoDo) Omit(cols ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p photoDo) Join(table schema.Tabler, on ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p photoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p photoDo) RightJoin(table schema.Tabler, on ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p photoDo) Group(cols ...field.Expr) IPhotoDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p photoDo) Having(conds ...gen.Condition) IPhotoDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p photoDo) Limit(limit int) IPhotoDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p photoDo) Offset(offset int) IPhotoDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p photoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPhotoDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p photoDo) Unscoped() IPhotoDo {
	return p.withDO(p.DO.Unscoped())
}

func (p photoDo) Create(values ...*po.Photo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p photoDo) CreateInBatches(values []*po.Photo, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p photoDo) Save(values ...*po.Photo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p photoDo) First() (*po.Photo, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.Photo), nil
	}
}

func (p photoDo) Take() (*po.Photo, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.Photo), nil
	}
}

func (p photoDo) Last() (*po.Photo, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.Photo), nil
	}
}

func (p photoDo) Find() ([]*po.Photo, error) {
	result, err := p.DO.Find()
	return result.([]*po.Photo), err
}

func (p photoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.Photo, err error) {
	buf := make([]*po.Photo, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p photoDo) FindInBatches(result *[]*po.Photo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p photoDo) Attrs(attrs ...field.AssignExpr) IPhotoDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p photoDo) Assign(attrs ...field.AssignExpr) IPhotoDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p photoDo) Joins(fields ...field.RelationField) IPhotoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p photoDo) Preload(fields ...field.RelationField) IPhotoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p photoDo) FirstOrInit() (*po.Photo, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.Photo), nil
	}
}

func (p photoDo) FirstOrCreate() (*po.Photo, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.Photo), nil
	}
}

func (p photoDo) FindByPage(offset int, limit int) (result []*po.Photo, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p photoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p photoDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p photoDo) Delete(models ...*po.Photo) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *photoDo) withDO(do gen.Dao) *photoDo {
	p.DO = *do.(*gen.DO)
	return p
}
