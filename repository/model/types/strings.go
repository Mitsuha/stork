package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
)

type StringArray []string

func (j *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var result StringArray
	err := json.Unmarshal(bytes, &result)

	*j = result

	return err
}

func (j StringArray) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}
