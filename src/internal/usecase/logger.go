package usecase

import (
	"github.com/sirupsen/logrus"
	"projects_template/rimport"
)

type LoggerUsecase struct {
	log *logrus.Logger
	rimport.RepositoryImports
}

func NewLogger(log *logrus.Logger,
	ri rimport.RepositoryImports,
) *LoggerUsecase {
	return &LoggerUsecase{
		log:               log,
		RepositoryImports: ri,
	}
}

func (u *LoggerUsecase) SaveLog(logText string) error {
	ts := u.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		u.log.Errorln("не удается стартовать транзакцию", err)
		return err
	}
	defer ts.Rollback()

	err := u.Repository.Logger.SaveLog(ts, logText)
	if err != nil {
		u.log.Errorln(err)
		return err
	}

	if err := ts.Commit(); err != nil {
		u.log.Errorln("не удается зафиксировать транзакцию", err)
		return err
	}

	return nil
}
