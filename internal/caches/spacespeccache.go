package caches

import (
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/pkg/cache"
	"strconv"
	"sync"
)

// 加载mysql中的SpaceSpec到内存中，数据量不大

type SpecCache struct {
	cache *cache.Cache
	dao   *dao.SpaceTemplateDao
	once  sync.Once
}

func newSpecCache(dao *dao.SpaceTemplateDao) *SpecCache {
	return &SpecCache{
		cache: cache.New("space-spec"),
		dao:   dao,
	}
}

func (c *SpecCache) LoadCache() {
	specs, err := c.dao.GetAllSpec()
	if err != nil {
		panic(err)
	}

	// 巨坑
	for i, _ := range specs {
		spec := specs[i]
		c.cache.Set(strconv.Itoa(int(spec.Id)), &spec)
	}
}

func (c *SpecCache) Get(key uint32) *model.SpaceSpec {
	item, ok := c.cache.GetByInt(int(key))
	if !ok {
		return nil
	}
	// 复制一份，防止外面把缓存数据给修改了
	s := item.(*model.SpaceSpec)
	spec := *s
	return &spec
}
