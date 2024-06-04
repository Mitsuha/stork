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

func newSong(db *gorm.DB, opts ...gen.DOOption) song {
	_song := song{}

	_song.songDo.UseDB(db, opts...)
	_song.songDo.UseModel(&model.Song{})

	tableName := _song.songDo.TableName()
	_song.ALL = field.NewAsterisk(tableName)
	_song.ID = field.NewString(tableName, "id")
	_song.AlbumID = field.NewInt(tableName, "album_id")
	_song.Title = field.NewString(tableName, "title")
	_song.Length = field.NewFloat64(tableName, "length")
	_song.Track = field.NewInt(tableName, "track")
	_song.Disc = field.NewInt(tableName, "disc")
	_song.Lyrics = field.NewString(tableName, "lyrics")
	_song.Path = field.NewString(tableName, "path")
	_song.Mtime = field.NewInt(tableName, "mtime")
	_song.CreatedAt = field.NewTime(tableName, "created_at")
	_song.UpdatedAt = field.NewTime(tableName, "updated_at")
	_song.ArtistID = field.NewInt(tableName, "artist_id")
	_song.Year = field.NewInt(tableName, "year")
	_song.Genre = field.NewString(tableName, "genre")
	_song.Interaction = songHasOneInteraction{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Interaction", "model.Interaction"),
	}

	_song.Album = songBelongsToAlbum{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Album", "model.Album"),
		Artist: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Album.Artist", "model.Artist"),
		},
	}

	_song.Artist = songBelongsToArtist{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Artist", "model.Artist"),
	}

	_song.fillFieldMap()

	return _song
}

type song struct {
	songDo songDo

	ALL         field.Asterisk
	ID          field.String
	AlbumID     field.Int
	Title       field.String
	Length      field.Float64
	Track       field.Int
	Disc        field.Int
	Lyrics      field.String
	Path        field.String
	Mtime       field.Int
	CreatedAt   field.Time
	UpdatedAt   field.Time
	ArtistID    field.Int
	Year        field.Int
	Genre       field.String
	Interaction songHasOneInteraction

	Album songBelongsToAlbum

	Artist songBelongsToArtist

	fieldMap map[string]field.Expr
}

func (s song) Table(newTableName string) *song {
	s.songDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s song) As(alias string) *song {
	s.songDo.DO = *(s.songDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *song) updateTableName(table string) *song {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.AlbumID = field.NewInt(table, "album_id")
	s.Title = field.NewString(table, "title")
	s.Length = field.NewFloat64(table, "length")
	s.Track = field.NewInt(table, "track")
	s.Disc = field.NewInt(table, "disc")
	s.Lyrics = field.NewString(table, "lyrics")
	s.Path = field.NewString(table, "path")
	s.Mtime = field.NewInt(table, "mtime")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.ArtistID = field.NewInt(table, "artist_id")
	s.Year = field.NewInt(table, "year")
	s.Genre = field.NewString(table, "genre")

	s.fillFieldMap()

	return s
}

func (s *song) WithContext(ctx context.Context) ISongDo { return s.songDo.WithContext(ctx) }

func (s song) TableName() string { return s.songDo.TableName() }

func (s song) Alias() string { return s.songDo.Alias() }

func (s song) Columns(cols ...field.Expr) gen.Columns { return s.songDo.Columns(cols...) }

func (s *song) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *song) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 17)
	s.fieldMap["id"] = s.ID
	s.fieldMap["album_id"] = s.AlbumID
	s.fieldMap["title"] = s.Title
	s.fieldMap["length"] = s.Length
	s.fieldMap["track"] = s.Track
	s.fieldMap["disc"] = s.Disc
	s.fieldMap["lyrics"] = s.Lyrics
	s.fieldMap["path"] = s.Path
	s.fieldMap["mtime"] = s.Mtime
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["artist_id"] = s.ArtistID
	s.fieldMap["year"] = s.Year
	s.fieldMap["genre"] = s.Genre

}

func (s song) clone(db *gorm.DB) song {
	s.songDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s song) replaceDB(db *gorm.DB) song {
	s.songDo.ReplaceDB(db)
	return s
}

type songHasOneInteraction struct {
	db *gorm.DB

	field.RelationField
}

func (a songHasOneInteraction) Where(conds ...field.Expr) *songHasOneInteraction {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a songHasOneInteraction) WithContext(ctx context.Context) *songHasOneInteraction {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a songHasOneInteraction) Session(session *gorm.Session) *songHasOneInteraction {
	a.db = a.db.Session(session)
	return &a
}

func (a songHasOneInteraction) Model(m *model.Song) *songHasOneInteractionTx {
	return &songHasOneInteractionTx{a.db.Model(m).Association(a.Name())}
}

type songHasOneInteractionTx struct{ tx *gorm.Association }

func (a songHasOneInteractionTx) Find() (result *model.Interaction, err error) {
	return result, a.tx.Find(&result)
}

func (a songHasOneInteractionTx) Append(values ...*model.Interaction) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a songHasOneInteractionTx) Replace(values ...*model.Interaction) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a songHasOneInteractionTx) Delete(values ...*model.Interaction) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a songHasOneInteractionTx) Clear() error {
	return a.tx.Clear()
}

func (a songHasOneInteractionTx) Count() int64 {
	return a.tx.Count()
}

type songBelongsToAlbum struct {
	db *gorm.DB

	field.RelationField

	Artist struct {
		field.RelationField
	}
}

func (a songBelongsToAlbum) Where(conds ...field.Expr) *songBelongsToAlbum {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a songBelongsToAlbum) WithContext(ctx context.Context) *songBelongsToAlbum {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a songBelongsToAlbum) Session(session *gorm.Session) *songBelongsToAlbum {
	a.db = a.db.Session(session)
	return &a
}

func (a songBelongsToAlbum) Model(m *model.Song) *songBelongsToAlbumTx {
	return &songBelongsToAlbumTx{a.db.Model(m).Association(a.Name())}
}

type songBelongsToAlbumTx struct{ tx *gorm.Association }

func (a songBelongsToAlbumTx) Find() (result *model.Album, err error) {
	return result, a.tx.Find(&result)
}

func (a songBelongsToAlbumTx) Append(values ...*model.Album) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a songBelongsToAlbumTx) Replace(values ...*model.Album) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a songBelongsToAlbumTx) Delete(values ...*model.Album) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a songBelongsToAlbumTx) Clear() error {
	return a.tx.Clear()
}

func (a songBelongsToAlbumTx) Count() int64 {
	return a.tx.Count()
}

type songBelongsToArtist struct {
	db *gorm.DB

	field.RelationField
}

func (a songBelongsToArtist) Where(conds ...field.Expr) *songBelongsToArtist {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a songBelongsToArtist) WithContext(ctx context.Context) *songBelongsToArtist {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a songBelongsToArtist) Session(session *gorm.Session) *songBelongsToArtist {
	a.db = a.db.Session(session)
	return &a
}

func (a songBelongsToArtist) Model(m *model.Song) *songBelongsToArtistTx {
	return &songBelongsToArtistTx{a.db.Model(m).Association(a.Name())}
}

type songBelongsToArtistTx struct{ tx *gorm.Association }

func (a songBelongsToArtistTx) Find() (result *model.Artist, err error) {
	return result, a.tx.Find(&result)
}

func (a songBelongsToArtistTx) Append(values ...*model.Artist) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a songBelongsToArtistTx) Replace(values ...*model.Artist) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a songBelongsToArtistTx) Delete(values ...*model.Artist) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a songBelongsToArtistTx) Clear() error {
	return a.tx.Clear()
}

func (a songBelongsToArtistTx) Count() int64 {
	return a.tx.Count()
}

type songDo struct{ gen.DO }

type ISongDo interface {
	gen.SubQuery
	Debug() ISongDo
	WithContext(ctx context.Context) ISongDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISongDo
	WriteDB() ISongDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISongDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISongDo
	Not(conds ...gen.Condition) ISongDo
	Or(conds ...gen.Condition) ISongDo
	Select(conds ...field.Expr) ISongDo
	Where(conds ...gen.Condition) ISongDo
	Order(conds ...field.Expr) ISongDo
	Distinct(cols ...field.Expr) ISongDo
	Omit(cols ...field.Expr) ISongDo
	Join(table schema.Tabler, on ...field.Expr) ISongDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISongDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISongDo
	Group(cols ...field.Expr) ISongDo
	Having(conds ...gen.Condition) ISongDo
	Limit(limit int) ISongDo
	Offset(offset int) ISongDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISongDo
	Unscoped() ISongDo
	Create(values ...*model.Song) error
	CreateInBatches(values []*model.Song, batchSize int) error
	Save(values ...*model.Song) error
	First() (*model.Song, error)
	Take() (*model.Song, error)
	Last() (*model.Song, error)
	Find() ([]*model.Song, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Song, err error)
	FindInBatches(result *[]*model.Song, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Song) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISongDo
	Assign(attrs ...field.AssignExpr) ISongDo
	Joins(fields ...field.RelationField) ISongDo
	Preload(fields ...field.RelationField) ISongDo
	FirstOrInit() (*model.Song, error)
	FirstOrCreate() (*model.Song, error)
	FindByPage(offset int, limit int) (result []*model.Song, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISongDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindByID(id string) (result *model.Song, err error)
	IdIn(ids []string) (result []*model.Song, err error)
	CountAndLength() (result *model.CountAndLength, err error)
	MostPlayed(uid int, limit int) (result []*model.Song, err error)
	RecentlyPlayed(uid int, limit int) (result []*model.Song, err error)
	RecentlyAdded(uid int, limit int) (result []*model.Song, err error)
	FindByPlaylist(pid int) (result []*model.Song, err error)
	Favorite(uid int) (result []*model.Song, err error)
}

// FindByID SELECT * FROM @@table where id = @id
func (s songDo) FindByID(id string) (result *model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM songs where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// IdIn SELECT * FROM @@table WHERE id IN (@ids)
func (s songDo) IdIn(ids []string) (result []*model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT * FROM songs WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// CountAndLength SELECT COUNT(*) AS count, SUM(length) AS length FROM @@table
func (s songDo) CountAndLength() (result *model.CountAndLength, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT COUNT(*) AS count, SUM(length) AS length FROM songs ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// MostPlayed SELECT @@table .*, albums.name, artists.name, interactions.liked, interactions.play_count
// FROM @@table
// LEFT JOIN interactions ON interactions.song_id = @@table .id AND interactions.user_id = @uid
// JOIN albums ON @@table .album_id = albums.id
// JOIN artists ON @@table .artist_id = artists.id
// WHERE interactions.play_count > 0
// ORDER BY interactions.play_count DESC
// LIMIT @limit
func (s songDo) MostPlayed(uid int, limit int) (result []*model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	params = append(params, limit)
	generateSQL.WriteString("SELECT songs .*, albums.name, artists.name, interactions.liked, interactions.play_count FROM songs LEFT JOIN interactions ON interactions.song_id = songs .id AND interactions.user_id = ? JOIN albums ON songs .album_id = albums.id JOIN artists ON songs .artist_id = artists.id WHERE interactions.play_count > 0 ORDER BY interactions.play_count DESC LIMIT ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// RecentlyPlayed SELECT @@table .* FROM @@table
// LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = @uid
// ORDER BY interactions.last_played_at DESC LIMIT @limit
func (s songDo) RecentlyPlayed(uid int, limit int) (result []*model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	params = append(params, limit)
	generateSQL.WriteString("SELECT songs .* FROM songs LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = ? ORDER BY interactions.last_played_at DESC LIMIT ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// RecentlyAdded SELECT @@table .* FROM @@table LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = @uid
// ORDER BY songs.created_at DESC LIMIT @limit
func (s songDo) RecentlyAdded(uid int, limit int) (result []*model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	params = append(params, limit)
	generateSQL.WriteString("SELECT songs .* FROM songs LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = ? ORDER BY songs.created_at DESC LIMIT ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// FindByPlaylist SELECT @@table .* FROM @@table LEFT JOIN playlist_song ON playlist_song.song_id = songs.id WHERE playlist_song.playlist_id = @pid
func (s songDo) FindByPlaylist(pid int) (result []*model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, pid)
	generateSQL.WriteString("SELECT songs .* FROM songs LEFT JOIN playlist_song ON playlist_song.song_id = songs.id WHERE playlist_song.playlist_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// Favorite SELECT @@table .* FROM @@table LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = @uid
// AND interactions.liked = 1
func (s songDo) Favorite(uid int) (result []*model.Song, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	generateSQL.WriteString("SELECT songs .* FROM songs LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = ? AND interactions.liked = 1 ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s songDo) Debug() ISongDo {
	return s.withDO(s.DO.Debug())
}

func (s songDo) WithContext(ctx context.Context) ISongDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s songDo) ReadDB() ISongDo {
	return s.Clauses(dbresolver.Read)
}

func (s songDo) WriteDB() ISongDo {
	return s.Clauses(dbresolver.Write)
}

func (s songDo) Session(config *gorm.Session) ISongDo {
	return s.withDO(s.DO.Session(config))
}

func (s songDo) Clauses(conds ...clause.Expression) ISongDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s songDo) Returning(value interface{}, columns ...string) ISongDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s songDo) Not(conds ...gen.Condition) ISongDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s songDo) Or(conds ...gen.Condition) ISongDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s songDo) Select(conds ...field.Expr) ISongDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s songDo) Where(conds ...gen.Condition) ISongDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s songDo) Order(conds ...field.Expr) ISongDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s songDo) Distinct(cols ...field.Expr) ISongDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s songDo) Omit(cols ...field.Expr) ISongDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s songDo) Join(table schema.Tabler, on ...field.Expr) ISongDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s songDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISongDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s songDo) RightJoin(table schema.Tabler, on ...field.Expr) ISongDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s songDo) Group(cols ...field.Expr) ISongDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s songDo) Having(conds ...gen.Condition) ISongDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s songDo) Limit(limit int) ISongDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s songDo) Offset(offset int) ISongDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s songDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISongDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s songDo) Unscoped() ISongDo {
	return s.withDO(s.DO.Unscoped())
}

func (s songDo) Create(values ...*model.Song) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s songDo) CreateInBatches(values []*model.Song, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s songDo) Save(values ...*model.Song) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s songDo) First() (*model.Song, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Song), nil
	}
}

func (s songDo) Take() (*model.Song, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Song), nil
	}
}

func (s songDo) Last() (*model.Song, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Song), nil
	}
}

func (s songDo) Find() ([]*model.Song, error) {
	result, err := s.DO.Find()
	return result.([]*model.Song), err
}

func (s songDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Song, err error) {
	buf := make([]*model.Song, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s songDo) FindInBatches(result *[]*model.Song, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s songDo) Attrs(attrs ...field.AssignExpr) ISongDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s songDo) Assign(attrs ...field.AssignExpr) ISongDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s songDo) Joins(fields ...field.RelationField) ISongDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s songDo) Preload(fields ...field.RelationField) ISongDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s songDo) FirstOrInit() (*model.Song, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Song), nil
	}
}

func (s songDo) FirstOrCreate() (*model.Song, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Song), nil
	}
}

func (s songDo) FindByPage(offset int, limit int) (result []*model.Song, count int64, err error) {
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

func (s songDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s songDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s songDo) Delete(models ...*model.Song) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *songDo) withDO(do gen.Dao) *songDo {
	s.DO = *do.(*gen.DO)
	return s
}
