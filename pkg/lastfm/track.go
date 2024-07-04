package lastfm

import (
	"errors"
	"github.com/goccy/go-json"
)

func (l *Lastfm) TrackSearch(track string, option Options) ([]Track, error) {
	params := Options{
		"method":  {"track.search"},
		"track":   {track},
		"api_key": {l.apiKey},
		"format":  {"json"},
	}

	req := newGetRequest(params.Merge(option))

	resp, err := l.request(req)
	if err != nil {
		return nil, err
	}

	var result trackSearchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Results.TrackMatches.Track, nil
}

func (l *Lastfm) TrackGetInfo(option Options) (*TrackInfo, error) {
	params := Options{
		"method":  {"track.getInfo"},
		"api_key": {l.apiKey},
		"format":  {"json"},
	}

	req := newGetRequest(params.Merge(option))

	resp, err := l.request(req)
	if err != nil {
		return nil, err
	}

	var result trackInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	if result.Track == nil {
		return nil, errors.New("track not found")
	}

	return result.Track, nil
}
