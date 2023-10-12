package bridge

import (
	"projects_template/internal/entity/template"
	"projects_template/internal/transaction"
)



type Template interface{
	AwesomePublicMethod(ts transaction.Session) (data template.TemplateObject, err error)
}
