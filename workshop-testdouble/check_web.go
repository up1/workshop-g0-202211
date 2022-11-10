package mock

import "net/http"

type Process2 func(func(int) bool) 

type Process func(url string) bool

func CheckAllWebs(p Process,urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = p(url)
	}
	return results
}

func CheckWeb(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}
