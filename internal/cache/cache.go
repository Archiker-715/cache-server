package cache

import "time"

type reqPort string
type reqURL string

type reqEntity struct {
	reqURL  reqURL
	reqBody string
	method  string
}

type cacheItem struct {
	response []byte
	ttl      *time.Timer
}

type Cache struct {
	cache map[reqPort]map[reqEntity]cacheItem
}

func InitCache() *Cache {
	return &Cache{cache: make(map[reqPort]map[reqEntity]cacheItem, 0)}
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
	cacheItem := c.cache[port][req]
	return cacheItem.response
}

func (c *Cache) SaveCache(port reqPort, url reqURL, requestBody, method string, responseBody []byte) {
	req := reqEntity{
		reqURL:  url,
		reqBody: requestBody,
		method:  method,
	}
	if c.cache[port] == nil {
		c.cache[port] = make(map[reqEntity]cacheItem)
	}
	c.cache[port][req] = cacheItem{
		response: responseBody,
		ttl: time.AfterFunc(5*time.Second, func() {
			c.deleteCachedReq(port, req)
		}),
	}
}

func (c *Cache) ClearCache() {
	c.cache = make(map[reqPort]map[reqEntity]cacheItem, 0)
}

func (c *Cache) deleteCachedReq(port reqPort, req reqEntity) {
	if r, ok := c.cache[port]; ok {
		if cacheItem, ok := r[req]; ok {
			cacheItem.ttl.Stop()
			delete(r, req)
		}
	}
}

func (c *Cache) ReflectReqPort(port string) reqPort {
	return reqPort(port)
}

func (c *Cache) ReflectReqURL(URL string) reqURL {
	return reqURL(URL)
}
