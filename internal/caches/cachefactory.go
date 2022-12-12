package caches

import (
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"reflect"
	"sync"
)

type CacheFactory struct {
	caches map[reflect.Type]interface{}
	lock   sync.Mutex
}

func NewCacheFactory() *CacheFactory {
	return &CacheFactory{caches: make(map[reflect.Type]interface{})}
}

func (f *CacheFactory) TmplCache(dao *dao.SpaceTemplateDao) *TmplCache {
	t := reflect.TypeOf(&TmplCache{})
	f.lock.Lock()
	defer f.lock.Unlock()
	if c, ok := f.caches[t]; ok {
		return c.(*TmplCache)
	}

	cache := newTmplCache(dao)
	f.caches[t] = cache

	cache.LoadCache()

	return cache
}

func (f *CacheFactory) SpecCache(dao *dao.SpaceTemplateDao) *SpecCache {
	t := reflect.TypeOf(&SpecCache{})
	f.lock.Lock()
	defer f.lock.Unlock()
	if c, ok := f.caches[t]; ok {
		return c.(*SpecCache)
	}

	cache := newSpecCache(dao)
	f.caches[t] = cache
	cache.LoadCache()

	return cache
}
