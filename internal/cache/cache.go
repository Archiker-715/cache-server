package cache

type reqPort string
type reqURL string
type response []byte

type reqEntity struct {
	reqURL  reqURL
	reqBody string
	method  string
}

type Cache struct {
	cache map[reqPort]map[reqEntity]response
}

func InitCache() *Cache {
	return &Cache{cache: make(map[reqPort]map[reqEntity]response, 0)}
}

func (c Cache) Cached(port reqPort, url reqURL, requestBody, method string) bool {
	req := reqEntity{
		reqURL:  url,
		reqBody: requestBody,
		method:  method,
	}
	if _, ok := c.cache[port][req]; ok {
		return true
	}
	return false
}

func (c Cache) GetCache(port reqPort, url reqURL, requestBody, method string) []byte {
	req := reqEntity{
		reqURL:  url,
		reqBody: requestBody,
		method:  method,
	}
	return c.cache[port][req]
}

func (c *Cache) SaveCache(port reqPort, url reqURL, requestBody, method string, responseBody []byte) {
	req := reqEntity{
		reqURL:  url,
		reqBody: requestBody,
		method:  method,
	}
	if c.cache[port] == nil {
		c.cache[port] = make(map[reqEntity]response)
	}
	c.cache[port][req] = responseBody
}

func (c *Cache) ClearCache() {
	c.cache = make(map[reqPort]map[reqEntity]response, 0)
}

func (c *Cache) ReflectReqPort(port string) reqPort {
	return reqPort(port)
}

func (c *Cache) ReflectReqURL(URL string) reqURL {
	return reqURL(URL)
}
