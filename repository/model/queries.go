package model

import "gorm.io/gen"

func ApplyQueries(g *gen.Generator) {
	var tables = []any{
		Album{}, Artist{}, Setting{}, User{}, PersonalAccessToken{}, Songs{}, QueueState{},
	}
	g.ApplyBasic(tables...)

	g.ApplyInterface(func(CommonQueries) {}, tables...)

	g.ApplyInterface(func(TokenQueries) {}, PersonalAccessToken{})

	g.ApplyInterface(func(SongQueries) {}, Songs{})

	g.ApplyInterface(func(QueueStateQueries) {}, QueueState{})
}

type CommonQueries interface {
	// FindAll SELECT * FROM @@table
	FindAll() ([]*gen.T, error)

	// FindByID SELECT * FROM @@table WHERE id = @id
	FindByID(id int) (*gen.T, error)
}

type TokenQueries interface {
	// WhereIDAndToken SELECT * FROM @@table WHERE id = @id AND token = @token
	WhereIDAndToken(id string, token string) (*gen.T, error)
}

type SongQueries interface {
	//IdIn SELECT * FROM @@table WHERE id IN (@ids)
	IdIn(ids []string) ([]*gen.T, error)
}

type QueueStateQueries interface {
	//WhereUser SELECT * FROM @@table WHERE user_id = @uid
	WhereUser(uid int) (*gen.T, error)
}
