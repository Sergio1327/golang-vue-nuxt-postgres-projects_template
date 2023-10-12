package repository

import (
	"projects_template/internal/entity/template"
	"projects_template/internal/transaction"
)

type Logger interface {
	SaveLog(ts transaction.Session, logText string) error
}

type Template interface {
	TemplateMethod(ts transaction.Session) (templateData template.TemplateObject, err error)
}
