package gengin

import (
	"projects_template/internal/entity/global"
	"projects_template/internal/entity/rpc"
	"projects_template/internal/transaction"
	"projects_template/tools/logger"

	"github.com/sirupsen/logrus"
)

type LoadDataFunc[T any] func(ts transaction.Session) (T, error)

func LoadData[T any](
	rc *rpc.Context,
	log *logrus.Logger,
	sessionManager transaction.SessionManager,
	f LoadDataFunc[T],
	stateName string,
	canErrNoData bool,
	needCommit bool,
) {
	ts := sessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		log.Errorln("ошибка открытия транзакции; ошибка", err)
		rc.ReturnError(global.ErrInternalError)
		return
	}
	defer ts.Rollback()

	data, err := f(ts)
	if err != nil {
		if err == global.ErrNoData && canErrNoData {
			rc.ReturnStateResult(stateName, nil)
			return
		}

		rc.GinContext.Error(logger.ErrLog(err))
		rc.ReturnError(err)
		return
	}

	if needCommit {
		err = ts.Commit()
		if err != nil {
			log.Errorln("ошибка при коммите", err)
			rc.ReturnError(global.ErrInternalError)
			return
		}
	}

	rc.ReturnStateResult(stateName, data)
}
