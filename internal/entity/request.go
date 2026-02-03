package entity

import "github.com/Archiker-715/cache-server/internal/cache"

type Request struct {
	Port   string
	Method string
	Url    string
	Body   string
	Cache  *cache.Cache
}
