package ytsearch

type YTSearchResult struct {
	// Title of the video
	Title       string `json:"title"`

	// Video ID
	// https://www.youtube.com/watch?v=<VIDEO_ID>
	VideoId     string `json:"videoId"`

	// Published time
	// Example: 1 month ago
	PublishTime string `json:"publishTime"`

	// Channel name
	Channel     string `json:"channel"`

	// Channel ID
	// Example: UC-lHJZR3Gqxm24_Vd_AJ5Yw
	ChannelId   string `json:"channelId"`

	// View count in short form
	// Example: 1.2M
	Views       string `json:"views"`

	// Duration in long form
	// Example: 1 hour, 2 minutes
	Duration    string `json:"duration"`

	// Best quality thumbnail
	Thumbnail   string `json:"thumbnail"`
}
