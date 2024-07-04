package lastfm

type trackSearchResponse struct {
	Results struct {
		TrackMatches struct {
			Track []Track `json:"track"`
		} `json:"trackmatches"`
		Attr struct{} `json:"@attr"`
	} `json:"results"`
}

type Track struct {
	Name       string `json:"name"`
	Artist     string `json:"artist"`
	Url        string `json:"url"`
	Streamable string `json:"streamable"`
	Listeners  string `json:"listeners"`
	Image      []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Mbid string `json:"mbid"`
}

type trackInfoResponse struct {
	Track *TrackInfo `json:"track"`
}
type TrackInfo struct {
	Name      string `json:"name"`
	Mbid      string `json:"mbid"`
	Url       string `json:"url"`
	Duration  string `json:"duration"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Artist    struct {
		Name string `json:"name"`
		Mbid string `json:"mbid"`
		Url  string `json:"url"`
	} `json:"artist"`
	Album   TrackAlbum `json:"album"`
	Toptags struct {
		Tag []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"tag"`
	} `json:"toptags"`
}

type TrackAlbum struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Mbid   string `json:"mbid"`
	Url    string `json:"url"`
	Image  []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Attr struct {
		Position string `json:"position"`
	} `json:"@attr"`
}

func (t TrackAlbum) GetImage() string {
	if len(t.Image) == 0 {
		return ""
	}
	return t.Image[len(t.Image)-1].Text
}
