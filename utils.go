package ytsearch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

var headers = map[string]string{
	"Accept":       "*/*",
	"Content-Type": "application/json",
	"Host":         "www.youtube.com",
	"Referer":      "https://www.youtube.com/",
	"Origin":       "https://www.youtube.com",
	"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
}

func GetResults(q string) []map[string]interface{} {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://www.youtube.com/youtubei/v1/search?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8", strings.NewReader(`{"query":"`+q+`","context":{"client":{"hl":"en","gl":"US","clientName":"MWEB","clientVersion":"2.20211109.01.00"}}}`))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	contents, _, _, _ := jsonparser.Get(body, "contents", "sectionListRenderer", "contents", "[0]", "itemSectionRenderer", "contents")
	var data []map[string]interface{}
	json.Unmarshal(contents, &data)
	return data
}

func GetBestThumb(list []byte) string {
	var list_ []interface{}
	var hq int
	json.Unmarshal(list, &list_)
	hq = len(list_) - 1
	return list_[hq].(map[string]interface{})["url"].(string)
}

func ParseData(data map[string]interface{}) (YTSearchResult, error) {
	vdata, _ := json.Marshal(data)
	videoId, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "videoId")
	if videoId == "" {
		return YTSearchResult{}, errors.New("videoId not found")
	}
	title, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "title", "runs", "[0]", "text")
	publishedTime, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "publishedTimeText", "runs", "[0]", "text")
	viewCount, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "shortViewCountText", "runs", "[0]", "text")
	duration, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "lengthText", "accessibility", "accessibilityData", "label")
	channel, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "longBylineText", "runs", "[0]", "text")
	channelId, _ := jsonparser.GetString(vdata, "compactVideoRenderer", "longBylineText", "runs", "[0]", "navigationEndpoint", "browseEndpoint", "browseId")
	tlist, _, _, _ := jsonparser.Get(vdata, "compactVideoRenderer", "thumbnail", "thumbnails")

	return YTSearchResult{Title: title, VideoId: videoId, PublishTime: publishedTime, Channel: channel, ChannelId: channelId, Views: viewCount, Duration: duration, Thumbnail: GetBestThumb(tlist)}, nil
}
