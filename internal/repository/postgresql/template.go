package postgresql

import (
	"projects_template/internal/entity/template"
	"projects_template/internal/repository"
	"projects_template/internal/transaction"
)

type templateRepository struct {
}

func NewTemplate() repository.Template {
	return &templateRepository{}
}

func (r *templateRepository) TemplateMethod(ts transaction.Session) (templateData template.TemplateObject, err error) {
	return templateData, err
}
