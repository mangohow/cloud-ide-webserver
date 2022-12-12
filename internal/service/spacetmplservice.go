package service

import (
	"github.com/mangohow/cloud-ide-webserver/internal/dao"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
)

type SpaceTmplService struct {
	dao *dao.SpaceTemplateDao
}

func NewSpaceTmplService() *SpaceTmplService {
	return &SpaceTmplService{
		dao: dao.NewSpaceTemplateDao(),
	}
}

func (s *SpaceTmplService) GetAllUsingTmpl() (tmpls []model.SpaceTemplate, err error) {
	return s.dao.GetAllUsingTmpl()
}

func (s *SpaceTmplService) GetAllSpec() (specs []model.SpaceSpec, err error) {
	return s.dao.GetAllSpec()
}
