// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/mitsuha/stork/repository/model"
)

func newArtist(db *gorm.DB, opts ...gen.DOOption) artist {
	_artist := artist{}

	_artist.artistDo.UseDB(db, opts...)
	_artist.artistDo.UseModel(&model.Artist{})

	tableName := _artist.artistDo.TableName()
	_artist.ALL = field.NewAsterisk(tableName)
	_artist.ID = field.NewInt(tableName, "id")
	_artist.Name = field.NewString(tableName, "name")
	_artist.Image = field.NewString(tableName, "image")
	_artist.CreatedAt = field.NewTime(tableName, "created_at")
	_artist.UpdatedAt = field.NewTime(tableName, "updated_at")

	_artist.fillFieldMap()

	return _artist
}

type artist struct {
	artistDo artistDo

	ALL       field.Asterisk
	ID        field.Int
	Name      field.String
	Image     field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a artist) Table(newTableName string) *artist {
	a.artistDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a artist) As(alias string) *artist {
	a.artistDo.DO = *(a.artistDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *artist) updateTableName(table string) *artist {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.Name = field.NewString(table, "name")
	a.Image = field.NewString(table, "image")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *artist) WithContext(ctx context.Context) IArtistDo { return a.artistDo.WithContext(ctx) }

func (a artist) TableName() string { return a.artistDo.TableName() }

func (a artist) Alias() string { return a.artistDo.Alias() }

func (a artist) Columns(cols ...field.Expr) gen.Columns { return a.artistDo.Columns(cols...) }

func (a *artist) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *artist) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["image"] = a.Image
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a artist) clone(db *gorm.DB) artist {
	a.artistDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a artist) replaceDB(db *gorm.DB) artist {
	a.artistDo.ReplaceDB(db)
	return a
}

type artistDo struct{ gen.DO }

type IArtistDo interface {
	gen.SubQuery
	Debug() IArtistDo
	WithContext(ctx context.Context) IArtistDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArtistDo
	WriteDB() IArtistDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArtistDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArtistDo
	Not(conds ...gen.Condition) IArtistDo
	Or(conds ...gen.Condition) IArtistDo
	Select(conds ...field.Expr) IArtistDo
	Where(conds ...gen.Condition) IArtistDo
	Order(conds ...field.Expr) IArtistDo
	Distinct(cols ...field.Expr) IArtistDo
	Omit(cols ...field.Expr) IArtistDo
	Join(table schema.Tabler, on ...field.Expr) IArtistDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArtistDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArtistDo
	Group(cols ...field.Expr) IArtistDo
	Having(conds ...gen.Condition) IArtistDo
	Limit(limit int) IArtistDo
	Offset(offset int) IArtistDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArtistDo
	Unscoped() IArtistDo
	Create(values ...*model.Artist) error
	CreateInBatches(values []*model.Artist, batchSize int) error
	Save(values ...*model.Artist) error
	First() (*model.Artist, error)
	Take() (*model.Artist, error)
	Last() (*model.Artist, error)
	Find() ([]*model.Artist, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Artist, err error)
	FindInBatches(result *[]*model.Artist, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Artist) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArtistDo
	Assign(attrs ...field.AssignExpr) IArtistDo
	Joins(fields ...field.RelationField) IArtistDo
	Preload(fields ...field.RelationField) IArtistDo
	FirstOrInit() (*model.Artist, error)
	FirstOrCreate() (*model.Artist, error)
	FindByPage(offset int, limit int) (result []*model.Artist, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArtistDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindAll() (result []*model.Artist, err error)
	FindByID(id int) (result *model.Artist, err error)
	FindByUserID(uid int) (result []*model.Artist, err error)
	MostPlayed(uid int, limit int) (result []*model.Artist, err error)
}

// FindAll SELECT * FROM @@table
func (a artistDo) FindAll() (result []*model.Artist, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM artists ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// FindByID SELECT * FROM @@table WHERE id = @id
func (a artistDo) FindByID(id int) (result *model.Artist, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM artists WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// FindByUserID SELECT * FROM @@table WHERE user_id = @uid
func (a artistDo) FindByUserID(uid int) (result []*model.Artist, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	generateSQL.WriteString("SELECT * FROM artists WHERE user_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// MostPlayed SELECT @@table .*
// FROM @@table
// LEFT JOIN songs ON @@table .id = songs.artist_id
// LEFT JOIN interactions ON interactions.song_id = songs.id AND interactions.user_id = @uid
// GROUP BY artists.id, play_count, artists.name, artists.image, artists.created_at, artists.updated_at
// ORDER BY play_count DESC
// LIMIT @limit
func (a artistDo) MostPlayed(uid int, limit int) (result []*model.Artist, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	params = append(params, limit)
	generateSQL.WriteString("SELECT artists .* FROM artists LEFT JOIN songs ON artists .id = songs.artist_id LEFT JOIN interactions ON interactions.song_id = songs.id AND interactions.user_id = ? GROUP BY artists.id, play_count, artists.name, artists.image, artists.created_at, artists.updated_at ORDER BY play_count DESC LIMIT ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a artistDo) Debug() IArtistDo {
	return a.withDO(a.DO.Debug())
}

func (a artistDo) WithContext(ctx context.Context) IArtistDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a artistDo) ReadDB() IArtistDo {
	return a.Clauses(dbresolver.Read)
}

func (a artistDo) WriteDB() IArtistDo {
	return a.Clauses(dbresolver.Write)
}

func (a artistDo) Session(config *gorm.Session) IArtistDo {
	return a.withDO(a.DO.Session(config))
}

func (a artistDo) Clauses(conds ...clause.Expression) IArtistDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a artistDo) Returning(value interface{}, columns ...string) IArtistDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a artistDo) Not(conds ...gen.Condition) IArtistDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a artistDo) Or(conds ...gen.Condition) IArtistDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a artistDo) Select(conds ...field.Expr) IArtistDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a artistDo) Where(conds ...gen.Condition) IArtistDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a artistDo) Order(conds ...field.Expr) IArtistDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a artistDo) Distinct(cols ...field.Expr) IArtistDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a artistDo) Omit(cols ...field.Expr) IArtistDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a artistDo) Join(table schema.Tabler, on ...field.Expr) IArtistDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a artistDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArtistDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a artistDo) RightJoin(table schema.Tabler, on ...field.Expr) IArtistDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a artistDo) Group(cols ...field.Expr) IArtistDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a artistDo) Having(conds ...gen.Condition) IArtistDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a artistDo) Limit(limit int) IArtistDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a artistDo) Offset(offset int) IArtistDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a artistDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArtistDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a artistDo) Unscoped() IArtistDo {
	return a.withDO(a.DO.Unscoped())
}

func (a artistDo) Create(values ...*model.Artist) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a artistDo) CreateInBatches(values []*model.Artist, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a artistDo) Save(values ...*model.Artist) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a artistDo) First() (*model.Artist, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Artist), nil
	}
}

func (a artistDo) Take() (*model.Artist, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Artist), nil
	}
}

func (a artistDo) Last() (*model.Artist, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Artist), nil
	}
}

func (a artistDo) Find() ([]*model.Artist, error) {
	result, err := a.DO.Find()
	return result.([]*model.Artist), err
}

func (a artistDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Artist, err error) {
	buf := make([]*model.Artist, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a artistDo) FindInBatches(result *[]*model.Artist, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a artistDo) Attrs(attrs ...field.AssignExpr) IArtistDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a artistDo) Assign(attrs ...field.AssignExpr) IArtistDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a artistDo) Joins(fields ...field.RelationField) IArtistDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a artistDo) Preload(fields ...field.RelationField) IArtistDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a artistDo) FirstOrInit() (*model.Artist, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Artist), nil
	}
}

func (a artistDo) FirstOrCreate() (*model.Artist, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Artist), nil
	}
}

func (a artistDo) FindByPage(offset int, limit int) (result []*model.Artist, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a artistDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a artistDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a artistDo) Delete(models ...*model.Artist) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *artistDo) withDO(do gen.Dao) *artistDo {
	a.DO = *do.(*gen.DO)
	return a
}
