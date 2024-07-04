package lastfm

type artistGetInfoResponse struct {
	Artist *ArtistInfo `json:"artist"`
}
type ArtistInfo struct {
	Name  string `json:"name"`
	Mbid  string `json:"mbid"`
	Url   string `json:"url"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Streamable string `json:"streamable"`
	Ontour     string `json:"ontour"`
	Tags       struct {
		Tag []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"tag"`
	} `json:"tags"`
	Bio struct {
		Published string `json:"published"`
		Summary   string `json:"summary"`
		Content   string `json:"content"`
	} `json:"bio"`
}

func (a ArtistInfo) GetImage() string {
	if len(a.Image) == 0 {
		return ""
	}
	return a.Image[len(a.Image)-1].Text

}
