package queue

type UpdateQueueStateReq struct {
	Songs []string `json:"songs"`
}

type UpdatePlaybackStatusReq struct {
	Song     string `json:"song"`
	Position int    `json:"position"`
}
