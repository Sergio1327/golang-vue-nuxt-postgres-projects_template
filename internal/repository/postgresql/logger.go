package postgresql

import (
	"projects_template/internal/repository"
	"projects_template/internal/transaction"
)

type loggerRepository struct {
}

func NewLogger() repository.Logger {
	return &loggerRepository{}
}

func (r *loggerRepository) SaveLog(ts transaction.Session, logText string) error {
	query := `
	insert into logs
	(log_text)
	values( $1 )`

	_, err := SqlxTx(ts).Exec(query, logText)

	return err
}
