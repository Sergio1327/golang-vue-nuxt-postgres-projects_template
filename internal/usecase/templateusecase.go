package usecase

import (
	"projects_template/bimport"
	"projects_template/internal/entity/global"
	"projects_template/internal/entity/template"
	"projects_template/internal/transaction"
	"projects_template/rimport"

	"github.com/sirupsen/logrus"
)

type TemplateUsecase struct {
	// вначале системные объекты - логи, конфиги
	log   *logrus.Logger
	dblog *logrus.Logger
	// далее репозитории
	rimport.RepositoryImports
	*bimport.BridgeImports
}

func NewTemplate(
	log, dblog *logrus.Logger,
	ri rimport.RepositoryImports,
	bi *bimport.BridgeImports,
) *TemplateUsecase {
	return &TemplateUsecase{
		log:               log,
		dblog:             dblog,
		RepositoryImports: ri,
		BridgeImports:     bi,
	}
}

func (u *TemplateUsecase) AwesomePublicMethod(ts transaction.Session) (data template.TemplateObject, err error) {

	data, err = u.Repository.Template.TemplateMethod(ts)
	if err != nil {
		u.dblog.Errorln("не удалось получить объект шаблона", err)
		err = global.ErrInternalError
		return
	}

	return
}
