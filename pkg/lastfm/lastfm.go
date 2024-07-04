package lastfm

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Lastfm struct {
	apiKey string
	secret string
	client http.Client
}

func New(apiKey, secret string) *Lastfm {
	return &Lastfm{
		apiKey: apiKey,
		secret: secret,
		client: http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func newGetRequest(query map[string][]string) *http.Request {
	u := url.URL{
		Scheme:   "https",
		Host:     "ws.audioscrobbler.com",
		Path:     "/2.0/",
		RawQuery: url.Values(query).Encode(),
	}

	return &http.Request{
		Method:     "GET",
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}
}

func (l *Lastfm) request(r *http.Request) ([]byte, error) {
	response, err := l.client.Do(r)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to search the track: " + response.Status)
	}

	return io.ReadAll(response.Body)
}
