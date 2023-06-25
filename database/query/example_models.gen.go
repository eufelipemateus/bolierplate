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

	"github.com/eufelipemateus/bolierplate-go/models"
)

func newExampleModel(db *gorm.DB, opts ...gen.DOOption) exampleModel {
	_exampleModel := exampleModel{}

	_exampleModel.exampleModelDo.UseDB(db, opts...)
	_exampleModel.exampleModelDo.UseModel(&models.ExampleModel{})

	tableName := _exampleModel.exampleModelDo.TableName()
	_exampleModel.ALL = field.NewAsterisk(tableName)
	_exampleModel.ID = field.NewUint(tableName, "id")
	_exampleModel.CreatedAt = field.NewTime(tableName, "created_at")
	_exampleModel.UpdatedAt = field.NewTime(tableName, "updated_at")
	_exampleModel.Name = field.NewFloat64(tableName, "name")

	_exampleModel.fillFieldMap()

	return _exampleModel
}

type exampleModel struct {
	exampleModelDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	Name      field.Float64

	fieldMap map[string]field.Expr
}

func (e exampleModel) Table(newTableName string) *exampleModel {
	e.exampleModelDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e exampleModel) As(alias string) *exampleModel {
	e.exampleModelDo.DO = *(e.exampleModelDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *exampleModel) updateTableName(table string) *exampleModel {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewUint(table, "id")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.Name = field.NewFloat64(table, "name")

	e.fillFieldMap()

	return e
}

func (e *exampleModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *exampleModel) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 4)
	e.fieldMap["id"] = e.ID
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["name"] = e.Name
}

func (e exampleModel) clone(db *gorm.DB) exampleModel {
	e.exampleModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e exampleModel) replaceDB(db *gorm.DB) exampleModel {
	e.exampleModelDo.ReplaceDB(db)
	return e
}

type exampleModelDo struct{ gen.DO }

type IExampleModelDo interface {
	gen.SubQuery
	Debug() IExampleModelDo
	WithContext(ctx context.Context) IExampleModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IExampleModelDo
	WriteDB() IExampleModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IExampleModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IExampleModelDo
	Not(conds ...gen.Condition) IExampleModelDo
	Or(conds ...gen.Condition) IExampleModelDo
	Select(conds ...field.Expr) IExampleModelDo
	Where(conds ...gen.Condition) IExampleModelDo
	Order(conds ...field.Expr) IExampleModelDo
	Distinct(cols ...field.Expr) IExampleModelDo
	Omit(cols ...field.Expr) IExampleModelDo
	Join(table schema.Tabler, on ...field.Expr) IExampleModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IExampleModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IExampleModelDo
	Group(cols ...field.Expr) IExampleModelDo
	Having(conds ...gen.Condition) IExampleModelDo
	Limit(limit int) IExampleModelDo
	Offset(offset int) IExampleModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IExampleModelDo
	Unscoped() IExampleModelDo
	Create(values ...*models.ExampleModel) error
	CreateInBatches(values []*models.ExampleModel, batchSize int) error
	Save(values ...*models.ExampleModel) error
	First() (*models.ExampleModel, error)
	Take() (*models.ExampleModel, error)
	Last() (*models.ExampleModel, error)
	Find() ([]*models.ExampleModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ExampleModel, err error)
	FindInBatches(result *[]*models.ExampleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ExampleModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IExampleModelDo
	Assign(attrs ...field.AssignExpr) IExampleModelDo
	Joins(fields ...field.RelationField) IExampleModelDo
	Preload(fields ...field.RelationField) IExampleModelDo
	FirstOrInit() (*models.ExampleModel, error)
	FirstOrCreate() (*models.ExampleModel, error)
	FindByPage(offset int, limit int) (result []*models.ExampleModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IExampleModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e exampleModelDo) Debug() IExampleModelDo {
	return e.withDO(e.DO.Debug())
}

func (e exampleModelDo) WithContext(ctx context.Context) IExampleModelDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e exampleModelDo) ReadDB() IExampleModelDo {
	return e.Clauses(dbresolver.Read)
}

func (e exampleModelDo) WriteDB() IExampleModelDo {
	return e.Clauses(dbresolver.Write)
}

func (e exampleModelDo) Session(config *gorm.Session) IExampleModelDo {
	return e.withDO(e.DO.Session(config))
}

func (e exampleModelDo) Clauses(conds ...clause.Expression) IExampleModelDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e exampleModelDo) Returning(value interface{}, columns ...string) IExampleModelDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e exampleModelDo) Not(conds ...gen.Condition) IExampleModelDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e exampleModelDo) Or(conds ...gen.Condition) IExampleModelDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e exampleModelDo) Select(conds ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e exampleModelDo) Where(conds ...gen.Condition) IExampleModelDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e exampleModelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IExampleModelDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e exampleModelDo) Order(conds ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e exampleModelDo) Distinct(cols ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e exampleModelDo) Omit(cols ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e exampleModelDo) Join(table schema.Tabler, on ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e exampleModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e exampleModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e exampleModelDo) Group(cols ...field.Expr) IExampleModelDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e exampleModelDo) Having(conds ...gen.Condition) IExampleModelDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e exampleModelDo) Limit(limit int) IExampleModelDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e exampleModelDo) Offset(offset int) IExampleModelDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e exampleModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IExampleModelDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e exampleModelDo) Unscoped() IExampleModelDo {
	return e.withDO(e.DO.Unscoped())
}

func (e exampleModelDo) Create(values ...*models.ExampleModel) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e exampleModelDo) CreateInBatches(values []*models.ExampleModel, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e exampleModelDo) Save(values ...*models.ExampleModel) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e exampleModelDo) First() (*models.ExampleModel, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ExampleModel), nil
	}
}

func (e exampleModelDo) Take() (*models.ExampleModel, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ExampleModel), nil
	}
}

func (e exampleModelDo) Last() (*models.ExampleModel, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ExampleModel), nil
	}
}

func (e exampleModelDo) Find() ([]*models.ExampleModel, error) {
	result, err := e.DO.Find()
	return result.([]*models.ExampleModel), err
}

func (e exampleModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ExampleModel, err error) {
	buf := make([]*models.ExampleModel, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e exampleModelDo) FindInBatches(result *[]*models.ExampleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e exampleModelDo) Attrs(attrs ...field.AssignExpr) IExampleModelDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e exampleModelDo) Assign(attrs ...field.AssignExpr) IExampleModelDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e exampleModelDo) Joins(fields ...field.RelationField) IExampleModelDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e exampleModelDo) Preload(fields ...field.RelationField) IExampleModelDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e exampleModelDo) FirstOrInit() (*models.ExampleModel, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ExampleModel), nil
	}
}

func (e exampleModelDo) FirstOrCreate() (*models.ExampleModel, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ExampleModel), nil
	}
}

func (e exampleModelDo) FindByPage(offset int, limit int) (result []*models.ExampleModel, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e exampleModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e exampleModelDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e exampleModelDo) Delete(models ...*models.ExampleModel) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *exampleModelDo) withDO(do gen.Dao) *exampleModelDo {
	e.DO = *do.(*gen.DO)
	return e
}
