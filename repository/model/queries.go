package model

import "gorm.io/gen"

func ApplyQueries(g *gen.Generator) {
	var tables = []any{
		Album{}, Artist{}, Setting{}, User{}, PersonalAccessToken{}, Artist{}, Album{}, Songs{}, QueueState{},
	}
	g.ApplyBasic(tables...)

	g.ApplyInterface(func(CommonQueries) {}, tables...)

	g.ApplyInterface(func(TokenQueries) {}, PersonalAccessToken{})

	g.ApplyInterface(func(ArtistQueries) {}, Artist{})

	g.ApplyInterface(func(AlbumQueries) {}, Album{})

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

type ArtistQueries interface {
	//MostPlayed SELECT @@table .*
	//FROM @@table
	//LEFT JOIN songs ON @@table .id = songs.artist_id
	//LEFT JOIN interactions ON interactions.song_id = songs.id AND interactions.user_id = @uid
	//GROUP BY artists.id, play_count, artists.name, artists.image, artists.created_at, artists.updated_at
	//ORDER BY play_count DESC
	//LIMIT @limit
	MostPlayed(uid int, limit int) ([]*gen.T, error)
}

type AlbumQueries interface {
	//MostPlayed SELECT @@table .*
	//	FROM @@table
	//	LEFT JOIN songs ON @@table .id = songs.album_id
	//	LEFT JOIN interactions ON songs.id = interactions.song_id AND interactions.user_id = @uid
	//	ORDER BY interactions.play_count DESC
	//	LIMIT @limit
	MostPlayed(uid int, limit int) ([]*gen.T, error)

	//RecentlyAdded SELECT * FROM @@table WHERE id != 0 ORDER BY created_at DESC LIMIT @limit
	RecentlyAdded(limit int) ([]*gen.T, error)
}

type SongQueries interface {
	//IdIn SELECT * FROM @@table WHERE id IN (@ids)
	IdIn(ids []string) ([]*gen.T, error)

	//CountAndLength SELECT COUNT(*) AS count, SUM(length) AS length FROM @@table
	CountAndLength() (*CountAndLength, error)

	//MostPlayed SELECT @@table .*, albums.name, artists.name, interactions.liked, interactions.play_count
	//FROM @@table
	//LEFT JOIN interactions ON interactions.song_id = @@table .id AND interactions.user_id = @uid
	//JOIN albums ON @@table .album_id = albums.id
	//JOIN artists ON @@table .artist_id = artists.id
	//WHERE interactions.play_count > 0
	//ORDER BY interactions.play_count DESC
	//LIMIT @limit
	MostPlayed(uid int, limit int) ([]*gen.T, error)

	//RecentlyPlayed SELECT @@table .* FROM @@table
	//LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = @uid
	//ORDER BY interactions.last_played_at DESC LIMIT @limit
	RecentlyPlayed(uid int, limit int) ([]*gen.T, error)

	//RecentlyAdded SELECT @@table .* FROM @@table LEFT JOIN interactions ON interactions.song_id = songs.id WHERE interactions.user_id = @uid
	//ORDER BY songs.created_at DESC LIMIT @limit
	RecentlyAdded(uid int, limit int) ([]*gen.T, error)
}

type QueueStateQueries interface {
	//WhereUser SELECT * FROM @@table WHERE user_id = @uid
	WhereUser(uid int) (*gen.T, error)
}
