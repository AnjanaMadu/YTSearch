package ytsearch

import (
	"errors"
	"strings"
)

func Search(q string) ([]YTSearchResult, error) {
	contents := GetResults(strings.ReplaceAll(q, " ", "+"))
	results := make([]YTSearchResult, len(contents))
	
	// Check if there are no results
	if len(contents) == 0 {
		return results, errors.New("NO RESULTS FOUND")
	}

	// Parse the results into YTSearchResult structs
	for i, result := range contents {
		obj, err := ParseData(result)
		if err != nil {
			continue
		}
		results[i] = obj
	}
	return results, nil
}
