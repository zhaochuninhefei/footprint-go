// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package operator

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gitee.com/zhaochuninhefei/footprint-go/db/model"
)

func newBroodDbVersionCtl(db *gorm.DB, opts ...gen.DOOption) broodDbVersionCtl {
	_broodDbVersionCtl := broodDbVersionCtl{}

	_broodDbVersionCtl.broodDbVersionCtlDo.UseDB(db, opts...)
	_broodDbVersionCtl.broodDbVersionCtlDo.UseModel(&model.BroodDbVersionCtl{})

	tableName := _broodDbVersionCtl.broodDbVersionCtlDo.TableName()
	_broodDbVersionCtl.ALL = field.NewAsterisk(tableName)
	_broodDbVersionCtl.ID = field.NewInt64(tableName, "id")
	_broodDbVersionCtl.BusinessSpace = field.NewString(tableName, "business_space")
	_broodDbVersionCtl.MajorVersion = field.NewInt64(tableName, "major_version")
	_broodDbVersionCtl.MinorVersion = field.NewInt64(tableName, "minor_version")
	_broodDbVersionCtl.PatchVersion = field.NewInt64(tableName, "patch_version")
	_broodDbVersionCtl.ExtendVersion = field.NewInt64(tableName, "extend_version")
	_broodDbVersionCtl.Version = field.NewString(tableName, "version")
	_broodDbVersionCtl.CustomName = field.NewString(tableName, "custom_name")
	_broodDbVersionCtl.VersionType = field.NewString(tableName, "version_type")
	_broodDbVersionCtl.ScriptFileName = field.NewString(tableName, "script_file_name")
	_broodDbVersionCtl.ScriptDigestHex = field.NewString(tableName, "script_digest_hex")
	_broodDbVersionCtl.Success = field.NewInt64(tableName, "success")
	_broodDbVersionCtl.ExecutionTime = field.NewInt64(tableName, "execution_time")
	_broodDbVersionCtl.InstallTime = field.NewString(tableName, "install_time")
	_broodDbVersionCtl.InstallUser = field.NewString(tableName, "install_user")

	_broodDbVersionCtl.fillFieldMap()

	return _broodDbVersionCtl
}

type broodDbVersionCtl struct {
	broodDbVersionCtlDo

	ALL             field.Asterisk
	ID              field.Int64  // 数据库版本ID
	BusinessSpace   field.String // 业务空间
	MajorVersion    field.Int64  // 主版本号
	MinorVersion    field.Int64  // 次版本号
	PatchVersion    field.Int64  // 补丁版本号
	ExtendVersion   field.Int64  // 扩展版本号
	Version         field.String // 版本号,V[major].[minor].[patch].[extend_version]
	CustomName      field.String // 脚本自定义名称
	VersionType     field.String // 版本类型:SQL/BaseLine
	ScriptFileName  field.String // 脚本文件名
	ScriptDigestHex field.String // 脚本内容摘要(16进制)
	Success         field.Int64  // 是否执行成功
	ExecutionTime   field.Int64  // 脚本安装耗时
	InstallTime     field.String // 脚本安装时间,格式:[yyyy-MM-dd HH:mm:ss]
	InstallUser     field.String // 脚本安装用户

	fieldMap map[string]field.Expr
}

func (b broodDbVersionCtl) Table(newTableName string) *broodDbVersionCtl {
	b.broodDbVersionCtlDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b broodDbVersionCtl) As(alias string) *broodDbVersionCtl {
	b.broodDbVersionCtlDo.DO = *(b.broodDbVersionCtlDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *broodDbVersionCtl) updateTableName(table string) *broodDbVersionCtl {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.BusinessSpace = field.NewString(table, "business_space")
	b.MajorVersion = field.NewInt64(table, "major_version")
	b.MinorVersion = field.NewInt64(table, "minor_version")
	b.PatchVersion = field.NewInt64(table, "patch_version")
	b.ExtendVersion = field.NewInt64(table, "extend_version")
	b.Version = field.NewString(table, "version")
	b.CustomName = field.NewString(table, "custom_name")
	b.VersionType = field.NewString(table, "version_type")
	b.ScriptFileName = field.NewString(table, "script_file_name")
	b.ScriptDigestHex = field.NewString(table, "script_digest_hex")
	b.Success = field.NewInt64(table, "success")
	b.ExecutionTime = field.NewInt64(table, "execution_time")
	b.InstallTime = field.NewString(table, "install_time")
	b.InstallUser = field.NewString(table, "install_user")

	b.fillFieldMap()

	return b
}

func (b *broodDbVersionCtl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *broodDbVersionCtl) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 15)
	b.fieldMap["id"] = b.ID
	b.fieldMap["business_space"] = b.BusinessSpace
	b.fieldMap["major_version"] = b.MajorVersion
	b.fieldMap["minor_version"] = b.MinorVersion
	b.fieldMap["patch_version"] = b.PatchVersion
	b.fieldMap["extend_version"] = b.ExtendVersion
	b.fieldMap["version"] = b.Version
	b.fieldMap["custom_name"] = b.CustomName
	b.fieldMap["version_type"] = b.VersionType
	b.fieldMap["script_file_name"] = b.ScriptFileName
	b.fieldMap["script_digest_hex"] = b.ScriptDigestHex
	b.fieldMap["success"] = b.Success
	b.fieldMap["execution_time"] = b.ExecutionTime
	b.fieldMap["install_time"] = b.InstallTime
	b.fieldMap["install_user"] = b.InstallUser
}

func (b broodDbVersionCtl) clone(db *gorm.DB) broodDbVersionCtl {
	b.broodDbVersionCtlDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b broodDbVersionCtl) replaceDB(db *gorm.DB) broodDbVersionCtl {
	b.broodDbVersionCtlDo.ReplaceDB(db)
	return b
}

type broodDbVersionCtlDo struct{ gen.DO }

type IBroodDbVersionCtlDo interface {
	gen.SubQuery
	Debug() IBroodDbVersionCtlDo
	WithContext(ctx context.Context) IBroodDbVersionCtlDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBroodDbVersionCtlDo
	WriteDB() IBroodDbVersionCtlDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBroodDbVersionCtlDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBroodDbVersionCtlDo
	Not(conds ...gen.Condition) IBroodDbVersionCtlDo
	Or(conds ...gen.Condition) IBroodDbVersionCtlDo
	Select(conds ...field.Expr) IBroodDbVersionCtlDo
	Where(conds ...gen.Condition) IBroodDbVersionCtlDo
	Order(conds ...field.Expr) IBroodDbVersionCtlDo
	Distinct(cols ...field.Expr) IBroodDbVersionCtlDo
	Omit(cols ...field.Expr) IBroodDbVersionCtlDo
	Join(table schema.Tabler, on ...field.Expr) IBroodDbVersionCtlDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBroodDbVersionCtlDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBroodDbVersionCtlDo
	Group(cols ...field.Expr) IBroodDbVersionCtlDo
	Having(conds ...gen.Condition) IBroodDbVersionCtlDo
	Limit(limit int) IBroodDbVersionCtlDo
	Offset(offset int) IBroodDbVersionCtlDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBroodDbVersionCtlDo
	Unscoped() IBroodDbVersionCtlDo
	Create(values ...*model.BroodDbVersionCtl) error
	CreateInBatches(values []*model.BroodDbVersionCtl, batchSize int) error
	Save(values ...*model.BroodDbVersionCtl) error
	First() (*model.BroodDbVersionCtl, error)
	Take() (*model.BroodDbVersionCtl, error)
	Last() (*model.BroodDbVersionCtl, error)
	Find() ([]*model.BroodDbVersionCtl, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BroodDbVersionCtl, err error)
	FindInBatches(result *[]*model.BroodDbVersionCtl, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BroodDbVersionCtl) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBroodDbVersionCtlDo
	Assign(attrs ...field.AssignExpr) IBroodDbVersionCtlDo
	Joins(fields ...field.RelationField) IBroodDbVersionCtlDo
	Preload(fields ...field.RelationField) IBroodDbVersionCtlDo
	FirstOrInit() (*model.BroodDbVersionCtl, error)
	FirstOrCreate() (*model.BroodDbVersionCtl, error)
	FindByPage(offset int, limit int) (result []*model.BroodDbVersionCtl, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBroodDbVersionCtlDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b broodDbVersionCtlDo) Debug() IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Debug())
}

func (b broodDbVersionCtlDo) WithContext(ctx context.Context) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b broodDbVersionCtlDo) ReadDB() IBroodDbVersionCtlDo {
	return b.Clauses(dbresolver.Read)
}

func (b broodDbVersionCtlDo) WriteDB() IBroodDbVersionCtlDo {
	return b.Clauses(dbresolver.Write)
}

func (b broodDbVersionCtlDo) Session(config *gorm.Session) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Session(config))
}

func (b broodDbVersionCtlDo) Clauses(conds ...clause.Expression) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b broodDbVersionCtlDo) Returning(value interface{}, columns ...string) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b broodDbVersionCtlDo) Not(conds ...gen.Condition) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b broodDbVersionCtlDo) Or(conds ...gen.Condition) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b broodDbVersionCtlDo) Select(conds ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b broodDbVersionCtlDo) Where(conds ...gen.Condition) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b broodDbVersionCtlDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBroodDbVersionCtlDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b broodDbVersionCtlDo) Order(conds ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b broodDbVersionCtlDo) Distinct(cols ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b broodDbVersionCtlDo) Omit(cols ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b broodDbVersionCtlDo) Join(table schema.Tabler, on ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b broodDbVersionCtlDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b broodDbVersionCtlDo) RightJoin(table schema.Tabler, on ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b broodDbVersionCtlDo) Group(cols ...field.Expr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b broodDbVersionCtlDo) Having(conds ...gen.Condition) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b broodDbVersionCtlDo) Limit(limit int) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b broodDbVersionCtlDo) Offset(offset int) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b broodDbVersionCtlDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b broodDbVersionCtlDo) Unscoped() IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Unscoped())
}

func (b broodDbVersionCtlDo) Create(values ...*model.BroodDbVersionCtl) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b broodDbVersionCtlDo) CreateInBatches(values []*model.BroodDbVersionCtl, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b broodDbVersionCtlDo) Save(values ...*model.BroodDbVersionCtl) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b broodDbVersionCtlDo) First() (*model.BroodDbVersionCtl, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BroodDbVersionCtl), nil
	}
}

func (b broodDbVersionCtlDo) Take() (*model.BroodDbVersionCtl, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BroodDbVersionCtl), nil
	}
}

func (b broodDbVersionCtlDo) Last() (*model.BroodDbVersionCtl, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BroodDbVersionCtl), nil
	}
}

func (b broodDbVersionCtlDo) Find() ([]*model.BroodDbVersionCtl, error) {
	result, err := b.DO.Find()
	return result.([]*model.BroodDbVersionCtl), err
}

func (b broodDbVersionCtlDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BroodDbVersionCtl, err error) {
	buf := make([]*model.BroodDbVersionCtl, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b broodDbVersionCtlDo) FindInBatches(result *[]*model.BroodDbVersionCtl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b broodDbVersionCtlDo) Attrs(attrs ...field.AssignExpr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b broodDbVersionCtlDo) Assign(attrs ...field.AssignExpr) IBroodDbVersionCtlDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b broodDbVersionCtlDo) Joins(fields ...field.RelationField) IBroodDbVersionCtlDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b broodDbVersionCtlDo) Preload(fields ...field.RelationField) IBroodDbVersionCtlDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b broodDbVersionCtlDo) FirstOrInit() (*model.BroodDbVersionCtl, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BroodDbVersionCtl), nil
	}
}

func (b broodDbVersionCtlDo) FirstOrCreate() (*model.BroodDbVersionCtl, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BroodDbVersionCtl), nil
	}
}

func (b broodDbVersionCtlDo) FindByPage(offset int, limit int) (result []*model.BroodDbVersionCtl, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b broodDbVersionCtlDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b broodDbVersionCtlDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b broodDbVersionCtlDo) Delete(models ...*model.BroodDbVersionCtl) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *broodDbVersionCtlDo) withDO(do gen.Dao) *broodDbVersionCtlDo {
	b.DO = *do.(*gen.DO)
	return b
}
