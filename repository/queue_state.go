package repository

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/mitsuha/stork/internal/container"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type QueueStates struct {
	ctx context.Context
}

func NewQueueStates(ctx context.Context) *QueueStates {
	return &QueueStates{ctx: ctx}
}

type QueueStateRepo struct {
	*model.QueueState `json:",inline"`
	Songs             []*model.Song `json:"songs"`
	CurrentSong       *model.Song   `json:"current_song"`
}

func (q *QueueStates) UsersQueueState(uid int) (*QueueStateRepo, error) {
	state, err := dao.QueueState.WithContext(q.ctx).WhereUser(uid)

	if err != nil {
		return nil, err
	}

	songs, err := dao.Song.WithContext(q.ctx).IdIn(state.SongIds)
	if err != nil {
		return nil, err
	}

	var currentSong *model.Song
	for _, song := range songs {
		if song.ID == state.CurrentSongID {
			currentSong = song
			break
		}
	}

	return &QueueStateRepo{
		QueueState:  state,
		Songs:       songs,
		CurrentSong: currentSong,
	}, nil
}

func (q *QueueStates) UserQueueStateExist(uid int) bool {
	err := container.Singled.DB.Select("id").Where("user_id = ?", uid).First(&model.QueueState{}).Error

	return err == nil
}

func (q *QueueStates) UpdateQueueState(uid int, songs []string) error {
	exist := q.UserQueueStateExist(uid)
	if !exist {
		return dao.QueueState.WithContext(q.ctx).Create(&model.QueueState{
			UserID:        uid,
			SongIds:       songs,
			CurrentSongID: songs[0],
		})
	}

	marshal, _ := json.Marshal(songs)
	_, err := dao.QueueState.WithContext(q.ctx).Where(dao.QueueState.UserID.Eq(uid)).Updates(map[string]any{
		"song_ids": string(marshal),
	})
	return err
}

func (q *QueueStates) UpdatePlayState(uid int, song string, position int) error {
	exist := q.UserQueueStateExist(uid)
	if !exist {
		return dao.QueueState.WithContext(q.ctx).Create(&model.QueueState{
			UserID:           uid,
			SongIds:          []string{song},
			CurrentSongID:    song,
			PlaybackPosition: position,
		})
	}

	_, err := dao.QueueState.WithContext(q.ctx).Where(dao.QueueState.UserID.Eq(uid)).Updates(map[string]any{
		"current_song_id":   song,
		"playback_position": position,
	})
	return err
}
