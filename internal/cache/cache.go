package cache

type reqPort string
type reqURL string
type response []byte

var cache = make(map[reqPort]map[reqURL]response, 0)

func Cache(port, url string, responseBody []byte) (wasCached bool) {
	requestPort := reqPort(port)
	requestURL := reqURL(url)

	if cached(requestPort, requestURL) {
		return true
	} else {
		saveCache(requestPort, requestURL, responseBody)
		return false
	}
}

func cached(port reqPort, url reqURL) bool {
	if _, ok := cache[port][url]; ok {
		return true
	}
	return false
}

func saveCache(port reqPort, url reqURL, responseBody []byte) {
	cache[port][url] = responseBody
}

func ClearCache() {
	cache = make(map[reqPort]map[reqURL]response, 0)
}
