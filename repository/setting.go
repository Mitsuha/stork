package repository

import (
	"context"
	"encoding/json"
	"github.com/mitsuha/stork/repository/model/dao"
	"sync"
	"time"
)

type AppSettings map[string]any

func (s *AppSettings) MediaPath() string {
	return (*s)["media_path"].(string)
}

var _settings = AppSettings{}

var _settingOnce sync.Once

func Settings() AppSettings {
	_settingOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		defer cancel()

		models, err := dao.Setting.WithContext(ctx).FindAll()
		if err != nil {
			panic(err)
		}

		for _, model := range models {
			var value any

			if err := json.Unmarshal([]byte(model.Value), &value); err != nil {
				panic(err)
			}
			_settings[model.Key] = value
		}
	})
	return _settings
}
