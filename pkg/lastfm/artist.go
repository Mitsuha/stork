package lastfm

import "github.com/goccy/go-json"

func (l *Lastfm) ArtistGetInfo(option Options) (*ArtistInfo, error) {
	params := Options{
		"method":  {"artist.getinfo"},
		"api_key": {l.apiKey},
		"format":  {"json"},
	}

	req := newGetRequest(params.Merge(option))

	resp, err := l.request(req)
	if err != nil {
		return nil, err
	}

	var result artistGetInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Artist, nil
}
