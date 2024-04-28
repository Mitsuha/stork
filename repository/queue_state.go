package repository

import (
	"context"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
)

type QueueStateRepo struct {
	*model.QueueState `json:",inline"`
	Songs             []*model.Songs `json:"songs"`
	CurrentSong       *model.Songs   `json:"current_song"`
}

func UsersQueueState(ctx context.Context, uid int) (*QueueStateRepo, error) {
	state, err := dao.QueueState.WithContext(ctx).WhereUser(uid)

	if err != nil {
		return nil, err
	}

	songs, err := dao.Songs.WithContext(ctx).IdIn(state.SongIds)
	if err != nil {
		return nil, err
	}

	var currentSong *model.Songs
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
