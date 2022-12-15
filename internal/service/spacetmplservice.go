package service

import (
	"github.com/mangohow/cloud-ide-webserver/internal/caches"
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
)

type SpaceTmplService struct {
	dao *dao.SpaceTemplateDao
	tmplCache *caches.TmplCache
	specCache *caches.SpecCache
}

func NewSpaceTmplService() *SpaceTmplService {
	d := dao.NewSpaceTemplateDao()
	return &SpaceTmplService{
		dao: d,
		tmplCache: caches.CacheFactory().TmplCache(d),
		specCache: caches.CacheFactory().SpecCache(d),
	}
}

func (s *SpaceTmplService) GetAllUsingTmpl() (tmpls []*model.SpaceTemplate, err error) {
	return s.tmplCache.GetAll(), nil
}

func (s *SpaceTmplService) GetAllSpec() (specs []*model.SpaceSpec, err error) {
	return s.specCache.GetAll(), nil
}
