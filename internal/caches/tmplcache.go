package caches

import (
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/pkg/cache"
	"strconv"
	"sync"
)

// 加载mysql中的SpaceTemplate到内存中，数据量不大

type TmplCache struct {
	cache *cache.Cache
	dao   *dao.SpaceTemplateDao
	once  sync.Once
}

func newTmplCache(dao *dao.SpaceTemplateDao) *TmplCache {
	return &TmplCache{
		cache: cache.New("spaceTmpl"),
		dao:   dao,
	}
}

func (t *TmplCache) LoadCache() {
	tmpls, err := t.dao.GetAllUsingTmpl()
	if err != nil {
		panic(err)
	}

	// 巨坑，不能 使用_, item := range tmlpls 然后添加, item不是指针
	for i, _ := range tmpls {
		tmpl := tmpls[i]
		t.cache.Set(strconv.Itoa(int(tmpl.Id)), &tmpl)
	}
}

func (t *TmplCache) Get(key uint32) *model.SpaceTemplate {
	item, ok := t.cache.GetByInt(int(key))
	if !ok {
		return nil
	}
	// 复制一份，防止外面把缓存数据给修改了
	tp := item.(*model.SpaceTemplate)
	tmpl := *tp
	return &tmpl
}
