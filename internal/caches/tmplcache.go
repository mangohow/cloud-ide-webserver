package caches

import (
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/pkg/cache"
)

// 加载mysql中的SpaceTemplate到内存中，数据量不大

type TmplCache struct {
	cache *cache.Cache
	dao   *dao.SpaceTemplateDao
}

func newTmplCache(dao *dao.SpaceTemplateDao) *TmplCache {
	return &TmplCache{
		cache: cache.New("spaceTmpl"),
		dao:   dao,
	}
}

const (
	TmplsKey = "tmpls"
	KindsKey = "kinds"
)

func (t *TmplCache) LoadCache() {
	tmpls, err := t.dao.GetAllUsingTmpl()
	if err != nil {
		panic(err)
	}

	tpls := make([]*model.SpaceTemplate, len(tmpls))

	for i := 0; i < len(tmpls); i++ {
		tmpl := tmpls[i]
		tpls[i] = &tmpl
	}
	t.cache.Set(TmplsKey, tpls)

	kinds, err := t.dao.GetAllTmplKind()
	if err != nil {
		panic(err)
	}
	kds := make([]*model.TmplKind, len(kinds))
	for i := 0; i < len(kinds); i++ {
		k := kinds[i]
		kds[i] = &k
	}
	t.cache.Set(KindsKey, kds)
}

func (t *TmplCache) GetTmpl(key uint32) *model.SpaceTemplate {
	item, ok := t.cache.Get(TmplsKey)
	if !ok {
		return nil
	}

	tps, ok := item.([]*model.SpaceTemplate)

	tmpl := *tps[key]
	return &tmpl
}

func (t *TmplCache) GetAllTmpl() []*model.SpaceTemplate {
	get, ok := t.cache.Get(TmplsKey)
	if !ok {
		return nil
	}
	items := get.([]*model.SpaceTemplate)

	// 拷贝一份
	tmpls := make([]*model.SpaceTemplate, len(items))
	for i := 0; i < len(items); i++ {
		tp := *items[i]
		tmpls[i] = &tp
	}

	return tmpls
}

func (t *TmplCache) GetAllKinds() []*model.TmplKind {
	get, ok := t.cache.Get(KindsKey)
	if !ok {
		return nil
	}
	items := get.([]*model.TmplKind)

	kinds := make([]*model.TmplKind, len(items))
	for i := 0; i < len(items); i++ {
		k := *items[i]
		items[i] = &k
	}

	return kinds
}