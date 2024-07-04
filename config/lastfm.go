package config

var Lastfm *lastfm

type lastfm struct {
	Enable bool   `json:"enable" yaml:"enable"`
	APIKey string `json:"api_key" yaml:"api_key"`
	Secret string `json:"secret" yaml:"secret"`
}
