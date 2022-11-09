package mock

import "net/http"

func CheckAllWebs(urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = checkWeb(url)
	}
	return results
}

func checkWeb(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}
