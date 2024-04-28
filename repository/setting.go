package repository

import (
	"context"
	"encoding/json"
	"github.com/mitsuha/stork/repository/model/dao"
	"sync"
	"time"
)

type settings map[string]any

var _settings = settings{}

var _settingOnce sync.Once

func Settings() settings {
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
