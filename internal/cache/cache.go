package cache

type reqPort string
type reqURL string
type response []byte

type Cache struct {
	cache  map[reqPort]map[reqURL]response
	cached bool
}

func InitCache() *Cache {
	return &Cache{cache: make(map[reqPort]map[reqURL]response, 0)}
}

func (c Cache) Cached(port reqPort, url reqURL) bool {
	if _, ok := c.cache[port][url]; ok {
		return true
	}
	return false
}

func (c Cache) GetCache(port reqPort, url reqURL) []byte {
	return c.cache[port][url]
}

func (c Cache) SaveCache(port reqPort, url reqURL, responseBody []byte) {
	c.cache[port][url] = responseBody
}

func (c *Cache) ClearCache() {
	c.cache = make(map[reqPort]map[reqURL]response, 0)
}

func (c *Cache) ReflectReqPort(port string) reqPort {
	return reqPort(port)
}

func (c *Cache) ReflectReqURL(URL string) reqURL {
	return reqURL(URL)
}
