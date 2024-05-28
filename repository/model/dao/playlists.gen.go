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

func newPlaylist(db *gorm.DB, opts ...gen.DOOption) playlist {
	_playlist := playlist{}

	_playlist.playlistDo.UseDB(db, opts...)
	_playlist.playlistDo.UseModel(&model.Playlist{})

	tableName := _playlist.playlistDo.TableName()
	_playlist.ALL = field.NewAsterisk(tableName)
	_playlist.ID = field.NewInt(tableName, "id")
	_playlist.UserID = field.NewInt(tableName, "user_id")
	_playlist.Name = field.NewString(tableName, "name")
	_playlist.Rules = field.NewString(tableName, "rules")
	_playlist.FolderID = field.NewString(tableName, "folder_id")
	_playlist.CreatedAt = field.NewTime(tableName, "created_at")
	_playlist.UpdatedAt = field.NewTime(tableName, "updated_at")

	_playlist.fillFieldMap()

	return _playlist
}

type playlist struct {
	playlistDo playlistDo

	ALL       field.Asterisk
	ID        field.Int
	UserID    field.Int
	Name      field.String
	Rules     field.String
	FolderID  field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p playlist) Table(newTableName string) *playlist {
	p.playlistDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p playlist) As(alias string) *playlist {
	p.playlistDo.DO = *(p.playlistDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *playlist) updateTableName(table string) *playlist {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt(table, "id")
	p.UserID = field.NewInt(table, "user_id")
	p.Name = field.NewString(table, "name")
	p.Rules = field.NewString(table, "rules")
	p.FolderID = field.NewString(table, "folder_id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *playlist) WithContext(ctx context.Context) IPlaylistDo { return p.playlistDo.WithContext(ctx) }

func (p playlist) TableName() string { return p.playlistDo.TableName() }

func (p playlist) Alias() string { return p.playlistDo.Alias() }

func (p playlist) Columns(cols ...field.Expr) gen.Columns { return p.playlistDo.Columns(cols...) }

func (p *playlist) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *playlist) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["name"] = p.Name
	p.fieldMap["rules"] = p.Rules
	p.fieldMap["folder_id"] = p.FolderID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p playlist) clone(db *gorm.DB) playlist {
	p.playlistDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p playlist) replaceDB(db *gorm.DB) playlist {
	p.playlistDo.ReplaceDB(db)
	return p
}

type playlistDo struct{ gen.DO }

type IPlaylistDo interface {
	gen.SubQuery
	Debug() IPlaylistDo
	WithContext(ctx context.Context) IPlaylistDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPlaylistDo
	WriteDB() IPlaylistDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPlaylistDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPlaylistDo
	Not(conds ...gen.Condition) IPlaylistDo
	Or(conds ...gen.Condition) IPlaylistDo
	Select(conds ...field.Expr) IPlaylistDo
	Where(conds ...gen.Condition) IPlaylistDo
	Order(conds ...field.Expr) IPlaylistDo
	Distinct(cols ...field.Expr) IPlaylistDo
	Omit(cols ...field.Expr) IPlaylistDo
	Join(table schema.Tabler, on ...field.Expr) IPlaylistDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPlaylistDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPlaylistDo
	Group(cols ...field.Expr) IPlaylistDo
	Having(conds ...gen.Condition) IPlaylistDo
	Limit(limit int) IPlaylistDo
	Offset(offset int) IPlaylistDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPlaylistDo
	Unscoped() IPlaylistDo
	Create(values ...*model.Playlist) error
	CreateInBatches(values []*model.Playlist, batchSize int) error
	Save(values ...*model.Playlist) error
	First() (*model.Playlist, error)
	Take() (*model.Playlist, error)
	Last() (*model.Playlist, error)
	Find() ([]*model.Playlist, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Playlist, err error)
	FindInBatches(result *[]*model.Playlist, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Playlist) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPlaylistDo
	Assign(attrs ...field.AssignExpr) IPlaylistDo
	Joins(fields ...field.RelationField) IPlaylistDo
	Preload(fields ...field.RelationField) IPlaylistDo
	FirstOrInit() (*model.Playlist, error)
	FirstOrCreate() (*model.Playlist, error)
	FindByPage(offset int, limit int) (result []*model.Playlist, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPlaylistDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindAll() (result []*model.Playlist, err error)
	FindByID(id int) (result *model.Playlist, err error)
	FindByUserID(uid int) (result []*model.Playlist, err error)
}

// FindAll SELECT * FROM @@table
func (p playlistDo) FindAll() (result []*model.Playlist, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM playlists ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// FindByID SELECT * FROM @@table WHERE id = @id
func (p playlistDo) FindByID(id int) (result *model.Playlist, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM playlists WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// FindByUserID SELECT * FROM @@table WHERE user_id = @uid
func (p playlistDo) FindByUserID(uid int) (result []*model.Playlist, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	generateSQL.WriteString("SELECT * FROM playlists WHERE user_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p playlistDo) Debug() IPlaylistDo {
	return p.withDO(p.DO.Debug())
}

func (p playlistDo) WithContext(ctx context.Context) IPlaylistDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p playlistDo) ReadDB() IPlaylistDo {
	return p.Clauses(dbresolver.Read)
}

func (p playlistDo) WriteDB() IPlaylistDo {
	return p.Clauses(dbresolver.Write)
}

func (p playlistDo) Session(config *gorm.Session) IPlaylistDo {
	return p.withDO(p.DO.Session(config))
}

func (p playlistDo) Clauses(conds ...clause.Expression) IPlaylistDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p playlistDo) Returning(value interface{}, columns ...string) IPlaylistDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p playlistDo) Not(conds ...gen.Condition) IPlaylistDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p playlistDo) Or(conds ...gen.Condition) IPlaylistDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p playlistDo) Select(conds ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p playlistDo) Where(conds ...gen.Condition) IPlaylistDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p playlistDo) Order(conds ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p playlistDo) Distinct(cols ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p playlistDo) Omit(cols ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p playlistDo) Join(table schema.Tabler, on ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p playlistDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p playlistDo) RightJoin(table schema.Tabler, on ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p playlistDo) Group(cols ...field.Expr) IPlaylistDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p playlistDo) Having(conds ...gen.Condition) IPlaylistDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p playlistDo) Limit(limit int) IPlaylistDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p playlistDo) Offset(offset int) IPlaylistDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p playlistDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPlaylistDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p playlistDo) Unscoped() IPlaylistDo {
	return p.withDO(p.DO.Unscoped())
}

func (p playlistDo) Create(values ...*model.Playlist) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p playlistDo) CreateInBatches(values []*model.Playlist, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p playlistDo) Save(values ...*model.Playlist) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p playlistDo) First() (*model.Playlist, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Playlist), nil
	}
}

func (p playlistDo) Take() (*model.Playlist, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Playlist), nil
	}
}

func (p playlistDo) Last() (*model.Playlist, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Playlist), nil
	}
}

func (p playlistDo) Find() ([]*model.Playlist, error) {
	result, err := p.DO.Find()
	return result.([]*model.Playlist), err
}

func (p playlistDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Playlist, err error) {
	buf := make([]*model.Playlist, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p playlistDo) FindInBatches(result *[]*model.Playlist, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p playlistDo) Attrs(attrs ...field.AssignExpr) IPlaylistDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p playlistDo) Assign(attrs ...field.AssignExpr) IPlaylistDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p playlistDo) Joins(fields ...field.RelationField) IPlaylistDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p playlistDo) Preload(fields ...field.RelationField) IPlaylistDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p playlistDo) FirstOrInit() (*model.Playlist, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Playlist), nil
	}
}

func (p playlistDo) FirstOrCreate() (*model.Playlist, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Playlist), nil
	}
}

func (p playlistDo) FindByPage(offset int, limit int) (result []*model.Playlist, count int64, err error) {
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

func (p playlistDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p playlistDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p playlistDo) Delete(models ...*model.Playlist) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *playlistDo) withDO(do gen.Dao) *playlistDo {
	p.DO = *do.(*gen.DO)
	return p
}
