package ytsearch

type YTSearchResult struct {
	Title       string `json:"title"`
	VideoId     string `json:"videoId"`
	PublishTime string `json:"publishTime"`
	Channel     string `json:"channel"`
	ChannelId   string `json:"channelId"`
	Views       string `json:"views"`
	Duration    string `json:"duration"`
	Thumbnail   string `json:"thumbnail"`
}
